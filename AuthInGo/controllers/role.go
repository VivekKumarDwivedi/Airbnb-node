package controllers

import (
	"AuthInGo/services"
	"AuthInGo/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type RoleController struct {
	RoleService services.RoleService
}

func NewRoleController(roleService services.RoleService) *RoleController {
	return &RoleController{
		RoleService: roleService,
	}
}

func (rc *RoleController) GetRoleById(w http.ResponseWriter, r *http.Request) {
	roleId := chi.URLParam(r, "id")

	if roleId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("missing role ID"))
		return
	}
	id, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid Role ID", fmt.Errorf("role ID must be a valid integer"))
		return
	}
	role, err := rc.RoleService.GetRoleById(id)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to get role", err)
		return
	}
	if role == nil {
		utils.WriteJsonErrorResponse(w, http.StatusNotFound, "Role not found", fmt.Errorf("role with ID %s not found", roleId))
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role fetched successfully", role)
}

func (rc *RoleController) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := rc.RoleService.GetAllRoles()
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to get roles", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Roles fetched successfully", roles)
}

func (rc *RoleController) CreateRole(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	description := r.FormValue("description")
	if name == "" || description == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Name and Description are required", fmt.Errorf("missing name or description"))
		return
	}
	role, err := rc.RoleService.CreateRole(name, description)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to create role", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusCreated, "Role created successfully", role)
}

func (rc *RoleController) DeleteRoleById(w http.ResponseWriter, r *http.Request) {
	roleId := chi.URLParam(r, "id")
	if roleId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("missing role ID"))
		return
	}

	id, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid Role ID", fmt.Errorf("role ID must be a valid integer"))
		return
	}
	err = rc.RoleService.DeleteRoleById(id)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to delete role", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role deleted successfully", nil)
}

func (rc *RoleController) UpdateRole(w http.ResponseWriter, r *http.Request) {
	roleId := chi.URLParam(r, "id")
	if roleId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("missing role ID"))
		return
	}
	id, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid Role ID", fmt.Errorf("role ID must be a valid integer"))
		return
	}
	name := r.FormValue("name")
	description := r.FormValue("description")
	if name == "" || description == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Name and Description are required", fmt.Errorf("missing name or description"))
		return
	}
	role, err := rc.RoleService.UpdateRole(id, name, description)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to update role", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role updated successfully", role)
}

func (rc *RoleController) GetRolePermissions(w http.ResponseWriter, r *http.Request) {
	roleId := chi.URLParam(r, "id")
	if roleId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("missing role ID"))
		return
	}
	id, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid Role ID", fmt.Errorf("role ID must be a valid integer"))
		return
	}
	permissions, err := rc.RoleService.GetRolePermission(id)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to get role permissions", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role permissions fetched successfully", permissions)
}

func (rc *RoleController) AddPermissionToRole(w http.ResponseWriter, r *http.Request) {
	roleId := chi.URLParam(r, "id")
	if roleId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("missing role ID"))
		return
	}
	id, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid Role ID", fmt.Errorf("role ID must be a valid integer"))
		return
	}
	permissionIdStr := r.FormValue("permission_id")
	if permissionIdStr == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Permission ID is required", fmt.Errorf("missing permission ID"))
		return
	}
	permissionId, err := strconv.ParseInt(permissionIdStr, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid Permission ID", fmt.Errorf("permission ID must be a valid integer"))
		return
	}
	rolePermission, err := rc.RoleService.AddPermissionToRole(id, permissionId)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to add permission to role", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Permission added to role successfully", rolePermission)
}
