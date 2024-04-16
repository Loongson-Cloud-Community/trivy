# update sqlite to support loong64
#go get modernc.org/sqlite@v1.29.7
#go mod tidy
CGO_ENABLED=0 go build ./cmd/trivy
