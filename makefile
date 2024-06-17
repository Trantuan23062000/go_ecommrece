# Tạo ra binary từ source code Go
build:
	@go build -o bin/go_api_ecommerce cmd/main.go

# Chạy tất cả các bài test trong dự án
test:
	@go test -v ./...

# Chạy server sau khi build
run: build
	@./bin/go_api_ecommerce
