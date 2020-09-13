package domain

import (
	"ddd/pkg/assistant/datastruct"
)

type MemberRepo interface {
	Find(memberID string) (*Member, error)
	Add(m *Member) error
}

type Member struct {
	MemberID    string          `db:"member_id"`
	CreatedDate datastruct.Time `db:"created_date"`
	SelfIntro   *string         `db:"self_intro"`
}
