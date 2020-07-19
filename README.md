# MHLS: My Humble Leetcode Solutions

Welcome to MHLS!

This repository is created for the purpose of documenting my code solutions that have been accepted by Leetcode.
However, not all of my accepted solutions will be documented here, only those with difficulty categorised as medium
or above will be collected into this repository. Having that said, problems that are easy but sufficiently interesting
will as so be documented here.

The solutions in MHLS are mainly written in C++ and Go. However, solutions in other languages such as Rust, Java,
Javascript, Python, etc might be added in the future. Solutions implemented in different languages are placed in
their respectively directory. C++ solutions can be found in the cpp directory and Go solutions are in the go directory.

## Building C++ Solutions
Each subdirectory in the cpp directory is a solution to a particular Leetcode problem. Each of these subdirectory contains
a simple Makefile to build a test. The test binary will be output to the **build** directory.

For example to build and run the test for the **sametree** Leetcode problem, run

```sh
cd sametree
make
build/sametree_test
```

To clean the directory, run

```sh
cd sametree
make clean
```

## Building Go Solutions
Like the C++ solutions, each Go subdirectory contains the code solution. However, to build and run the test
we only need to run the go CLI in the solution directory.

To build and run the test for **solveq**, run

```sh
cd solveq
go test
```
To run benchmark test for **solveq**, run
```sh
cd solveq
go test -bench . -benchmem
```

## What are sametree and solveq?
**sametree** is the code solution for the [Same Tree](https://leetcode.com/problems/same-tree) problem on Leetcode and **solveq**
is for [Solve the Equation](https://leetcode.com/problems/solve-the-equation). I use short directory names for all the code
solutions and it's especially needed for Go because these directories are the Go packages. The mappings for these directories
to their Leetcode problems can be found in the [Wiki](https://github.com/dt3zr/leetcode/wiki) area of the repository.
