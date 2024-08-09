package model

import "time"

type Sex int

const (
	Other Sex = iota
	Male
	Female
)

type (
	Address struct {
		Country   string
		City      string
		Street    string
		House     string
		Apartment string
	}

	Profile struct {
		UserId    int
		Sex       Sex
		BirthDate time.Time
		ShoesSize int
		Address   Address
	}
)
