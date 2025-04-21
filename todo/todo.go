package todo

import "time"

type Todo struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    Priority  `json:"priority"`
	Progress    Progress  `json:"progress"`
	DueDate     time.Time `json:"due_date"`
	IsDeleted   bool      `json:"isDeleted"` //Added for soft deletion
}

type Priority int

const (
	LOW Priority = iota + 1
	MEDIUM
	HIGH
)

func (p Priority) String() string {
	return [...]string{"LOW", "MEDIUM", "HIGH"}[p-1]
}

func (p Priority) EnumIndex() int {
	return int(p)
}

type Progress int

const (
	TO_DO Progress = iota + 1
	IN_PROGRESS
	ON_HOLD
	COMPLETED
)

func (p Progress) String() string {
	return [...]string{"TO_DO", "IN_PROGRESS", "ON_HOLD", "COMPLETED"}[p-1]
}

func (p Progress) EnumIndex() int {
	return int(p)
}
