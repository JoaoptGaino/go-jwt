build:
	@go build -o bin/go-jwt
dev: build
	/home/joaoptgaino/go/bin/CompileDaemon --command="./bin/go-jwt"