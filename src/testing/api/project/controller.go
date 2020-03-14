// Code generated by mockery v1.0.0. DO NOT EDIT.

package project

import (
	context "context"

	models "github.com/goharbor/harbor/src/common/models"
	mock "github.com/stretchr/testify/mock"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, projectID
func (_m *Controller) Get(ctx context.Context, projectID int64) (*models.Project, error) {
	ret := _m.Called(ctx, projectID)

	var r0 *models.Project
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.Project); ok {
		r0 = rf(ctx, projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: ctx, projectName
func (_m *Controller) GetByName(ctx context.Context, projectName string) (*models.Project, error) {
	ret := _m.Called(ctx, projectName)

	var r0 *models.Project
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Project); ok {
		r0 = rf(ctx, projectName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, projectName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}