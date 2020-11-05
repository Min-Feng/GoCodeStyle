// +build integration

package mysql_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"ddd/pkg/domain"
	"ddd/pkg/infra/BC1/mysql"
	helperlog "ddd/pkg/infra/part"
	"ddd/pkg/technical/logger"
	"ddd/pkg/technical/mock"
	"ddd/pkg/technical/types"
)

func TestMemberRepo_Add(t *testing.T) {
	logger.DeveloperMode()
	db := helperlog.NewMySQL(&mock.Config.MySQL)
	repo := mysql.NewMemberRepoCQ(db)

	tests := []struct {
		name   string
		member *domain.Member
	}{
		{
			member: &domain.Member{
				MemberID:    "a1",
				CreatedDate: types.Time{Time: mock.TimeNowFunc("1988-05-14")()},
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
			_, err := repo.AppendMember(nil, tt.member)
			assert.NoError(t, err)
		})
	}
}

func TestMemberRepo_Find(t *testing.T) {
	logger.DeveloperMode()
	cfg := mock.Config
	db := helperlog.NewMySQL(&cfg.MySQL)
	repo := mysql.NewMemberRepoCQ(db)

	_, err := repo.FindByMemberID(nil, "c5", false)
	assert.NoError(t, err)
}
