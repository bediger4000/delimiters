# Daily Coding Problem: Problem #1044 [Hard] 

This problem was asked by Facebook.

Given a string and a set of delimiters,
reverse the words in the string while maintaining the relative order
of the delimiters.
For example, given "hello/world:here", return "here/world:hello"

Follow-up: Does your solution work for the following cases:
"hello/world:here/", "hello//world:here"

## Analysis

The problem statement claims one thing ("given a string and a set of delimiters"),
but doesn't follow through in the examples, which include no given set of delimiters,
but rather let you pick up the delimiters from context.
The statement also doesn't define "delimiter" at all well.
Is a delimiter a single character, like ':' or '/'?
Can it be a short string, like "//"?
The example implies that single characters are delimiters.
I'm going to assume single-character delimiters.

### My Solution

[Code](main.go)

```sh
$ go build .
$ ./delimiters '/:' 'hello/world:here'
Original:
hello/world:here
Words reversed:
here/world:hello
$ ./delimiters ':/' "hello/world:here/"
Original:
hello/world:here/
Words reversed:
here/world:hello/
$ ./delimiters ':/' "hello//world:here"
Original:
hello//world:here/
Words reversed:
here//world:hello/
```
