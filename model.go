package main

import "gorm.io/gorm"

type Massage struct {
	gorm.Model

	Name string `json:"name"`
	Age  int16
	// Time int16
}
type Massage1 struct {
	Name string `json:"name"`
	Age  int16
	// Time int16
}
