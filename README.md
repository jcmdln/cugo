Note: This is a for-fun project that may change rapidly, though is intended to implement an important part of any system: it's core utilities. Please use Cugo with caution until introspective tests have been included to ensure correctness.


## About
Cugo is a set of core utilities for Unix-like systems in the form of a multi-call binary with the aim of being standards compliant. Much of the direction of Cugo will take inspiration from Rob Landley's project Toybox such as choice of reference `man` pages and utilities to implement, though special consideration of any missing features will be made when comparing `man` pages of various implementations of utilities.


## Design Methodology
All utilities should be simply correct, meaning that the contained implementations should focus on being written in a simple, readable manner and be reviewed for correctness. Additional focus on debugging and performance optimization is still in progress, though should not obfuscate what is occurring arbitrarily.

Go's standard library is feature complete enough to make many utilities trivial to implement and should be used as often as possible in place of third-party applications or libraries.


## Standards Compliance
While the aim of Cugo is to be standards compliant, there are many features included with utilities that are pointless in the context of a multi-call binary such as printing versions of each utility. Flags that serve a ritualistic rather than pragmatic use-case will not be included.

Example of ritualism: `rm` in the GNU specification has `--interactive=INTERVAL` for specifying how many times a user must say yes or no before it then proceeds with the rest of the provided input. Having these types of flags where a user is meant to be prompted X number of times before it proceeds are conditions that should have remained in the scope of the user (or developer) to instead handle this "problem" within their own application. Is this easy to implement? Well, yes but this is an example of the sort of flags that will not be present in Cugo.


## TODO
- Better integration with Delve to ensure correctness via introspection.
- Couple Delve's hooks with automated tests whenever it makes sense to do so.
- A ton of utilities. Too many to even list at this point and there are many that probably fall outside the scope of Cugo. A formal list will be made at a later date once more of the common utilities have been completed.
