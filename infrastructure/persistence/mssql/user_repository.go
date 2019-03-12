package mssql

import (
	"context"
	"database/sql"
	"fmt"
	irepo "github.com/tozastation/go-grpc-ddd-example/domain/repository"
	"github.com/tozastation/go-grpc-ddd-example/infrastructure/persistence/model/db"
	"github.com/tozastation/go-grpc-ddd-example/interfaces/auth"
	"github.com/tozastation/go-grpc-ddd-example/interfaces/error/irepoerror"
)

// UserRepository is
type UserRepository struct {
	*sql.DB
	irepoerror.IRepositoryError
}

// NewUserRepository is ...
func NewUserRepository(Conn *sql.DB, e irepoerror.IRepositoryError) irepo.IUserRepository {
	return &UserRepository{Conn, e}
}

// FindUserByUserToken is ...
func (repo *UserRepository) FindUserByUserToken(ctx context.Context, token string) (*db.User, error) {
	dbUser := db.User{}
	err := repo.DB.QueryRow("SELECT CityName FROM [Weather].[dbo].[Users] WHERE AccessToken = " + token).Scan(&dbUser.CityName)
	if err != nil {
		return nil, repo.IRepositoryError.SelectFailed(err)
	}
	return &dbUser, nil
}

// CreateUser is ...
func (repo *UserRepository) CreateUser(user *db.User) (string, error) {
	stmt, err := repo.DB.Prepare("INSERT INTO [Weather].[dbo].[Users](CityName, Name, Password, AccessToken) VALUES(?, ?, ?, ?)")
	if err != nil {
		return "", repo.IRepositoryError.PrepareStatementFailed(err)
	}
	defer func() {
		err := stmt.Close()
		if err != nil {
			fmt.Println("can't close statement!!", err)
		}
	}()
	_, err = stmt.Exec(user.Name, user.CityName, user.Password, user.AccessToken)
	if err != nil {
		return "", repo.IRepositoryError.ExecuteQueryFailed(err)
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
