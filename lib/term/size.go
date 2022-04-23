// SPDX-License-Identifier: ISC

package term

import (
	"golang.org/x/sys/unix"
)

func Size(fd int) (width, height int, err error) {
	ws, err := unix.IoctlGetWinsize(fd, unix.TIOCGWINSZ)
	if err != nil {
		return -1, -1, err
	}
	return int(ws.Col), int(ws.Row), nil
}
