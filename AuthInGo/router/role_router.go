package router

import (
	"AuthInGo/controllers"
	"AuthInGo/middlewares"

	"github.com/go-chi/chi/v5"
)

type RoleRouter struct {
	RoleController controllers.RoleController
}

func NewRoleRouter(_roleController *controllers.RoleController) Router {
	return &RoleRouter{
		RoleController: *_roleController,
	}
}

func (rr *RoleRouter) Register(r chi.Router) {
	r.Get("/roles/{id}", rr.RoleController.GetRoleById)
	r.Get("/roles", rr.RoleController.GetAllRoles)
	r.With(middlewares.CreateRoleRequestValidator).Post("/roles", rr.RoleController.CreateRole)
	r.With(middlewares.UpdateRoleRequestValidator).Put("/roles/{id}", rr.RoleController.UpdateRole)
	r.Delete("/roles/{id}", rr.RoleController.DeleteRoleById)
	r.With(middlewares.AddPermissionToRoleRequestValidator).Get("/roles/{id}/permissions", rr.RoleController.GetRolePermissions)
	r.With(middlewares.AddPermissionToRoleRequestValidator).Post("/roles/{id}/permissions", rr.RoleController.AddPermissionToRole)
	r.With(middlewares.RemovePermissionFromRoleRequestValidator).Delete("/roles/{id}/permissions", rr.RoleController.RemovePermissionFromRole)
	r.With(middlewares.JWTAuthMiddleware, middlewares.RequireAllRoles("admin")).Post("/roles/{userId}/assign/{roleId}", rr.RoleController.AssignRoleToUser)
}
