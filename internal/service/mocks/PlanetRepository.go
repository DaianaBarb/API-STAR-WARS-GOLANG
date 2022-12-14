// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"
	model "api-star-wars-golang/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// PlanetRepository is an autogenerated mock type for the PlanetRepository type
type PlanetRepository struct {
	mock.Mock
}

// DeleteById provides a mock function with given fields: ctx, id
func (_m *PlanetRepository) DeleteById(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindById provides a mock function with given fields: ctx, id
func (_m *PlanetRepository) FindById(ctx context.Context, id string) (*model.Planet, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.Planet
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Planet); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Planet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByParam provides a mock function with given fields: ctx, param
func (_m *PlanetRepository) FindByParam(ctx context.Context, param *model.PlanetIn) ([]model.Planet, error) {
	ret := _m.Called(ctx, param)

	var r0 []model.Planet
	if rf, ok := ret.Get(0).(func(context.Context, *model.PlanetIn) []model.Planet); ok {
		r0 = rf(ctx, param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Planet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.PlanetIn) error); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: parentContext, planet
func (_m *PlanetRepository) Save(parentContext context.Context, planet *model.Planet) (string, error) {
	ret := _m.Called(parentContext, planet)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *model.Planet) string); ok {
		r0 = rf(parentContext, planet)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.Planet) error); ok {
		r1 = rf(parentContext, planet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, p, id
func (_m *PlanetRepository) Update(ctx context.Context, p *model.Planet, id string) error {
	ret := _m.Called(ctx, p, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Planet, string) error); ok {
		r0 = rf(ctx, p, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
