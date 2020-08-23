// +build experiment

package experiment

import (
	"fmt"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"

	"ddd/pkg/domain"
	"ddd/pkg/helper/helperlog"
	"ddd/pkg/helper/helpertest/mock"
	"ddd/pkg/helper/helpertype"
	"ddd/pkg/repository/mysql"
)

func init() {
	helperlog.DevelopSetting()
}

// 實驗區 測試想法
func TestExperiment(t *testing.T) {
	var err error

	config := mock.Config()
	db := mysql.NewDB(&config.MySQL)

	sqlString, args, err := JoinSQL().ToSql()
	assert.NoError(t, err)
	fmt.Println(sqlString)

	data := new(Row)
	err = db.Get(data, sqlString, args...)
	assert.NoError(t, err)

	spew.Dump(data)

}

// type Row struct
// 	MemberID    string       `db:"member_id"`
// 	CreatedDate helpertype.StandardTime `db:"created_date"`
// 	SelfIntro   *string      `db:"self_intro"`
// 	ShopID      string       `db:"shop_id"`
// }

type Row struct {
	CreatedDate helpertype.Time `db:"created_date"`
	domain.Member
	ShopID string `db:"shop_id"`
}

func JoinSQL() sq.Sqlizer {
	builder := sq.
		Select("m.*", "s.shop_id").
		From("member AS m").
		Join("shop AS s ON m.created_date = s.created_date")
	return builder
}
