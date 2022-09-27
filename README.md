# go-isgorun

## Description

There is no easy way to detect if an executable was startet by \"go run\"<br>

All solutions I've found:
- are not multiplatform
- sometimes they work on Linux
- sometimes they work on Windows
- ...<br>

If an executable was called:
- directly the parent process IS NOT \"go run\"
- by \"go run\" the parent process IS \"go run\"

This program demonstrates the safe detection by the parent process.<br>
It uses https://github.com/t-ru/go-utils/tree/master/pkg/procinfo

Testcases:
- 1: run this program with \"go run\".
- 2: build this program with \"go build\" and execute it.


