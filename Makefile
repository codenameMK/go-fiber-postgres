build:
	go build -o bin/go-fiber-postgres

run: build
	./bin/go-fiber-postgres
test:
	go test -v ./..
	