package service

import (
	"go-micro-http/internal/app/micro-http/repository"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	apiTimeout time.Duration = 5 * time.Second
)

// DummyAccesser is interface of dummyAccesser.
type DummyAccesser interface {
	GetEmployee(ctx echo.Context) error
	GetEmployees(ctx echo.Context) error
}

type dummyAccesser struct {
	dummyRepository repository.DummyEmployee
}

func (d *dummyAccesser) GetEmployee(c echo.Context) error {
	id := c.Param("id")

	employee, err := d.dummyRepository.Find(id)
	if err != nil {
		c.Logger().Error(err)
		c.NoContent(http.StatusInternalServerError)
		return err
	}

	c.JSON(http.StatusOK, employee)

	return nil
}

func (d *dummyAccesser) GetEmployees(c echo.Context) error {
	employees, err := d.dummyRepository.List()
	if err != nil {
		c.Logger().Error(err)
		c.NoContent(http.StatusInternalServerError)
		return err
	}

	c.JSON(http.StatusOK, employees)

	return nil
}

// NewDummyAccesser is constructer of DummyAccesser interface.
func NewDummyAccesser(repo repository.DummyEmployee) DummyAccesser {
	return &dummyAccesser{
		dummyRepository: repo,
	}
}

// UseDummyRouting is routing for dummy.
func UseDummyRouting(e *echo.Echo) {
	da := NewDummyAccesser(repository.NewDummyEmployee())

	g := e.Group("/employee")
	{
		g.GET("", da.GetEmployees)
		g.GET("/:id", da.GetEmployee)
	}
}
