package models

type Answer struct {
	ID      int `db:"id" json:"id"`
	Approve int `db:"approve" json:"approve"`
}
