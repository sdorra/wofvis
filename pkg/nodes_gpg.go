package pkg

import (
	"os/exec"

	"bufio"

	"regexp"

	"strings"

	"github.com/pkg/errors"
)

var (
	pubExp         = regexp.MustCompile("^pub\\s+.*\\/([^\\s ]+)\\s+.*$")
	fingerprintExp = regexp.MustCompile("^.+=\\s+((:?[0-9A-Z]{4}\\s*){10})$")
	uidExp         = regexp.MustCompile("^uid\\s+\\[.*\\]\\s+([^<]+)\\s+<(.*)>$")
	sigExp         = regexp.MustCompile("^sig[\\s]+[0-9]?\\s+([^\\s]+).*$")
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

	nodes := make([]Node, 0)
	var node *Node = nil

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()

		if pubExp.MatchString(line) {
			if node != nil {
				nodes = append(nodes, *node)
			}
			node = &Node{
				ID: pubExp.FindStringSubmatch(line)[1],
			}
		} else if fingerprintExp.MatchString(line) {
			if node != nil {
				node.Fingerprint = fingerprintExp.FindStringSubmatch(line)[1]
			}
		} else if uidExp.MatchString(line) {
			if node != nil {
				uidParts := uidExp.FindStringSubmatch(line)
				node.Name = uidParts[1]
				node.Email = uidParts[2]
			}
		} else if sigExp.MatchString(line) {
			if node != nil {
				node.addSignature(sigExp.FindStringSubmatch(line)[1])
			}
		} else if strings.HasPrefix(line, "sub") {
			if node != nil {
				nodes = append(nodes, *node)
				node = nil
			}
		}
	}

	if err = cmd.Wait(); err != nil {
		return nil, errors.Wrap(err, "failed to wait for gpg command")
	}

	return nodes, nil
}
