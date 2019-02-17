// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package true

import "os"

// True always exists with a zero exit code.
func True() {
	os.Exit(0)
}
