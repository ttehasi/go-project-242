package code

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSize(t *testing.T) {
	res, err := GetSize("testdata/testfile1.txt")
	if err != nil {
		t.FailNow()
	}
	require.Equal(t, res, int64(9))

	res, err = GetSize("testdata/textfile2.txt")
	if err != nil {
		t.FailNow()
	}
	require.Equal(t, res, int64(0))

	res, err = GetSize("testdata/testfolder")
	if err != nil {
		t.FailNow()
	}
	require.Equal(t, res, int64(4))

	res, err = GetSize("testdata/testfolderempty")
	if err != nil {
		t.FailNow()
	}
	require.Equal(t, res, int64(0))
}
