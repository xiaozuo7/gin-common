package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SoftDeleteModel struct {
	DeleteStatus bool           `gorm:"not null;default:0" json:"delete_status"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

func Delete(session *gorm.DB, id uint, m interface{}) (tx *gorm.DB) {
	return session.Model(m).Where("id=?", id).Updates(map[string]interface{}{"delete_status": true, "deleted_at": time.Now()})
}

func Recover(session *gorm.DB, id uint, m interface{}) (tx *gorm.DB) {
	return session.Model(m).Where("id=?", id).Updates(map[string]interface{}{"delete_status": false, "deleted_at": nil})
}

type ActivateModel struct {
	ActivateStatus bool      `gorm:"not null;default:1" json:"activate_status"`
	ActivateAt     time.Time `json:"activate_at"`
}

func Activate(session *gorm.DB, id uint, m interface{}) (tx *gorm.DB) {
	return session.Model(m).Where("id=?", id).Updates(map[string]interface{}{"activate_status": true, "activate_at": time.Now()})
}

func Deactivate(session *gorm.DB, id uint, m interface{}) (tx *gorm.DB) {
	return session.Model(m).Where("id=?", id).Updates(map[string]interface{}{"activate_status": false, "activate_at": time.Now()})
}
