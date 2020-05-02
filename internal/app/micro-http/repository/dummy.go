package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"go-micro-http/internal/app/micro-http/entity"
)

var (
	apiTimeout time.Duration = 5 * time.Second
)

// DummyEmployee repository interface
type DummyEmployee interface {
	List() (employees entity.DummyEmployees, err error)
	Find(id string) (employee *entity.DummyEmployee, err error)
}

type dummyEmployee struct {
}

func (d *dummyEmployee) List() (employees entity.DummyEmployees, err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), apiTimeout)
	defer cancelFunc()

	getEmployeesRes := struct {
		Status string                `json:"status"`
		Data   entity.DummyEmployees `json:"data"`
	}{}

	if err := getJSONDecodedBodyByHTTPWithContext(
		ctx,
		http.MethodGet,
		"http://dummy.restapiexample.com/api/v1/employees",
		nil,
		&getEmployeesRes,
	); err != nil {
		return nil, err
	}

	if getEmployeesRes.Status != "success" {
		return nil, fmt.Errorf("status is not success: %s", getEmployeesRes.Status)
	}

	return getEmployeesRes.Data, nil
}

func (d *dummyEmployee) Find(id string) (employee *entity.DummyEmployee, err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), apiTimeout)
	defer cancelFunc()

	getEmployeeRes := struct {
		Status string               `json:"status"`
		Data   entity.DummyEmployee `json:"data"`
	}{}

	url := fmt.Sprintf("http://dummy.restapiexample.com/api/v1/employee/%s", id)

	if err := getJSONDecodedBodyByHTTPWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
		&getEmployeeRes,
	); err != nil {
		return nil, err
	}

	if getEmployeeRes.Status != "success" {
		return nil, fmt.Errorf("status is not success: %s", getEmployeeRes.Status)
	}

	return &getEmployeeRes.Data, nil
}

// NewDummyEmployee is contructer of DummyEmployee interface.
func NewDummyEmployee() DummyEmployee {
	return &dummyEmployee{}
}

func getJSONDecodedBodyByHTTPWithContext(
	ctx context.Context,
	method string,
	url string,
	body io.Reader,
	resBody interface{},
) error {
	client := http.DefaultClient

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(resBody); err != nil {
		return err
	}

	return nil
}
