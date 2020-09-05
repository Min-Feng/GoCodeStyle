// +build integration

package mysql_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"ddd/pkg/domain"
	"ddd/pkg/helper/helperlog"
	"ddd/pkg/helper/helpertest/mock"
	"ddd/pkg/helper/helpertype"
	"ddd/pkg/repository/mysql"
)

func TestMemberRepo_Add(t *testing.T) {
	helperlog.DeveloperMode()
	db := mysql.NewDB(&mock.Config.MySQL)
	repo := mysql.NewMemberRepo(db)

	tests := []struct {
		name   string
		member *domain.Member
	}{
		{
			member: &domain.Member{
				MemberID:    "a1",
				CreatedDate: helpertype.Time{Time: mock.TimeNowFunc("1988-05-14")()},
			},
		},
		// {
		// 	member: &domain.Member{
		// 		MemberID:    "f5",
		// 		CreatedDate: mock.TimeNowFunc("2099-10-17T00:31:20")(),
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
	helperlog.DeveloperMode()
	cfg := mock.Config
	db := mysql.NewDB(&cfg.MySQL)
	repo := mysql.NewMemberRepo(db)

	_, err := repo.Find("c5")
	assert.NoError(t, err)
}
