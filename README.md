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

This code creates an array of items,
each item contains a string and a type.
Type value is one  of "delimiter" or "word".
The code walks the array with 2 indexes, one increasing from 0,
the other decreasing from the last index in the array.
When it finds 2 word-typed items, one at each index, it swaps them.
When the decreasing index is less than the increasing index it stops.

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
Does this solution work for the following cases:
"hello/world:here/", "hello//world:here"

I see what they're getting at: does the solution code
put a zero-length word between '/' and '/'
in the case of "hello//world:here"?
My solution does not.

"hello//world:here" should arguably transform to "here/world/hello:"

What should "here/world/hello:"  transform to?
My code transforms it to "hello/world/here:", because it does not assume
a zero-length word after the final ':' delimiter.
The follow-up case of "hello/world:here/" should prompt exactly that discussion.

## Interview Analysis

If this problem were tightened up by defining single- or multi-character delimiters,
it would be a decent problem to prompt discussion.
There's some coding involved,
with a chance to for an interviewer to see some reasoning about array indexes,
and some easy software engineering practices (symbolic constants, variable naming).
The candidate could distinguish themselves by pointing out opportunities
for good practices, and why they wrote the pieces they did.

"[Hard]" seems unreasonable for this as a programming problem.
The "follow-on" test cases reveal what the posers of this problem wanted:
discussion, probably about pitfalls of informal requirements,
but possibly also about the limits of invertable transformations or some such.
