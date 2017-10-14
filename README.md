# Aurora

[![Coverage Status](https://coveralls.io/repos/github/Itay2805/Aurora/badge.svg?branch=master)](https://coveralls.io/github/Itay2805/Aurora?branch=master)

This is a compiler for the Aurora language.

Right now the compiler is written in Go.

# Build
To build you must have go installed

1. Clone the contents of the repo to the `src` folder of your Go workspace
2. Run `go-get.cmd` to get all the libraries the compiler uses
3. Run `build-compiler.cmd` and it will generate an exe file called `compiler.exe`

That is it, run the compiler with the first argument being a file to compile

# Editing the Grammars file

If you use VSCode and want to edit the ANTLR4 file, please use the ANTLR4 plugin with the following workplace config
```JSON
{
    "antlr4.generation.language": "Go",
    "antlr4.generation.mode": "external"
}
```

The following will set ANTLR to generate Go code instead of the default language (Java)
