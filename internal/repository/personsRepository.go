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
	"github.com/sirupsen/logrus"
)

type personsRepository struct {
	db     *sqlx.DB
	logger *logrus.Logger
}

const (
	personsTableName = "persons"
)

func NewPersonsRepository(db *sqlx.DB, logger *logrus.Logger) *personsRepository {
	return &personsRepository{db: db, logger: logger}
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

func (r *personsRepository) GetPersons(ctx context.Context, ids []int32, limit, offset int32) ([]Person, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "personsRepository.GetPersons")
	defer span.Finish()

	var err error
	defer span.SetTag("error", err != nil)

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=ANY($1) ORDER BY id LIMIT %d OFFSET %d",
		personsTableName, limit, offset)

	var persons []Person
	err = r.db.SelectContext(ctx, &persons, query, ids)
	if err != nil {
		r.logger.Errorf("%v query: %s args: %v ", err.Error(), query, ids)
		return []Person{}, err
	} else if len(persons) == 0 {
		return persons, ErrNotFound
	}

	return persons, nil
}

func (r *personsRepository) IsPersonAlreadyExists(ctx context.Context, person SearchPersonParam) (bool, []int32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "personsRepository.IsPersonAlreadyExists")
	defer span.Finish()

	var err error
	defer span.SetTag("error", err != nil)

	whereStatement, args := r.getWhereStatement(person)
	query := fmt.Sprintf("SELECT id FROM %s %s", personsTableName, whereStatement)
	if len(args) == 0 {
		return false, []int32{}, ErrInvalidArgument
	}
	var ids []int32
	err = r.db.SelectContext(ctx, &ids, query, args...)
	if err != nil {
		r.logger.Errorf("%v query: %s args: %v", err.Error(), query, args)
		return false, []int32{}, err
	} else if len(ids) == 0 {
		return false, []int32{}, nil
	}

	return true, ids, nil
}

func (r *personsRepository) SearchPersonByName(ctx context.Context, name string, limit, offset int32) ([]Person, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "personsRepository.SearchPersonByName")
	defer span.Finish()

	var err error
	defer span.SetTag("error", err != nil)

	query := fmt.Sprintf("SELECT * FROM %s WHERE LOWER(fullname_ru) LIKE($1)"+
		" OR LOWER(fullname_en) LIKE($1) ORDER BY id LIMIT %d OFFSET %d;", personsTableName, limit, offset)

	var persons []Person
	err = r.db.SelectContext(ctx, &persons, query, strings.ToLower(name)+"%")
	if errors.Is(err, sql.ErrNoRows) {
		return []Person{}, ErrNotFound
	} else if err != nil {
		r.logger.Errorf("%v query: %s args: %v", err.Error(), query, name)
		return []Person{}, err
	}
	return persons, nil
}

func (r *personsRepository) SearchPerson(ctx context.Context, person SearchPersonParam, limit, offset int32) ([]Person, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "personsRepository.SearchPerson")
	defer span.Finish()

	var err error
	defer span.SetTag("error", err != nil && !errors.Is(err, sql.ErrNoRows))

	whereStatement, args := r.getWhereStatement(person)
	query := fmt.Sprintf("SELECT * FROM %s %s ORDER BY id LIMIT %d OFFSET %d",
		personsTableName, whereStatement, limit, offset)
	if len(args) == 0 {
		return []Person{}, ErrInvalidArgument
	}

	var persons []Person
	err = r.db.SelectContext(ctx, &persons, query, args...)
	if errors.Is(err, sql.ErrNoRows) {
		return []Person{}, ErrNotFound
	} else if err != nil {
		r.logger.Errorf("%v query: %s", err.Error(), query)
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
	if err != nil {
		r.logger.Errorf("%v query: %s", err.Error(), query)
		return []Person{}, err
	} else if len(persons) == 0 {
		return []Person{}, ErrNotFound
	}

	return persons, nil
}

func (r *personsRepository) DeletePersons(ctx context.Context, ids []int32) ([]int32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "personsRepository.DeletePersons")
	defer span.Finish()

	var err error
	defer span.SetTag("error", err != nil)

	query := fmt.Sprintf("DELETE FROM %s WHERE id=ANY($1) RETURNING id", personsTableName)

	var deletedIDs = make([]int32, len(ids))
	err = r.db.SelectContext(ctx, &deletedIDs, query, ids)
	if errors.Is(err, sql.ErrNoRows) {
		return []int32{}, ErrNotFound
	} else if err != nil {
		r.logger.Errorf("%v query: %s args: %v", err.Error(), query, ids)
		return []int32{}, err
	}

	return deletedIDs, nil
}

func (r *personsRepository) IsPersonsExists(ctx context.Context, ids []int32) ([]int32, bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "personsRepository.IsPersonsExists")
	defer span.Finish()

	var err error
	defer span.SetTag("error", err != nil)

	query := fmt.Sprintf("SELECT id FROM %s WHERE id=ANY($1);",
		personsTableName)

	var foundIDs []int32
	err = r.db.SelectContext(ctx, &foundIDs, query, ids)
	if err != nil {
		r.logger.Errorf("%v query: %s args: %v", err.Error(), query, ids)
		return []int32{}, false, err
	}

	return foundIDs, len(ids) == len(foundIDs), nil
}

func (r *personsRepository) CreatePerson(ctx context.Context, person CreatePersonParam) (int32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "personsRepository.CreatePerson")
	defer span.Finish()

	var err error
	defer span.SetTag("error", err != nil)

	args, fields, values := r.getInsertStatement(person)
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s) RETURNING id", personsTableName, fields, values)

	var id int32
	err = r.db.GetContext(ctx, &id, query, args...)
	if err != nil {
		r.logger.Errorf("%v query: %s args: %v", err.Error(), query, args)
		return 0, err
	}
	return id, nil
}

func (r *personsRepository) IsPersonWithIDExist(ctx context.Context, id int32) (bool, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "personsRepository.IsPersonWithIDExist")
	defer span.Finish()

	var err error
	defer span.SetTag("error", err != nil && !errors.Is(err, sql.ErrNoRows))

	query := fmt.Sprintf("SELECT id FROM %s WHERE id=$1 LIMIT 1;", personsTableName)

	err = r.db.GetContext(ctx, &id, query, id)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	} else if err != nil {
		r.logger.Errorf("%v query: %s args: %v", err.Error(), query, id)
		return false, err
	}

	return true, nil
}

func (r *personsRepository) UpdatePerson(ctx context.Context, id int32,
	toUpdate UpdatePersonParam, excludeDefaultValues bool) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "personsRepository.UpdatePerson")
	defer span.Finish()

	var err error
	defer span.SetTag("error", err != nil)

	setStatement, args := r.getSetStatement(toUpdate, []any{id}, excludeDefaultValues)
	if len(args) == 1 {
		return nil
	}

	query := fmt.Sprintf("UPDATE %s %s WHERE id=$1", personsTableName, setStatement)
	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		r.logger.Errorf("%v query: %s args: %v", err.Error(), query, id)
		return err
	}
	return nil
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
