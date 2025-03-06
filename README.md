# Delve Debug Child Process

Reproduce that setting breakpoints in child process does not work.

Steps to reproduce:

```bash
./make.bash # to build driver and compile
./debug.sh # use delve-init.txt to debug switch target and debug child process
s # after executing commands in delve-init.txt, press s to switch to the child process
# `locals` in delve-init.txt will show the value of local variables `a` and `b`
# but set breakpoints will give the misleading error
#   Command failed: location ... not found
#   Set a suspended breakpoint (Delve will try to set this breakpoint when a plugin is loaded) [Y/n]
b main.compile # give misleading error but actually set yes can set the breakpoint correctly
b build.go:9 # give misleading error but actually set yes can set the breakpoint correctly
funcs # `main.compile` defined in utils/build.go is detected in `funcs` command. But unable to set breakpoints
```

Video:

[![asciicast](https://asciinema.org/a/706784.svg)](https://asciinema.org/a/706784)

# vscode-go debug multiple process

TBD
