package rm

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type RM struct {
	Dir         bool
	Force       bool
	Interactive bool
	Recursive   bool
	Verbose     bool
}

func Rm(args []string) {
	rm := &RM{}

	Empty := func(name string) bool {
		t, err := os.Open(name)
		defer t.Close()
		_, err = t.Readdirnames(1)
		if err == io.EOF {
			return true
		} else {
			return false
		}
	}

	Prompt := func(text string) bool {
		if rm.Interactive {
			fmt.Printf(text + " [Yes/No]: ")
			var a string
			_, err := fmt.Scan(&a)
			if err != nil {
				fmt.Println(err)
				return false
			}
			if a == "y" || a == "Y" || a == "yes" || a == "Yes" {
				return true
			}
			return false
		} else {
			return true
		}
	}

	Verbose := func(tgt string) {
		if rm.Verbose {
			fmt.Println("cugo: rm: removed", tgt)
		}
	}

	Remove := func(t string) {
		if rm.Force {
			os.Remove(t)
			Verbose(t)
		} else if Prompt("Remove '" + t + "'?") {
			os.Remove(t)
			Verbose(t)
		}
	}

	for _, target := range args {
		t, err := os.Stat(target)
		if os.IsNotExist(err) {
			fmt.Println("cugo: rm: Can't remove", "'"+target+"':",
				"no such file or directory")
			return
		}

		if t.IsDir() {
			if rm.Dir && Empty(target) == true {
				Remove(target)
				return
			}

			if rm.Recursive {
				for Empty(target) == false {
					filepath.Walk(target,
						func(t string, info os.FileInfo, err error) error {
							if info.IsDir() && Empty(t) == true {
								Remove(t)
							}
							if !info.IsDir() {
								Remove(t)
							}
							return nil
						},
					)
				}

				if Empty(target) == true {
					Remove(target)
				}
			} else {
				fmt.Println("cugo: rm: Can't remove directory '" +
					target + "'")
				return
			}
		} else {
			Remove(target)
		}
	}
}
