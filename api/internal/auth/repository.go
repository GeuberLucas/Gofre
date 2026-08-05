package auth

import (
	"database/sql"
	"log"
	"time"
)

type ResetTokensRepository struct {
	db *sql.DB
}

func NewResetTokensRepository(db *sql.DB) *ResetTokensRepository {
	return &ResetTokensRepository{db: db}
}

func (r *ResetTokensRepository) CreateResetToken(token *ResetToken) error {
	var sqlCommand string = "insert into auth.reset_tokens (user_id, hash_token, expires_at) values ($1,$2,$3)"
	statement, err := r.db.Prepare(sqlCommand)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(token.UserID, token.TokenHash, token.ExpiresAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *ResetTokensRepository) GetResetTokenByTokenHash(tokenHash string) (ResetToken, error) {
	var resetToken ResetToken
	var sqlCommand string = "select id, user_id, hash_token, expires_at from auth.reset_tokens where hash_token = $1"
	row := r.db.QueryRow(sqlCommand, tokenHash)
	err := row.Scan(&resetToken.ID, &resetToken.UserID, &resetToken.TokenHash, &resetToken.ExpiresAt)
	if err != nil {
		return resetToken, err
	}
	return resetToken, nil
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user User) int64 {
	sqlCommand := `INSERT INTO auth.users (name, last_name, cell_phone, username, email, password, created_at, updated_at)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id
					`

	statement, err := r.db.Prepare(sqlCommand)
	if err != nil {
		log.Fatal(err)
		return 0
	}

	var idLastInsert int64
	err = statement.QueryRow(user.Name, user.LastName, user.Cellphone, user.Username, user.Email, user.Password, time.Now(), time.Now()).Scan(&idLastInsert)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	defer statement.Close()
	return idLastInsert
}
func (r *UserRepository) GetUsers() ([]User, error) {
	var sqlCommand string = "select id, username, name, last_name, cell_phone, email, password, created_at, updated_at from auth.users"

	rows, err := r.db.Query(sqlCommand)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Name, &user.LastName, &user.Cellphone, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) GetUserByUsername(username string) (User, error) {
	var user User
	var sqlCommand string = "select id, username, name, last_name, cell_phone, email, password, created_at, updated_at from auth.users where username = $1"

	row := r.db.QueryRow(sqlCommand, username)
	err := row.Scan(&user.ID, &user.Username, &user.Name, &user.LastName, &user.Cellphone, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}
func (r *UserRepository) GetUserByEmail(email string) (User, error) {
	var user User
	var sqlCommand string = "select id, username, name, last_name, cell_phone, email, password, created_at, updated_at from auth.users where email = $1"

	row := r.db.QueryRow(sqlCommand, email)
	err := row.Scan(&user.ID, &user.Username, &user.Name, &user.LastName, &user.Cellphone, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByID(id int64) (User, error) {
	var user User
	var sqlCommand string = "select id, username, name, last_name, cell_phone, email, password, created_at, updated_at from auth.users where id = $1"

	row := r.db.QueryRow(sqlCommand, id)
	err := row.Scan(&user.ID, &user.Username, &user.Name, &user.LastName, &user.Cellphone, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user User) {
	var sqlCommand string = "update auth.users set name=$1, last_name=$2, cell_phone=$3, username=$4, email=$5, password=$6, updated_at=$7 where id=$8"

	statement, err := r.db.Prepare(sqlCommand)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer statement.Close()
	_, err = statement.Exec(user.Name, user.LastName, user.Cellphone, user.Username, user.Email, user.Password, time.Now(), user.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("User updated successfully")

}

func (r *UserRepository) UpdateUserPassword(userId int64, password []byte) {
	var sqlCommand string = "update auth.users set  password=$1 ,updated_at=$2 where id=$3"

	statement, err := r.db.Prepare(sqlCommand)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer statement.Close()
	_, err = statement.Exec(password, time.Now(), userId)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("User updated successfully")
}

func (r *UserRepository) DeleteUser(id int64) {
	var sqlCommand string = "delete from auth.users where id=$1"

	statement, err := r.db.Prepare(sqlCommand)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer statement.Close()
	_, err = statement.Exec(id)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("User deleted successfully")
}
