package mssql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	irepo "github.com/tozastation/gRPC-Training-Golang/domain/repository"
	"github.com/tozastation/gRPC-Training-Golang/infrastructure/persistence/model/db"
)

// UserRepository is
type UserRepository struct {
	*sql.DB
}

// NewUserRepository is ...
func NewUserRepository(Conn *sql.DB) irepo.IUserRepository {
	return &UserRepository{Conn}
}

// FindUserByUserToken is ...
func (repo *UserRepository) FindUserByUserToken(ctx context.Context, token string) (*db.User, error) {
	dbUser := db.User{}
	if err := repo.DB.QueryRow("SELECT CityName FROM [Weather].[dbo].[Users] WHERE AccessToken = " + token).Scan(&dbUser.CityName); err != nil {
		return nil, err
	}
	return &dbUser, nil
}

// CreateUser is ...
func (repo *UserRepository) CreateUser(user *db.User) (string, error) {
	stmt, err := repo.DB.Prepare("INSERT INTO [Weather].[dbo].[Users](CityName, Name, Password, AccessToken) VALUES(?, ?, ?, ?)")
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Name, user.CityName, user.Password, user.AccessToken)
	if err != nil {
		return "", err
	}
	return user.AccessToken, nil
}

// Login is ...
func (repo *UserRepository) Login(uID string, password []byte) (string, error) {
	dbUser := db.User{}
	if err := repo.DB.QueryRow("SELECT CityName, Password FROM [Weather].[dbo].[Users] WHERE Id = "+uID).Scan(&dbUser.CityName, &dbUser.Password); err != nil {
		return "", err
	}
	fmt.Println(string(password))
	fmt.Println(string(dbUser.Password))
	if string(dbUser.Password) != string(password) {
		err := errors.New("Not Found User")
		return "", err
	}
	return dbUser.CityName, nil
}

func userMapped(rows *sql.Rows) ([]*db.User, error) {
	dbUsers := []*db.User{}
	for rows.Next() {
		dbUser := db.User{}
		if err := rows.Scan(&dbUser.CityName); err != nil {
			return nil, err
		}
		dbUsers = append(dbUsers, &dbUser)
	}
	return dbUsers, nil
}
