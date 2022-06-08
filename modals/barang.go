package modals

import (
	db "restapi-tugas/database"
	"time"
)

type Barang struct {
	Id          uint64    `json:"BarangId" sql:"AUTO_INCREMENT" gorm:"primaryKey"`
	BaranngName string    `json:"BarangName"`
	Stock       string    `json:"Stock"`
	Status      uint64    `json:"Status"  gorm:"foreignKey:IdUser"`
	IdUser      uint64    `json:"IdUser" gorm:"foreignKey:IdUser"`
	IdGudang    uint64    `json:"IdGudang" gorm:"foreignKey:IdGudang"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FindAllBarangs() []Barang {
	Barangs := []Barang{}
	db.DBconnect.Find(&Barangs)
	return Barangs
}

// FindByBarangId
func FindByBarangId(BarangId string) Barang {
	Barang := Barang{}
	db.DBconnect.Where("id = ?", BarangId).First(&Barang)
	return Barang
}

// CreateBarang
func CreateBarang(barangs Barang) Barang {
	db.DBconnect.Create(&barangs)
	return barangs
}

// DeleteBarang
func DeleteBarang(BarangId string) bool {
	Barang := Barang{}
	result := db.DBconnect.Where("id = ?", BarangId).Delete(&Barang)
	return result.RowsAffected > 0
}

// UpdateBarang
func UpdateBarang(BarangId string, barang Barang) Barang {
	db.DBconnect.Model(&barang).Where("BarangId = ?", barang).Updates(barang)
	return barang
}
