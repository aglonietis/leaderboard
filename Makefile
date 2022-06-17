# Starts project in hot-reload mode
develop:
	air

# Runs migrations
migrate:
	./bin/migrate

# Seeds database
seed:
	./bin/seed 3000

# Sets the environment up for development or testing
setup:
	go mod download
	cp .env.example .env
	docker compose down -v;
	docker compose up -d

# Runs database relating actions and starts the server
setup_server:
	make migrate
	make seed
	make run

# Builds migrations executable
build_migrations:
	go build -o bin/migrate cmd/migrations/main.go

# Builds seeds executable
build_seeds:
	go build -o bin/seed cmd/seeds/main.go

# Builds API server executable
build_api:
	go build -o main cmd/api/main.go

# Starts the API server
run:
	./bin/api


