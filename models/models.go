package models

import "time"

type Student struct {
	Id int
	Name string
	Surname string
	Code string
	Grade string
	Section string
	Birthdate time.Time
	PublicId string
	Photo string
}

type IdResponse struct {
	Id int
}

var ImageTypes = []string{".webp", ".svg", ".png", ".apng", ".avif", ".gif", ".jpg", ".jpeg", ".jfif", ".pjpeg", ".pjp"}
