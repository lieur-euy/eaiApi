package modals

import (
	db "restapi-tugas/database"
)

type Status struct {
	Id     uint64 `json:"StatusId" sql:"AUTO_INCREMENT" `
	Status string `json:"Status"`
}

// ======================================= status ===================================

// FindAllStatuss
func FindAllStatuss() []Status {
	Statuss := []Status{}
	db.DBconnect.Find(&Statuss)
	return Statuss
}

// FindByStatusId
func FindByStatusId(StatusId string) Status {
	Status := Status{}
	db.DBconnect.Where("id = ?", StatusId).First(&Status)
	return Status
}

// CreateStatus
func CreateStatus(Status Status) Status {
	db.DBconnect.Create(&Status)
	return Status
}

// DeleteStatus
func DeleteStatus(StatusId string) bool {
	Status := Status{}
	result := db.DBconnect.Where("id = ?", StatusId).Delete(&Status)
	return result.RowsAffected > 0
}

// UpdateStatus
func UpdateStatus(StatusId string, Status Status) Status {
	db.DBconnect.Model(&Status).Where("id = ?", StatusId).Updates(Status)
	return Status
}
