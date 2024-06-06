package postgres

import (
	"app/models"
	repoi "app/storage/repoI"
	"context"
	"log"

	"github.com/google/uuid"
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
	_, err := u.conn.Exec(
		ctx, 
		query, 
		user.UserID, 
		user.Fullname, 
		user.Username, 
		user.Gmail, 
		user.Password,
	)
	if err != nil {
		log.Println("error on CreateUser ", err)
		return err
	}
	return nil
}

func (u *UserRepo) GetUsers(ctx context.Context) ([]models.User, error) {
	var (
		users []models.User
		user models.User
	)

	query := `
		SELECT * 
			FROM users_todo;`
	rows, err := u.conn.Query(ctx, query)
	if err != nil {
		log.Println("Error on get USers", err)
	}
	for rows.Next() {

		err := rows.Scan(
			&user.UserID, 
			&user.Fullname, 
			&user.Username, 
			&user.Gmail, 
			&user.Password)
		if err != nil {
			log.Println("Error on Scan all users", err)
			return users, err
		}
		users = append(users, user)
	}

	defer rows.Close()

	return users, nil
}

func (u *UserRepo) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	var user models.User
	query := `
		SELECT * 
			FROM users_todo 
			WHERE 
				username = $1;`
	err := u.conn.QueryRow(
		ctx, 
		query, 
		username,
		).
	Scan(
		&user.UserID, 
		&user.Fullname, 
		&user.Username, 
		&user.Gmail, 
		&user.Password,
	)
	if err != nil {
		log.Println("err on GetUsersByUsername ", err)
		return user, err
	}

	return user, nil
}

func (u *UserRepo) UpdateUser(ctx context.Context, username string, user models.User) error {
	query := `UPDATE users_todo 
	SET username = $1,
		fullname = $2,
		gmail = $3,
		password = $4
	WHERE username = $5;`

	_, err := u.conn.Exec(
		ctx, 
		query, 
		user.Username, 
		user.Fullname, 
		user.Gmail, 
		user.Password, 
		username,
	)
	if err != nil {
		log.Println("error on Updating User ", err)
		return err
	}
	return nil

}

func (u *UserRepo) DeleteUserByUsername(ctx context.Context, UserId string) error {
	user_id, uerr := uuid.Parse(UserId)

	if uerr != nil {
		log.Println("incorrect user ID:", uerr)
		return uerr
	}

	todo_query := `
		DELETE
			FROM todos
			WHERE
				user_id=$1;`

	_, terr := u.conn.Exec(ctx, todo_query, user_id)

	if terr != nil {
		log.Println("Error on deleting todos when delete user:",terr)
		return terr
	}

	query := `
		DELETE 
			FROM users_todo 
			WHERE 
				user_id=$1;`

	_, err := u.conn.Exec(ctx, query, user_id)

	if err != nil {
		log.Println("Error on Deleting user", err)
		return err
	}

	return nil
}
