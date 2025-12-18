package services

import (
	repositories "AuthInGo/db/repositories"
	"AuthInGo/models"
)

type RoleService interface {
	GetRoleById(id int64) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
	GetAllRoles() ([]*models.Role, error)
	CreateRole(name string, description string) (*models.Role, error)
	DeleteRoleById(id int64) error
	UpdateRole(id int64, name string, description string) (*models.Role, error)
	GetRolePermission(roleId int64) ([]*models.RolePermission, error)
	AddPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error)
}

type RoleServiceImpl struct {
	roleRepository           repositories.RoleRepository
	rolePermissionRepository repositories.RolePermissionRepository
}

func NewRoleService(roleRepository repositories.RoleRepository) *RoleServiceImpl {
	return &RoleServiceImpl{
		roleRepository: roleRepository,
	}
}

func (r *RoleServiceImpl) GetRoleById(id int64) (*models.Role, error) {
	return r.roleRepository.GetRoleByID(id)
}

func (r *RoleServiceImpl) GetRoleByName(name string) (*models.Role, error) {
	return r.roleRepository.GetRoleByName(name)
}

func (r *RoleServiceImpl) GetAllRoles() ([]*models.Role, error) {
	return r.roleRepository.GetAllRoles()
}

func (r *RoleServiceImpl) CreateRole(name string, description string) (*models.Role, error) {
	return r.roleRepository.CreateRole(name, description)
}

func (r *RoleServiceImpl) DeleteRoleById(id int64) error {
	return r.roleRepository.DeleteRoleByID(id)
}

func (r *RoleServiceImpl) UpdateRole(id int64, name string, description string) (*models.Role, error) {
	return r.roleRepository.UpdateRole(id, name, description)
}

func (r *RoleServiceImpl) GetRolePermission(roleId int64) ([]*models.RolePermission, error) {
	return r.rolePermissionRepository.GetRolePermissionByRoleId(roleId)
}

func (r *RoleServiceImpl) AddPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error) {
	return r.rolePermissionRepository.AddPermissionToRole(roleId, permissionId)
}
