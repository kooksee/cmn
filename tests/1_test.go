package tests

import (
	"testing"
	"github.com/kooksee/cmn"
)

func TestName(t *testing.T) {
	cmn.P(cmn.Color.Blue("ffff"))
	cmn.P(cmn.Rand.RandStr(4))
}
