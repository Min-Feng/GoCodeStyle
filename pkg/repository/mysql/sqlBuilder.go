package mysql

import (
	sq "github.com/Masterminds/squirrel"
)

type GenericSQLBuilder struct{}

func (GenericSQLBuilder) IsTheRowExist(fieldName string, rowValue interface{}, tableName string) sq.Sqlizer {
	return sq.
		Select(fieldName).
		From(tableName).
		Where(sq.Eq{fieldName: rowValue}).
		Suffix("FOR UPDATE")
}

// startTIme or endTime 某一方沒值, 資料庫查詢可以少一次判斷
// https://kknews.cc/zh-tw/code/9zbqpjl.html
//
// startTime and endTime 可用的型別為 int family, string, Time
//
// squirrel 在進行 ToSql() 的時候
// 會呼叫可用型別所擁有的 interface: driver.Value
// 最後轉化為 ToSql() 的 args
func (GenericSQLBuilder) TimeRange(timeFieldName string, startTime interface{}, endTime interface{}) sq.Sqlizer {
	var timeBuilder sq.Sqlizer
	switch {
	case startTime != nil && endTime != nil:
		timeBuilder = sq.
			And{
			sq.GtOrEq{timeFieldName: startTime},
			sq.LtOrEq{timeFieldName: endTime},
		}
	case startTime != nil && endTime == nil:
		timeBuilder = sq.GtOrEq{timeFieldName: startTime}
	case startTime == nil && endTime != nil:
		timeBuilder = sq.LtOrEq{timeFieldName: endTime}
	}
	return timeBuilder
}
