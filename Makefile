ARTIFUCT := ./http_server
GO := `which go`
GOSRC := \
	./http_server_wrapper.go \
	./main.go \
	./request_handler.go \

build_and_run: $(ARTIFUCT)
	$(ARTIFUCT)

$(ARTIFUCT): $(GOSRC)
	$(GO) build

test: $(ARTIFUCT)
	$(GO) test ./...

clean:
	rm -f $(ARTIFUCT)
