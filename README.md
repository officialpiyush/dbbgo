<div align="center">
<img align="center" src="https://i.vgy.me/e0YicW.png" alt="Logo" />
<br>
<br>
<strong><i>An Unofficial API Wrapper around https://discordsbestbots.xyz/ </i> </strong>
</div>

---

# Installing

* Install it using

```
go get github.com/officialpiyush/dbbgo
```


# Initialization

* Build The Client for DiscordBestBots:

```go
var id = "" // Your Bot's ID
var token = "" // The discordsbestbots.xyz token

c := dbbgo.DBBClient(token, id)
```

# Methods

## func PostStats

```go
func (c *Client) PostStats(guilds int) error
```

* POST Bot's Stats to the API
* This function is only for Non Sharded Bots. Sharded Bots should use `PostShardedStats` func.

<details>
<summary>Example</summary>

```go
err := c.PostStats(guild_count) // where guild_count is your bot's guild count. INT Parameter
```
</details>

## func PostShardedStats

```go
func (c *Client) PostShardedStats(guilds int, shards []int) error
```

* POST Sharded Bot's Stats to the API

<details>
<summary>Example</summary>

```go
err := c.PostShardedStats(guild_count, []int{21,43}) // where guild_count is your bot's guild count and 2nd param is . INT Parameter
```
</details>

