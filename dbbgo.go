package dbbgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Client Create a new DBB Client to post stats
type Client struct {
	ClientID string
	Token    string
}

// Stats Stats structure for non-sharded bots
type Stats struct {
	Guilds int `json:"guilds"`
}

// ShardedStats Stats for Sharded Bots
type ShardedStats struct {
	Guilds int   `json:"guilds"`
	Shards []int `json:"shards"`
}

// Error Handling
type missingCredentialsError struct {
	err string
}

func (e *missingCredentialsError) Error() string {
	return fmt.Sprintf("Missing Credentials: %s", e.err)
}

var fetch = &http.Client{}

// DBBClient Create a new DBBClient to post stats
func DBBClient(token, clientID string) *Client {
	client := &Client{
		ClientID: clientID,
		Token:    token,
	}

	return client
}

// PostStats Post Stats to DBB. For Non Sharded Bots only
func (c *Client) PostStats(guilds int) error {
	if c.Token == "" {
		noToken := &missingCredentialsError{err: "PostStats function needs a token, but no token was provided."}
		log.Fatal(noToken)
		return noToken
	}

	stats := Stats{
		Guilds: guilds,
	}

	jsonV, err := json.Marshal(&stats)

	if err != nil {
		log.Fatal(err)
		return err
	}

	req, err := http.NewRequest("POST", "https://discordsbestbots.xyz/api/bots/"+c.ClientID+"/stats", bytes.NewBuffer(jsonV))

	if err != nil {
		log.Fatal(err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.Token)

	_, err = fetch.Do(req)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// PostShardedStats Post Stats to DBB. For Sharded Bots only
func (c *Client) PostShardedStats(guilds int, shards []int) error {
	if c.Token == "" {
		noToken := &missingCredentialsError{err: "PostShardedStats function needs a token, but no token was provided."}
		log.Fatal(noToken)
		return noToken
	}

	stats := ShardedStats{
		Guilds: guilds,
		Shards: shards,
	}

	jsonV, err := json.Marshal(&stats)

	if err != nil {
		log.Fatal(err)
		return err
	}

	req, err := http.NewRequest("POST", "https://discordsbestbots.xyz/api/bots/"+c.ClientID+"/stats", bytes.NewBuffer(jsonV))

	if err != nil {
		log.Fatal(err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.Token)

	_, err = fetch.Do(req)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
