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

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(userID), nil
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

func (repo UserRepository) FindById(userID uint64) (models.User, error) {
	lines, err := repo.db.Query(
		"select id, name, nick, email, password, createdAt from users where id = ?;",
		userID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repo UserRepository) FindByEmail(userEmail string) (models.User, error) {
	lines, err := repo.db.Query(
		"select id, name, nick, email, password, createdAt from users where email = ?;",
		userEmail,
	)
	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repo UserRepository) Update(
	userID uint64,
	user models.User,
) error {
	statement, err := repo.db.Prepare(
		"update users set name = ?, nick = ?, email = ? where id = ?;",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, userID); err != nil {
		return err
	}

	return nil
}

func (repo UserRepository) Delete(userID uint64) error {
	statement, err := repo.db.Prepare(
		"delete from users where id = ?;",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID); err != nil {
		return err
	}

	return nil
}

func (repo UserRepository) FollowUser(userID, followerID uint64) error {
	statement, err := repo.db.Prepare(
		"insert ignore into followers (user_id, follower_id) values (? , ?);",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

func (repo UserRepository) UnfollowUser(userID, followerID uint64) error {
	statement, err := repo.db.Prepare(
		"delete from followers where user_id = ? and follower_id = ?;",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

func (repo UserRepository) FindFollowers(userID uint64) ([]models.User, error) {
	lines, err := repo.db.Query(
		`
			select u.id, u.name, u.nick, u.email, u.createdAt
			from users u
			inner join  followers f on u.id = f.follower_id
			where f.user_id = ?
		`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User
		lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		)

		users = append(users, user)
	}

	return users, nil
}

func (repo UserRepository) FindFollowing(userID uint64) ([]models.User, error) {
	lines, err := repo.db.Query(
		`
			select u.id, u.name, u.nick, u.email, u.createdAt
			from users u
			inner join  followers f on u.id = f.user_id
			where f.follower_id = ?
		`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User
		lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		)

		users = append(users, user)
	}

	return users, nil
}

func (repo UserRepository) UpdatePassword(
	userID uint64,
	password string,
) error {
	statement, err := repo.db.Prepare(
		"update users set password = ? where id = ?;",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(password, userID); err != nil {
		return err
	}

	return nil
}
