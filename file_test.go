package xlsxreader

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGettingFileByNameSuccess(t *testing.T) {
	zipFiles := []*zip.File{
		{FileHeader: zip.FileHeader{Name: "Bill"}},
		{FileHeader: zip.FileHeader{Name: "Bobby"}},
		{FileHeader: zip.FileHeader{Name: "Bob"}},
		{FileHeader: zip.FileHeader{Name: "Ben"}},
	}

	file, err := getFileForName(zipFiles, "Bob")

	require.NoError(t, err)
	require.Equal(t, zipFiles[2], file)
}

func TestGettingFileByNameFailure(t *testing.T) {
	zipFiles := []*zip.File{}

	_, err := getFileForName(zipFiles, "OOPS")

	require.EqualError(t, err, "File not found: OOPS")

}

func TestOpeningMissingFile(t *testing.T) {
	_, err := OpenFile("/Users/zhangbob/git/golang/test/xls/1.xlsx")

	require.EqualError(t, err, "open this_doesnt_exist.zip: no such file or directory")
}

func TestToSlice(t *testing.T) {
	out, err := ToSlice("/Users/zhangbob/git/golang/test/xls/1.xlsx")
	if err != nil {
		t.Fatal("err:", err.Error())
	}

	for _, sheet := range out {
		for _, row := range sheet {
			for _, cell := range row {
				if cell == "" {
					fmt.Printf("%s ", "空值")
				} else {
					fmt.Printf("%s ", cell)
				}
			}
			fmt.Println("")
		}
	}
}

func TestOpeningXlsxFile(t *testing.T) {
	actual, err := OpenFile("./test/test-small.xlsx")
	defer actual.Close()

	require.NoError(t, err)
	require.Equal(t, []string{"datarefinery_groundtruth_400000"}, actual.Sheets)
}

func TestClosingFile(t *testing.T) {
	actual, err := OpenFile("./test/test-small.xlsx")
	require.NoError(t, err)
	err = actual.Close()
	require.NoError(t, err)
}

func TestNewReaderFromXlsxBytes(t *testing.T) {
	f, _ := os.Open("./test/test-small.xlsx")
	defer f.Close()

	b, _ := ioutil.ReadAll(f)

	actual, err := NewReader(b)

	require.NoError(t, err)
	require.Equal(t, []string{"datarefinery_groundtruth_400000"}, actual.Sheets)
}
