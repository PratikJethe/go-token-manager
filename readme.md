# Go Token Manager

## Description
Go Token Manager is a service designed to manage a pool of unique tokens. It allows for the generation, assignment, unblocking, and deletion of tokens, with a mechanism to automatically release tokens after a specified time.

This project aims to provide a robust solution for token management, suitable for use cases where a scalable and reliable service is required.

## Features
- Manage a pool of unique tokens.
- Randomly assign and release tokens.
- Support for automatic expiration of tokens.
- Designed for scalability and efficiency.
----
## How to use
- Clone this repository.
```sh
git clone https://github.com/PratikJethe/go-token-manager
```
----

## How to setup
- To set up docker containers of postgres and token_service from docker-compose file, run.

```sh
docker compose up -d
```
- To take down docker containers, run.
```sh
docker compose down
```

## Enviroment Variables
- create your own `.env` file in the root of the project.

| ENV_VARIABLE | Description |
| :-------- | :------------------------- |
| `DB_HOST` | Host for database connection 
| `DB_USER` | Database user
| `DB_PASSWORD` | Database password
| `DB_PORT` | Database port
| `DB_NAME` | Database name
| `SERVER_PORT` | Server port
| `SERVER_HOST` | Server Host
| `DB_MIGRATION_FILE` | DB migration file path
| `TOKEN_LENGTH` | Length of the token to be generated in bytes
| `TOKEN_EXPIRATION_DURATION` | Token expiration duration in seconds
| `TOKEN_ACTIVE_DURATION` | Token active duration in seconds


----
## Migrations
- Migration scripts are in the `db/migrations` folder.
- Migrations are run using [golang-migrate](https://pkg.go.dev/github.com/golang-migrate/migrate/v4) package.
- Migrations run during the initialisation of server.
```sh
# command to add a migration file
migrate create -ext sql -dir <directory_path> -seq <migration_name>
```
----

## API Reference
#### Postman Setup
- [Postman](https://www.postman.com/) is an API platform for building and using APIs.
- Postman collection is in  `Idfy.postman_collection.json` file in the root of the project.
- Open postman and import this collection. After importing you can see all the requests under `Idfy` collection.


#### Create

```http
  POST /tokens/create
```

#### response
```
{
    "id": 1,
    "token": "9666026ea9edc1c32103ac6e4b300e62",
    "last_activation_time": null,
    "is_deleted": false,
    "created_at": "2024-10-20T13:51:49.782294Z",
    "updated_at": "2024-10-20T13:51:49.782294Z"
}
```
---

### Assign
```http
  GET /tokens/assign
```

#### response
```
{
    "token": "9666026ea9edc1c32103ac6e4b300e62"
}
```
---

### Delete
```http
  DELETE /tokens/delete?token={token_id}
```

#### response
```
{
    "message": "Token deleted successfully"
}
```


---
### Unblock
```http
  POST /tokens/unblock
```
| Body Parameters | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `token` | `string` | **Required**. |

#### response
```
{
    "message": "Token unblocked successfully"
}
```

---
### Keep Alive
```http
  POST /tokens/keep-alive
```
| Body Parameters | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `token` | `string` | **Required**. |

#### response
```
{
    "message": "Token kept alive"
}
```

---
