package services

import (
	env "AuthInGo/config/env"
	db "AuthInGo/db/repositories"
	"AuthInGo/dto"
	"AuthInGo/models"
	"AuthInGo/utils"
	"fmt"

	jwt "github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserById(id int64) (*models.User, error)
	CreateUser(payload *dto.CreateUserRequestDTO) (*models.User, error)
	LoginUser(payload *dto.LoginUserrequestDTO) (string, error)
	DeleteUserById() error
	GetAllUser() ([]*models.User, error)
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (u *UserServiceImpl) GetUserById(id int64) (*models.User, error) {
	fmt.Println("Fetching user in UserService")
	user, err := u.userRepository.GetByID(id)
	if err != nil {
		fmt.Println("Error fetching user:", err)
		return nil, err
	}
	return user, nil
}

func (u *UserServiceImpl) GetAllUser() ([]*models.User, error) {
	fmt.Println("Fetching all user in UserService")
	row, err := u.userRepository.GetAll()

	if err != nil {
		fmt.Println("No User found")
	}
	return row, nil

}

func (u *UserServiceImpl) DeleteUserById() error {
	fmt.Println("Fetching user in UserService")
	id := 2
	err := u.userRepository.DeleteByID(int64(id))

	if err != nil {
		fmt.Println("User Should not found")
		return err
	}

	return nil
}

func (u *UserServiceImpl) CreateUser(payload *dto.CreateUserRequestDTO) (*models.User, error) {
	fmt.Println("Creating user in UserService")
	// Step 1. Hash the password using utils.HashPassword
	hashedPassword, err := utils.HashPassword(payload.Password)

	if err != nil {
		fmt.Println("error hashing password:", err)
		return nil, err
	}
	// Step 2. Call the repository to create the user
	user, err := u.userRepository.Create(payload.Username, payload.Email, hashedPassword)

	if err != nil {
		fmt.Println("Error craeting user:", err)
		return nil, err
	}
	// Step 3. Return the created user
	return user, nil
}

func (u *UserServiceImpl) LoginUser(payload *dto.LoginUserrequestDTO) (string, error) {
	//Pre-requsite This function will be given email and password as parameter,which we can hardcode for now.
	email := payload.Email
	password := payload.Password
	//step 1: Make a repo call to get the user by email
	user, err := u.userRepository.GetByEmail(email)

	if err != nil {
		fmt.Println("Error Fetching user by email:", err)
		return "", err
	}
	//step 2: If user exists , or not. If not exists ,return error
	if user == nil {
		fmt.Println("No user found with given email")
		return "", fmt.Errorf("no user found with given email: %s", email)
	}

	//step 3: If user exists, check the password using utils.CheckPasswordHash()
	isPasswordVaild := utils.CheckPasswordHash(password, user.Password)

	if !isPasswordVaild {
		fmt.Println("Password does not match")
		return "", nil
	}
	//step 4: if password matches, print a JWT tokens , else return error saying password does not match
	jwtPayload := jwt.MapClaims{
		"email": user.Email,
		"id":    user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPayload)

	tokenString, err := token.SignedString([]byte(env.GetString("JWT_SECRET", "TOKEN")))

	if err != nil {
		fmt.Println("Error signing token:", err)
		return "", err
	}
	fmt.Println("JWT Token:", tokenString)

	return tokenString, nil
}
