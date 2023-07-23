package psql

import (
	"context"
	"github.com/go-pg/pg/v10"
)

type User struct {
	//nolint:structcheck,gocritic,unused
	tableName struct{} `sql:"users"`
	ID        string   `sql:"id,notnull"`
	UserName  string   `sql:"user_name,notnull"`
	LastName  string   `sql:"last_name,notnull"`
	FirstName string   `sql:"first_name,notnull"`
	Password  string   `sql:"password,notnull"`
	Email     string   `sql:"email,notnull"`
	Mobile    string   `sql:"mobile,notnull"`
	ASZF      bool     `sql:"aszf,notnull"`
}

type UserPSQLRepository struct {
	db *pg.DB
}

func NewUserPSQLRepository(db *pg.DB) *UserPSQLRepository {
	return &UserPSQLRepository{db: db}
}

func (r *UserPSQLRepository) GetUserCount(
	ctx context.Context) (int, error) {
	user := new(User)
	count, err := r.db.WithContext(ctx).
		Model(user).
		Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *UserPSQLRepository) CreateUser(ctx context.Context, user User) (*User, error) {
	_, err := r.db.WithContext(ctx).
		Model(&user).
		Returning("*").
		Insert()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserPSQLRepository) GetUsers(ctx context.Context) ([]User, error) {
	users := []User{}
	err := r.db.WithContext(ctx).
		Model(&users).
		Select()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserPSQLRepository) UpdateUser(ctx context.Context, user User) (*User, error) {
	_, err := r.db.WithContext(ctx).
		Model(&user).
		Returning("*").
		Update()
	if err != nil {
		return nil, err
	}
	return &user, nil
}
