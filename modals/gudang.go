package modals

import (
	db "restapi-tugas/database"
)

type Gudang struct {
	Id         uint64 `json:"GudangId" sql:"AUTO_INCREMENT" `
	GudangName string `json:"GudangName"`
	Alamat     string `json:"Alamat"`
}

// FindAllGudangs
func FindAllGudangs() []Gudang {
	Gudangs := []Gudang{}
	db.DBconnect.Find(&Gudangs)
	return Gudangs
}

// FindByGudangId
func FindByGudangId(GudangId string) Gudang {
	Gudang := Gudang{}
	db.DBconnect.Where("id = ?", GudangId).First(&Gudang)
	return Gudang
}

// CreateGudang
func CreateGudang(Gudang Gudang) Gudang {
	db.DBconnect.Create(&Gudang)
	return Gudang
}

// DeleteGudang
func DeleteGudang(GudangId string) bool {
	Gudang := Gudang{}
	result := db.DBconnect.Where("GudangId = ?", GudangId).Delete(&Gudang)
	return result.RowsAffected > 0
}

// UpdateGudang
func UpdateGudang(GudangId string, Gudang Gudang) Gudang {
	db.DBconnect.Model(&Gudang).Where("id = ?", GudangId).Updates(Gudang)
	return Gudang
}
