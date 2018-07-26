package cmn

import (
	"bytes"
	"sync"
)

func NewKBuffer() *kBuffer {
	return &kBuffer{dmt: []byte("\n")}
}

type kBuffer struct {
	buf  []byte
	dmt  []byte
	lock sync.RWMutex
}

func (t *kBuffer) SetDmt(dmt []byte) *kBuffer {
	t.dmt = dmt
	return t
}

func (t *kBuffer) Next(b []byte) [][]byte {
	t.lock.Lock()
	defer t.lock.Unlock()

	if b == nil {
		return nil
	}

	t.buf = append(t.buf, b...)

	if len(t.buf) < 1 {
		return nil
	}

	if !bytes.Contains(t.buf, t.dmt) {
		return nil
	}

	d := bytes.Split(t.buf, t.dmt)
	if len(d) < 1 {
		return nil
	}

	t.buf = d[len(d)-1]
	return d[:len(d)-1]
}
