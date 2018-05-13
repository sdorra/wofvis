package pkg

import (
	"fmt"

	"encoding/json"

	"github.com/urfave/cli"
)

func PrintJSON(ctx *cli.Context) error {
	nodes, err := createNodes()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(nodes, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	return nil
}
