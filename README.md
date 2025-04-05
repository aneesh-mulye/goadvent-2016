# Advent of Code 2016, in Go to learn Go

## Status

- [X] Day  1
    - [X] Part 1
    - [X] Part 2
- [X] Day  2
    - [X] Part 1
    - [X] Part 2
- [X] Day  3
    - [X] Part 1
    - [X] Part 2
- [ ] Day  4
    - [ ] Part 1
    - [ ] Part 2
- [ ] Day  5
    - [ ] Part 1
    - [ ] Part 2
- [ ] Day  6
    - [ ] Part 1
    - [ ] Part 2
- [ ] Day  7
    - [ ] Part 1
    - [ ] Part 2
- [ ] Day  8
    - [ ] Part 1
    - [ ] Part 2
- [ ] Day  9
    - [ ] Part 1
    - [ ] Part 2
- [ ] Day  10
    - [ ] Part 1
    - [ ] Part 2
- [ ] Day  11
    - [ ] Part 1
    - [ ] Part 2
- [ ] Day  12
    - [ ] Part 1
    - [ ] Part 2
- [ ] Day  13
    - [ ] Part 1
    - [ ] Part 2
- [ ] Day  14
    - [ ] Part 1
    - [ ] Part 2
- [ ] Day  15
    - [ ] Part 1
    - [ ] Part 2
- [ ] Day  16
    - [ ] Part 1
    - [ ] Part 2
- [ ] Day  17
    - [ ] Part 1
    - [ ] Part 2
- [ ] Day  18
    - [ ] Part 1
    - [ ] Part 2
- [ ] Day  19
    - [ ] Part 1
    - [ ] Part 2
- [ ] Day  20
    - [ ] Part 1
    - [ ] Part 2
- [ ] Day  21
    - [ ] Part 1
    - [ ] Part 2
- [ ] Day  22
    - [ ] Part 1
    - [ ] Part 2
- [ ] Day  23
    - [ ] Part 1
    - [ ] Part 2
- [ ] Day  24
    - [ ] Part 1
    - [ ] Part 2
- [ ] Day  25
    - [ ] Part 1
    - [ ] Part 2
- [ ] Day  26
    - [ ] Part 1
    - [ ] Part 2
- [ ] Day  27
    - [ ] Part 1
    - [ ] Part 2
- [ ] Day  28
    - [ ] Part 1
    - [ ] Part 2
- [ ] Day  29
    - [ ] Part 1
    - [ ] Part 2
- [ ] Day  30
    - [ ] Part 1
    - [ ] Part 2

## Notes on specific days

### Day 1

Kinda surprised that there isn't a function in the standard library to read a
text by line into a slice of strings, with the newlines stripped out.

The lower level machinery is good; still, slightly surprising.

TODO: any existing utility libraries that provide all these things one
sorta-kinda expects from a language flirting with dynamism, implied by ∃GC?

### Day 2

Complex numbers turning out surprisingly useful here.

As a friend said, a language with a *proper* numeric stack has not yet been
made.

### Day 3

TIL that Go has accreted *two* sets of functions for sorting (and almost
certainly other things as well): one from the Before \[T\]imes, when you didn't
have generics and had to have different sort functions for each type, such as
`sort.Ints`, and another from after generics, from the `slices` package, which
contains a generic `slices.Sort` that operates on any ordered type.

I'm increasingly tempted to use `https://github.com/samber/lo`, or something
like it, for solving these problems. It's just so much more congenial to think
of them in terms of the functional sequence transforms rather than basically
building up a variant of that machinery every single time for every problem
using manual loops and whatnot.

Given that I want to learn it, may just do that once I've solved enough problems
that I'm confident in my ability to work idiomatically within the language.
