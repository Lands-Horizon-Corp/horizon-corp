# UseServer

`UseServer` is a utility class that simplifies handling HTTP requests using Axios. It provides features like error handling, retries, logging, and customizable configurations. This class is designed to work seamlessly with various frameworks such as Vite, React, Nuxt, Next.js, and Node.js Express.

## Features

-   **HTTP Methods**: Supports `GET`, `POST`, `PUT`, and `DELETE` requests.
-   **Automatic Retries**: Retries failed requests based on customizable conditions.
-   **Error Logging**: Logs errors to a specified endpoint.
-   **Customizable Headers**: Automatically includes an API key in the request headers.
-   **Request/Response Interceptors**: Allows adding custom logic for requests and responses.

## Installation

coming soon in npm, pnpm, and bun

Here's a sample README for your `UseServer` class implementation, including usage instructions and details about the environment variable configuration.

## Usage

### Basic Setup

To set up an instance of `UseServer`, import the class and create an instance:

```typescript
import UseServer from '@/server/UseServer'

const server = new UseServer()
```

### Environment Variables

The `UseServer` class automatically retrieves the `HORIZON_CORP_API_KEY` from your environment variables. Depending on your framework, set the environment variable as follows:

-   **Vite**: In your `.env` file:

    ```plaintext
    VITE_HORIZON_CORP_API_KEY=your_api_key_here
    ```

-   **Node.js Express**: In your `.env` file:

    ```plaintext
    HORIZON_CORP_API_KEY=your_api_key_here
    ```
