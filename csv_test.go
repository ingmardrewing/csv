package csv

import (
	"path"
	"runtime"
	"testing"

	"github.com/ingmardrewing/fs"
)

func currentDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

func TestReadCsv(t *testing.T) {
	filepath := path.Join(currentDir(), "testResources/read/")
	filename := "test.csv"

	data, err := ReadCsv(filepath, filename)
	if err != nil {
		t.Error("An error occured:", err)
	}

	expected := "test line one"
	actual := data[0][0]

	if actual != expected {
		t.Error("Expected", expected, ", but got", actual)
	}

	expected = "test, with comma"
	actual = data[2][2]

	if actual != expected {
		t.Error("Expected", expected, ", but got", actual)
	}
}

func TestWrapContent(t *testing.T) {
	expected := `"whassup, man?"`
	actual := wrapContent(`"whassup, man?"`)

	if actual != expected {
		t.Error("Expected", expected, "dtos, but got", actual)
	}
}

func TestWriteCsv(t *testing.T) {
	record1 := []string{"test1", `"another test"`}
	record2 := []string{"test2", `"yet another test"`}
	data := [][]string{record1, record2}

	filepath := path.Join(currentDir(), "testResources/write/")
	filename := "writetest.csv"
	err := WriteCsv(filepath, filename, data)

	if err != nil {
		t.Error(err)
	}

	expected := `"test1","another test"
"test2","yet another test"
`
	actual := fs.ReadFileAsString(filepath + "/" + filename)

	if actual != expected {
		t.Error("Expected", expected, "dtos, but got", actual)
	}
	fs.RemoveFile(filepath, filename)
}
