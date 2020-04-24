package service

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	url        string        = "http://dummy.restapiexample.com/api/v1/employees"
	apiTimeout time.Duration = 5 * time.Second
)

// DummyAccesser is interface of dummyAccesser.
type DummyAccesser interface {
	GetEmployees(ctx echo.Context) error
}

type dummyAccesser struct {
}

func (d *dummyAccesser) GetEmployees(c echo.Context) error {
	client := http.DefaultClient

	ctx, cancelFunc := context.WithTimeout(context.Background(), apiTimeout)
	defer cancelFunc()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	ch := make(chan struct{})

	go func() {
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			ch <- struct{}{}
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			c.Logger().Error(err)
			c.NoContent(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, string(body))

		ch <- struct{}{}
	}()

	<-ch
	return nil
}

// NewDummyAccesser is constructer of DummyAccesser interface.
func NewDummyAccesser() DummyAccesser {
	return &dummyAccesser{}
}

// UseDummyRouting is routing for dummy.
func UseDummyRouting(e *echo.Echo) {
	da := NewDummyAccesser()
	g := e.Group("/employee")
	{
		g.GET("", da.GetEmployees)
	}
}
