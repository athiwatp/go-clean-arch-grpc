// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import domain "github.com/bxcodec/go-clean-arch-grpc/domain"
import mock "github.com/stretchr/testify/mock"

// ArticleUsecase is an autogenerated mock type for the ArticleUsecase type
type ArticleUsecase struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *ArticleUsecase) Delete(id int64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: cursor, num
func (_m *ArticleUsecase) Fetch(cursor string, num int64) ([]domain.Article, string, error) {
	ret := _m.Called(cursor, num)

	var r0 []domain.Article
	if rf, ok := ret.Get(0).(func(string, int64) []domain.Article); ok {
		r0 = rf(cursor, num)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Article)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string, int64) string); ok {
		r1 = rf(cursor, num)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, int64) error); ok {
		r2 = rf(cursor, num)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: id
func (_m *ArticleUsecase) GetByID(id int64) (domain.Article, error) {
	ret := _m.Called(id)

	var r0 domain.Article
	if rf, ok := ret.Get(0).(func(int64) domain.Article); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.Article)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByTitle provides a mock function with given fields: title
func (_m *ArticleUsecase) GetByTitle(title string) (domain.Article, error) {
	ret := _m.Called(title)

	var r0 domain.Article
	if rf, ok := ret.Get(0).(func(string) domain.Article); ok {
		r0 = rf(title)
	} else {
		r0 = ret.Get(0).(domain.Article)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: _a0
func (_m *ArticleUsecase) Store(_a0 *domain.Article) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Article) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ar
func (_m *ArticleUsecase) Update(ar *domain.Article) error {
	ret := _m.Called(ar)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Article) error); ok {
		r0 = rf(ar)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
