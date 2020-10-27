package mysql

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/davecgh/go-spew/spew"
	"github.com/jmoiron/sqlx"
	"github.com/morikuni/failure"
	"github.com/rs/zerolog/log"

	"ddd/pkg/domain"
)

const TableNameMember = "member"

func NewMemberRepo(db *sqlx.DB) *MemberRepo {
	return &MemberRepo{db: db}
}

type MemberRepo struct {
	db         *sqlx.DB
	sqlBuilder MemberRepoSQLBuilder
}

func (repo *MemberRepo) Find(memberID string) (*domain.Member, error) {
	member := new(domain.Member)

	sqlString, args, _ := repo.sqlBuilder.Find(memberID).ToSql()
	err := repo.db.Get(member, sqlString, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, failure.Translate(err, domain.ErrNotFound, failure.Message("mysql select"))
		}
		return nil, failure.Wrap(err, failure.Message("mysql select"))
	}

	if log.Debug().Enabled() {
		log.Debug().Msgf("MemberRepo find memberID=%v\n%v", memberID, spew.Sdump(member))
	}
	return member, err
}

func (repo *MemberRepo) Add(m *domain.Member) error {
	sqlString, args, _ := repo.sqlBuilder.Add(m).ToSql()
	result, err := repo.db.Exec(sqlString, args...)
	if err != nil {
		return failure.Wrap(err, failure.Message("mysql insert"))
	}

	if log.Debug().Enabled() {
		id, _ := result.LastInsertId()
		// 由於 memberID 不是用 AUTO INCREMENT, 所以返回零
		log.Debug().Int64("member_id", id).Msg("MemberRepo mysql insert row:")
	}
	return nil
}

type MemberRepoSQLBuilder struct{}

func (MemberRepoSQLBuilder) Find(memberID string) sq.Sqlizer {
	return sq.
		Select("*").
		From(TableNameMember).
		Where(sq.Eq{"member_id": memberID}).
		OrderBy("created_date DESC")
}

func (MemberRepoSQLBuilder) Add(m *domain.Member) sq.Sqlizer {
	return sq.
		Insert(TableNameMember).
		Columns("member_id", "created_date", "self_intro").
		Values(m.MemberID, m.CreatedDate, m.SelfIntro)
}
