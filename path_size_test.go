package code

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSize(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	parentDir := filepath.Dir(wd)

	res, err := GetSize(parentDir + "/hexlet-go-1/testdata/testfile1.txt")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	require.Equal(t, res, int64(9))

	res, err = GetSize(parentDir + "/hexlet-go-1/testdata/textfile2.txt")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	require.Equal(t, res, int64(0))

	res, err = GetSize(parentDir + "/hexlet-go-1/testdata/testfolder")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	require.Equal(t, res, int64(4))

	res, err = GetSize(parentDir + "/hexlet-go-1/testdata/testfolderempty")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	require.Equal(t, res, int64(0))
}
