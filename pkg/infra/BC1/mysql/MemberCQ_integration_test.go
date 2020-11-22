// +build integration

package mysql_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"ddd/pkg/domain"
	"ddd/pkg/infra/BC1/mysql"
	"ddd/pkg/infra/part"
	"ddd/pkg/technical/logger"
	"ddd/pkg/technical/mock"
	"ddd/pkg/technical/types"
)

func TestMemberRepoCQ_Append(t *testing.T) {
	logger.DeveloperMode()
	db := part.NewMySQL(&mock.Config.MySQL)
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
			_, err := repo.AppendMember(context.Background(), tt.member)
			assert.NoError(t, err)
		})
	}
}

func TestMemberRepoCQ_Query(t *testing.T) {
	logger.DeveloperMode()
	cfg := mock.Config
	db := part.NewMySQL(&cfg.MySQL)
	repo := mysql.NewMemberRepoCQ(db)

	_, err := repo.QueryByMemberID(context.Background(), "c5", false)
	assert.NoError(t, err)
}
