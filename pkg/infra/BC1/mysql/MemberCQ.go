package mysql

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/davecgh/go-spew/spew"
	"github.com/jmoiron/sqlx"
	"github.com/morikuni/failure"
	"github.com/rs/zerolog/log"

	"ddd/pkg/domain"
	"ddd/pkg/domain/basic"
	"ddd/pkg/technical/uow"
)

const TableNameMember = "member"

func NewMemberRepoCQ(db *sqlx.DB) *MemberRepoCQ {
	return &MemberRepoCQ{db: db}
}

type MemberRepoCQ struct {
	db         *sqlx.DB
	sqlBuilder MemberSQLBuilder
}

func (repo *MemberRepoCQ) FindByMemberID(ctx context.Context, memberID string, isUpdate bool) (member domain.Member, err error) {
	sqlString, args, _ := repo.sqlBuilder.FindByMemberID(memberID, isUpdate).ToSql()

	sqlExec := uow.ExecWithTxOrDB(ctx, repo.db)
	if err = sqlx.GetContext(ctx, sqlExec, member, sqlString, args...); err != nil {
		if err == sql.ErrNoRows {
			err = failure.Translate(err, basic.ErrNotFound)
			return
		}
		err = failure.Wrap(err)
		return
	}

	if log.Debug().Enabled() {
		log.Debug().Msgf("Find By MemberID=%v\n%v", memberID, spew.Sdump(member))
	}
	return member, nil
}

func (repo *MemberRepoCQ) AppendMember(ctx context.Context, m *domain.Member) (id int64, Err error) {
	sqlString, args, _ := repo.sqlBuilder.AppendMember(m).ToSql()
	result, err := repo.db.Exec(sqlString, args...)
	if err != nil {
		return 0, failure.Wrap(err)
	}

	id, _ = result.LastInsertId()
	if log.Debug().Enabled() {
		// 由於 memberID 不是用 AUTO INCREMENT, 所以返回零
		log.Debug().Int64("member_id", id).Msg("MemberRepoCQ mysql insert row:")
	}
	return id, nil
}

type MemberSQLBuilder struct{}

func (MemberSQLBuilder) FindByMemberID(memberID string, isUpdate bool) (selectBuilder sq.SelectBuilder) {
	selectBuilder = sq.Select("*").From(TableNameMember).Where(sq.Eq{"member_id": memberID})
	if isUpdate {
		// selectBuilder.Suffix("FOR UPDATE")
		selectBuilder.Suffix("LOCK IN SHARE MODE")
	}
	return
}

func (MemberSQLBuilder) AppendMember(m *domain.Member) sq.Sqlizer {
	return sq.
		Insert(TableNameMember).
		Columns("member_id", "created_date", "self_intro").
		Values(m.MemberID, m.CreatedDate, m.SelfIntro)
}
