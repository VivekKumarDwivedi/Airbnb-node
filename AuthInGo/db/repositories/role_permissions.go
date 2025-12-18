package db

import (
	"AuthInGo/models"
	"database/sql"
)

type RolePermissionRepository interface {
	GetRolePermissionById(id int64) (*models.RolePermission, error)
	GetRolePermissionByRoleId(roleId int64) ([]*models.RolePermission, error)
	AddPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error)
	RemovePermissionFromRole(roleId int64, permissionId int64) error
	GetAllRolePermission() ([]*models.RolePermission, error)
}
type RolePermissionRepoImpl struct {
	db *sql.DB
}

func NewRolePermissionRepository(_db *sql.DB) RolePermissionRepository {
	return &RolePermissionRepoImpl{
		db: _db,
	}
}

func (r *RolePermissionRepoImpl) GetRolePermissionById(id int64) (*models.RolePermission, error) {

	query := "SELECT id, role_id, permission_id, created_at, updated_at FROM role_permissions WHERE id = ?"

	row := r.db.QueryRow(query, id)

	rolePermission := &models.RolePermission{}
	err := row.Scan(
		&rolePermission.Id,
		&rolePermission.RoleID,
		&rolePermission.PermissionID,
		&rolePermission.CreatedAt,
		&rolePermission.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return rolePermission, nil
}

func (r *RolePermissionRepoImpl) GetRolePermissionByRoleId(roleId int64) ([]*models.RolePermission, error) {

	query := "SELECT id, role_id, permission_id, created_at, updated_at FROM role_permissions WHERE role_id = ?"

	rows, err := r.db.Query(query, roleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var rolePermissions []*models.RolePermission
	for rows.Next() {
		rolePermission := &models.RolePermission{}
		err := rows.Scan(
			&rolePermission.Id,
			&rolePermission.RoleID,
			&rolePermission.PermissionID,
			&rolePermission.CreatedAt,
			&rolePermission.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		rolePermissions = append(rolePermissions, rolePermission)
	}

	return rolePermissions, nil
}

func (r *RolePermissionRepoImpl) AddPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error) {

	query := "INSERT INTO role_permissions (role_id, permission_id, created_at, updated_at) VALUES (?, ?, datetime('now'), datetime('now'))"
	result, err := r.db.Exec(query, roleId, permissionId)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return r.GetRolePermissionById(id)
}

func (r *RolePermissionRepoImpl) RemovePermissionFromRole(roleId int64, permissionId int64) error {

	query := "DELETE FROM role_permissions WHERE role_id = ? AND permission_id = ?"
	_, err := r.db.Exec(query, roleId, permissionId)

	if err != nil {
		return err
	}
	return nil
}
func (r *RolePermissionRepoImpl) GetAllRolePermission() ([]*models.RolePermission, error) {

	query := "SELECT id, role_id, permission_id, created_at, updated_at FROM role_permissions"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err

	}
	defer rows.Close()
	var rolePermissions []*models.RolePermission

	for rows.Next() {
		rolePermission := &models.RolePermission{}
		err := rows.Scan(
			&rolePermission.Id,
			&rolePermission.RoleID,
			&rolePermission.PermissionID,
			&rolePermission.CreatedAt,
			&rolePermission.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		rolePermissions = append(rolePermissions, rolePermission)
	}
	return rolePermissions, nil
}
