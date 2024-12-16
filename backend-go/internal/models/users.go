package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	HashedPassword []byte    `json:"hashed_password"`
	Created        time.Time `json:"created_at"`
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Authenticate(name, password string) (int, error) {
	var id int
	var hashedPassword []byte
	fmt.Println("name used for query:", name)
	stmt := "SELECT id, hashed_password FROM users WHERE name = ?"
	row := m.DB.QueryRow(stmt, name)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("fuck me")
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}
	return id, nil
}

func (m *UserModel) Insert(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, hashed_password) VALUES (?, ?)`

	_, err = m.DB.Exec(stmt, username, hashedPassword)
	if err != nil {
		fmt.Println(err)
		var mySQLError *mysql.MySQLError
		if errors.As(err, &mySQLError) {
			if mySQLError.Number == 1062 { //&& strings.Contains(mySQLError.Message, "users_uc_email") {
				return ErrDuplicateUsername
			}
		}
		return err
	}
	return nil
}


func (m *UserModel) UpdateFavoriteTeam(userID int, teamName string) error {
	stmt := `UPDATE users SET favteam = ? WHERE id = ?`
	_, err := m.DB.Exec(stmt, teamName, userID)
	return err
}

func (m *UserModel) RemoveFavoriteTeam(userID int) error {
	stmt := `UPDATE users SET favteam = NULL WHERE id = ?`
	_, err := m.DB.Exec(stmt, userID)
	return err
}

func (m *UserModel) GetFavoriteTeam(userID int) (string, error) {
	var favTeam sql.NullString
	stmt := `SELECT favteam FROM users WHERE id = ?`
	err := m.DB.QueryRow(stmt, userID).Scan(&favTeam)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrUserNotFound
		}
		return "", err
	}

	if favTeam.Valid {
		return favTeam.String, nil
	}
	return "", nil 
}
