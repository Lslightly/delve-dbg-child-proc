vscode-go可能没有相应的支持，已搜索失败的路径
- [debugging · golang/vscode-go Wiki](https://github.com/golang/vscode-go/wiki/debugging)
- [debug in vscode - Reddit Search!](https://www.reddit.com/r/golang/search/?q=debug+in+vscode&cId=51c36998-7f0e-4ef3-b041-1c1050a83203&iId=013990b5-3dcc-424e-b4b9-0940bdd13278)
- [golang/vscode-go · Discussions · GitHub](https://github.com/golang/vscode-go/discussions)
- [Issues · golang/vscode-go](https://github.com/golang/vscode-go/issues)
- [Go in Visual Studio Code](https://code.visualstudio.com/docs/languages/go#_debugging)
- https://github.com/Microsoft/vscode-cpptools/issues/1211
一些提问：
- https://discord.com/channels/118456055842734083/1346141590011252766/1346141590011252766
- [Any support for `target follow-exec`? · golang/vscode-go · Discussion #3704](https://github.com/golang/vscode-go/discussions/3704)

搜索work list:
- vscode go debug child process
	- [Is it possible to debug a go program in vscode as a child/debugee of a visual studio debugger session? - Stack Overflow](https://stackoverflow.com/questions/79277132/is-it-possible-to-debug-a-go-program-in-vscode-as-a-child-debugee-of-a-visual-st)
- [Debugging spawned child processes with breakpoints when developing an extension : r/vscode](https://www.reddit.com/r/vscode/comments/xf3f2h/debugging_spawned_child_processes_with/) 
- [albertziegenhagel/childdebugger-vscode: VS Code extension to auto-attach a debugger to child processes for `cppvsdbg`.](https://github.com/albertziegenhagel/childdebugger-vscode?tab=readme-ov-file) 只在windows上对C++有子进程支持的debugger.
- [Debugging spawned child processes with breakpoints when developing an extension : r/vscode](https://www.reddit.com/r/vscode/comments/xf3f2h/debugging_spawned_child_processes_with/)
- [[Node] 如何使用 VSCode 调试 child_process - 简书](https://www.jianshu.com/p/b521e15dbab0)通过双开窗口以及attach方式实现子进程调试
- [Debugging spawned child processes with breakpoints when developing an extension : r/vscode](https://www.reddit.com/r/vscode/comments/xf3f2h/debugging_spawned_child_processes_with/)
- [Attach and debug any child process fork'ed by the target process · Issue #2551 · go-delve/delve](https://github.com/go-delve/delve/issues/2551)，里面出现了编译器开发者mdempsky。

现在的问题有两个：
- 虽然dlv命令行可以使用`target follow-exec -on`开启child process调试，但是vscode不可以。这是因为vscode无法输入其他功能导致的。使用dlv的`--init`选项配置也不行。
	- [ ] 这个问题可能解决不掉了
- 虽然dlv命令行可以开启child process调试，但是没法为child process设置断点。这个可能需要看一下其他选项，估计是可以的。比如说load source。
	- [c - How to set breakpoint in child process after set follow-fork mode child? - Stack Overflow](https://stackoverflow.com/questions/11003287/how-to-set-breakpoint-in-child-process-after-set-follow-fork-mode-child)
	- [Forks (Debugging with GDB)](https://sourceware.org/gdb/current/onlinedocs/gdb.html/Forks.html)

