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
	db *sqlx.DB
}

func (repo *MemberRepoCQ) QueryByMemberID(ctx context.Context, memberID string, inWriteMode bool) (member domain.Member, err error) {
	queryBuilder := sq.Select("*").From(TableNameMember).Where(sq.Eq{"member_id": memberID})
	if inWriteMode {
		queryBuilder = queryBuilder.Suffix("LOCK IN SHARE MODE")
	}
	sqlString, args, _ := queryBuilder.ToSql()

	rdbms := uow.GetDBOrTxByContext(repo.db, ctx)
	if err = sqlx.GetContext(ctx, rdbms, member, sqlString, args...); err != nil {
		if err == sql.ErrNoRows {
			err = failure.Translate(err, basic.ErrNotFound)
			return
		}
		err = failure.Translate(err, basic.ErrDB)
		return
	}

	if log.Debug().Enabled() {
		log.Debug().Msgf("QueryByMemberID MemberID=%v: Dump VarInfo member=\n%v", memberID, spew.Sdump(member))
	}
	return member, nil
}

func (repo *MemberRepoCQ) AppendMember(ctx context.Context, m *domain.Member) (id int64, Err error) {
	sqlString, args, _ := sq.
		Insert(TableNameMember).
		Columns("member_id", "created_date", "self_intro").
		Values(m.MemberID, m.CreatedDate, m.SelfIntro).
		ToSql()

	rdbms := uow.GetDBOrTxByContext(repo.db, ctx)
	result, err := rdbms.ExecContext(ctx, sqlString, args...)
	if err != nil {
		return 0, failure.Translate(err, basic.ErrDB)
	}

	id, _ = result.LastInsertId()
	if log.Debug().Enabled() {
		// 由於 memberID 不是用 AUTO INCREMENT, 所以返回零
		log.Debug().Int64("member_id", id).Msg("MemberRepoCQ mysql insert row:")
	}
	return id, nil
}
