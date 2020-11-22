package domain

import (
	"context"

	"ddd/pkg/technical/types"
)

type MemberRepoCQ interface {
	AppendMember(context.Context, *Member) (id int64, Err error)
	QueryByMemberID(ctx context.Context, memberID string, inWriteMode bool) (member Member, err error)
}

type Member struct {
	MemberID    string     `db:"member_id"`
	CreatedDate types.Time `db:"created_date"`
	SelfIntro   *string    `db:"self_intro"`
}
