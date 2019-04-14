RELEASE_DIR=releases
.PHONY $(RELEASE_DIR)

deps:
	@echo "Downloading deps..."
	@go mod download
test: deps
	go test -v ./...
build: deps
	gox -output="$(RELEASE_DIR)/{{.Dir}}_{{.OS}}_{{.Arch}}"
$(RELEASE_DIR):
	@mkdir -p $(RELEASE_DIR)