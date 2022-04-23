// SPDX-License-Identifier: ISC

package basename

import (
	"testing"
)

func BenchmarkBasename(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Basename("/wew/lad", "")
	}
}

func TestBasename(t *testing.T) {
	result, err := Basename("/wew/lad", "ad")
	if err != nil {
		t.Errorf("%s", err)
	}

	if result != "l" {
		t.Errorf("got %s, expected l", result)
	}
}

func TestBasenameDuplicateSlashes(t *testing.T) {
	result, err := Basename("//wew//lad//", "ad")
	if err != nil {
		t.Errorf("%s", err)
	}

	if result != "l" {
		t.Errorf("got %s, expected l", result)
	}
}

func TestBasenameSuffixEmpty(t *testing.T) {
	result, err := Basename("//wew//lad//", "")
	if err != nil {
		t.Errorf("%s", err)
	}

	if result != "lad" {
		t.Errorf("got %s, expected lad", result)
	}
}

func TestBasenameSuffixInvalid(t *testing.T) {
	result, err := Basename("/wew/lad", "l")
	if err != nil {
		t.Errorf("%s", err)
	}

	if result != "lad" {
		t.Errorf("got %s, expected lad", result)
	}
}

func TestBasenameSuffixValid(t *testing.T) {
	result, err := Basename("/wew/lad/foo/bar", "r")
	if err != nil {
		t.Errorf("%s", err)
	}

	if result != "ba" {
		t.Errorf("got %s, expected /wew/lad", result)
	}
}
