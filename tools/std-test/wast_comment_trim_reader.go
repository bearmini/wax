package main

import (
	"bufio"
	"io"
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
	nest := 0

	for {
		line, err = wr.r.ReadString('\n')
		if err != nil {
			if err != io.EOF || line == "" {
				return "", err
			}
		}

		line, nest = trimComments(line, nest)
		if line != "" {
			return line, nil
		}
	}
}

func trimComments(s string, nest int) (string, int) {
	for i := range s {
		if i == 0 {
			continue
		}
		if s[i] == ';' {
			if s[i-1] == ';' {
				if nest == 0 {
					return s[:i-1], nest
				}
			}
			if s[i-1] == '(' {
				if nest == 0 {
					pre := s[:i-1]
					post, newNest := trimComments(s[i+1:], nest+1)
					return pre + post, newNest
				}
				return trimComments(s[i+1:], nest+1)
			}
		}
		if s[i] == ')' {
			if s[i-1] == ';' {
				if nest == 0 {
					return trimComments(s[i+1:], 0)
				}
				return trimComments(s[i+1:], nest-1)
			}
		}
	}

	if nest == 0 {
		return s, nest
	}

	return "", nest
}

func searchBlockCommentEnd(s string) int {
	for i := range s {
		if i == 0 {
			continue
		}
		if s[i-1] == ';' && s[i] == ')' {
			return i - 1
		}
	}
	return -1
}
