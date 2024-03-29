// SPDX-License-Identifier: ISC

// delay for a specified amount of time.
//
// SYNOPSIS
//
//    sleep DURATION ...
//
// DESCRIPTION
//
// Sleep suspends execution for a minimum of the specified duration of
// time.  By default, sleep uses seconds as it's unit of time though
// supports seconds (s), minutes (m), and hours(m) as characters at the
// end of the provided interval.  This number must be positive and may
// contain a decimal.
//
// SEE ALSO
//
//     https://golang.org/pkg/time/#Sleep
//
// REFERENCES
//
//     https://pubs.opengroup.org/onlinepubs/9699919799/utilities/sleep.html
//
package sleep
