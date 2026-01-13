package code

import (
	"fmt"
	// "os"
	// "path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSize(t *testing.T) {
	// wd, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }

	// parentDir := filepath.Dir(wd)

	res, err := GetSize("testdata/testfile1.txt", true, false)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	require.Equal(t, res, int64(9))

	res, err = GetSize("testdata/textfile2.txt", true, false)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	require.Equal(t, res, int64(0))

	res, err = GetSize("testdata/testfolder", true, false)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	require.Equal(t, res, int64(10))

	res, err = GetSize("testdata/testfolder", true, true)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	require.Equal(t, res, int64(14))
}

func TestFormatSize(t *testing.T) {
	testNum1 := float64(10)
	testNum2 := float64(1024 * 10)
	testNum3 := float64(1024 * 1024 * 10)
	testNum4 := float64(1024 * 1024 * 1024 * 10)
	testNum5 := float64(1024 * 1024 * 1024 * 1024 * 10)
	testNum6 := float64(1024 * 1024 * 1024 * 1024 * 1024 * 10)

	res := FormatSize(int64(testNum1), true)
	require.Equal(t, res, "10B")

	res = FormatSize(int64(testNum2), true)
	require.Equal(t, res, "10.0KB")

	res = FormatSize(int64(testNum3), true)
	require.Equal(t, res, "10.0MB")

	res = FormatSize(int64(testNum4), true)
	require.Equal(t, res, "10.0GB")

	res = FormatSize(int64(testNum5), true)
	require.Equal(t, res, "10.0TB")

	res = FormatSize(int64(testNum6), true)
	require.Equal(t, res, "10.0PB")
}
