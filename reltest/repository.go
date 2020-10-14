package reltest

import (
	"context"
	"runtime"
	"testing"

	"github.com/go-rel/rel"
	"github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	repo    rel.Repository
	mock    mock.Mock
	ctxData ctxData
}

var _ rel.Repository = (*Repository)(nil)

// Adapter provides a mock function with given fields:
func (r *Repository) Adapter(ctx context.Context) rel.Adapter {
	return r.repo.Adapter(ctx)
}

// Instrumentation provides a mock function with given fields: instrumenter
func (r *Repository) Instrumentation(instrumenter rel.Instrumenter) {
	r.repo.Instrumentation(instrumenter)
}

// Ping database.
func (r *Repository) Ping(ctx context.Context) error {
	return r.repo.Ping(ctx)
}

// Iterate through a collection of records from database in batches.
// This function returns iterator that can be used to loop all records.
// Limit, Offset and Sort query is automatically ignored.
func (r *Repository) Iterate(ctx context.Context, query rel.Query, options ...rel.IteratorOption) rel.Iterator {
	ret := r.mock.Called(fetchContext(ctx), query, options).Get(0)
	return (*iterator)(ret.(*Iterate))
}

// ExpectIterate apply mocks and expectations for Iterate
func (r *Repository) ExpectIterate(query rel.Query, options ...rel.IteratorOption) *Iterate {
	return ExpectIterate(r, query, options)
}

// Aggregate provides a mock function with given fields: query, aggregate, field
func (r *Repository) Aggregate(ctx context.Context, query rel.Query, aggregate string, field string) (int, error) {
	r.repo.Aggregate(ctx, query, aggregate, field)
	ret := r.mock.Called(fetchContext(ctx), query, aggregate, field)
	return ret.Int(0), ret.Error(1)
}

// MustAggregate provides a mock function with given fields: query, aggregate, field
func (r *Repository) MustAggregate(ctx context.Context, query rel.Query, aggregate string, field string) int {
	result, err := r.Aggregate(ctx, query, aggregate, field)
	must(err)
	return result
}

// ExpectAggregate apply mocks and expectations for Aggregate
func (r *Repository) ExpectAggregate(query rel.Query, aggregate string, field string) *Aggregate {
	return ExpectAggregate(r, query, aggregate, field)
}

// Count provides a mock function with given fields: collection, queriers
func (r *Repository) Count(ctx context.Context, collection string, queriers ...rel.Querier) (int, error) {
	r.repo.Count(ctx, collection, queriers...)
	ret := r.mock.Called(fetchContext(ctx), collection, queriers)
	return ret.Int(0), ret.Error(1)
}

// MustCount provides a mock function with given fields: collection, queriers
func (r *Repository) MustCount(ctx context.Context, collection string, queriers ...rel.Querier) int {
	count, err := r.Count(ctx, collection, queriers...)
	must(err)
	return count
}

// ExpectCount apply mocks and expectations for Count
func (r *Repository) ExpectCount(collection string, queriers ...rel.Querier) *Aggregate {
	return ExpectCount(r, collection, queriers)
}

// Find provides a mock function with given fields: record, queriers
func (r *Repository) Find(ctx context.Context, record interface{}, queriers ...rel.Querier) error {
	r.repo.Find(ctx, record, queriers...)
	return r.mock.Called(fetchContext(ctx), record, queriers).Error(0)
}

// MustFind provides a mock function with given fields: record, queriers
func (r *Repository) MustFind(ctx context.Context, record interface{}, queriers ...rel.Querier) {
	must(r.Find(ctx, record, queriers...))
}

// ExpectFind apply mocks and expectations for Find
func (r *Repository) ExpectFind(queriers ...rel.Querier) *Find {
	return ExpectFind(r, queriers)
}

// FindAll provides a mock function with given fields: records, queriers
func (r *Repository) FindAll(ctx context.Context, records interface{}, queriers ...rel.Querier) error {
	r.repo.FindAll(ctx, records, queriers...)
	return r.mock.Called(fetchContext(ctx), records, queriers).Error(0)
}

// MustFindAll provides a mock function with given fields: records, queriers
func (r *Repository) MustFindAll(ctx context.Context, records interface{}, queriers ...rel.Querier) {
	must(r.FindAll(ctx, records, queriers...))
}

// ExpectFindAll apply mocks and expectations for FindAll
func (r *Repository) ExpectFindAll(queriers ...rel.Querier) *FindAll {
	return ExpectFindAll(r, queriers)
}

// FindAndCountAll provides a mock function with given fields: records, queriers
func (r *Repository) FindAndCountAll(ctx context.Context, records interface{}, queriers ...rel.Querier) (int, error) {
	r.repo.FindAndCountAll(ctx, records, queriers...)
	ret := r.mock.Called(fetchContext(ctx), records, queriers)
	return ret.Int(0), ret.Error(1)
}

// MustFindAndCountAll provides a mock function with given fields: records, queriers
func (r *Repository) MustFindAndCountAll(ctx context.Context, records interface{}, queriers ...rel.Querier) int {
	count, err := r.FindAndCountAll(ctx, records, queriers...)
	must(err)
	return count
}

// ExpectFindAndCountAll apply mocks and expectations for FindAndCountAll
func (r *Repository) ExpectFindAndCountAll(queriers ...rel.Querier) *FindAndCountAll {
	return ExpectFindAndCountAll(r, queriers)
}

// Insert provides a mock function with given fields: record, mutators
func (r *Repository) Insert(ctx context.Context, record interface{}, mutators ...rel.Mutator) error {
	ret := r.mock.Called(fetchContext(ctx), record, mutators)

	r.repo.Insert(ctx, record, mutators...)
	return ret.Error(0)
}

// MustInsert provides a mock function with given fields: record, mutators
func (r *Repository) MustInsert(ctx context.Context, record interface{}, mutators ...rel.Mutator) {
	must(r.Insert(ctx, record, mutators...))
}

// ExpectInsert apply mocks and expectations for Insert
func (r *Repository) ExpectInsert(mutators ...rel.Mutator) *Mutate {
	return ExpectInsert(r, mutators)
}

// InsertAll records.
func (r *Repository) InsertAll(ctx context.Context, records interface{}) error {
	ret := r.mock.Called(fetchContext(ctx), records)

	r.repo.InsertAll(ctx, records)
	return ret.Error(0)
}

// MustInsertAll records.
func (r *Repository) MustInsertAll(ctx context.Context, records interface{}) {
	must(r.InsertAll(ctx, records))
}

// ExpectInsertAll records.
func (r *Repository) ExpectInsertAll() *Mutate {
	return ExpectInsertAll(r)
}

// Update provides a mock function with given fields: record, mutators
func (r *Repository) Update(ctx context.Context, record interface{}, mutators ...rel.Mutator) error {
	ret := r.mock.Called(fetchContext(ctx), record, mutators)

	if err := r.repo.Update(ctx, record, mutators...); err != nil {
		return err
	}

	return ret.Error(0)
}

// MustUpdate provides a mock function with given fields: record, mutators
func (r *Repository) MustUpdate(ctx context.Context, record interface{}, mutators ...rel.Mutator) {
	must(r.Update(ctx, record, mutators...))
}

// ExpectUpdate apply mocks and expectations for Update
func (r *Repository) ExpectUpdate(mutators ...rel.Mutator) *Mutate {
	return ExpectUpdate(r, mutators)
}

// UpdateAll provides a mock function with given fields: query
func (r *Repository) UpdateAll(ctx context.Context, query rel.Query, mutates ...rel.Mutate) error {
	return r.mock.Called(fetchContext(ctx), query, mutates).Error(0)
}

// MustUpdateAll provides a mock function with given fields: query
func (r *Repository) MustUpdateAll(ctx context.Context, query rel.Query, mutates ...rel.Mutate) {
	must(r.UpdateAll(ctx, query, mutates...))
}

// ExpectUpdateAll apply mocks and expectations for UpdateAll
func (r *Repository) ExpectUpdateAll(query rel.Query, mutates ...rel.Mutate) *MutateAll {
	return ExpectUpdateAll(r, query, mutates)
}

// Delete provides a mock function with given fields: record
func (r *Repository) Delete(ctx context.Context, record interface{}, options ...rel.Cascade) error {
	return r.mock.Called(fetchContext(ctx), record, options).Error(0)
}

// MustDelete provides a mock function with given fields: record
func (r *Repository) MustDelete(ctx context.Context, record interface{}, options ...rel.Cascade) {
	must(r.Delete(ctx, record, options...))
}

// ExpectDelete apply mocks and expectations for Delete
func (r *Repository) ExpectDelete(options ...rel.Cascade) *Delete {
	return ExpectDelete(r, options)
}

// DeleteAll provides a mock function with given fields: query
func (r *Repository) DeleteAll(ctx context.Context, query rel.Query) error {
	return r.mock.Called(fetchContext(ctx), query).Error(0)
}

// MustDeleteAll provides a mock function with given fields: query
func (r *Repository) MustDeleteAll(ctx context.Context, query rel.Query) {
	must(r.DeleteAll(ctx, query))
}

// ExpectDeleteAll apply mocks and expectations for DeleteAll
func (r *Repository) ExpectDeleteAll(query rel.Query) *MutateAll {
	return ExpectDeleteAll(r, query)
}

// Preload provides a mock function with given fields: records, field, queriers
func (r *Repository) Preload(ctx context.Context, records interface{}, field string, queriers ...rel.Querier) error {
	return r.mock.Called(fetchContext(ctx), records, field, queriers).Error(0)
}

// MustPreload provides a mock function with given fields: records, field, queriers
func (r *Repository) MustPreload(ctx context.Context, records interface{}, field string, queriers ...rel.Querier) {
	must(r.Preload(ctx, records, field, queriers...))
}

// ExpectPreload apply mocks and expectations for Preload
func (r *Repository) ExpectPreload(field string, queriers ...rel.Querier) *Preload {
	return ExpectPreload(r, field, queriers)
}

// Transaction provides a mock function with given fields: fn
func (r *Repository) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	ctxData := fetchContext(ctx)
	r.mock.Called(ctxData)

	var err error
	func() {
		defer func() {
			if p := recover(); p != nil {
				switch e := p.(type) {
				case runtime.Error:
					panic(e)
				case error:
					err = e
				default:
					panic(e)
				}
			}
		}()

		ctxData.txDepth++
		ctx = wrapContext(ctx, ctxData)
		err = fn(ctx)
	}()

	return err
}

// ExpectTransaction declare expectation inside transaction.
func (r *Repository) ExpectTransaction(fn func(*Repository)) {
	r.mock.On("Transaction", r.ctxData).Once()

	r.ctxData.txDepth++
	fn(r)
	r.ctxData.txDepth--
}

// Exec raw sql.
func (r *Repository) Exec(ctx context.Context, stmt string, args ...interface{}) (int64, error) {
	ret := r.mock.Called(ctx, stmt, args)
	return (int64)(ret.Int(0)), ret.Error(1)
}

// AssertExpectations asserts that everything was in fact called as expected. Calls may have occurred in any order.
func (r *Repository) AssertExpectations(t *testing.T) bool {
	return r.mock.AssertExpectations(t)
}

// New test repository.
func New() *Repository {
	return &Repository{
		repo: rel.New(&nopAdapter{}),
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
