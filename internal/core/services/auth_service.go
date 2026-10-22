package services

import (
	"errors"
	"fiber-ecommerce-api/internal/core/domain/entities"
	"fiber-ecommerce-api/internal/core/domain/ports/repositories"
	"fiber-ecommerce-api/pkg/utils"
)

type AuthServiceImpl struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) *AuthServiceImpl  {
	return &AuthServiceImpl {
		userRepo: userRepo,
	}
}

func (s *AuthServiceImpl) Register(req entities.RegisterRequest) (*entities.User, error) {
	// check exist email
	exist, _ := s.userRepo.GetByEmail(req.Email)
	if exist != nil {
		return nil, errors.New("Email ซ้ำ")
	}

	// ตรวจความซับซ้อนของ password
	if err := utils.ValidatePassword(req.Password); err != nil {
		return nil, err
	}

	// เข้ารหัส password
	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// สร้าง user
	user := &entities.User{
		Email: req.Email,
		Password: hashPassword,
		FirstName: req.FirstName,
		LastName: req.LastName,
		Role: entities.RoleUser,
		IsActive: true,
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthServiceImpl) AdminRegister(req entities.AdminRegisterRequest) (*entities.User, error) {
	// เช็คถ้ามี user ที่มีอีเมลเดียวกันอยู่แล้ว
	existingUser, _ := s.userRepo.GetByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	// ตรวจความซับซ้อนของ password
	if err := utils.ValidatePassword(req.Password); err != nil {
		return nil, err
	}

	// Hash รหัสผ่าน
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// สร้าง user ใหม่ด้วย role ที่กำหนด
	user := &entities.User{
		Email:     req.Email,
		Password:  hashedPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      req.Role, // ใช้ role ที่ส่งมา
		IsActive:  true,
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthServiceImpl) Login(req entities.LoginRequest) (*entities.LoginResponse, error) {
	// check exist email
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("Invalid email or password")
	}

	/// check active
	if !user.IsActive {
		return nil, errors.New("User is not active")
	}

	// check password
	if err := utils.CheckPassword(req.Password, user.Password); err != nil {
		return nil, errors.New("Invalid email or password")
	}

	// generate JWT
	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return nil, errors.New("Failed to generate JWT")
	}

	// สร้าง Login Response
	response := &entities.LoginResponse{
		Token: token,
		User: *user,
	}

	return response, nil
}

func (s *AuthServiceImpl) GetUserByID(id uint) (*entities.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *AuthServiceImpl) UpdateUser(user *entities.User) error {
	return s.userRepo.Update(user)
}