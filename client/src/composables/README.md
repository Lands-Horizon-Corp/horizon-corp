# User Server

-   The `baseURL` defaults to `import.meta.env.VITE_CLIENT_SERVER_URL`.
-   The `maxRetryCount` defaults to `3`.
-   The `errorUrl` defaults to `'/error-handler'`.
-   The `errorLogger`, `retryCondition`, and `errorEnhancer` will use their default implementations.

```ts
const server = new UseServer()
```

This will create an instance using the default configurations. You can then use it to make API requests, as shown in the previous examples:

```ts
server
    .get('/api/users')
    .then((response) => {
        console.log('User data:', response.data)
    })
    .catch((error) => {
        console.error('Error fetching users:', error)
    })
```

If you want to customize any part of the configuration (e.g., different retry logic or error handling), you can pass the necessary arguments when instantiating the class.
