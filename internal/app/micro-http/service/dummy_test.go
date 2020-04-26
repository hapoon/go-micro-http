package service

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-micro-http/internal/app/micro-http/entity"
	"go-micro-http/internal/app/micro-http/mock_repository"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestDummyAccesserGetEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()

	mock := mock_repository.NewMockDummyEmployee(ctrl)
	da := NewDummyAccesser(mock)

	scenario := "Getting employee is failed"
	{
		req := httptest.NewRequest(http.MethodGet, "/employee/1", nil)
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")
		mock.EXPECT().Find("1").Return(&entity.DummyEmployee{}, fmt.Errorf("ID: 1 is not exist"))

		actual := da.GetEmployee(ctx)
		assert.Error(t, actual, scenario)
		assert.Equal(t, http.StatusInternalServerError, res.Code, scenario)
	}

	scenario = "Getting employee is succeeded"
	{
		req := httptest.NewRequest(http.MethodGet, "/employee/2", nil)
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)
		ctx.SetParamNames("id")
		ctx.SetParamValues("2")
		mock.EXPECT().Find("2").Return(&entity.DummyEmployee{}, nil)

		actual := da.GetEmployee(ctx)
		assert.NoError(t, actual, scenario)
		assert.Equal(t, http.StatusOK, res.Code, scenario)
	}
}

func TestDummyAccesserGetEmployees(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()

	mock := mock_repository.NewMockDummyEmployee(ctrl)
	da := NewDummyAccesser(mock)

	scenario := "Getting employees is failed"
	{
		req := httptest.NewRequest(http.MethodGet, "/employees", nil)
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)

		mock.EXPECT().List().Return(nil, fmt.Errorf("employees are not exist"))

		actual := da.GetEmployees(ctx)
		assert.Error(t, actual, scenario)
		assert.Equal(t, http.StatusInternalServerError, res.Code, scenario)
	}

	scenario = "Getting employees is succeeded"
	{
		req := httptest.NewRequest(http.MethodGet, "/employees", nil)
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)

		mock.EXPECT().List().Return(entity.DummyEmployees{}, nil)

		actual := da.GetEmployees(ctx)
		assert.NoError(t, actual, scenario)
		assert.Equal(t, http.StatusOK, res.Code, scenario)
	}
}
