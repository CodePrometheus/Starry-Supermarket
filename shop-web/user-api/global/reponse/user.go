package reponse

import (
	"fmt"
	"time"
)

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	time := fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02"))
	return []byte(time), nil
}

type UserResponse struct {
	Id       int32    `json:"id"`
	Nickname string   `json:"name"`
	Birthday JsonTime `json:"birthday"`
	Gender   string   `json:"gender"`
	Email    string   `json:"email"`
}
