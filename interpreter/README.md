# Interpreter Built in Go for Monkey-Lang*

## What is an Interpreter?

An interpreter is a program that reads and executes code written in a specific programming language. Unlike a compiler, which translates the entire source code into machine code before execution, an interpreter processes the code line by line, allowing for immediate feedback and flexibility.

---

## Why Do We Need One?

Interpreters allow developers to:
- Experiment with new languages or features.
- Provide immediate execution of code for scripting or learning purposes.
- Build tools to understand the structure and semantics of code.

---

# Steps to Build the Interpreter
The interpretation process follows a structured pipeline:

**Source Code → Tokenization → Abstract Syntax Tree → Evaluation**

## **0. Source Code**
The source code is the raw input written by the programmer. For example:
```javascript
let x = 10;
if (x > 5) {
    puts("x is greater than 5");
}
```

Essentially, each little piece of variable, keyword and syntax belongs to some token type. Below you find the different tokens used in this language:


## 1. Tokenization (Lexing)

**Tokenization** (or **lexing**) is the first step in interpreting the source code. It involves breaking the raw input into meaningful symbols, called tokens. These tokens represent the smallest units of the language, such as keywords, identifiers, literals, and operators.

For example, given the following Monkey code:
```javascript
let x = 10;
if (x > 5) { 
    puts("x is greater than 5");
}
```
The tokenization process would produce:
```
[LET, IDENT(x), ASSIGN, INT(10), SEMICOLON, IF, LPAREN, IDENT(x), GT, INT(5), RPAREN, LBRACE, IDENT(puts), LPAREN, STRING("x is greater than 5"), RPAREN, SEMICOLON, RBRACE]
```

## 2. AST: Token to Tree

## 3. Evaluation
****
# The monkey language
The language includes a variety of different tokens: operators, 


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