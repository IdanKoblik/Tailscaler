# Tailscaler

A tailscale CLI application that allow to organize In a table tailscale users data.
* Usefully for managing tailscale users data from multiple routes.
* This project is for a private usage only and learning purpose.

## Setup
#### API setup
Inside ```/tailscaler/api``` create ```config.json```
file with the next information:
```json
{
    "ip": "domain/api",
    "port": "port"
}

```
#### CLI setup
Inside ```/tailscaler/cli``` create ```config.json```
file with the next information:
```json
{
    "apiURL": "the api url that you configured inside `/tailscaler/api/config.json`"
}
```

## API Reference

#### Add user to the API

```http
 POST /tailscaler/create_user
```

| Parameter      | Type    | Description                          |
|------------|---------|--------------------------------------|
| Router     | `string`  | **Required** Router sender name.                  |
| ID         | `string`  | **Required** User ID.                             |
| HostName   | `string`  | **Required** Hostname for the user.               |
| OS         | `string`  | Operating system of the user.        |
| AllowedIPs | `string array`   | List of allowed IPs for the user.    |
| CurAddr    | `string`  | Current address of the user.         |
| Active     | `string`  | Status of the user (active/inactive).|


#### Get all users

```http
  GET /tailscaler/get_users
```

#### Lookup user

```http
  GET /tailscaler/find_user_by_id/{hostName}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `hostName`      | `string` | **Required**. Tailscale user hostname |
