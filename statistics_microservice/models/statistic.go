package models

type Statistic struct {
	BookName   string `json:"bookName"`
	Author     string `json:"author"`
	Percentage int    `json:"percentage"`
}
