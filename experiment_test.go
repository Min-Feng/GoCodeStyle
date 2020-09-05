// +build experiment

package experiment

import (
	"encoding/json"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/davecgh/go-spew/spew"

	"ddd/pkg/domain"
	"ddd/pkg/helper/helperlog"
	"ddd/pkg/helper/helpertype"
)

func init() {
	helperlog.DeveloperMode()
}

// 實驗區 測試想法
func TestExperiment(t *testing.T) {
	b := []byte(`{
  "age": 12.2,
  "money": null
}`)

	m := make(map[string]interface{})
	json.Unmarshal(b, &m)

	spew.Dump(m)
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
