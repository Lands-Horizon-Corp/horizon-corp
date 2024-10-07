# UseServer Documentation

## Overview

The `UseServer` class provides a customizable interface for handling HTTP requests and errors using Axios. It supports configurable options for base URL, error logging, request retries, and error enhancement.

## Installation

Ensure you have Axios installed in your project. You can add it via npm:

```bash
npm install axios
```

## TypeScript Definitions

The class makes use of several TypeScript types for improved type safety. Here are the main types you'll interact with:

-   `CustomAxiosRequestConfig`: Custom configuration for Axios requests.
-   `RequestParams`: Type for request query parameters.
-   `ErrorDetails`: Type for structured error information.

## Usage

### Importing UseServer

To use the `UseServer` class in your TypeScript file, import it as follows:

```typescript
import UseServer from '@/path/to/UseServer'
```

### Creating an Instance

You can create an instance of the `UseServer` class with optional configurations:

```typescript
const server = new UseServer(
    'http://localhost:8080/api/v1', // Base URL (optional)
    { timeout: 10000 }, // Default Axios configuration (optional)
    3, // Maximum retry count (optional)
    '/error-handler', // Error logging URL (optional)
    myCustomErrorLogger, // Custom error logger function (optional)
    myCustomRetryCondition, // Custom retry condition (optional)
    myCustomErrorEnhancer // Custom error enhancer function (optional)
)
```

### Making Requests

You can make GET, POST, PUT, and DELETE requests using the methods provided by the `UseServer` instance.

#### GET Request

```typescript
const response = await server.get('/endpoint', { param1: 'value1' })
console.log(response.data)
```

#### POST Request

```typescript
const data = { key: 'value' }
const response = await server.post('/endpoint', data, { param1: 'value1' })
console.log(response.data)
```

#### PUT Request

```typescript
const updateData = { key: 'newValue' }
const response = await server.put('/endpoint', updateData, { param1: 'value1' })
console.log(response.data)
```

#### DELETE Request

```typescript
const response = await server.delete('/endpoint', { param1: 'value1' })
console.log(response.data)
```

### Customizing Error Handling

#### Error Logger

You can provide a custom error logger function that logs errors in your preferred manner:

```typescript
const myCustomErrorLogger = async (errorDetails: ErrorDetails) => {
    console.error('Logging error:', errorDetails)
    // Additional logging logic here
}

const server = new UseServer(
    'http://localhost:8080/api/v1',
    undefined,
    undefined,
    undefined,
    myCustomErrorLogger
)
```

#### Retry Condition

You can define a custom retry condition to control when requests should be retried:

```typescript
const myCustomRetryCondition = (error: AxiosError, retryCount: number) => {
    return retryCount < 3 && error.response?.status === 503 // Retry on service unavailable
}

const server = new UseServer(
    'http://localhost:8080/api/v1',
    undefined,
    undefined,
    undefined,
    undefined,
    myCustomRetryCondition
)
```

#### Error Enhancer

You can also provide a custom error enhancer to modify error messages before they are rejected:

```typescript
const myCustomErrorEnhancer = (error: AxiosError) => {
    error.message = 'Custom error message: ' + error.message
    return error
}

const server = new UseServer(
    'http://localhost:8080/api/v1',
    undefined,
    undefined,
    undefined,
    undefined,
    undefined,
    myCustomErrorEnhancer
)
```

### Adding Interceptors

You can add request and response interceptors for more control over the HTTP request/response cycle.

#### Adding a Request Interceptor

```typescript
server.addRequestInterceptor(
    (config) => {
        // Modify request config here
        console.log('Request Interceptor:', config)
        return config
    },
    (error) => {
        console.error('Request Interceptor Error:', error)
        return Promise.reject(error)
    }
)
```

#### Adding a Response Interceptor

```typescript
server.addResponseInterceptor(
    (response) => {
        console.log('Response Interceptor:', response)
        return response
    },
    (error) => {
        console.error('Response Interceptor Error:', error)
        return Promise.reject(error)
    }
)
```

### Retrieving the Internal Axios Instance

If you need direct access to the Axios instance, you can retrieve it using:

```typescript
const axiosInstance = server.getHttpClient()
```

## Error Handling

The `UseServer` class handles errors gracefully by retrying failed requests based on the defined retry condition, logging errors using the specified logger, and enhancing errors as needed. The default implementations can be overridden by providing custom functions.

## Conclusion

The `UseServer` class is a powerful tool for managing HTTP requests and errors in a TypeScript Node.js application. By utilizing its customizable features, you can create a robust and resilient HTTP client tailored to your specific needs.
