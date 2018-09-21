package exists

import "os"

func Exists(target string) bool {
	_, err := os.Stat(target)
	if os.IsNotExist(err) {
		return false
	}

	return true
}
