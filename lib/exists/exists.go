package exists

import "os"

// Exists checks whether the provided target does in fact exist, and returns a
// boolean. Exists returns 'true' if the target exists, and 'false' if it does
// not.
func Exists(target string) bool {
	_, err := os.Stat(target)
	if os.IsNotExist(err) {
		return false
	}

	return true
}
