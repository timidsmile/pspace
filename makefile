format:
	find . -name "*.go" ! -path "./vendor*" | xargs -I{} go fmt {}
build:
	go build -o pspace main.go
watch:
	@if [ -f $(GOPATH)/bin/fswatch ]; \
			then \
			echo "fswatch already installed"; \
		else \
			echo "install fswatch ..." && \
			go get -u -v github.com/codeskyblue/fswatch; \
		fi
		$(GOPATH)/bin/fswatch
doc:
	open http://localhost:8650/pkg/github.com/timidsmile/pspace/
	godoc -http=:8650

