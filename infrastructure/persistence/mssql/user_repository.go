package mssql

import (
	"context"
	"database/sql"
	"fmt"
	irepo "github.com/tozastation/gRPC-Training-Golang/domain/repository"
	"github.com/tozastation/gRPC-Training-Golang/infrastructure/persistence/model/db"
	"github.com/tozastation/gRPC-Training-Golang/interfaces/auth"
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
	err := repo.DB.QueryRow("SELECT CityName FROM [Weather].[dbo].[Users] WHERE AccessToken = " + token).Scan(&dbUser.CityName)
	if err != nil {
		return nil, fmt.Errorf("select query failed: %v", err)
	}
	return &dbUser, nil
}

// CreateUser is ...
func (repo *UserRepository) CreateUser(user *db.User) (string, error) {
	stmt, err := repo.DB.Prepare("INSERT INTO [Weather].[dbo].[Users](CityName, Name, Password, AccessToken) VALUES(?, ?, ?, ?)")
	if err != nil {
		return "", fmt.Errorf("prepare query failed: %v", err)
	}
	defer func() {
		err := stmt.Close()
		if err != nil {
			fmt.Println("can't close statement!!", err)
		}
	}()
	_, err = stmt.Exec(user.Name, user.CityName, user.Password, user.AccessToken)
	if err != nil {
		return "", fmt.Errorf("execute query failed: %v", err)
	}
	return user.AccessToken, nil
}

// Login is ...
func (repo *UserRepository) Login(uID, password string) (string, error) {
	dbUser := db.User{}
	err := repo.DB.QueryRow("SELECT CityName, Password FROM [Weather].[dbo].[Users] WHERE Id = "+uID).Scan(&dbUser.CityName, &dbUser.Password)
	if err != nil {
		return "", fmt.Errorf("select query failed: %v", err)
	}
	err = auth.CheckHash(dbUser.Password, password)
	if err != nil {
		return "", fmt.Errorf("not correct password: %v", err)
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
