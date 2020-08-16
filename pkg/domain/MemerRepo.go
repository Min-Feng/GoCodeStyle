package domain

import (
	"ddd/pkg/adapter"
)

type MemberRepo interface {
	Find(memberID string) (*Member, error)
	Add(m *Member) error
}

type Member struct {
	MemberID    string       `db:"member_id"`
	CreatedDate adapter.Time `db:"created_date"`
	SelfIntro   *string      `db:"self_intro"`
}
