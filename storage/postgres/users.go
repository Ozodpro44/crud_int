package postgres

import (
	"app/models"
	repoi "app/storage/repoI"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type UserRepo struct {
	conn *pgx.Conn
}

func NewUserRepo(conn *pgx.Conn) repoi.UserRepoI {

	userRepo := &UserRepo{
		conn: conn,
	}

	return userRepo
}

func (u *UserRepo) Login(ctx context.Context, username string) (models.User, error) {
	var user models.User
	query := `
		SELECT * FROM users_todo
			WHERE username = $1;`
	err := u.conn.QueryRow(ctx, query, username).Scan(&user.UserID, &user.Fullname, &user.Username, &user.Gmail, &user.Password)
	if err != nil {
		log.Println("Error user to login", err)
		return user, err
	}
	return user, nil
}

func (u *UserRepo) CreateUser(ctx context.Context, user models.User) error {
	query := `
		INSERT INTO 
		users_todo (
			user_id,
			fullname,
			username,
			gmail,
			password
			) VALUES (
				$1, $2, $3, $4, $5
			)`
	_, err := u.conn.Exec(ctx, query, user.UserID, user.Fullname, user.Username, user.Gmail, user.Password)
	if err != nil {
		log.Println("error on CreateUser ", err)
		return err
	}
	return nil
}

func (u *UserRepo) GetUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	var user models.User
	query := "SELECT * FROM users_todo;"
	row, err := u.conn.Query(ctx, query)
	if err != nil {
		log.Println("Error on get USers", err)
	}
	for row.Next() {

		err := row.Scan(&user.UserID, &user.Fullname, &user.Username, &user.Gmail, &user.Password)
		if err != nil {
			log.Println("Error on Scan all users", err)
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *UserRepo) GetUserByUsername(ctx context.Context, userId string) (*models.User, error) {
	return nil, nil
}

func (u *UserRepo) UpdateUser(ctx context.Context, username string, user models.User) error {
	query := "UPDATE users SET username = $1,fullname = $2,gmail = $3,password = $4 WHERE username = $5;"
	_, err := u.conn.Exec(ctx, query, user.Username, user.Fullname, user.Gmail, user.Password, username)
	if err != nil {
		log.Println("error on Updating User ", err)
		return err
	}
	return nil

}

func (u *UserRepo) DeleteUserByUsername(ctx context.Context, username string) error {
	query := "DELETE FROM users_todo WHERE username=$1"

	_, err := u.conn.Exec(ctx, query, username)

	if err != nil {
		log.Println("Error on Deleting user", err)
		return err
	}

	return nil
}
