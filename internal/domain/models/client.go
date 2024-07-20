package models

type Client struct {
	ID         int    `bd:"id" json:"id"`
	Surname    string `bd:"surname" json:"surname" binding:"required"`
	Name       string `bd:"name" json:"name"`
	Patronymic string `bd:"patronymic" json:"patronymic"`
	Email      string `bd:"email" json:"email" binding:"required"`
	Approve    int    `bd:"approve" json:"approve" default:"0"`
}
