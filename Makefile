GOTEST=grc go test -count=1
GORUN=go run
GOBUILD=go build
GOFMT=go fmt

format:
	$(GOFMT) ./...

lint:
	$(LINT) ./...

run-watch:
	find . -name '*.go' | entr -r $(GORUN) cmd/server/main.go

run:
	$(GORUN) cmd/server/main.go

pre-build: clean
	mkdir -p deploy

build: pre-build
	$(GOBUILD) -o deploy/api cmd/server/main.go

clean:
	rm -rf deploy

test:
	$(GOTEST) -v -race -cover ./...

to_pr_final_run: format
	$(LINT)
	$(GOTEST) ./...