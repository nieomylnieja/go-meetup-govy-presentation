---
theme: ./theme.json
author: Mateusz Hawrus
date: February 13, 2025
paging: Slide %d / %d
---

# How I learned to stop worrying and love generics in Go

Exploring validation API solutions in modern Go.

---

## Who am I?

👨 💼

### Where can you find me?

1. https://github.com/nieomylnieja/
2. https://www.linkedin.com/in/mateusz-hawrus-038463174/

---

## Legend

1. The basics, how would one go about writing validation code.
2. Getting fancier, previous solutions.
3. Enter generics.
4. Beyond type safety. What govy brings to the table?

---

## Simple checks

```go
~~~./code-to-slide.awk code/01-simple/main.go

~~~
```

---

## Simple checks with aggregate

```go
~~~./code-to-slide.awk code/02-simple-aggregate/main.go

~~~
```

---

## Simple checks with slice of students

```go
~~~./code-to-slide.awk code/02-simple-aggregate/main.go

~~~
```

---

## Simple approach summary

Things we could've done to improve it:

1. DRY, extract common functions for validating specific fields out.
2. 

Summary:

1. Works perfectly fine with simple structures.
2. Keeping consistent messages across multiple APIs can become a challenge.
3. What if our models' names change? Example: `Teacher.Name` --> `Teacher.FirstName`. Can we make sure messages reflect this change?
4. 

---

## Wrapping up

The presentation can be found at https://github.com/nieomylnieja/go-meetup-govy-presentation.
