### Step-by-Step Usage Guide

#### 1. **Import the `UseServer` class**

To start using the class, import it into your file where you'll make your HTTP requests.

```typescript
import UseServer from '@/services/useServer'
```

#### 2. **Initialize the `UseServer` instance**

You can initialize the `UseServer` class by passing a `baseUrl` and optional configurations like retry count, error URL, and custom error handling functions.

```typescript
const server = new UseServer(
    'https://api.example.com', // The base URL for your API
    { timeout: 5000 }, // Optional default Axios config, e.g., timeout
    3, // Max retry count (optional)
    '/log-error' // Error logging URL (optional)
)
```

#### 3. **Making HTTP Requests**

You can use the provided methods (`get`, `post`, `put`, `delete`) to interact with your API.

##### Example: **GET Request**

You can fetch data using the `get` method by passing the API endpoint and optional query parameters.

```typescript
async function fetchData() {
    try {
        const response = await server.get('/users', { role: 'admin' })
        console.log('Fetched data:', response.data)
    } catch (error) {
        console.error('Error fetching data:', error.message)
    }
}
```

##### Example: **POST Request**

To send data to the server, use the `post` method by passing the endpoint, data, and optional query parameters.

```typescript
async function createUser(userData: { name: string; email: string }) {
    try {
        const response = await server.post('/users', userData)
        console.log('User created:', response.data)
    } catch (error) {
        console.error('Error creating user:', error.message)
    }
}
```

##### Example: **PUT Request**

To update data, use the `put` method by passing the endpoint, data, and optional query parameters.

```typescript
async function updateUser(
    userId: string,
    userData: { name?: string; email?: string }
) {
    try {
        const response = await server.put(`/users/${userId}`, userData)
        console.log('User updated:', response.data)
    } catch (error) {
        console.error('Error updating user:', error.message)
    }
}
```

##### Example: **DELETE Request**

To delete data, use the `delete` method by passing the endpoint and optional query parameters.

```typescript
async function deleteUser(userId: string) {
    try {
        const response = await server.delete(`/users/${userId}`)
        console.log('User deleted:', response.data)
    } catch (error) {
        console.error('Error deleting user:', error.message)
    }
}
```

#### 4. **Customizing Axios Interceptors**

You can add custom request or response interceptors using the provided `addRequestInterceptor` and `addResponseInterceptor` methods.

##### Example: **Request Interceptor**

You can add a request interceptor to modify requests before they are sent.

```typescript
server.addRequestInterceptor((config) => {
    // Add an Authorization header to all requests
    config.headers['Authorization'] = `Bearer ${localStorage.getItem('token')}`
    return config
})
```

##### Example: **Response Interceptor**

You can add a response interceptor to handle specific behaviors based on responses.

```typescript
server.addResponseInterceptor(
    (response) => {
        // Log response status code
        console.log('Response status:', response.status)
        return response
    },
    (error) => {
        // Handle errors globally
        console.error('Error during response:', error.message)
        return Promise.reject(error)
    }
)
```

#### 5. **Handling Errors and Retries**

The `UseServer` class has built-in error logging, retries, and custom error handling. The default behavior retries requests on server errors (`5xx` status codes) up to the configured `maxRetryCount`.

You can customize error logging and retry logic by passing your own `errorLogger`, `retryCondition`, and `errorEnhancer` during instantiation.

##### Example: **Custom Error Logger**

You can log errors in a custom way by defining your own error logger.

```typescript
const customErrorLogger: ErrorLogger = async (errorDetails) => {
    // Send error details to a logging service
    console.error('Logging custom error:', errorDetails)
}

const customServer = new UseServer(
    'https://api.example.com',
    undefined,
    3,
    '/log-error',
    customErrorLogger // Use custom error logger
)
```

##### Example: **Custom Retry Condition**

You can customize when to retry requests by providing a `retryCondition` function.

```typescript
const customRetryCondition: RetryCondition = (error, retryCount) => {
    // Retry on network errors only, and stop after 5 retries
    return retryCount < 5 && !error.response
}

const customServerWithRetry = new UseServer(
    'https://api.example.com',
    undefined,
    5, // Max retries
    '/log-error',
    undefined,
    customRetryCondition // Custom retry condition
)
```

### Example Usage in a Vue Component

Here’s how you could use this in a Vue component (assuming Vue 3 with Composition API):

```typescript
<script setup lang="ts">
import { ref, onMounted } from 'vue';
import UseServer from '@/services/useServer';

const server = new UseServer('https://api.example.com');
const users = ref([]);

async function loadUsers() {
  try {
    const response = await server.get('/users');
    users.value = response.data;
  } catch (error) {
    console.error('Error loading users:', error.message);
  }
}

onMounted(() => {
  loadUsers();
});
</script>

<template>
  <ul>
    <li v-for="user in users" :key="user.id">{{ user.name }}</li>
  </ul>
</template>
```
