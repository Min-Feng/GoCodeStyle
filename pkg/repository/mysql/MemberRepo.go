package mysql

import (
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/davecgh/go-spew/spew"
	"github.com/jmoiron/sqlx"
	"github.com/morikuni/failure"
	"github.com/rs/zerolog/log"

	"ddd/pkg/domain"
)

const MemberTableName = "member"

func NewMemberRepo(db *sqlx.DB) *MemberRepo {
	return &MemberRepo{db: db}
}

type MemberRepo struct {
	db         *sqlx.DB
	sqlBuilder MemberRepoSQLBuilder
}

func (repo *MemberRepo) Find(memberID string) (*domain.Member, error) {
	member := new(domain.Member)

	sqlString, args := repo.sqlBuilder.Find(memberID)
	err := repo.db.Get(member, sqlString, args...)
	if err != nil {
		if strings.Contains(err.Error(), ErrMsgNotFound) {
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
	sqlString, args := repo.sqlBuilder.Add(m)
	result, err := repo.db.Exec(sqlString, args...)
	if err != nil {
		return failure.Wrap(err, failure.Message("mysql insert"))
	}

	if log.Debug().Enabled() {
		id, _ := result.LastInsertId()
		log.Debug().Int64("member_id", id).Msg("MemberRepo mysql insert row:")
	}
	return nil
}

type MemberRepoSQLBuilder struct{}

func (MemberRepoSQLBuilder) Find(memberID string) (string, []interface{}) {
	sqlString, args, err := sq.
		Select("*").
		From(MemberTableName).
		Where(sq.Eq{"member_id": memberID}).
		OrderBy("created_date DESC").
		ToSql()
	if err != nil {
		log.Error().Err(err).Msg("Build sql string 'MemberRepo.Find' failed:")
	}
	return sqlString, args
}

func (MemberRepoSQLBuilder) Add(m *domain.Member) (string, []interface{}) {
	sqlString, args, err := sq.
		Insert(MemberTableName).
		Columns("member_id", "created_date", "self_intro").
		// Values(m.MemberID, m.CreatedDate, m.SelfIntro).
		Values(m.MemberID, m.CreatedDate, m.SelfIntro).
		ToSql()
	if err != nil {
		log.Error().Err(err).Msg("Build sql string 'MemberRepo.Add' failed:")
	}
	return sqlString, args
}
