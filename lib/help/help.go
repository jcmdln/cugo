// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

package help

import (
	"fmt"
	"os"

	"github.com/jcmdln/flagger"
)

func Help(name, usage, description string, flags *flagger.Flags) {
	fmt.Printf("Usage: %s %s\n%s\n\n", name, usage, description)
	flags.Print("Options:")
	os.Exit(0)
}
