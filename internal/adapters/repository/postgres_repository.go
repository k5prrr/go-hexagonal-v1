package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type UserRepositoryPostgres struct {
	conn *pgx.Conn
}

func NewUserRepositoryPostgres() *UserRepositoryPostgres {
	return &UserRepositoryPostgres{}
}

func (u *UserRepositoryPostgres) Connect(name, user, password, address, port string) error {
	if address == "" {
		address = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	connectString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, address, port, name)

	var err error
	ctx := context.Background()

	u.conn, err = pgx.Connect(ctx, connectString)
	if err != nil {
		return fmt.Errorf("Unable to connect to database: %w\n", err)
	}

	return nil
}

func (u *UserRepositoryPostgres) CloseConnection() {
	if u.conn == nil {
		return
	}
	ctx := context.Background()
	u.conn.Close(ctx)
}
func (u *UserRepositoryPostgres) Create(data string) (string, error)    {}
func (u *UserRepositoryPostgres) Read(uuid string) (string, error)      {}
func (u *UserRepositoryPostgres) Update(uuid string, data string) error {}
func (u *UserRepositoryPostgres) Delete(uuid string) error              {}
