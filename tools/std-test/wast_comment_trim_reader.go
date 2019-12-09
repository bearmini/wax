package main

import (
	"bufio"
	"io"
	"strings"
)

type WastCommentTrimLineReader struct {
	r *bufio.Reader
}

func NewWastCommentTrimLineReader(r io.Reader) *WastCommentTrimLineReader {
	return &WastCommentTrimLineReader{
		r: bufio.NewReader(r),
	}
}

func (wr *WastCommentTrimLineReader) NextLine() (string, error) {
	var err error
	line := ""

	for {
		line, err = wr.r.ReadString('\n')
		if err != nil {
			if err != io.EOF || line == "" {
				return "", err
			}
		}

		line = trimComment(line)
		if line != "" {
			return line, nil
		}
	}
}

func trimComment(s string) string {
	p := strings.Index(s, ";;")
	if p >= 0 {
		s = s[:p]
	}

	return strings.Trim(s, " ")
}
