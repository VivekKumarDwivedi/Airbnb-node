package services

import (
	env "AuthInGo/config/env"
	db "AuthInGo/db/repositories"
	"AuthInGo/utils"
	"fmt"

	jwt "github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserById() error
	CreateUser() error
	LoginUser() (string, error)
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (u *UserServiceImpl) GetUserById() error {
	fmt.Println("Fetching user in UserService")
	// u.userRepository.Create()
	u.userRepository.GetByID()
	return nil
}

func (u *UserServiceImpl) CreateUser() error {
	fmt.Println("Creating user in UserService")
	password := "example_password"
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	u.userRepository.Create(
		"username_example_2",
		"user2@example.com",
		hashedPassword,
	)
	return nil
}

func (u *UserServiceImpl) LoginUser() (string, error) {
	//Pre-requsite This function will be given email and password as parameter,which we can hardcode for now.
	email := "user2@example.com"
	password := "example_password"
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
	payload := jwt.MapClaims{
		"email": user.Email,
		"id":    user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString([]byte(env.GetString("JWT_SECRET", "TOKEN")))

	if err != nil {
		fmt.Println("Error signing token:", err)
		return "", err
	}
	fmt.Println("JWT Token:", tokenString)

	return tokenString, nil
}
