package main

import (
	"io"

	"github.com/bearmini/sexp"
	"github.com/pkg/errors"
)

type WastReader struct {
	r *WastCommentTrimLineReader
}

func NewWastReader(r io.Reader) *WastReader {
	return &WastReader{
		r: NewWastCommentTrimLineReader(r),
	}
}

func (wr *WastReader) NextSexp() (*sexp.Sexp, error) {
	buf := ""

	for {
		line, err := wr.r.NextLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		buf += line
		se, consumed, err := tryFindSexp(buf)
		if err != nil {
			return nil, err
		}

		if se != nil {
			buf = buf[consumed:]
			return se, nil
		}
	}

	return nil, nil
}

func tryFindSexp(s string) (*sexp.Sexp, int, error) {
	buf := []rune{}
	n := 0
	for _, r := range s {
		buf = append(buf, r)

		if r == '(' {
			n++
			continue
		}
		if r == ')' {
			n--
			if n == 0 {
				break
			}
			if n < 0 {
				return nil, 0, errors.New("unmatched parentheses")
			}
		}
	}

	se, err := sexp.Parse(string(buf))
	if err != nil {
		return nil, 0, err
	}
	return se, len(string(buf)), nil
}
