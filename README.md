# Go: A Language Exploration
* Authors: Connor Runyan & Jordan Siaha<br>
* Compiler Used: [golang.org](https://golang.org/dl/)

## History
Go started originally as a personal side-project of a group of Google Engineers named Robert Griesemer, Ken Thompson, and Rob Pike.  They first began thinking about seriously about creating a new language in September 2007, and they had some working compilers by January 2008.  In November 2009 Go became open source, and all continued development since then has been managed by a group called the _Go Programming Language Project_.

## Paradigm
Go is an imperative language.  Since it was written with low-level systems programming in mind, it looks and feels a lot like C from a distance, so much so that it is often viewed as what C would have looked like if it was written in the current century.  While Go does pilfer many tools and ideas from other language paradigms, it is most definitely still an imperative language.

## Typing System
Go is a strongly typed language, and type declarations are required.  Here are a couple of examples.  Note that the type of a variable is given after its name.
```Go
var radius float64   // declare a variable named radius of type float64
var width, height int = 50, 60   // allows multiple declarations on one line like this
```
Go also allows a programmer to use the following shorthand variable declaration/initialization operator, which is smart enough to figure out the type on it's own if it can.
```Go
message := "stringy literal"   // since this is a string literal, Go knows I mean:
                               // var message string = "stringy literal"
```
A programmer can also create new types.  This is in contrast to C, where the closest you could get to a new type would have been just declaring a struct.  Go extends that idea by letting a function definition specifiy what is called a _reciver_, in other words, the type that a function should operate on. Here is an example type definition, and then two different ways to write a function that makes use of it.
```Go
type Point struct {
  X, Y float64
}

// If you were in C, this sort of function would be your only option.
// It has to take the point you want as an argument.
func sum(p Point) float64 {
  return p.X + p.Y
}
// To invoke the above:
p := Point{5,6}
sum(p) // returns 11

// In Go, you can specifically bind a function to a type, called the reciever.
// This makes the function into a method, with a specific type it belongs to.
func (p Point) sum() float64 {
  return p.X + p.Y
}
// To invoke the above:
p := Point{5,6}
p.sum() // returns 11
```
Go also allows a programmer to treat functions as first class objects (or in this case, first class types).  They work like any other value.  Here is an example:
```Go
func printHelloWorld() {
  fmt.Println("Hello, World!")
}

v := printHelloWorld
v()
```
The output of that above snippet would be "Hello, World!" printed out to console.  In Go, a function is like any other value.

## Control Structures
Go provides a fairly standard suite of control structures.  Naturally, we have the classic for loop.  While generally not interesting, it shows us some proof that programmers will still need their semicolon keys.
```Go
for i:= 0; i < 10; i++ {
  sum += i
}
```
The classic if/else/else if branching structure works like one would expect, as does the oft-overlooked switch statement.  Here is an example of those two:
```Go
if 5*5 == 25 {
  // this code is skipped if 5*5 isn't 25.
} else if 4 * 4 == 16 {
  // else if chaining works just like in C.
} else {
  // final catch-all else statement.
}

h := 7
switch h {
case 5:
  // if h had the value 5
case 6:
  // if h had the value 6
case 7,8:
  // Go lets you list multiple possibilities separated by commas
}
```
Go does not have a separate while keyword.  A for with a blank condition will loop eternally, so you'll have to include your own escape condition.  Go also has goto statements and labels, for some reason.  One interesting addition is the _defer_ statement.  A statement with the modifier _defer_ on the front will evaluate any needed arguments immediately, but not actually execute until the function it is in finished executing first.  The following function will output "first second" to console.
```Go
func printStuff() {
  defer fmt.Print("second")
  fmt.Print("first ")
}
```

## Support For Data Abstractions // TODO
As far as data abstractions, Go has examples of all 3 data abstraction types(Basic, Structured, Unit).  On the basic level, it has everything you would expect out of traditional variables.  You can manipulate data like integers or strings without having to deal with the individual bits and bytes directly.  For structured data abstractions, it has a similar suite as C. // TODO check textbook
Support for Data Abstractions – which abstractions are provided, how can a programmer create
new ones?

## Syntax
Go has an interesting syntax, since it mixes a lot a C style with the "backwards" declarations from the Pascal family of languages.  Personally, we're not huge fans of the 'backwards' typing.  While this is likely due to the history we have with C-family languages, we also think that the type of a variable can often be more important than its name, and should come first as such.

## Semantics
Go is a statically scoped language.  As many systems languages do, Go prioritizes execution speed and does everything it can at compile time to reach that goal.  Go supports only a few types of constants: characters, strings, booleans, and numeric values.  The key here is that the idea of 'numeric values' is actually pretty smart in Go.  While compiling, it doesn't make the numeric constant into a specific type until it hits an instruction that would force it to.  By doing this, the value will only gain a type when a programmer uses it.  Storage allocation (static vs stack vs heap) is unfortunately inplementation specific.  The general rule, however, is that the compiler will do what is called _escape analysis_ to try and decide if the entity can live on the stack or if it has to be allocated on the heap.  The general rule is that variables that can live on the stack will, with the heap being a backup.  Heap allocation is used for most of the reference types, since their usage might extend outside of the function they were created in.  Go is fully garbage collected.

## Desirable Language Characteristics
Desirable Language Characteristics – In Topic 2, we covered four categories of language characteristics
that are generally thought of as ‘desirable’: (i) Efficiency, (ii) Regularity, (iii) Security/Reliability,
and (iv) Extensibility. Choose any three of these four, and discuss features of your
language that support (or limit!) them.


## Example Program
#### Un-Weighted _k_ Nearest Neighbors
```Java
if(code.isWritten()){
  this.replace(code);
} else {
  code.write();
}
```
## References
* [The Go Language Project: Frequently Asked Questions](https://golang.org/doc/faq)
* [Go at Google: Language Design in the Service of Software Engineering](https://talks.golang.org/2012/splash.article)
* Reference 3
