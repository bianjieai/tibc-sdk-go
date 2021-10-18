format:
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./lite/*/statik.go" -not -path "*.pb.go" | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./lite/*/statik.go" -not -path "*.pb.go" | xargs misspell -w
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./lite/*/statik.go" -not -path "*.pb.go" | xargs goimports -w -local github.com/bianjieai/irita
