package models

import "time"

type Student struct {
	Id int
	Name string
	Surname string
	Code string
	Grade string
	Birthdate time.Time
	PublicId string
	Photo string
}