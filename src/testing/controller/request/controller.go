// Code generated by mockery v2.9.4. DO NOT EDIT.

package request

import (
	context "context"

	models "github.com/goharbor/harbor/src/pkg/request/models"
	mock "github.com/stretchr/testify/mock"

	q "github.com/goharbor/harbor/src/lib/q"

	request "github.com/goharbor/harbor/src/controller/request"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// Approve provides a mock function with given fields: ctx, project
func (_m *Controller) Approve(ctx context.Context, project *models.Request) error {
	ret := _m.Called(ctx, project)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Request) error); ok {
		r0 = rf(ctx, project)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Count provides a mock function with given fields: ctx, query
func (_m *Controller) Count(ctx context.Context, query *q.Query) (int64, error) {
	ret := _m.Called(ctx, query)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) int64); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, project
func (_m *Controller) Create(ctx context.Context, project *models.Request) (int64, error) {
	ret := _m.Called(ctx, project)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *models.Request) int64); ok {
		r0 = rf(ctx, project)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Request) error); ok {
		r1 = rf(ctx, project)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Controller) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exists provides a mock function with given fields: ctx, requestIDOrName
func (_m *Controller) Exists(ctx context.Context, requestIDOrName interface{}) (bool, error) {
	ret := _m.Called(ctx, requestIDOrName)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) bool); ok {
		r0 = rf(ctx, requestIDOrName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, requestIDOrName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, requestIDOrName, options
func (_m *Controller) Get(ctx context.Context, requestIDOrName interface{}, options ...request.Option) (*models.Request, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, requestIDOrName)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *models.Request
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...request.Option) *models.Request); ok {
		r0 = rf(ctx, requestIDOrName, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Request)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...request.Option) error); ok {
		r1 = rf(ctx, requestIDOrName, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: ctx, requestName, options
func (_m *Controller) GetByName(ctx context.Context, requestName string, options ...request.Option) (*models.Request, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, requestName)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *models.Request
	if rf, ok := ret.Get(0).(func(context.Context, string, ...request.Option) *models.Request); ok {
		r0 = rf(ctx, requestName, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Request)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...request.Option) error); ok {
		r1 = rf(ctx, requestName, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, query, options
func (_m *Controller) List(ctx context.Context, query *q.Query, options ...request.Option) ([]*models.Request, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*models.Request
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query, ...request.Option) []*models.Request); ok {
		r0 = rf(ctx, query, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Request)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *q.Query, ...request.Option) error); ok {
		r1 = rf(ctx, query, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Reject provides a mock function with given fields: ctx, project
func (_m *Controller) Reject(ctx context.Context, project *models.Request) error {
	ret := _m.Called(ctx, project)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Request) error); ok {
		r0 = rf(ctx, project)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
