# Language Syntax
this file contains the syntax reference for the language

## variable
```
var <var_name> : [<type>] [= <expression>];
```
	
## constant
```
const [export] <var_name> : [<type>] : <expression>;
```

if the export is defined then the constant will be exported and can be accessed
from the outside
	
## function

a function can be defined inside a module or part of a class
when it is a part of a class it will be a static method, meaning
it is not a method and does not require you to use it on
an object, but it can not access `this`

```
func [export] [native] ([<arg_name>: <type>...]) [-> <type>] { ... }
```

if the native is defined then there will be no need for function body,
it will search for the implementation as a C external function
	
if the export is defined then the function will be exported and can be accessed
from the outside
	
## module

a module is a container for functions and classes

```
module <module_name> { ... }
```

# expressions

## expression precedence

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
<function_name>([[<name>=]<value>...])
```

args can be named, so you can do either `new(Object)` or `new(class=Object)` for example
 
# example
```
// everything under module lang is staticly imported by default
module lang {
	
	// Class and Object are already defined in AuroraBootstrap.h
	// since they can not be represented using normal stuff
	// the compiler just knows them by default
	
	func export native new(class: Class&) -> Object*
	func export native gcnew(class: Class&) -> Object^
	func export native delete(obj: Object&)
	
	func export native ClassOf(object: Object&) -> Class*
	
	func export SizeOf(object: Any&) -> Int {
		var class := ClassOf(object);
		return class.size;
	}
	
}
```
