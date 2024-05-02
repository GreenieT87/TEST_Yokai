package model

import (
    "gorm.io/gorm"
)

type ExampleModel struct {
    gorm.Model
	CustomerID int
    Name string
}