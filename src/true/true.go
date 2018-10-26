// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// return true value
//
// true always exists with a zero exit code.
package true

import "os"

func True() {
	os.Exit(0)
}
