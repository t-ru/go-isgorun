/*
MIT License

Copyright (c) 2022 Thomas Rudolph

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"fmt"
	"os"

	"github.com/t-ru/go-utils/pkg/procinfo"
)

func main() {

	fmt.Println()
	fmt.Println("There is no easy way to detect if an executable was startet by \"go run\"")
	fmt.Println()

	fmt.Println("All solutions I've found:")
	fmt.Println()
	fmt.Println("   - are not multiplatform")
	fmt.Println("   - sometimes they work on Linux")
	fmt.Println("   - sometimes they work on Windows")
	fmt.Println("   - ...")
	fmt.Println()

	fmt.Println("If an executable was called:")
	fmt.Println()
	fmt.Println("    - directly the parent process IS NOT \"go run\"")
	fmt.Println("    - by \"go run\" the parent process IS \"go run\"")
	fmt.Println()

	fmt.Println("This program demonstrates the safe detection by the parent process.")
	fmt.Println()

	fmt.Println("Testcases:")
	fmt.Println()
	fmt.Println("    - 1: run this program with \"go run\".")
	fmt.Println("    - 2: build this program with \"go build\" and execute it.")
	fmt.Println()

	if isGoRun() {
		fmt.Println("Executed by \"go run\"? YES")
	} else {
		fmt.Println("Executed by \"go run\"? NO")
	}

	fmt.Println()

}

func isGoRun() bool {
	pi, _ := procinfo.New(os.Getpid())

	return pi.ParentProcessIsGoRun()
}
