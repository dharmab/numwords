package numwords

import (
	"bufio"
	"strings"
)

func explode(s string) (out []string) {
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, ",", "")

	r := strings.NewReader(s)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	return
}
