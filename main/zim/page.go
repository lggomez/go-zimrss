package zim

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func ParsePage(path string, info os.FileInfo) (*PageMetadata, error) {
	// Read the first 6 lines which compose the page header
	lines := make([]string, 0, 6)
	f, err := os.Open(path)
	if err != nil {
		return &PageMetadata{}, err
	}
	scanner := bufio.NewScanner(f)
	for i := 0; i < 6 && scanner.Scan(); i++ {
		lines = append(lines, scanner.Text())
	}

	if len(lines) < 6 {
		return &PageMetadata{}, fmt.Errorf("zim: error parsing page - invalid header length %v", len(lines))
	}

	creationDateTokens := strings.Split(lines[2], " ")
	creationDate, err := time.Parse(time.RFC3339, creationDateTokens[1])
	if err != nil {
		creationDate = time.Now()
	}

	// NOTE: This assumes the language on which the content created date is english
	// i.e: Created Saturday 02 January 2021
	// ref for format tokens: https://github.com/golang/go/blob/a9cc1051c11f821cb03d63fb9e05930f9e2f9fa5/src/time/format.go#L92
	var contentDate time.Time
	if lines[5] != "" {
		contentDateTokens := strings.SplitN(lines[5], " ", 2)
		contentDate, err = time.Parse("Monday 01 January 2006", contentDateTokens[1])
		if err != nil {
			contentDate = time.Now()
		}
	}

	contentTypeTokens := strings.Split(lines[0], " ")
	contentType := contentTypeTokens[1]

	wikiFormatTokens := strings.SplitN(lines[1], " ", 2)
	wikiFormat := wikiFormatTokens[1]

	metadata := &PageMetadata{
		Title:        strings.TrimSpace(strings.ReplaceAll(lines[4], "======", "")),
		ContentType:  contentType,
		WikiFormat:   wikiFormat,
		CreationDate: creationDate,
		ContentDate:  &contentDate,
		Path:         path,
	}

	return metadata, nil
}
