set dotenv-load := true

run:
    go run main.go

docker-run:
    docker-compose up --build