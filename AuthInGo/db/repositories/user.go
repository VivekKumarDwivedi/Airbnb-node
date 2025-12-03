package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	GetByID() (*models.User, error)
	Create(username string, email string, hashedPassword string) error
	GetAll() ([]*models.User, error)
	DeleteByID(id int64) error
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(_db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: _db,
	}
}

func (u *UserRepositoryImpl) GetAll() ([]*models.User, error) {
	return nil, nil
}
func (u *UserRepositoryImpl) DeleteByID(id int64) error {
	return nil
}
func (u *UserRepositoryImpl) Create(username string, email string, hashedPassword string) error {

	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"

	result, err := u.db.Exec(query, username, email, hashedPassword)

	if err != nil {
		fmt.Println("Error inserting user:", err)
		return err
	}

	rowAffected, rowErr := result.RowsAffected()

	if rowErr != nil {
		fmt.Println("Error getting rows affected:", rowErr)
		return rowErr
	}
	if rowAffected == 0 {
		fmt.Println("No rows were affected, user not created")
		return nil
	}

	fmt.Println("User created successfully, rows affected:", rowAffected)

	return nil
}
func (u *UserRepositoryImpl) GetByID() (*models.User, error) {
	fmt.Println("Fetching user in UserRepository")

	//step 1: Prepare the query
	query := "SELECT id,username,email,password,created_at,updated_at FROM users WHERE id =?"

	//step 2: execute the query
	row := u.db.QueryRow(query, 1)

	//step 3: process the row data
	user := &models.User{}

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdateAt)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user Found with the given Id")
			return nil, err
		} else {
			fmt.Println("Error scaning users:", err)
			return nil, err
		}
	}

	// step 4: return  the user details
	fmt.Println("User Fetched successfully:", user)

	return user, nil
}
