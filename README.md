# Aurora

This is a compiler for the Aurora language.

Right now the compiler is written in Go.

# Language Features
Currently, the compiler can only parse and print function declarations

But, in the future the language will support:
1. Both GC Managed memory and User Managed memory, A std library which does *not* use the GC, so if you don't use it aswell, you can just disable it.
2. Classes
3. Null Safety, you must initialize an object and you can not initialize it to `null` (it is possible to set a variable nullable, but by default a variable can not be null)
3. Smart pointers, for helping managing normal pointers
4. `unsage` blocks, for performing pointer math on GC pointers.
5. Operator overloading, including custom operators (?)

much more :)

# Build
To build you must have go installed

1. Create in your workspace under `src` and folder named `Aurora`
2. Clone the contents of the repo to that folder (so you will have `%GOPATH%/src/Aurora/Compiler` and so on...)
3. Run `go-get.cmd` to get all the libraries the compiler uses
4. Run `build-compiler.cmd` and it will generate an exe file called `compiler.exe`

That is it, run the compiler with the first argument being a file to compile
