// +build experiment

package mysql_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"ddd/pkg/adapter"
	"ddd/pkg/domain"
	"ddd/pkg/loghelper"
	"ddd/pkg/mock"
	"ddd/pkg/repository/mysql"
	"ddd/pkg/testtool"
)

func TestMemberRepo_Add(t *testing.T) {
	loghelper.DevelopSetting()
	cfg := testtool.Config()
	db := mysql.NewDB(&cfg.MySQL)
	repo := mysql.NewMemberRepo(db)

	tests := []struct {
		name   string
		member *domain.Member
	}{
		{
			member: &domain.Member{
				MemberID:    "c5",
				CreatedDate: adapter.Time{Time: mock.NewTimeNowFunc("1988-05-14T07:33:25")()},
			},
		},
		// {
		// 	member: &domain.Member{
		// 		MemberID:    "f5",
		// 		CreatedDate: mock.NewTimeNowFunc("2099-10-17T00:31:20")(),
		// 	},
		// },
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := repo.Add(tt.member)
			assert.NoError(t, err)
		})
	}
}

func TestMemberRepo_Find(t *testing.T) {
	loghelper.DevelopSetting()
	cfg := testtool.Config()
	db := mysql.NewDB(&cfg.MySQL)
	repo := mysql.NewMemberRepo(db)

	_, err := repo.Find("c5")
	assert.NoError(t, err)
}
