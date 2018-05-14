package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPubExp(t *testing.T) {
	line := "pub   rsa4096/0xABCD123456789A1 2018-04-22 [SC]"

	assert.True(t, pubExp.MatchString(line))
	assert.Equal(t, "0xABCD123456789A1", pubExp.FindStringSubmatch(line)[1])
}

func TestFingerprintExp(t *testing.T) {
	line := "  Schl.-Fingerabdruck = 9B78 C44E 0625 644D 56B3  027F 14B1 5D4C 8C93 50A2"

	assert.True(t, fingerprintExp.MatchString(line))
	assert.Equal(t, "9B78 C44E 0625 644D 56B3  027F 14B1 5D4C 8C93 50A2", fingerprintExp.FindStringSubmatch(line)[1])
}

func TestUidExp(t *testing.T) {
	line := "uid                   [ultimate] Tricia McMillian <tricia.mcmillian@hitchhiker.com>"

	assert.True(t, uidExp.MatchString(line))
	assert.Equal(t, "Tricia McMillian", uidExp.FindStringSubmatch(line)[1])
	assert.Equal(t, "tricia.mcmillian@hitchhiker.com", uidExp.FindStringSubmatch(line)[2])
}

func TestSigExp(t *testing.T) {
	line := "sig 3        0xABCD123456789A1 2018-05-08  Tricia McMillian <tricia.mcmillian@hitchhiker.com>"

	assert.True(t, sigExp.MatchString(line))
	assert.Equal(t, "0xABCD123456789A1", sigExp.FindStringSubmatch(line)[1])
}
