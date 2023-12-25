package client

import (
	"cf-tool/util"
	"encoding/json"
	"fmt"
)

type Contest struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	Type                string `json:"type"`
	Phase               string `json:"phase"`
	Frozen              bool   `json:"frozen"`
	DurationSeconds     int    `json:"durationSeconds"`
	StartTimeSeconds    int    `json:"startTimeSeconds"`
	RelativeTimeSeconds int    `json:"relativeTimeSeconds"`
}

// CList get list of 10 contests
func (c *Client) CList(info Info) (contests []Contest, err error) {
	URL, err := info.ContestListURL(c.host)
	if err != nil {
		return
	}

	body, err := util.GetBody(c.client, URL)
	if err != nil {
		return
	}
	
	// marshal body to json
	data := struct {
        Status string `json:"status"`
        Result []Contest `json:"result"`
    }{}
	if err = json.Unmarshal(body, &data); err != nil {
		return
	}


	if status := data.Status; status != "OK" {
		return nil, fmt.Errorf("cannot get any contests")
	}
	
	contests = data.Result
	return contests, nil
}
