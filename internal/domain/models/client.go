package models

type Client struct {
	ID         int    `bd:"id" json:"id"`
	Surname    string `bd:"surname" json:"surname" binding:"required"`
	Name       string `bd:"name" json:"name"`
	Patronymic string `bd:"patronymic" json:"patronymic"`
	Email      string `bd:"email" json:"email" binding:"required"`
	Approval    int    `bd:"approval" json:"approval" default:"0"`
}

type Statistic struct {
	TotalClients int `json:"total_clients"`
	Approved     int `json:"approved"`
	Unapproved   int `json:"unapproved"`
	Waiting      int `json:"waiting"`
}
