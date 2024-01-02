TARGET=main
PACKAGE_PATH=.

.PHONY: build
build:
	go build -o ${TARGET} ${PACKAGE_PATH}

.PHONY: run
run: build
	./${TARGET}

.PHONY: clean
clean:
	go clean

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	golangci-lint run --enable-all

.PHONE: deps
deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2
	go install github.com/cosmtrek/air@latest

.PHONY: up
up: build
	air --build.cmd "make build" --build.bin "./${TARGET}" --build.delay "100" \
	--build.exclude_dir "" \
	--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
	--misc.clean_on_exit "true"

.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -race -buildvcs -vet=off ./...