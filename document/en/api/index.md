# Horizon Corps REST API

This page introduces the REST API for Horizon Corps and provides an overview of how to interact with it. The Horizon Corps API is powered by Go, ensuring high performance and blazing-fast response times for all backend operations.

## Introduction to REST API

The REST API enables developers to interact programmatically with Horizon Corps systems, performing tasks like user authentication, data retrieval, and management of cooperative-related processes. The API adheres to RESTful principles, ensuring a stateless and standardized interaction.

### Key Features

- **Blazing Fast**: Powered by Go, the API ensures high-performance processing, capable of handling large-scale requests.
- **Scalability**: Designed to grow with your needs, allowing seamless expansion as your cooperative operations increase.
- **JSON Responses**: All data is transmitted in JSON format, making it easy to integrate with various systems and platforms.
- **Secure**: Implements OAuth2 for secure authentication and authorization across all endpoints.

## Example Request

Below is an example of a simple request using the Horizon Corps REST API to fetch user data:

::: code-group

```bash [cURL]
# cURL
curl -X GET "https://api.horizoncorps.com/v1/users" \
-H "Authorization: Bearer YOUR_ACCESS_TOKEN"
```

```js [Javascript]
// Fetch (JavaScript)
fetch("https://api.horizoncorps.com/v1/users", {
  method: "GET",
  headers: {
    Authorization: "Bearer YOUR_ACCESS_TOKEN",
  },
})
  .then((response) => response.json())
  .then((data) => console.log(data))
  .catch((error) => console.error("Error:", error));
```

```py [python 🐍]
# Python (requests library)
import requests

url = "https://api.horizoncorps.com/v1/users"
headers = {
    "Authorization": "Bearer YOUR_ACCESS_TOKEN"
}

response = requests.get(url, headers=headers)
if response.status_code == 200:
    print(response.json())
else:
    print(f"Error: {response.status_code}")
```

:::

### Parameters

| name |   type   | data type | description                   |
| ---- | :------: | --------: | ----------------------------- |
| `id` | required |      uuid | The unique identified of user |

### Responses

| http code |   content-type   |                                  response |
| --------- | :--------------: | ----------------------------------------: |
| `200`     | application/json |                                  UserBase |
| `400`     | application/json | `{"code":"400", "message":"Bad Request"}` |

<details>
<summary>UserBase Type</summary>

| Field                |    Type    |                                              Description |
| -------------------- | :--------: | -------------------------------------------------------: |
| `id`                 |   string   |                           Unique identifier for the user |
| `email`              |   string   |                                     User's email address |
| `username`           |   string   |                             Unique username for the user |
| `firstName`          |   string   |                                        User's first name |
| `middleName`         |  string?   |                            (Optional) User's middle name |
| `lastName`           |   string   |                                         User's last name |
| `contactNumber`      |   string   |                                    User's contact number |
| `permanentAddress`   |   string   |                                 User's permanent address |
| `description`        |   string   |                  A short description or bio for the user |
| `birthDate`          |    Date    |                                     User's date of birth |
| `validEmail`         |  boolean   |                Indicates if the user's email is verified |
| `validContactNumber` |  boolean   |       Indicates if the user's contact number is verified |
| `mediaId`            |  string?   |              (Optional) ID of the user's profile picture |
| `profilePicture`     |   Media?   | (Optional) Media object representing the profile picture |
| `roles`              |   Role[]   |                      Array of roles assigned to the user |
| `status`             | UserStatus |                               Current status of the user |

</details>
