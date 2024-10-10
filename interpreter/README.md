# Interpreter built in GO for monkey-lang*

## What is an interpreter...

## And why do we need one?

****
# Steps
*Source code -> Tokenization -> Abstract syntax -> Evaluation*
## 0. Source code

## 1. Tokenization & Lexing

## 2. Token to Tree

## 3. Evaluation
****
# The monkey language

```
// Monkey supports variable declarations
let age = 25;
let name = "Jonathan";

// Simple math and functions
let add = fn(a, b) { 
    return a + b;
};

let result = add(5, 10);
puts(result);  // Outputs: 15

// Conditionals
if (age > 18) {
    puts("You are an adult");
} else {
    puts("You are a minor");
}

// While- and for loops
while (age > 10) {
    puts(age)
}

for (i = 0; i < 15; i = i + 1) {
    puts(i)
}

// First-class functions and closures
let outer = fn(x) {
    let inner = fn(y) {
        return x + y;
    };
    return inner;
};

let adder = outer(10);
puts(adder(5));  // Outputs: 15

```