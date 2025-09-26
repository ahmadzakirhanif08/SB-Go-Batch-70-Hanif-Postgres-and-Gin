package models

import "gorm.io/gorm"

type Bioskop struct {
    gorm.Model
    Nama       string  `json:"nama" gorm:"not null"`
    Lokasi     string  `json:"lokasi" gorm:"not null"`
    Rating     float64 `json:"rating"`
    FilmTayang []Film `json:"film_tayang"`
}

type Film struct {
    gorm.Model
    NamaFilm   string `json:"nama_film"`
    DurasiFilm int    `json:"durasi_film"`
    BioskopID uint `json:"bioskop_id"`
    Jadwal     []JadwalFilm `json:"jadwal"`
}

type JadwalFilm struct {
    gorm.Model
    WaktuMulai string `json:"waktu_mulai"`
    FilmID     uint   `json:"film_id"`
}