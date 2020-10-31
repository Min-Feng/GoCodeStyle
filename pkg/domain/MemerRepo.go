package domain

import (
	"context"

	"ddd/pkg/technical/types"
)

type MemberRepo interface {
	FindByID(ctx context.Context, memberID string) (*Member, error)
	Append(context.Context, *Member) (id int64, Err error)
}

type Member struct {
	MemberID    string     `db:"member_id"`
	CreatedDate types.Time `db:"created_date"`
	SelfIntro   *string    `db:"self_intro"`
}
