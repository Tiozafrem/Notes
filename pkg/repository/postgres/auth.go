package postgres

import (
	"fmt"
	"notes/model"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	err := row.Scan(&id)

	return id, err
}

func (r *AuthPostgres) GetUser(username, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", userTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}

func (r *AuthPostgres) GetUserByDeviceId(deviceId int) (model.User, error) {
	var user model.User
	query := fmt.Sprintf(
		`SELECT id 
		FROM %s users_table 
		INNER JOIN %s devices_users_table ON 
			users_table.id = devices_users_table.user_id
		WHERE devices_users_table = $1`,
		userTable, deviceUserTable)
	err := r.db.Get(&user, query, deviceId)

	return user, err
}

func (r *AuthPostgres) GetDeviceByRefreshToken(refreshToken string) (model.DeviceUser, error) {
	var deviceUser model.DeviceUser
	query := fmt.Sprintf(
		`SELECT *
		FROM %s device_user_table
		WHERE device_user_table.refresh_token = $1`,
		deviceUserTable)
	err := r.db.Get(&deviceUser, query, refreshToken)
	return deviceUser, err
}

func (r *AuthPostgres) CreateDevice(deviceUser model.DeviceUser) (int, error) {
	var id int
	query := fmt.Sprintf(
		`INSERT INTO %s (name, user_id, description, refresh_token, expire) 
		VALUES ($1, $2, $3, $4, $5) RETURNING id`, deviceUserTable)
	row := r.db.QueryRow(query, deviceUser.Name, deviceUser.UserId,
		deviceUser.Description, deviceUser.RefreshToken, deviceUser.Expire)
	err := row.Scan(&id)
	return id, err
}

func (r *AuthPostgres) DeleteDeviceByDeviceId(deviceId int) error {
	query := fmt.Sprintf(
		`DELETE FROM %s device_user_table
		WHERE device_user_table.id = $1`,
		deviceUserTable)
	_, err := r.db.Exec(query, deviceId)
	return err
}

func (r *AuthPostgres) UpdateRefreshTokenByDevice(deviceUser model.DeviceUser) error {
	query := fmt.Sprintf(
		`UPDATE %s device_user_table
		SET refresh_token = $2, expire = $3
		WHERE device_user_table.id = $1`,
		deviceUserTable)

	_, err := r.db.Exec(query, deviceUser.Id,
		deviceUser.RefreshToken, deviceUser.Expire)
	return err
}
