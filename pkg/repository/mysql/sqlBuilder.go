package mysql

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/rs/zerolog/log"
)

type GenericSQLBuilder struct{}

func (GenericSQLBuilder) IsTheRowExist(fieldName string, rowValue interface{}, tableName string) (string, []interface{}) {
	sqlString, args, err := sq.
		Select(fieldName).
		From(tableName).
		Where(sq.Eq{fieldName: rowValue}).
		Suffix("FOR UPDATE").
		ToSql()
	if err != nil {
		log.Error().Err(err).Msg("Build sql string 'IsTheRowExist' failed:")
	}

	return sqlString, args
}
