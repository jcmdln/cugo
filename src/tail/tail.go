// SPDX-License-Identifier: ISC
//go:build testing

package tail

type Option struct {
	Count  int
	Follow bool
	Number int
}

func (opt *Option) Tail(operands []string) error {
	return nil
}
