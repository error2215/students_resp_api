package models

type Student struct {
	Id         int32  `json:"id"`
	FirstName  string `json:"name"`
	SecondName string `json:"second_name"`
	Group      int32  `json:"group"`
	Course     int32  `json:"course"`
	YearRate   int32  `json:"year_rate"`
}
