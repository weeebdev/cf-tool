package client

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"cf-tool/util"

	"github.com/fatih/color"
)

// Register contest
func (c *Client) Register(info Info) (err error) {
	color.Cyan("Register " + info.ContestID)

	URL, err := info.RegisterURL(c.host)
	if err != nil {
		return
	}

	body, err := util.GetBody(c.client, URL)
	if err != nil {
		return
	}

	handle, err := findHandle(body)
	if err != nil {
		return
	}

	fmt.Printf("Current user: %v\n", handle)

	csrf, err := findCsrf(body)
	if err != nil {
		return
	}

	body, err = util.PostBody(c.client, fmt.Sprintf("%v?csrf_token=%v", URL, csrf), url.Values{
		"csrf_token":          {csrf},
		"ftaa":                {c.Ftaa},
		"bfaa":                {c.Bfaa},
		"action":              {"formSubmitted"},
		"tabSize":             {"4"},
		"_tta":                {"394"},
		"sourceCodeConfirmed": {"true"},
		"takePartAs":          {"personal"},
	})
	if err != nil {
		return
	}

	errMsg, err := findErrorMessage(body)
	if err == nil {
		return errors.New(errMsg)
	}

	msg, err := findMessage(body)
	if err != nil {
		return errors.New("Register failed")
	}
	if !strings.Contains(msg, "successfully registered") {
		return errors.New(msg)
	}

	color.Green("Registered")

	return
}
