package models

type Answer struct {
	ID      int `db:"id" json:"id" binding:"required"`
	Approve int `db:"approve" json:"approve" binding:"required"`
}
