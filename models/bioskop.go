package models

import "gorm.io/gorm"

type Bioskop struct {
	gorm.Model
	Nama   string  `json:"nama" gorm:"not null"`
	Lokasi string  `json:"lokasi" gorm:"not null"`
	Rating float64 `json:"rating"`
}