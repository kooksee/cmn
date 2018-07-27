package cmn

import "log"

var Assert = myAssert{}

type myAssert struct{}

func (m myAssert) StrEqual(a, b string) {
	if a != b {
		log.Fatalf(F("%s != %s", a, b))
	}
}
