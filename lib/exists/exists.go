// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

package exists

import (
	//"errors"
	"os"
)

// Exists checks whether the provided target does in fact exist, and returns a
// boolean. Exists returns 'true' if the target exists, and 'false' if it does
// not.
func Exists(target string) error {
	_, err := os.Stat(target)
	if os.IsNotExist(err) {
		return err
	}

	return nil
}
