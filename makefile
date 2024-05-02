

build: ## 🛠️  Build Project
	@go build -o ./bin/convertor

run: build ## 🏃 Build and Run the project
	@go build -o ./bin/convertor && ./bin/convertor

test: ## 🧪 Test the project
	@go test -v ./...
