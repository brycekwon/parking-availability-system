GO_FLAGS = CGO_ENABLED=1
GO_START = ./cmd/website/main.go
GO_TARGET = ./main
HUGO_FLAGS = --gc --cleanDestinationDir


build: clean ui server

run: build
	$(GO_TARGET)

ui:
	hugo build $(HUGO_FLAGS) --templateMetrics --templateMetricsHints

dev-ui:
	hugo server $(HUGO_FLAGS) --disableFastRender --watch

server:
	$(GO_FLAGS) go build -o $(GO_TARGET) $(GO_START)

dev-server:
	$(GO_FLAGS) go run $(GO_START)

clean:
	rm -rf public/ resources/ .hugo_build.lock hugo_stats.json $(GO_TARGET)

clean-all: clean
	rm -rf *.db