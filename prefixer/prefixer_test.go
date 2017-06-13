package prefixer

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestPrefixLines(t *testing.T) {
	prefix := `a go {{"template"}}`
	inputString := "lala\nbaba"

	scanner := bufio.NewScanner(strings.NewReader(inputString))
	var buffer bytes.Buffer

	prefixer := NewPrefixer(prefix, scanner, &buffer)

	prefixer.PrefixLines()

	lines := strings.Split(strings.TrimSpace(buffer.String()), "\n")

	t.Logf("actual output: %q", buffer.String())

	for _, line := range lines {
		if !strings.Contains(line, "a go template") {
			t.Errorf("expected line %q to start with %q", line, prefix)
		}
	}

}
