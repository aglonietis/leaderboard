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
	cp .env.example .env
	docker compose up -d;
	make migrate
	make seed
	make run

# Refreshes the environment
refresh:
	docker compose down -v;
	make setup

# Builds migrations executable
build_migrations:
	go build -o bin/migrate cmd/migrations/main.go

# Builds seeds executable
build_seeds:
	go build -o bin/seed cmd/seeds/main.go

# Builds API server executable
build_api:
	go build -o bin/api cmd/api/main.go

# Starts the API server
run:
	./bin/api


