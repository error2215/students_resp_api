package student

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type Student struct {
	Id         int32  `json:"id"`
	FirstName  string `json:"name"`
	SecondName string `json:"second_name"`
	Group      int32  `json:"group"`
	Course     int32  `json:"course"`
	YearRate   int32  `json:"year_rate"`
}

func SliceToJson(students []*Student) (json.RawMessage, error) {
	encoded, err := json.Marshal(students)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return encoded, nil
}
