package controllers

import (
	"AuthInGo/dto"
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
	uc.UserService.GetUserById()
	w.Write([]byte("User Fetching endpoint done"))
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateUser called in UserController")
	uc.UserService.CreateUser()
	w.Write([]byte("User CreateUser endpoint done"))
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LoginUser called in UserController")

	var payload dto.LoginUserrequestDTO

	if jsonErr := utils.ReadJsonBody(r, &payload); jsonErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Something went wrong while logging in", jsonErr)
		return
	}

	if validationErr := utils.Validator.Struct(payload); validationErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Input invalid data", validationErr)
		fmt.Println("validationErr :", validationErr)
		return
	}

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
