package controllers

import (
	"AuthInGo/dto"
	"AuthInGo/middlewares"
	"AuthInGo/services"
	"AuthInGo/utils"
	"fmt"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(_userService services.UserService) *UserController {
	return &UserController{
		UserService: _userService,
	}
}

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetUserById called in UserController")
	// Read id from query ?id=123
	userId := r.URL.Query().Get("id")

	// If not in query, try context
	if userId == "" {
		userId = r.Context().Value(middlewares.UserIDKey).(string)

	}

	fmt.Println("User ID from context or query:", userId)

	// If still empty → error
	if userId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "User ID is required", fmt.Errorf("missing user ID"))
		return
	}

	// Convert string → int

	user, err := uc.UserService.GetUserById(userId)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to fetch user", err)
		return
	}
	if user == nil {
		utils.WriteJsonErrorResponse(w, http.StatusNotFound, "User not found", fmt.Errorf("user with ID %s not found", userId))
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User fetched successfully", user)
	fmt.Println("User fetched successfully:", user)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("CreateUser called in UserController\n")
	payload, ok := r.Context().Value(middlewares.PayloadKey).(dto.CreateUserRequestDTO)
	if !ok {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	fmt.Println("Payload received:", payload)

	user, err := uc.UserService.CreateUser(&payload)

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to create user", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusCreated, "User created successfully", user)
	fmt.Println("User created successfully:", user)
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LoginUser called in UserController")

	payload := r.Context().Value(middlewares.PayloadKey).(dto.LoginUserrequestDTO)

	fmt.Println("Pyload recived:", payload)

	jwtToken, err := uc.UserService.LoginUser(&payload)

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to login user", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User Logged in successfully", jwtToken)
}

func (uc *UserController) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteUserById called in UserController")
	uc.UserService.DeleteUserById()
	w.Write([]byte("User delete endpoint done"))
}

func (uc *UserController) GetAllUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllUser called in UserController")
	uc.UserService.GetAllUser()
	w.Write([]byte("All user getting endpoint done"))
}
