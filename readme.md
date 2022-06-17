# Leaderboard

Test project by Valdis Aglonietis

# How to set up project

### Run commands in the console in project root 

The first commands will download dependencies, copy .env.example to .env , start docker servers and 
the second command will run migrations, run seeds and start the API server

```
make setup
make setup_server
```

# Known Problems:

1. Validation is show as string, not properly as JSON
2. No Tests
3. Did not specify before, but have ignored the following about now showing if already in previous pages:
   f the name of the player is passed and the player is not in this list of results *(and their result is not in any of the previous pages)*, a list of players around the current player should be returned.

# How to test

By default application can be tested at localhost:8080 . Port can be changed with API_PORT in .env file

There are 4 endpoints available:

### Home - GET /

Can be checked to see if API is working properly

### Login - POST /api/v1/login

Can be used to get token by logging to the API server. Default credentials are provided in the example:

```json
{
   "username": "leader",
   "password": "leader"
}
```

### Score - POST /api/v1/scores

Can be used to store a score. Score Request body:

```json
{
    "name": "greatestPlayerOfAllTime",
    "score":999
}
```

### Leaderboard - GET /api/v1/leaderboard/:type?page=<page>&name=<name>

Can be used to request Leaderboard Pages. There are different parameters options:

type:
1. monthly (/api/v1/leaderboard/monthly) - returns monthly leaderboard
2. empty (/api/v1/leaderboard) - retursn all time leaderboard

page - pagination for data

name - filter by username

# Notes

Golang version: 1.18

Servers: PostgreSQL, API

There is Postman Collection: "LeaderboardAPI.postman_collection.json" and Environment example: "Leaderboard-dev.postman_environment.json"

There are additional commands:

* make develop - starts in live-reload mode
* make migrate - executes migrations
* make seed - executes seeds . It has a parameter 3000 and that indicates the amount of scores generated
* make build_migrations - builds migration executable
* make build_seeds - builds seed executable
* make api - build API server executable
* make run - runs the API server
