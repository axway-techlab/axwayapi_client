package main

import (
	"encoding/json"
	"fmt"
	"github.com/axway-techlab/axwayapi_client/axwayapi"
)

func main() {
	host := "https://manager.testing.axway-techlabs.com/api/portal/v1.4"
	username := "apiadmin"
	password := "changeme"
	c, err := axwayapi.NewClient(&host, &username, &password)
	if err != nil {
		panic(err)
	}

	config, err := c.GetConfig()
	if err != nil {
		panic(err)
	}

	config.MinimumPasswordLength = 5
	config.LockUserAccount.Attempts = 10

	config, err = c.UpdateConfig(config)
	if err != nil {
		panic(err)
	}

	config, err = c.GetConfig()
	if err != nil {
		panic(err)
	} else {
		b, _ := json.Marshal(&config)
		fmt.Println(string(b))
	}
}
