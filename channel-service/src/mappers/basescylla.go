package mappers

import (
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"

	"context"
)

// ** scylladb repo definition **
type ScyllaDBMapper[T any] struct {
	dbSession *gocqlx.Session
	ctx       context.Context
	tableName string
}

func (m *ScyllaDBMapper[T]) GetByKeys(
	// Get an item by its associated partition/clustering keys
	item *T, keyCount int, key ...string,
) error {
	comparisons := make([]qb.Cmp, keyCount)
	for index, key := range key {
		comparisons[index] = qb.Eq(key)
	}

	q := m.dbSession.Query(
		qb.Select(m.tableName).Where(comparisons...).ToCql(),
	).BindStruct(item)

	err := q.GetRelease(item)
	return err
}

func (m *ScyllaDBMapper[T]) Filter(
	item *T, keyCount int, key ...string,
) ([]T, error) {
	// Find all the items with provided filter values
	comparisons := make([]qb.Cmp, keyCount)
	for index, key := range key {
		comparisons[index] = qb.Eq(key)
	}

	q := m.dbSession.Query(
		qb.Select(m.tableName).Where(comparisons...).ToCql(),
	).BindStruct(item)
	var items []T
	items = make([]T, 0) // Used to avoid returning nil when no items in the array
	err := q.SelectRelease(&items)
	return items, err
}

func (m *ScyllaDBMapper[T]) Patch(item *T, setCount, columnTitleSet []string, keyCount int, key ...string) error {
	err := m.GetByKeys(item, keyCount, key...)
	if err != nil {
		return err
	}

	comparisons := make([]qb.Cmp, keyCount)
	for index, key := range key {
		comparisons[index] = qb.Eq(key)
	}
	q := m.dbSession.Query(
		qb.Update(m.tableName).Set(columnTitleSet...).Where(comparisons...).ToCql(),
	).BindStruct(item)

	return q.ExecRelease()
}

func (m *ScyllaDBMapper[T]) Delete(
	item *T, keyCount int, key ...string,
) error {
	err := m.GetByKeys(item, keyCount, key...)
	if err != nil {
		return err
	}

	comparisons := make([]qb.Cmp, keyCount)
	for index, key := range key {
		comparisons[index] = qb.Eq(key)
	}

	// Confirmed channel exists, continue with the delete
	q := m.dbSession.Query(
		qb.Delete(m.tableName).Where(comparisons...).ToCql(),
	).BindStruct(item)

	return q.ExecRelease()
}
