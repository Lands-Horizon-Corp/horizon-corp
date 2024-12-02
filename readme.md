# Docker Requirements Documentation

## Setup

- **Initialize**:

  ```sh
  docker compose up --build -d
  ```

  Builds and starts the containers in detached mode.

- **Start**:

  ```sh
  docker compose up -d
  ```

  Starts the containers in detached mode.

- **Stop**:

  ```sh
  docker compose down -v
  ```

  Stops and removes the containers.

- **Reset**:
  ```sh
  docker compose down -v; docker compose up --build -d
  ```
  Stops and removes the containers, then rebuilds and starts them.

## Migration

1. **Seeding & Migration (Initializing Database)**:

   ```sh
   docker compose run --rm flyway migrate
   ```

   Executes migrations to update the database schema.

2. **Baseline**:

   ```sh
   docker compose run --rm flyway baseline -v 1 -d "Initial baseline"
   ```

   Creates a baseline version of the database schema.

3. **Clean**:

   ```sh
   docker compose run --rm flyway clean
   ```

   Drops all objects from the database.

4. **Info**:

   ```sh
   docker compose run --rm flyway info
   ```

   Shows the status of all migrations.

5. **Validate**:

   ```sh
   docker compose run --rm flyway validate
   ```

   Validates the applied migrations against the available ones.

6. **Undo (Teams Edition)**:

   ```sh
   docker compose run --rm flyway undo
   ```

   Undoes the last applied migration. (Available in Flyway Teams Edition)

7. **Repair**:

   ```sh
   docker compose run --rm flyway repair
   ```

   Repairs the metadata table.

8. **Generate**:
   ```sh
   docker compose run --rm flyway generate
   ```
   Generates new migration scripts.

---

## AWS Deployment

---

# Local Setup

- Next 3 : Done
- Tailwind : Done
- Eslint : Done
- Prettier : Done
- Prettier Tailwind : Done
- ShadCN : Done
- Vee validate : Done
- Zod : Done
- Zustand : Done

# Docker

- Setup docker : Done
- Setup db connection : Done
- Setup hot reload for frontend : Done
- hot reload for backend : Done
- setup db migration : Done
- local storage : Done
- emailing: Not Done
- SMS: Not Done

# setup github pipeline workflow

    > create mysql instance
    > create s3 storage instance
    > migrate flyway to mysql aws
    > deploy backend
    > deploy backend

- github workflow : Not Done
- deploy database and server to EC2 : Not Done
- deploy frontend to vercel nuxt : Not Done
- deployment storage : Not Done

# setup toastify

# setup error handling

# setup authentication

```ts
// interfaces.ts
export interface Filter {
  field: string;
  operator: string;
  value: any;
}

export interface Pagination {
  limit: number;
  page: number;
  sortBy?: string;
  sortOrder?: string;
  total?: number;
  totalPages?: number;
  prevPage?: number;
  nextPage?: number;
}

export interface ListRequest {
  filters: Filter[];
  pagination: Pagination;
}

export interface ListResponse<T> {
  data: T[];
  pagination: Pagination;
}
```

```ts
// api.ts
import axios from "axios";
import { ListRequest, ListResponse } from "./interfaces";

export async function listEntities<T>(
  endpoint: string,
  request: ListRequest
): Promise<ListResponse<T>> {
  const response = await axios.post<ListResponse<T>>(endpoint, request);
  return response.data;
}
```

```ts
// UserList.vue (similar to the previous example)
import { listEntities } from "./api";
import { User, Pagination, Filter } from "./interfaces";

const fetchUsers = async () => {
  // Build filters and request
  const request = {
    filters: filters.value,
    pagination: pagination.value,
  };

  const response = await listEntities<User>("/users", request);
  users.value = response.data;
  pagination.value = response.pagination;
};
```

```
can you create me a filtering on my repository. here a sample data i accepts as array
filters: [
  {column: 'name',  "mode": "equal", dataType: 'text', "value": "helo world"},
  {column: 'bday',  "mode": "equal", dataType: 'date',   "from": "2024-11-03T16:00:00.000Z", "to": "2024-11-20T16:00:00.000Z"},
  {column: 'Age',"mode": "gt", dataType: 'number',   "from": "2024-11-03T16:00:00.000Z", "to": "2024-11-20T16:00:00.000Z"},
  {column: 'Gender',"mode": "gt", dataType: 'number',   "from": "2024-11-03T16:00:00.000Z", "to": "2024-11-20T16:00:00.000Z"},
]
preload: ["Gender"]

this are the modes
text: [
        { value: 'equal', label: 'Equals' },
        { value: 'nequal', label: 'Does Not Equal' },
        { value: 'contains', label: 'Contains' },
        { value: 'ncontains', label: 'Does Not Contain' },
        { value: 'startswith', label: 'Starts With' },
        { value: 'endswith', label: 'Ends With' },
        { value: 'isempty', label: 'Is Empty' },
        { value: 'isnotempty', label: 'Is Not Empty' },
    ],
    number: [
        { value: 'equal', label: '(=) Equal' },
        { value: 'nequal', label: 'Not equal' },
        { value: 'gt', label: '(>) Greater Than' },
        { value: 'gte', label: '(>=) Greater Than or Equal' },
        { value: 'lt', label: '(<) Less than' },
        { value: 'lte', label: '(<=) Less than or Equal to' },
        { value: 'range', label: 'Range' },
    ],
    date: [
        { value: 'equal', label: 'On' },
        { value: 'nequal', label: 'Not On' },
        { value: 'before', label: 'Before' },
        { value: 'after', label: 'After' },
        { value: 'gte', label: 'On or After' },
        { value: 'lte', label: 'On or Before' },
        { value: 'range', label: 'Range' },
    ],

this returns something like this
interface Pages {
  page: string // get fetch request to get data
  pageIndex: string
}

interface DatatableDesyo<T>{
  data: T[],
  pageIndex: number
  totalPage: number
  pageSize: number
  totalSize: number // Total size of all data including all pages
  pages: Pages[]
}

const sampleData: DatatableDesyo<{ id: number, name: string }> = {
  data: [
    { id: 1, name: 'John Doe' },
    { id: 2, name: 'Jane Smith' },
    { id: 3, name: 'Tom Brown' },
    { id: 4, name: 'Lisa White' },
    { id: 5, name: 'James Black' }
  ],
  pageIndex: 1,
  totalPage: 5,
  pageSize: 5,
  totalSize: 25,
  pages: [
    { page: '/api/data?pageIndex=1', pageIndex: '1' },
    { page: '/api/data?pageIndex=2', pageIndex: '2' },
    { page: '/api/data?pageIndex=3', pageIndex: '3' },
    { page: '/api/data?pageIndex=4', pageIndex: '4' },
    { page: '/api/data?pageIndex=5', pageIndex: '5' }
  ]
}
```
