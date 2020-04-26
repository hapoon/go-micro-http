.PHONY: run test coverage mockgen

run:
	go run cmd/micro-http/main.go

test:
	go test ./...

coverage:
	go test -cover ./...

mockgen:
	mockgen -source=internal/app/micro-http/service/dummy.go -destination=internal/app/micro-http/mock_service/dummy.go
	mockgen -source=internal/app/micro-http/repository/dummy.go -destination=internal/app/micro-http/mock_repository/dummy.go
