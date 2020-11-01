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

func NewMemberRepo(db *sqlx.DB) *MemberRepo {
	return &MemberRepo{db: db}
}

type MemberRepo struct {
	db         *sqlx.DB
	sqlBuilder MemberRepoSQLBuilder
}

func (repo *MemberRepo) FindByID(ctx context.Context, memberID string) (domain.Member, error) {
	ext := uow.GetTxOrDB(ctx, repo.db)
	sqlString, args, _ := repo.sqlBuilder.FindByID(memberID).ToSql()
	member := new(domain.Member)

	err := sqlx.GetContext(ctx, ext, member, sqlString, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, failure.Translate(err, basic.ErrNotFound, failure.Message("mysql select"))
		}
		return nil, failure.Wrap(err, failure.Message("mysql select"))
	}

	if log.Debug().Enabled() {
		log.Debug().Msgf("MemberRepo find memberID=%v\n%v", memberID, spew.Sdump(member))
	}
	return member, err
}

func (repo *MemberRepo) Append(ctx context.Context, m *domain.Member) (id int64, Err error) {
	sqlString, args, _ := repo.sqlBuilder.Append(m).ToSql()
	result, err := repo.db.Exec(sqlString, args...)
	if err != nil {
		return 0, failure.Wrap(err, failure.Message("mysql insert"))
	}

	id, _ = result.LastInsertId()
	if log.Debug().Enabled() {
		// 由於 memberID 不是用 AUTO INCREMENT, 所以返回零
		log.Debug().Int64("member_id", id).Msg("MemberRepo mysql insert row:")
	}
	return id, nil
}

type MemberRepoSQLBuilder struct{}

func (MemberRepoSQLBuilder) FindByID(memberID string) sq.Sqlizer {
	return sq.
		Select("*").
		From(TableNameMember).
		Where(sq.Eq{"member_id": memberID}).
		OrderBy("created_date DESC")
}

func (MemberRepoSQLBuilder) Append(m *domain.Member) sq.Sqlizer {
	return sq.
		Insert(TableNameMember).
		Columns("member_id", "created_date", "self_intro").
		Values(m.MemberID, m.CreatedDate, m.SelfIntro)
}
