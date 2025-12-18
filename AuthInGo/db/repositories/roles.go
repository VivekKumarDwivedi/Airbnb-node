package db

import (
	"AuthInGo/models"
	"database/sql"
)

type RoleRepository interface {
	GetRoleByID(id int64) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
	GetAllRoles() ([]*models.Role, error)
	CreateRole(name string, description string) (*models.Role, error)
	DeleteRoleByID(id int64) error
	UpdateRole(id int64, name string, description string) (*models.Role, error)
}

type RoleRepositoryImpl struct {
	db *sql.DB
}

func NewRoleRepository(_db *sql.DB) RoleRepository {
	return &RoleRepositoryImpl{
		db: _db,
	}
}

func (r *RoleRepositoryImpl) GetRoleByID(id int64) (*models.Role, error) {

	query := "SELECT id, name, description, created_at, updated_at FROM roles WHERE id = ?"

	row := r.db.QueryRow(query, id)

	role := &models.Role{}
	err := row.Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (r *RoleRepositoryImpl) GetRoleByName(name string) (*models.Role, error) {

	query := "SELECT id, name, description, created_at, updated_at FROM roles WHERE name = ?"

	row := r.db.QueryRow(query, name)

	role := &models.Role{}
	err := row.Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (r *RoleRepositoryImpl) GetAllRoles() ([]*models.Role, error) {

	query := "SELECT id, name, description, created_at, updated_at FROM roles"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []*models.Role
	for rows.Next() {
		role := &models.Role{}
		err := rows.Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}

func (r *RoleRepositoryImpl) CreateRole(name string, description string) (*models.Role, error) {

	query := "INSERT INTO roles (name, description) VALUES (?, ?)"
	result, err := r.db.Exec(query, name, description)
	if err != nil {
		return nil, err
	}
	roleID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return r.GetRoleByID(roleID)
}

func (r *RoleRepositoryImpl) DeleteRoleByID(id int64) error {

	query := "DELETE FROM roles WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
func (r *RoleRepositoryImpl) UpdateRole(id int64, name string, description string) (*models.Role, error) {

	query := "UPDATE roles SET name = ?, description = ? WHERE id = ?"
	_, err := r.db.Exec(query, name, description, id)
	if err != nil {
		return nil, err
	}
	return r.GetRoleByID(id)
}
