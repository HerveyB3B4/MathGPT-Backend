package models

import "time"

type Message struct {
	ID        uint      `gorm:"primaryKey"`                                    // 消息唯一标识
	SessionID uint      `gorm:"not null"`                                      // 外键，关联 Session 表
	Session   Session   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // 会话信息
	Sender    string    `gorm:"type:enum('user', 'assistant');not null"`       // 发送者
	Text      string    `gorm:"type:text;not null"`                            // 消息内容
	CreatedAt time.Time `gorm:"autoCreateTime"`                                // 消息发送时间
}
