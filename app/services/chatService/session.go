package chatservice

import (
	"MATHB/app/models"
	"MATHB/config/database"
)

func CreateSession(userID uint) (*models.Session, error) {
	session := &models.Session{UserID: userID}
	result := database.DB.Create(session)
	if result.Error != nil {
		return nil, result.Error
	}
	return session, nil
}

func GetSession(userID uint, limit int) ([]models.Session, error) {
	var session []models.Session
	result := database.DB.Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).
		Find(&session)
	if result.Error != nil {
		return nil, result.Error
	}
	return session, nil
}

// UpdateSession 更新会话的最后活动时间
func UpdateSession(sessionID uint) error {
	result := database.DB.Model(&models.Session{}).
		Where("id = ?", sessionID).
		Update("last_activity_at", database.DB.NowFunc())
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteSession 删除一个指定会话
func DeleteSession(sessionID uint) error {
	result := database.DB.Where("id = ?", sessionID).Delete(&models.Session{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetSessionByID 根据会话 ID 查询会话详情
func GetSessionByID(sessionID uint) (*models.Session, error) {
	var session models.Session
	result := database.DB.Where("id = ?", sessionID).First(&session)
	if result.Error != nil {
		return nil, result.Error
	}
	return &session, nil
}
