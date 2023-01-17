package repositories

import (
	"database/sql"
	"fmt"

	"github.com/brunohubner/devbook-api/src/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (repo UserRepository) Create(user models.User) (uint64, error) {
	statement, err := repo.db.Prepare(
		"insert into users (name, nick, email, password) values(?, ?, ?, ?);",
	)

	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	userId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(userId), nil
}

func (repo UserRepository) FindMany(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%

	lines, err := repo.db.Query(
		"select id, name, nick, email, createdAt from users where name LIKE ? or nick LIKE ?;",
		nameOrNick,
		nameOrNick,
	)

	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
