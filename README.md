# Unhandled Errors in Go Program

This repository demonstrates a common error in Go programs: failing to handle errors returned by functions.  The `bug.go` file contains the buggy code, while `bugSolution.go` shows how to properly handle errors.

The original program simply prints a message and ignores potential errors.  This can lead to silent failures, making debugging difficult.

The solution demonstrates the correct approach: checking the returned error and handling it gracefully, either by logging the error, recovering from it, or terminating the program with an appropriate error message.