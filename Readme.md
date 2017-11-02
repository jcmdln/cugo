Note: This is a for-fun project that may change rapidly, though is intented to implement an important part of any system: it's core utilities. Please use Cugo with caution until introspective tests have been included to ensure correctness.


## About
Cugo is a set of core utilities for Unix-like systems in the form of a multi-call binary with the aim of being standards compliant without cruft or bloat. Much of the direction of Cugo will take inspiration from Rob Landley's project Toybox such as choice of reference `man` pages and utilities to implement, though special consideration of any missing features will be made when comparing `man` pages of various implementations of utilities.

When writing the core utilities only the Go standard library should be used as often as possible rather than importing away the promblem. The only external import that is currently in use by Cugo is for SPF13's Cobra, though this provides the frontend for the core utilities rather than implementing some part of them. Go already implements much of the functionality needed to build a large majority of the planned core utilities. In this way we may ensure that when the Go standard library updates, especially when part of some critical security flaw, that no external dependencies will need to be updated or would potentially cause it's own issues.

While the aim of Cugo is to be standards compliant, there are many features included with utilities that are outright useless in this context such as printing versions or how `rm` in the GNU specification has `-i`, `-I`, and `--interactive=INTERVAL` for specifiying how many times a user must say yes or no before it then proceeds with the rest of the provided input. Having these types of flags where a user is meant to be prompted X number of times before it proceeds are conditions that should have remained in the scope of the user (or developer) to instead handle this "problem" with shell scripting. Is this easy to implement? Well, yes but this is an example of the flags that will not be present in Cugo as they serve a ritualistic rather than pragmatic use-case.


## TODO
- Better integration with Delve to ensure correctness via introspection.
- Couple Delve's hooks with automated tests whenever it makes sense to do so.
- A ton of utilities. Too many to even list at this point and there are many that probably fall outside the scope of Cugo. A formal list will be made at a later date once more of the common utilities have been completed.
