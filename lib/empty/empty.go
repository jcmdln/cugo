package empty

import (
	"io"
	"os"
)

func Empty(dir string) bool {
	t, err := os.Open(dir)
	defer t.Close()

	_, err = t.Readdirnames(1)
	if err == io.EOF {
		return true
	}

	return false
}
