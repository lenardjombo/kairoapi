package models

import (
	"time"
)

type Project struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Createdat time.Time `json:"createdat"`
	Updatedat time.Time `json:"updatedat"`
}