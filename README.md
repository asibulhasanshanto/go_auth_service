# Go Auth Service
This is a simple auth service written in Go. It uses JWT for authentication. It also uses a MySQL database to store the users.

## How to run
### 1. Clone the repository
   ```bash
   git clone https://github.com/asibulhasanshanto/go_auth_service.git
   ```
### 2. Build and run the service
   -  Run the following command at first
      ```bash
      docker compose up -d
      ```
   -  Copy the `/internal/config/demo_api.yaml` file content
   - Go to `http://localhost:8585` and create a new file named `go/api` in the consul and paste the content there.
   - Open any database client and create a new database named `auth_service`
   - copy the content from the file `/internal/store/demo_db.sql` file and run it in the database to create the tables.
   - restart the docker containers along with the network
       ```bash
       docker compose down && docker compose up -d 
       ```