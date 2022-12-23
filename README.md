# SSI Web Events

This is an incredibly lightweight and impressive plugin that allows you to monitor user's javascript activities on a website.

## Features

- Track user's any javascript events on a website
- Comes with a server that you can quickly deploy with Docker
- Keeps data in MongoDB. In this way, data processing is faster.
- Allows you to store logged in user information.
- You can store specific fields or global variables.
- IP address information is stored to prevent spam data.

##  How It Works

This plugin requires a server to work. But you don't have to worry. Just say docker run!

Check out the examples to run the server with docker or docker-compose [here](https://github.com/ssibrahimbas/ssi-we/tree/main/examples)

Then you add javascript to your html page like adding bootstrap, a little configuration and boom. It's ready now.

## Setup Server

The server application was developed with golang and published to docker. The Docker image is approximately 25 mb.

View in Docker Hub: [https://hub.docker.com/r/ssibrahimbas/ssi-we](https://hub.docker.com/r/ssibrahimbas/ssi-we)

Check out the examples to run the server with docker or docker-compose [here](https://github.com/ssibrahimbas/ssi-we/tree/main/examples)

##  Setup Client

You can add the client to your html page with the following code.

```html
<script type="text/javascript">
  window.ssiWe = {
    url: "localhost:8080",
    page: "index.html",
    debug: false,
    user: {
      id: "123",
      name: "John Doe",
      email: "info@john.com",
    },
  };
</script>
<script
  src="https://cdn.jsdelivr.net/gh/ssibrahimbas/ssi-we@main/dist/ssi-we.min.js"
  type="text/javascript"
  defer
  async
></script>
```

Note:
Configuration must be before cdn. And be careful not to import before the DOM is loaded. Because this plugin uses DOM.

##  Configuration

### Server Configuration

| Name                   | Type   | Description             | Default | Required |
| ---------------------- | ------ | ----------------------- | ------- | -------- |
|  HOST                  | string | Server host             | x       | true     |
| PORT                   | string | Server port             | x       | true     |
| MONGO_HOST             | string | MongoDB host            | x       | true     |
| MONGO_PORT             | string | MongoDB port            | x       | true     |
| MONGO_USER             | string | MongoDB user            | x       | true     |
| MONGO_PASSWORD         | string | MongoDB password        | x       | true     |
| MONGO_DB               | string | MongoDB database name   | x       | true     |
| MONGO_COLLECTION       | string | MongoDB collection name | x       | true     |
| CORS_ALLOW_ORIGINS     | string | CORS allow origins      | x       | true     |
| CORS_ALLOW_METHODS     | string | CORS allow methods      | x       | true     |
| CORS_ALLOW_HEADERS     | string | CORS allow headers      | x       | true     |
| CORS_ALLOW_CREDENTIALS | string | CORS allow credentials  | x       | true     |

### Client Configuration

| Name              | Type    | Description                     | Default | Required |
| ----------------- | ------- | ------------------------------- | ------- | -------- |
| url               | string  | Server url                      | x       | true     |
| page              | string  | Page name                       | x       | true     |
| debug             | boolean | Debug mode                      | false   | false    |
| user              | any     | User information                | null    | false    |
| params            | any     | Any params                      | null    | false    |
| reconnectDelay    | number  | Reconnect delay in milliseconds | 1000    | false    |
| reconnectAttempts | number  | Reconnect attempts              | 5       | false    |
| reconnect         | boolean | Reconnect on disconnect         | true    | false    |
| openEvent         | boolean | Send open event on connect      | false   | false    |

### Element Configuration

| Name              | Type   | Description                      | Default | Required |
| ----------------- | ------ | -------------------------------- | ------- | -------- |
| data-we-event     | string | JavaScript event                 | click   | false    |
| data-we-name      | string | Name of the action, ex: Register | x       | true     |
| data-we-params-\* | any    | Any params                       | null    | false    |

## Examples

###  Basic Use Case

```html
<button type="button" data-we-name="register">Register</button>
```

Saved as:

```json
{
  "id": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
  "name": "register",
  "page": "index.html",
  "type": "click",
  "element": "button",
  "params": {},
  "user": {},
  "ip": "xxx.xxx.xxx.xxx",
  "createdAt": "2021-01-01T00:00:00Z"
}
```

###   With Params

```html
<button type="button" data-we-name="register" data-we-params-id="123">
  Register
</button>
```

Saved as:

```json
{
  "id": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
  "name": "register",
  "page": "index.html",
  "type": "click",
  "element": "button",
  "params": {
    "id": "123"
  },
  "user": {},
  "ip": "xxx.xxx.xxx.xxx",
  "createdAt": "2021-01-01T00:00:00Z"
}
```

###  With User

```html
<button type="button" data-we-name="buyProduct" data-we-params-id="123">
  Buy
</button>

<script type="text/javascript">
  window.ssiWe = {
    url: "localhost:8080",
    page: "index.html",
    debug: false,
    user: {
      id: "123",
      name: "John Doe",
      email: "info@john.com",
    },
  };
</script>
<script
  src="https://cdn.jsdelivr.net/gh/ssibrahimbas/ssi-we@main/dist/ssi-we.min.js"
  type="text/javascript"
  defer
  async
></script>
```

Saved as:

```json
{
  "id": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
  "name": "buyProduct",
  "page": "index.html",
  "type": "click",
  "element": "button",
  "params": {
    "id": "123"
  },
  "user": {
    "id": "123",
    "name": "John Doe",
    "email": "info@john.com"
  },
  "ip": "xxx.xxx.xxx.xxx",
  "createdAt": "2021-01-01T00:00:00Z"
}
```

### Other Events

```html
<button type="button" data-we-name="register" data-we-event="mouseover">
  Register
</button>
```

Saved as:

```json
{
  "id": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
  "name": "register",
  "page": "index.html",
  "type": "mouseover",
  "element": "button",
  "ip": "xxx.xxx.xxx.xxx",
  "createdAt": "2021-01-01T00:00:00Z"
}
```
