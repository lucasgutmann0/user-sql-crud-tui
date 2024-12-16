# Build the application
build:
	@echo "Building..."
	@go build -o main cmd/api/main.go

# Execute the program
run:
	@go run cmd/app/main.go

# Execute the tests in the project
test:
	@go test ./... -v

# Execute the tests in the project but, using colors to make it look good
testc:
	@set -o pipefail && go test ./... -v -json | tparse -all

# Execute the tests in the project
testc_old:
	@richgo test ./... -v


