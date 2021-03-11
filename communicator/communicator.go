package communicator

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"uno/engine"
)


type Communicator struct {
	Url   string
	Name  string
	token string
	matchId string
}

type startResp struct {
	MatchId string `json:"matchId"`
	Status uint8 `json:"status"`
}

type BoardResp struct {
	MatchId string `json:"matchId"`
	Status engine.GameStatus `json:"status"`
	MyMove bool `json:"myMove"`
	Hand []engine.Card `json:"hand"`
	CurrCard engine.Card `json:"currentCard"`
}

func (c *Communicator) Token() error {
	jd, err := json.Marshal(map[string]string{"name": c.Name})
	if err != nil {
		return err
	}

	resp, err := http.Post(
		fmt.Sprintf("%s/api/Match/token", c.Url),
		"application/json",
		bytes.NewBuffer(jd),
	)
	if err != nil {
		return err
	}

	var res map[string]string

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return err
	}

	c.token = res["token"]
	fmt.Printf("Getting token = %s\n", c.token)
	return nil
}

func (c *Communicator) StartGame(oToken string) error {
	for {
		m := map[string]string{"token": c.token}

		if oToken != "" {
			m["opponent"] = oToken
		}

		jd, err := json.Marshal(m)
		if err != nil {
			return err
		}

		resp, err := http.Post(
			fmt.Sprintf("%s/api/Match/start", c.Url),
			"application/json",
			bytes.NewBuffer(jd),
		)
		if err != nil {
			return err
		}
		if resp.StatusCode != http.StatusOK {
			return errors.New(fmt.Sprintf("start game exit with code = %d", resp.StatusCode))
		}

		var res startResp

		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return err
		}

		if res.Status == 0 {
			fmt.Printf("Cannot start game in queue, retry\n")
			continue
		}

		c.matchId = res.MatchId
		fmt.Printf("Game started\n")
		return nil
	}
}

func (c *Communicator) Board() (*BoardResp, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api/Game/board?Token=%s&MatchId=%s", c.Url, c.token, c.matchId))
	if err != nil {
		return nil, err
	}

	var br *BoardResp
	if err := json.NewDecoder(resp.Body).Decode(&br); err != nil {
		return nil, err
	}
	return br, nil
}
