package models

import "time"

type Student struct {
	Id int
	Name string
	Surname string
	Birthdate time.Time
	PublicId string
	Photo string
}