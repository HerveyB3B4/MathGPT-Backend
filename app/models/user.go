package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`        // 用户唯一标识
	Username  string    `gorm:"unique;not null"`   // 用户名
	Email     string    `gorm:"type:varchar(100)"` // 邮箱
	Phone     string    `gorm:"unique;not null"`   // 手机号
	CreatedAt time.Time // 创建时间
}
