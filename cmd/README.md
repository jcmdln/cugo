This namespace contains the frontend code for the command line interface
that drives `cugo`, which is `flagger`. If you import this namespace, you
will have access to all commands including their command line interface,
should you wish to do so. This will inflate your binary considerably and
is not an intended use-case.

If you would like to import a single utility, you may do so from the
`src` directory as outlined in
[README](https://github.com/jcmdln/cugo/blob/master/src/README.md)
