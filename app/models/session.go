package models

import "time"

type Session struct {
	ID             uint      `gorm:"primaryKey"`                                    // 会话唯一标识
	UserID         uint      `gorm:"not null"`                                      // 外键，关联 User 表
	User           User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // 用户信息
	StartedAt      time.Time `gorm:"autoCreateTime"`                                // 会话开始时间
	LastActivityAt time.Time `gorm:"autoUpdateTime"`                                // 最后活动时间
}
