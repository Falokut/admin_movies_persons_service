package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/jackc/pgx/v5"
	stdlib "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
)

type personsRepository struct {
	db *sqlx.DB
}

const (
	personsTableName = "persons"
)

func NewPersonsRepository(db *sqlx.DB) *personsRepository {
	return &personsRepository{db: db}
}

func NewPostgreDB(cfg DBConfig) (*sqlx.DB, error) {
	conStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)

	if !cfg.EnablePreparedStatements {
		stdlib.RegisterConnConfig(&pgx.ConnConfig{
			DefaultQueryExecMode: pgx.QueryExecModeSimpleProtocol,
		})
	}

	return sqlx.Connect("pgx", conStr)
}

func (r *personsRepository) Shutdown() {
	r.db.Close()
}

func (r *personsRepository) GetPersons(ctx context.Context, ids []string, limit, offset int32) ([]Person, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "personsRepository.GetPersons")
	defer span.Finish()

	var err error
	defer span.SetTag("error", err != nil)

	query, args, err := sqlx.In(fmt.Sprintf("SELECT * FROM %s WHERE id IN(?) ORDER BY id LIMIT %d OFFSET %d",
		personsTableName, limit, offset), ids)
	if err != nil {
		return []Person{}, err
	}
	query = sqlx.Rebind(sqlx.DOLLAR, query)

	var persons []Person
	err = r.db.SelectContext(ctx, &persons, query, args...)
	if err != nil {
		return []Person{}, err
	} else if len(persons) == 0 {
		return persons, ErrNotFound
	}

	return persons, nil
}

const defaultIsExistLimit = 2

func (r *personsRepository) IsPersonAlreadyExists(ctx context.Context, person SearchPersonParam) (bool, []string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "personsRepository.IsPersonAlreadyExists")
	defer span.Finish()

	var err error
	defer span.SetTag("error", err != nil && !errors.Is(err, sql.ErrNoRows))

	whereStatement, args := r.getWhereStatement(person)
	query := fmt.Sprintf("SELECT id FROM %s %s LIMIT %d", personsTableName, whereStatement, defaultIsExistLimit)
	if len(args) == 0 {
		return false, []string{}, ErrInvalidArgument
	}
	var ids []string
	err = r.db.SelectContext(ctx, &ids, query, args...)
	if err != nil {
		return false, []string{}, err
	} else if len(ids) == 0 {
		return false, []string{}, nil
	}

	return true, ids, nil
}

func (r *personsRepository) SearchPersonByName(ctx context.Context, name string, limit, offset int32) ([]Person, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "personsRepository.SearchPersonByName")
	defer span.Finish()

	var err error
	defer span.SetTag("error", err != nil)

	query := fmt.Sprintf("SELECT * FROM %s WHERE LOWER(fullname_ru) LIKE($1)"+
		" OR LOWER(fullname_en) LIKE('$1') ORDER BY id LIMIT %d OFFSET %d;", personsTableName, limit, offset)

	var persons []Person
	err = r.db.SelectContext(ctx, &persons, query, strings.ToLower(name)+"%")
	if errors.Is(err, sql.ErrNoRows) {
		return []Person{}, ErrNotFound
	} else if err != nil {
		return []Person{}, err
	}
	return persons, nil
}

func (r *personsRepository) SearchPerson(ctx context.Context, person SearchPersonParam, limit, offset int32) ([]Person, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "personsRepository.SearchPerson")
	defer span.Finish()

	var err error
	defer span.SetTag("error", err != nil)

	whereStatement, args := r.getWhereStatement(person)
	query := fmt.Sprintf("SELECT * FROM %s %s ORDER BY id LIMIT %d OFFSET %d", personsTableName, whereStatement, limit, offset)
	if len(args) == 0 {
		return []Person{}, ErrInvalidArgument
	}

	var persons []Person
	err = r.db.SelectContext(ctx, &persons, query, args...)
	if errors.Is(err, sql.ErrNoRows) {
		return []Person{}, ErrNotFound
	} else if err != nil {
		return []Person{}, err
	}
	return persons, nil
}

func (r *personsRepository) GetAllPersons(ctx context.Context, limit, offset int32) ([]Person, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "personsRepository.GetAllPersons")
	defer span.Finish()

	var err error
	defer span.SetTag("error", err != nil)
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id LIMIT %d OFFSET %d",
		personsTableName, limit, offset)

	var persons []Person
	err = r.db.SelectContext(ctx, &persons, query)
	if errors.Is(err, sql.ErrNoRows) {
		return []Person{}, ErrNotFound
	} else if err != nil {
		return []Person{}, err
	}

	return persons, nil
}

func (r *personsRepository) DeletePersons(ctx context.Context, ids []string) ([]string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "personsRepository.DeletePersons")
	defer span.Finish()

	var err error
	defer span.SetTag("error", err != nil)

	query, args, err := sqlx.In(fmt.Sprintf("DELETE FROM %s WHERE id IN(?) RETURNING id", personsTableName), ids)
	if err != nil {
		return []string{}, err
	}

	query = sqlx.Rebind(sqlx.DOLLAR, query)
	var deletedIDs = make([]string, len(ids))
	err = r.db.SelectContext(ctx, &deletedIDs, query, args...)
	if errors.Is(err, sql.ErrNoRows) {
		return []string{}, ErrNotFound
	} else if err != nil {
		return []string{}, err
	}

	return deletedIDs, nil
}

func (r *personsRepository) CreatePerson(ctx context.Context, person CreatePersonParam) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "personsRepository.CreatePerson")
	defer span.Finish()

	var err error
	defer span.SetTag("error", err != nil)

	args, fields, values := r.getInsertStatement(person)
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s) RETURNING id", personsTableName, fields, values)

	var id []string
	err = r.db.SelectContext(ctx, &id, query, args...)
	if err != nil {
		return "", err
	} else if len(id) == 0 {
		return "", errors.New("something went wrong")
	}

	return id[0], nil
}

func (r *personsRepository) IsPersonWithIDExist(ctx context.Context, id string) (bool, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "personsRepository.IsPersonWithIDExist")
	defer span.Finish()

	var err error
	defer span.SetTag("error", err != nil && !errors.Is(err, sql.ErrNoRows))

	query := fmt.Sprintf("SELECT id FROM %s WHERE id=$1 LIMIT 1;", personsTableName)

	_, err = r.db.ExecContext(ctx, query, id)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *personsRepository) UpdatePerson(ctx context.Context, id string, toUpdate UpdatePersonParam, excludeDefaultValues bool) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "personsRepository.UpdatePerson")
	defer span.Finish()

	var err error
	defer span.SetTag("error", err != nil)

	setStatement, args := r.getSetStatement(toUpdate, []any{id}, excludeDefaultValues)
	if len(args) == 1 {
		return nil
	}

	query := fmt.Sprintf("UPDATE %s %s WHERE id=$1", personsTableName, setStatement)
	res, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if errors.Is(err, sql.ErrNoRows) || affected == 0 {
		return ErrNotFound
	}

	return err
}

func (r *personsRepository) getWhereStatement(person SearchPersonParam) (string, []any) {
	rv := reflect.ValueOf(person)
	rt := rv.Type()

	statements := make([]string, 0, rt.NumField())
	args := make([]any, 0, rt.NumField())
	index := 1

	for i := 0; i < rt.NumField(); i++ {
		v := rv.Field(i).Interface()

		if isDefaultValue(v) {
			continue
		}

		statements = append(statements, fmt.Sprintf("%s=$%d", rt.Field(i).Tag.Get("db"), index))
		args = append(args, v)
		index++
	}

	return " WHERE " + strings.Join(statements, " AND "), args
}

func (r *personsRepository) getSetStatement(toUpdate UpdatePersonParam, args []any, excludeDefault bool) (string, []any) {
	rv := reflect.ValueOf(toUpdate)
	rt := rv.Type()

	index := len(args) + 1
	statements := make([]string, 0, rt.NumField())

	for i := 0; i < rt.NumField(); i++ {
		v := rv.Field(i).Interface()

		if excludeDefault && isDefaultValue(v) {
			continue
		}

		statements = append(statements, fmt.Sprintf("%s=$%d", rt.Field(i).Tag.Get("db"), index))
		args = append(args, v)
		index++
	}

	return " SET " + strings.Join(statements, ", "), args
}

func (r *personsRepository) getInsertStatement(person CreatePersonParam) ([]any, string, string) {
	rv := reflect.ValueOf(person)
	rt := rv.Type()
	var fields = make([]string, 0, rt.NumField())
	var args = make([]any, 0, rt.NumField())
	var values = make([]string, 0, rt.NumField())
	index := 1

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)

		v := rv.Field(i).Interface()
		if isDefaultValue(v) {
			continue
		}

		fields = append(fields, field.Tag.Get("db"))
		values = append(values, fmt.Sprintf("$%d", index))
		args = append(args, v)
		index++
	}

	return args, strings.Join(fields, ", "), strings.Join(values, ", ")
}

func isDefaultValue(field interface{}) bool {
	fieldVal := reflect.ValueOf(field)

	return !fieldVal.IsValid() || fieldVal.Interface() == reflect.Zero(fieldVal.Type()).Interface()
}
