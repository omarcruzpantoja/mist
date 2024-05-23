package repositories

import (
	"context"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
)

type BaseRepo interface {
	GetByKey(ctx context.Context, key string) (interface{}, error)
	GetAll(ctx context.Context) ([]interface{}, error)
	Create(ctx context.Context, obj interface{}) (interface{}, error)
	Update(ctx context.Context, obj interface{}) (interface{}, error)
	Delete(ctx context.Context, id string) error
}

// ** scylladb repo definition **
type ScyllaDBRepo[T any] struct {
	dbSession *gocqlx.Session
}

func NewScyllaDBRepo[T any](dbSession *gocqlx.Session) *ScyllaDBRepo[T] {
	return &ScyllaDBRepo[T]{dbSession: dbSession}
}

func (r *ScyllaDBRepo[T]) GetByKeys(ctx context.Context, tableName string, item *T, keyCount int, key ...string) (*T, error) {
	comparisons := make([]qb.Cmp, keyCount)
	for index, key := range key {
		comparisons[index] = qb.Eq(key)
	}

	q := r.dbSession.Query(
		qb.Select(tableName).Where(comparisons...).ToCql(),
	).BindStruct(item)

	err := q.GetRelease(item)
	return item, err
}

func (r *ScyllaDBRepo[T]) Filter(ctx context.Context, tableName string, item *T, keyCount int, key ...string) ([]*T, error) {
	comparisons := make([]qb.Cmp, keyCount)
	for index, key := range key {
		comparisons[index] = qb.Eq(key)
	}

	q := r.dbSession.Query(
		qb.Select(tableName).Where(comparisons...).ToCql(),
	).BindStruct(item)
	var items []*T
	err := q.SelectRelease(&items)
	return items, err
}

// type BaseMapper struct {
// 	Actions BaseRepo
// }
