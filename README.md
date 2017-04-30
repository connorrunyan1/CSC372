# Go: A Language Exploration
* Authors: Connor Runyan & Jordan Siaha<br>
* Compiler Used: [golang.org](https://golang.org/dl/)

## History
Go started originally as a personal side-project of a group of Google Engineers named Robert Griesemer, Ken Thompson, and Rob Pike.  They first began thinking about seriously about creating a new language in September 2007, and they had some working compilers by January 2008.  In November 2009 Go became open source, and all continued development since then has been managed by a group called the _Go Programming Language Project_.

## Paradigm
As far as a language paradigm is concerned, it is difficult to pin Go down.  It sits somewhere between Object Oriented and good 'ol Imperative, probably closer to the Object Oriented side.  While it has many of the conveniences and tools that we expect from an OO language, it is missing a few.  Namely, type inheritance.  In Go, types/classes don't inherit from eachother.
```
type car interface {
  isParked() bool
  drive() float64
```
## Typing System
Go is a strongly typed language, and type declarations are required.
Typing System – is the language strongly or weakly typed, are type declarations required, can a
programmer create new types, and are functions first-class objects?

## Control Structures
Control Structures – how can control flow be controlled? (That is, what are the languages selection
and repetition options?)

## Support For Data Abstractions
Support for Data Abstractions – which abstractions are provided, how can a programmer create
new ones?

## Syntax
Syntax – think about the syntax used by your language. What are the syntax choices that appeal
to you? Are there any syntactic choices that you’d like to see changed (why, and to what)?

## Semantics
Semantics – briefly explain how your language is scoped (static or dynamic?), which kinds of constants
are supported, how storage is allocated (which combination of static, stack-dynamic, and/or
heap-dynamic?), and how garbage is managed.

## Desirable Language Characteristics
Desirable Language Characteristics – In Topic 2, we covered four categories of language characteristics
that are generally thought of as ‘desirable’: (i) Efficiency, (ii) Regularity, (iii) Security/Reliability,
and (iv) Extensibility. Choose any three of these four, and discuss features of your
language that support (or limit!) them.

## Example Program
#### Un-Weighted _k_ Nearest Neighbors
```
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
