

# Docker Requirements Documentation

## Setup

- **Initialize**: 
  ```sh
  docker-compose up --build -d
  ```
  Builds and starts the containers in detached mode.

- **Start**: 
  ```sh
  docker-compose up -d
  ```
  Starts the containers in detached mode.

- **Stop**: 
  ```sh
  docker-compose down
  ```
  Stops and removes the containers.

- **Reset**: 
  ```sh
  docker-compose down; docker-compose up --build -d
  ```
  Stops and removes the containers, then rebuilds and starts them.

## Migration

1. **Seeding & Migration (Initializing Database)**: 
   ```sh
   docker-compose run --rm flyway migrate
   ```
   Executes migrations to update the database schema.

2. **Baseline**: 
   ```sh
   docker-compose run --rm flyway baseline -v 1 -d "Initial baseline"
   ```
   Creates a baseline version of the database schema.

3. **Clean**: 
   ```sh
   docker-compose run --rm flyway clean
   ```
   Drops all objects from the database.

4. **Info**: 
   ```sh
   docker-compose run --rm flyway info
   ```
   Shows the status of all migrations.

5. **Validate**: 
   ```sh
   docker-compose run --rm flyway validate
   ```
   Validates the applied migrations against the available ones.

6. **Undo (Teams Edition)**: 
   ```sh
   docker-compose run --rm flyway undo
   ```
   Undoes the last applied migration. (Available in Flyway Teams Edition)

7. **Repair**: 
   ```sh
   docker-compose run --rm flyway repair
   ```
   Repairs the metadata table.

8. **Generate**: 
   ```sh
   docker-compose run --rm flyway generate
   ```
   Generates new migration scripts.

---

# Local Setup
  - Nuxt 3 : Done
  - Tailwind : Done
  - Eslint : Done
  - Prettier : Done
  - Prettier Tailwind : Done
  - ShadCN : Done
  - Vee validate : Done
  - Zod : Done
  - Pinia store : Done

# Docker 
  - Setup docker : Done
  - Setup db connection : Done
  - Setup hot reload for frontend : Done
  - hot reload for backend : Done
  - setup db migration : Done 


# setup github pipeline workflow
  - deploy database and server to EC2 : Not Done
  - deploy frontend to vercel nuxt : Not Done

# setup toastify
# setup error handling
# setup authentication
