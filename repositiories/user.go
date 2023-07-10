package repositiories

import (
	"database/sql"
	"fmt"

	"github.com/aungkokoye/go_app/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) ContactSync() (map[string]models.SyncUser, error) {

	query := `
        SELECT
		    u.user_id,
			u.username,
			u.email,
			u.firstname,
			u.lastname,
			u.created,
			u.last_login
		FROM user u
		`
	users := make(map[string]models.SyncUser)
	rows, err := u.db.Query(query)

	if err != nil {
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		var user models.SyncUser
		err := rows.Scan(
			&user.UserID,
			&user.Username,
			&user.Email,
			&user.Firstname,
			&user.Lastname,
			&user.Created,
			&user.LastLogin,
		)

		if err != nil {
			fmt.Printf("Error in mapping data %s", err)
		} else {
			users[user.Email] = user
		}
	}

	return users, nil

}
