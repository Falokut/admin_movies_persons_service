package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type DBConfig struct {
	Host                     string `yaml:"host" env:"DB_HOST"`
	Port                     string `yaml:"port" env:"DB_PORT"`
	Username                 string `yaml:"username" env:"DB_USERNAME"`
	Password                 string `yaml:"password" env:"DB_PASSWORD"`
	DBName                   string `yaml:"db_name" env:"DB_NAME"`
	SSLMode                  string `yaml:"ssl_mode" env:"DB_SSL_MODE"`
	EnablePreparedStatements bool   `yaml:"enable_prepared_statements" env:"DB_ENABLE_PREPARED_STATEMENTS"`
}

type Person struct {
	ID         string         `db:"id"`
	FullnameRU string         `db:"fullname_ru"`
	FullnameEN sql.NullString `db:"fullname_en"`
	Birthday   sql.NullTime   `db:"birthday"`
	Sex        sql.NullString `db:"sex"`
	PhotoID    sql.NullString `db:"photo_id"`
}

type UpdatePersonParam struct {
	FullnameRU string    `db:"fullname_ru"`
	FullnameEN string    `db:"fullname_en"`
	Birthday   time.Time `db:"birthday"`
	Sex        string    `db:"sex"`
	PhotoID    string    `db:"photo_id"`
}

type SearchPersonParam struct {
	FullnameRU string    `db:"fullname_ru"`
	FullnameEN string    `db:"fullname_en"`
	Birthday   time.Time `db:"birthday"`
	Sex        string    `db:"sex"`
}

type CreatePersonParam struct {
	FullnameRU string    `db:"fullname_ru"`
	FullnameEN string    `db:"fullname_en"`
	Birthday   time.Time `db:"birthday"`
	Sex        string    `db:"sex"`
	PhotoID    string    `db:"photo_id"`
}

var ErrNotFound = errors.New("entity not found")
var ErrInvalidArgument = errors.New("invalid input data")

type PersonsRepository interface {
	GetPersons(ctx context.Context, ids []int32, limit, offset int32) ([]Person, error)
	GetAllPersons(ctx context.Context, limit, offset int32) ([]Person, error)
	DeletePersons(ctx context.Context, ids []int32) ([]int32, error)
	SearchPerson(ctx context.Context, person SearchPersonParam, limit, offset int32) ([]Person, error)
	UpdatePerson(ctx context.Context, id int32, toUpdate UpdatePersonParam, excludeDefaultValues bool) error
	CreatePerson(ctx context.Context, person CreatePersonParam) (int32, error)
	IsPersonWithIDExist(ctx context.Context, id int32) (bool, error)
	IsPersonAlreadyExists(ctx context.Context, person SearchPersonParam) (bool, []int32, error)
	IsPersonsExists(ctx context.Context, ids []int32) ([]int32, bool, error)
	SearchPersonByName(ctx context.Context, name string, limit, offset int32) ([]Person, error)
}
