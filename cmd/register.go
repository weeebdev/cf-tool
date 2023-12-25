package cmd

import (
	"cf-tool/client"
)

// Register command
func Register() (err error) {
	cln := client.Instance
	info := Args.Info

	if err = cln.Register(info); err != nil {
		if err = loginAgain(cln, err); err == nil {
			err = cln.Register(info)
		}
	}

	return
}
