package userservice

import (
	"MATHB/app/models"
	"MATHB/config/database"
)

// CreateUser 创建用户
func CreateUser(username, email, phone string) (*models.User, error) {
	user := &models.User{
		Username: username,
		Email:    email,
		Phone:    phone,
	}
	result := database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// GetUserByPhone 根据手机号查询用户
func GetUserByPhone(phone string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("phone = ?", phone).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByEmail 根据邮箱查询用户
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByID 根据用户 ID 查询用户
func GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	result := database.DB.First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// UpdateUser 更新用户信息
func UpdateUser(userID uint, updates map[string]interface{}) error {
	result := database.DB.Model(&models.User{}).Where("id = ?", userID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteUser 删除用户
func DeleteUser(userID uint) error {
	result := database.DB.Delete(&models.User{}, userID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
