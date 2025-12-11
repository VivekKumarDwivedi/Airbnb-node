package router

import (
	"AuthInGo/controllers"
	"AuthInGo/middlewares"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	UserController controllers.UserController
}

func NewUserRouter(_userController *controllers.UserController) Router {
	return &UserRouter{
		UserController: *_userController,
	}
}

func (ur *UserRouter) Register(r chi.Router) {

	r.Get("/profile", ur.UserController.GetUserById)
	r.With(middlewares.CreateUserRequestValidator).Post("/signup", ur.UserController.CreateUser)
	r.With(middlewares.UserLoginRequestValidator).Post("/login", ur.UserController.LoginUser)
	r.Delete("/delete", ur.UserController.DeleteUserById)
	r.Get("/users", ur.UserController.GetAllUser)
}
