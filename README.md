# Go Auth Service

This is a simple auth service written in Go. It uses JWT for authentication. It also uses a Postgres database to store the users.

## How to run

### 1. Clone the repository

```bash
git clone https://github.com/asibulhasanshanto/go_auth_service.git
```

### 2. Build and run the service

- Run the following command at first
  ```bash
  docker compose up -d
  ```
- Copy the `/internal/config/demo_api.yaml` file content
- Go to `http://localhost:8585` and create a new file named `go/api` in the consul and paste the content there.
- Open any database client and create a new database named `auth_service`
- copy the content from the file `/internal/store/demo_db.sql` file and run it in the database to create the tables.
- restart the docker containers along with the network
  ```bash
  docker compose down && docker compose up -d
  ```

## Endpoints

| Method | Path                              | Description          | Request Example                                                                      | Response Example                                                                                                                                                                                                                                                                                           |
|--------|-----------------------------------|---------------------|------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| POST   | /api/v1/auth/signup               | Register new user   | `{"email": "user@example.com", "password": "password123", "name": "testuser"}`      | `{"access_token": "jwt.token.here", "refresh_token": "refresh.token.here"}`                                                                                                                                                                                                                                  |
| POST   | /api/v1/auth/login                | Authenticate user   | `{"email": "user@example.com", "password": "password123"}`                          | `{"access_token": "jwt.token.here", "refresh_token": "refresh.token.here"}`                                                                                                                                                                                                                                  |
| GET    | /api/v1/auth/refresh-access-token | Refresh access token| Requires refresh token cookie                                                       | `{"access_token": "jwt.token.here", "refresh_token": "refresh.token.here"}`                                                                                                                                                                                                                                  |
| POST   | /api/v1/auth/logout               | Revoke tokens      | Requires access token                                                               | `{"message": "Successfully logged out"}`                                                                                                                                                                                                                                                                    |
| GET    | /api/v1/auth/me                   | Get current user   | Requires access token                                                               | `{"user": {"ID": 1, "Email": "user@example.com", "Name": "testuser", "Role": "user", "CreatedAt": "2025-03-29T10:05:29.376717Z", "UpdatedAt": "2025-03-29T10:05:29.376717Z", "DeletedAt": null, "PasswordChangedAt": null}}`                                                                                |
