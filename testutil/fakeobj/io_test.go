package fakeobj_test

import (
	"testing"

	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/testutil/fakeobj"
)

func TestNewReader(t *testing.T) {
	tr := fakeobj.NewReader()
	tr.Write([]byte("hello"))

	buf := make([]byte, 5)
	n, err := tr.Read(buf)
	assert.NoErr(t, err)
	assert.Eq(t, 5, n)
	assert.Eq(t, "hello", string(buf))

	tr.SetErrOnRead()
	_, err = tr.Read(buf)
	assert.Err(t, err)
}

func TestNewWriter(t *testing.T) {
	tw := fakeobj.NewBuffer()
	_, err := tw.Write([]byte("hello"))
	assert.NoErr(t, err)
	assert.Eq(t, "hello", tw.String())
	assert.NoErr(t, tw.Flush())
	assert.NoErr(t, tw.Sync())
	assert.Eq(t, "", tw.String())
	assert.NoErr(t, tw.Close())

	// write string
	_, err = tw.WriteString("hello")
	assert.NoErr(t, err)
	assert.Eq(t, "hello", tw.ResetGet())

	tw.SetErrOnWrite()
	_, err = tw.Write([]byte("hello"))
	assert.Err(t, err)
	assert.Eq(t, "", tw.String())

	tw.SetErrOnFlush()
	assert.Err(t, tw.Flush())

	tw.SetErrOnSync()
	assert.Err(t, tw.Sync())

	tw.SetErrOnClose()
	assert.Err(t, tw.Close())
}
