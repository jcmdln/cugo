This namespace contains the frontend code for the command line interface
that drives `cugo`, which is cobra. If you import this namespace, you
will have access to all commands including their flags should you wish to
do so. This will inflate your binary considerably.

If you would like to import a single command you may do so from the `src`
directory. This allows you to import the raw code itself, and you may
choose to use your own flag parsing or hard-code the values needed to
satisfy the command.
