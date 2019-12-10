package wax

import (
	"math"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func ParseHexfloat32(s string) (float32, error) {
	f, err := ParseHexfloat64(s)
	if err != nil {
		return 0, err
	}
	return float32(f), nil
}

func ParseHexfloat64(s string) (float64, error) {
	s = strings.Trim(s, " \r\n\t")
	sign := float64(1)
	if strings.HasPrefix(s, "-") {
		sign = -1
		s = strings.TrimPrefix(s, "-")
	}
	if !strings.HasPrefix(s, "0x") && !strings.HasPrefix(s, "0X") {
		return 0, errors.Errorf("invalid format: %s", s)
	}
	s = strings.TrimPrefix(s, "0x")
	s = strings.TrimPrefix(s, "0X")

	shexnumfrac := s
	spow := "0"
	i := strings.IndexAny(s, "pP")
	if i >= 0 {
		shexnumfrac = s[:i]
		spow = strings.Trim(s[i:], "pP")
	}

	shexnum := shexnumfrac
	shexfrac := "0"
	j := strings.IndexRune(shexnumfrac, '.')
	if j >= 0 {
		shexnum = shexnumfrac[:j]
		shexfrac = strings.Trim(shexnumfrac[j:], ".")
	}

	hexnum, err := strconv.ParseUint(shexnum, 16, 64)
	if err != nil {
		return 0, err
	}

	hexfrac, err := strconv.ParseUint(shexfrac, 16, 64)
	if err != nil {
		return 0, err
	}

	fhexfrac := float64(hexfrac) / float64(math.Pow(16, float64(len(shexfrac))))

	pow, err := strconv.ParseInt(spow, 10, 64)
	if err != nil {
		return 0, err
	}

	return sign * (float64(hexnum) + fhexfrac) * math.Pow(2, float64(pow)), nil
}
