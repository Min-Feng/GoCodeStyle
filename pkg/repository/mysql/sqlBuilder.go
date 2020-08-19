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

// TimeRange 呼叫端必須先確認 startTime 有值
// 沒有'小於等於', 資料庫查詢似乎可以少一次判斷?
// https://kknews.cc/zh-tw/code/9zbqpjl.html
func (GenericSQLBuilder) TimeRange(timeFieldName string, startTime interface{}, endTime interface{}) sq.Sqlizer {
	start := map[string]interface{}{timeFieldName: startTime}
	if endTime != nil {
		end := map[string]interface{}{timeFieldName: endTime}
		return sq.And{sq.GtOrEq(start), sq.LtOrEq(end)}
	}
	return sq.GtOrEq(start)
}
