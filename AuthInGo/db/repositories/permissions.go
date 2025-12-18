package db

import (
	"AuthInGo/models"
	"database/sql"
)

type PermissionRepository interface {
	GetPermissionByID(id int64) (*models.Permission, error)
	GetPermissionByName(name string) (*models.Permission, error)
	GetAllPermissions() ([]*models.Permission, error)
	CreatePermission(name string, description string, resource string, action string) (*models.Permission, error)
	DeletePermissionByID(id int64) error
	UpdatePermission(id int64, name string, description string, resource string, action string) (*models.Permission, error)
}

type PermissionRepositoryImpl struct {
	db *sql.DB
}

func NewPermissionRepository(_db *sql.DB) PermissionRepository {
	return &PermissionRepositoryImpl{
		db: _db,
	}
}

func (p *PermissionRepositoryImpl) GetPermissionByID(id int64) (*models.Permission, error) {

	query := "SELECT id, name, description, resource, action, created_at, updated_at FROM permissions WHERE id = ?"
	row := p.db.QueryRow(query, id)
	permission := &models.Permission{}
	err := row.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return permission, nil
}
func (p *PermissionRepositoryImpl) GetPermissionByName(name string) (*models.Permission, error) {

	query := "SELECT id, name, description, resource, action, created_at, updated_at FROM permissions WHERE name = ?"

	row := p.db.QueryRow(query, name)

	permission := &models.Permission{}
	err := row.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return permission, nil
}

func (p *PermissionRepositoryImpl) GetAllPermissions() ([]*models.Permission, error) {

	query := "SELECT id, name, description, resource, action, created_at, updated_at FROM permissions"
	rows, err := p.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []*models.Permission

	for rows.Next() {
		permission := &models.Permission{}
		err := rows.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt)
		if err != nil {
			return nil, err
		}

		permissions = append(permissions, permission)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}

func (p *PermissionRepositoryImpl) CreatePermission(name string, description string, resource string, action string) (*models.Permission, error) {
	query := "INSERT INTO permissions (name, description, resource, action) VALUES (?, ?, ?, ?)"

	result, err := p.db.Exec(query, name, description, resource, action)
	if err != nil {
		return nil, err
	}
	insertedID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return p.GetPermissionByID(insertedID)
}
func (p *PermissionRepositoryImpl) DeletePermissionByID(id int64) error {
	query := "DELETE FROM permissions WHERE id = ?"
	_, err := p.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *PermissionRepositoryImpl) UpdatePermission(id int64, name string, description string, resource string, action string) (*models.Permission, error) {
	query := "UPDATE permissions SET name = ?, description = ?, resource = ?, action = ? WHERE id = ?"
	_, err := p.db.Exec(query, name, description, resource, action, id)
	if err != nil {
		return nil, err
	}

	return p.GetPermissionByID(id)
}
