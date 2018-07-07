package csv

import (
	"bufio"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/ingmardrewing/fs"
)

func ReadCsv(filePath, fileName string) ([][]string, error) {
	completePath := path.Join(filePath, fileName)
	f, err := os.Open(completePath)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(bufio.NewReader(f))

	lines := [][]string{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		lines = append(lines, record)
	}
	return lines, nil
}

func WriteCsv(filepath, filename string, records [][]string) error {
	txt := ""
	if ok, err := fs.PathExists(filepath); !ok {
		if err != nil {
			return err
		}
		return errors.New("Path doesn't exist: " + filepath)
	}
	for _, record := range records {
		line := ""
		for _, field := range record {
			line = line + wrapContent(field) + ","
		}
		line = strings.TrimSuffix(line, ",")
		txt = txt + line + "\n"
	}
	fs.WriteStringToFS(filepath, filename, txt)
	return nil
}

func wrapContent(txt string) string {
	rx := regexp.MustCompile(`"`)
	return `"` + rx.ReplaceAllString(txt, ``) + `"`
}
