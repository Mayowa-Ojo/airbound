# commands

.PHONY: run run-migrate-down build clean run-build ent-gen ent-init ent-describe

PARSE_ARGS = $(filter-out $@,$(MAKECMDGOALS))

run:
	go run main.go

run-migrate-down:
	go run main.go -migrate-down

build: clean
	go build -o dist/airbound

run-build: build
	./dist/airbound

clean:
	rm -rf dist

ent-init:
	go run entgo.io/ent/cmd/ent init --target internal/ent/schema $(PARSE_ARGS)

ent-gen:
	go generate ./internal/ent

ent-describe:
	go run entgo.io/ent/cmd/ent describe ./internal/ent/schema

log:
	@echo $(PARSE_ARGS)

%:
	@: