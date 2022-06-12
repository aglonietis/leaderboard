services:
1. API server
2. database
3. cache server? (most likely outside of scope)

functionality:
1. authentication (pure jwt?)
2. store player score - name + score
   2.1. Update player score if exists and larger than previous
3. get leaderboard - paginated
   3.1. all time leaderboard
   3.2. monthly leaderboard
   3.3. player leaderboard and those around him
4. seeders


# How to start project

### Start needed containers - database
```
docker compose up
```

### Create an environment variable file
```
cp .env.example .env
```

### Start the server
```
make start
```

# Interesting things

### Possibilities to improve leaderboard rank check time

Relational database provides N*log2(N) speed, while Article: https://www.hindawi.com/journals/ijcgt/2018/3234873/
provides info that it is possible to reduce it to log2(N)
