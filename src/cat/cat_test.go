// SPDX-License-Identifier: ISC

package cat

import (
	"io"
	"os"
	"strings"
	"testing"
)

func BenchmarkCat(b *testing.B) {
	var err error

	u := &Option{}
	files := []*os.File{os.Stdin}

	for i := 0; i < b.N; i++ {
		if err = u.Cat(files); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkCatUnbuffered(b *testing.B) {
	var err error

	u := &Option{}
	u.Unbuffered = true

	files := []*os.File{os.Stdin}

	for i := 0; i < b.N; i++ {
		if err = u.Cat(files); err != nil {
			b.Error(err)
		}
	}
}

func TestCat(t *testing.T) {
	var (
		err  error
		file *os.File
	)

	u := &Option{}

	if file, err = os.CreateTemp(os.TempDir(), ""); err != nil {
		t.Error(err)
	}
	defer os.Remove(file.Name())

	if _, err = io.Copy(file, strings.NewReader("wewlad")); err != nil {
		t.Error(err)
	}

	files := []*os.File{file}
	if err = u.Cat(files); err != nil {
		t.Error(err)
	}
}
