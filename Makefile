default: all

all: fmt test build

fmt:
	$(info ******************** checking formatting ********************)
	@test -z $(shell gofmt -l $(SRC)) || (gofmt -d $(SRC); exit 1)

lint:
	$(info ******************** running lint tools ********************)
	golangci-lint run -v

test: 
	$(info ******************** running tests ********************)
	go test -v internal/...

generate:
	@echo "******************** generating ent schema ********************"
	go mod tidy
	go generate ./...
	go mod tidy
	@echo "******************** generating gqlgen ********************"
	go run github.com/99designs/gqlgen generate --verbose
	go mod tidy
	go run ./gen_schema.go
	@echo "******************* generating gqlgen client ********************)"
	go run github.com/Yamashou/gqlgenc generate --configdir schema

clean:
	$(info ******************** removing generated files from repo ********************)
	rm -rf internal/ent/generated/^doc.go
	rm -rf internal/api/^doc.go
	rm -rf internal/ent/schema/*.go
	rm -f schema/ent.graphql
	rm -f schema.graphql
	rm -rf internal/testclient/