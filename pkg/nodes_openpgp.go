package pkg

import (
	"encoding/hex"
	"os"
	"unicode"

	"github.com/pkg/errors"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/packet"
)

func identity(entity *openpgp.Entity) *openpgp.Identity {
	var identity *openpgp.Identity
	for _, identity = range entity.Identities {

	}
	return identity
}

func idOf(pubring openpgp.EntityList, sig *packet.Signature) string {
	signedBy := findKeyByID(pubring, *sig.IssuerKeyId)
	if signedBy != nil {
		return createId(signedBy)
	}
	return ""
}

func createId(entity *openpgp.Entity) string {
	return "0x" + entity.PrimaryKey.KeyIdString()
}

func createFingerprint(entity *openpgp.Entity) string {
	bytes := entity.PrimaryKey.Fingerprint

	ba := make([]byte, 0)
	for _, b := range bytes {
		ba = append(ba, b)
	}

	counter := 0

	fingerprint := ""
	for i, c := range hex.EncodeToString(ba) {
		if i%4 == 0 {
			if counter > 0 {
				fingerprint += " "
				if counter == 5 {
					fingerprint += " "
				}
			}
			counter++
		}
		fingerprint += string(unicode.ToUpper(rune(c)))
	}
	return fingerprint
}

func createNode(pubring openpgp.EntityList, entity *openpgp.Entity) Node {
	identity := identity(entity)

	node := Node{
		ID:          createId(entity),
		Name:        identity.UserId.Name,
		Email:       identity.UserId.Email,
		Fingerprint: createFingerprint(entity),
	}

	for _, sig := range identity.Signatures {
		signatureId := idOf(pubring, sig)
		if signatureId != "" {
			node.addSignature(signatureId)
		}
	}

	return node
}

func findKeyByID(pubring openpgp.EntityList, key uint64) *openpgp.Entity {
	for _, entity := range pubring {
		if entity.PrimaryKey.KeyId == key {
			return entity
		}
	}
	return nil
}

func CreateNodesWithOpenPGP() ([]Node, error) {
	pubringPath := os.ExpandEnv("$HOME/.gnupg/pubring.gpg")

	pubringFile, err := os.Open(pubringPath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open pubring %s", pubringPath)
	}

	pubring, err := openpgp.ReadKeyRing(pubringFile)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read keyring %s", pubringPath)
	}

	nodes := make([]Node, 0)
	for _, entity := range pubring {
		wofKey := createNode(pubring, entity)
		nodes = append(nodes, wofKey)
	}

	return nodes, nil
}
