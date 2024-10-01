main_package_path = cmd/api/
.DEFAULT_GOAL=run
.PHONY: run
run:
	go run ${main_package_path}main.go

.PHONY: build
build:
	go build -o cmd/build ${main_package_path}main.go

.PHONY: swag
swag:
	swag init -g ${main_package_path}main.go -o ${main_package_path}docs