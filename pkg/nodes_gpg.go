package pkg

import (
	"os/exec"

	"bufio"

	"fmt"

	"github.com/pkg/errors"
)

func CreateNodesWithGPG() ([]Node, error) {
	cmd := exec.Command("gpg", "--list-sigs")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create stdout pipe gpg command")
	}

	if err := cmd.Start(); err != nil {
		return nil, errors.Wrap(err, "failed to start gpg command")
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err = cmd.Wait(); err != nil {
		return nil, errors.Wrap(err, "failed to wait for gpg command")
	}

	return nil, nil
}
