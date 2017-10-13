# Language Syntax
this file contains the syntax reference for the language

## Name (Identifier)

regex = `[a-zA-Z_][a-zA-Z0-9_]*`

these are used to identify variables, constants functions and more

## Optionals
the language has built in support for optionals.
by default, you MUST initialize a value, either in the constructor of a class or right away, otherwise it will not compile.
when you define a variable as optional, if it is not initialized right away or in the constructor, it will initialize to a `nil`

you can check an optional against a nil to find if the optional is `some` or `none`

# Literals

## Integers

regex = `[0-9_]+`

integers are whole numbers, like 10, -56, 1546, -79

you can also use `_` for reading the number easily
like `10_000_000` instead of `10000000`

integer literals will return a primitive int

## Strings

regex = `"[^"]"`

*note that strings can be multilined

strings literals will return a string class

# Expressions

## Expression Precedence

| Precedence | Operator | Description |
|:----------:|------------------------------|--------------------------------------------------------------------------------|
| 1 | ()<br>[]<br>. | Function call<br> Array subscripting<br> member access | Left-to-right |
| 2 | + -<br> ! <br> * <br> & <br> | Unary plus and minus<br>Logical NOT<br>Indirection (dereference)<br>Address-of |
| 3 | * / % | multiplication, division, and remainder |
| 4 | + - | Addition and subtraction |
| 5 | && | Logical AND |
| 6 | \|\| | Logical OR |
| 7 | = | Simple assignment |

### Function call
```
<function_name>([<expression>[, <expression>]*])
``` 

# Statements

## variable
```
var [optional] <name> : [<type>] [= <expression>];
```

* if nullable is defined then the variable can be left uninited, otherwise you have to initialize it.
	
## constant
```
const [export] <name> : [<type>] : <expression>;
```

if the export is defined then the constant will be exported and can be accessed
from the outside

# Functions

a function can be defined inside a module or part of a class
when it is a part of a class it will be a class method, meaning
it can only be accessed when from the class instance

```
func [export] [native] <name>([[optional] [<name>=]<expression>[, [optional] [<name>=]<expression>]*]) [-> <type>] { <statment>* }
```

* if the native is defined then there will be no need for function body, it will search for the implementation as a C external function
* if the export is defined then the function will be exported and can be accessed from the outside
* if the method getts overrided by the a derived class, the function will be dynamically linked, otherwise it will be staticly linked
* optional argument can either not be provided at all (and then it will be just a `nil`)

# Modules

a module is a container for functions and classes

```
module <name> { <statment>* }
```

# Type definitions

primitive types are already pre-defined
these include i8, i16, i32, i64
unsigned version of these also exists as u8, u16, u32, u64
there is also int and uint (size is based on OS)

## Class

this allows you to define a class

```
class [export] <name> [extends <super>] { <statment>* }
```

**Example**<br>
```
class export SomeClass {
	var someInt : int = 0;
	
	func export init(someInt: int) {
		this.someInt = someInt;
	}
	
	func export getSomeInt() -> int {
		return this.someInt;
	}
}
```

## Extend

extend allows to extend an existing type and add more features to it like new functions.
these functions can not be overrided since they are staticly linked

```
extend <name> { <statement>* }
```

**Examples**<br>
add function to int
```
extend int {
	func max() -> int64 {
		// calculate the max based on size
		return maxNum
	}
}
```

add function to a class
```
class SomeClass {
	// ...
}

extend SomeClass {
	func saveToFile(path: string) {
		// save to file
	}
}
```
