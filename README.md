# Tailscaler

A Tailscale CLI application that allow to organize In a table Tailscale nodes data.
* Usefully for managing Tailscale nodes data from multiple routes.
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
    "apiURL": "the api url that you configured inside `/tailscaler/api/test_config.json`"
}
```

## API Reference

#### Add node to the API

```http
 POST /tailscaler/create_node
```

| Parameter      | Type    | Description                           |
|------------|---------|---------------------------------------|
| Router     | `string`  | **Required** Router sender name.      |
| ID         | `string`  | **Required** Node ID.                 |
| HostName   | `string`  | **Required** Hostname for the node.   |
| OS         | `string`  | Operating system of the node.         |
| AllowedIPs | `string array`   | List of allowed IPs for the node.     |
| CurAddr    | `string`  | Current address of the node.          |
| Active     | `string`  | Status of the node (active/inactive). |


#### Get all nodes

```http
  GET /tailscaler/get_nodes
```

#### Lookup nodes

```http
  GET /tailscaler/find_node_by_id/{hostName}
```

| Parameter | Type     | Description                           |
| :-------- | :------- |:--------------------------------------|
| `hostName`      | `string` | **Required**. Tailscale node hostname |
***

#### Routing system
![image](https://raw.githubusercontent.com/IdanKoblik/Tailscaler/main/assets/router.png?token=GHSAT0AAAAAACOISCMXKZAUIAG3SCSX4RUAZRCIARQ)
