package cmd

import (
	"cf-tool/client"
)

// Register command
func Register() (err error) {
	cln := client.Instance
	info := Args.Info
	
	if (Args.All) {
		if err = cln.RegisterAll(info); err != nil {
			if err = loginAgain(cln, err); err == nil {
				err = cln.RegisterAll(info)
			}
		}
		return
	}

	if err = cln.Register(info); err != nil {
		if err = loginAgain(cln, err); err == nil {
			err = cln.Register(info)
		}
	}

	return
}
