package pkg

import "github.com/urfave/cli"

type Node struct {
	ID          string   `json:"id"`
	Fingerprint string   `json:"fingerprint"`
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	SignedBy    []string `json:"signedBy"`
}

func (node *Node) addSignature(signature string) {
	if node.SignedBy == nil {
		node.SignedBy = make([]string, 0)
	}

	if node.ID == signature {
		return
	}

	for _, sig := range node.SignedBy {
		if sig == signature {
			return
		}
	}
	node.SignedBy = append(node.SignedBy, signature)
}

type NodeFactory func() ([]Node, error)

func CreateNodeFactory(ctx *cli.Context) NodeFactory {
	if ctx.GlobalBool("use-openpgp-api") {
		return CreateNodesWithOpenPGP
	}
	return CreateNodesWithGPG
}
