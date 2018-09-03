package rm

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var (
	Dir         bool
	Force       bool
	Interactive bool
	Recursive   bool
	Verbose     bool
)

func empty(name string) bool {
	t, err := os.Open(name)
	defer t.Close()
	_, err = t.Readdirnames(1)
	if err == io.EOF {
		return true
	}

	return false
}

func prompt(text string) bool {
	if Interactive {
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

func verbose(tgt string) {
	if Verbose {
		fmt.Println("cugo: rm: removed", tgt)
	}
}

func remove(t string) {
	if Force {
		os.Remove(t)
		verbose(t)
	} else if prompt("Remove '" + t + "'?") {
		os.Remove(t)
		verbose(t)
	}
}

func Rm(args []string) {
	for _, target := range args {
		t, err := os.Stat(target)
		if os.IsNotExist(err) {
			fmt.Println("cugo: rm: Can't remove", "'"+target+"':",
				"no such file or directory")
			return
		}

		if t.IsDir() {
			if Dir && empty(target) {
				remove(target)
				return
			}

			if Recursive {
				for !empty(target) {
					filepath.Walk(target,
						func(t string, info os.FileInfo, err error) error {
							if info.IsDir() && empty(t) {
								remove(t)
							}
							if !info.IsDir() {
								remove(t)
							}
							return nil
						},
					)
				}

				if empty(target) {
					remove(target)
				}
			} else {
				fmt.Println("cugo: rm: Can't remove directory '" +
					target + "'")
				return
			}
		} else {
			remove(target)
		}
	}
}
