package models

// import (
// 	"database/sql"
// 	"gorm.io/gorm"
// )

// type Company struct {
// 	ID    uint   `gorm:"primaryKey;autoIncrement;not null"`
// 	Name  string `gorm:"not null"`
// 	Users []User `gorm:"foreignKey:CompanyID"`
// }

// type User struct {
// 	gorm.Model
// 	Name         string         `gorm:"not null;column:nome"`
// 	Email        sql.NullString `gorm:"index;unique;not null;column:email"`
// 	Age          uint8          `gorm:"not null;check:age >= 18;column:idade"`
// 	Birthday     sql.NullTime   `gorm:"column:aniversario"`
// 	MemberNumber sql.NullString `gorm:"size:11;column:numero_membro"`
// 	ActivatedAt  sql.NullTime   `gorm:"column:ativado_em"`
// 	Score        float64        `gorm:"type:decimal;precision:10;scale:2;column:score"`
// 	CompanyID    uint
// 	Company      Company `gorm:"foreignKey:CompanyID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
// }
