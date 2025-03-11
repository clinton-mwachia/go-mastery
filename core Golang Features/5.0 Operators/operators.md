# Operators in Go

## Introduction

Operators are symbols that perform operations on variables and values. Go provides various types of operators, including **arithmetic, relational, logical, bitwise, assignment, and unary operators**.

---

## Code

```go
package main

import "fmt"

func main() {
    // Arithmetic Operators
    a, b := 10, 5
    fmt.Println("Arithmetic Operators:")
    fmt.Println("Addition:", a+b)  // 10 + 5 = 15
    fmt.Println("Subtraction:", a-b)  // 10 - 5 = 5
    fmt.Println("Multiplication:", a*b)  // 10 * 5 = 50
    fmt.Println("Division:", a/b)  // 10 / 5 = 2
    fmt.Println("Modulus:", a%b)  // 10 % 5 = 0

    // Relational (Comparison) Operators
    fmt.Println("\nRelational Operators:")
    fmt.Println("a == b:", a == b)  // false
    fmt.Println("a != b:", a != b)  // true
    fmt.Println("a > b:", a > b)  // true
    fmt.Println("a < b:", a < b)  // false
    fmt.Println("a >= b:", a >= b)  // true
    fmt.Println("a <= b:", a <= b)  // false

    // Logical Operators
    x, y := true, false
    fmt.Println("\nLogical Operators:")
    fmt.Println("x && y (AND):", x && y)  // false
    fmt.Println("x || y (OR):", x || y)  // true
    fmt.Println("!x (NOT):", !x)  // false

    // Bitwise Operators
    fmt.Println("\nBitwise Operators:")
    fmt.Println("a & b (AND):", a & b)  // 10 & 5 = 0
    fmt.Println("a | b (OR):", a | b)  // 10 | 5 = 15
    fmt.Println("a ^ b (XOR):", a ^ b)  // 10 ^ 5 = 15
    fmt.Println("a << 1 (Left Shift):", a << 1)  // 10 << 1 = 20
    fmt.Println("a >> 1 (Right Shift):", a >> 1)  // 10 >> 1 = 5

    // Assignment Operators
    fmt.Println("\nAssignment Operators:")
    c := 10
    fmt.Println("Initial value of c:", c)
    c += 5  // c = c + 5
    fmt.Println("c += 5:", c)
    c -= 3  // c = c - 3
    fmt.Println("c -= 3:", c)
    c *= 2  // c = c * 2
    fmt.Println("c *= 2:", c)
    c /= 4  // c = c / 4
    fmt.Println("c /= 4:", c)
    c %= 3  // c = c % 3
    fmt.Println("c %= 3:", c)

    // Unary Operators
    fmt.Println("\nUnary Operators:")
    num := 10
    fmt.Println("num:", num)
    num++  // Increment
    fmt.Println("num++:", num)
    num--  // Decrement
    fmt.Println("num--:", num)
}
```

---

## Explanation

### **1. Arithmetic Operators**

These operators perform basic mathematical operations:

| Operator | Description    | Example                 |
| -------- | -------------- | ----------------------- |
| `+`      | Addition       | `a + b` → `10 + 5 = 15` |
| `-`      | Subtraction    | `a - b` → `10 - 5 = 5`  |
| `*`      | Multiplication | `a * b` → `10 * 5 = 50` |
| `/`      | Division       | `a / b` → `10 / 5 = 2`  |
| `%`      | Modulus        | `a % b` → `10 % 5 = 0`  |

---

### **2. Relational (Comparison) Operators**

Used to compare values and return `true` or `false`:

| Operator | Description              | Example            |
| -------- | ------------------------ | ------------------ |
| `==`     | Equal to                 | `a == b` → `false` |
| `!=`     | Not equal to             | `a != b` → `true`  |
| `>`      | Greater than             | `a > b` → `true`   |
| `<`      | Less than                | `a < b` → `false`  |
| `>=`     | Greater than or equal to | `a >= b` → `true`  |
| `<=`     | Less than or equal to    | `a <= b` → `false` |

---

### **3. Logical Operators**

Used to combine multiple conditions:

| Operator | Description | Example            |     |     |            |      |     |      |
| -------- | ----------- | ------------------ | --- | --- | ---------- | ---- | --- | ---- |
| `&&`     | Logical AND | `x && y` → `false` |     |     |            |      |     |      |
| `        |             | `                  |     | `   | Logical OR | `x<2 |     | x<4` |
| `!`      | Logical NOT | `!x` → `false`     |     |     |            |      |     |      |

---

### **4. Bitwise Operators**

Operate at the binary level:

| Operator | Description | Example        |     |         |
| -------- | ----------- | -------------- | --- | ------- |
| `&`      | Bitwise AND | `10 & 5 = 0`   |     |         |
| `        | `           | Bitwise OR     | `10 | 5 = 15` |
| `^`      | Bitwise XOR | `10 ^ 5 = 15`  |     |         |
| `<<`     | Left shift  | `10 << 1 = 20` |     |         |
| `>>`     | Right shift | `10 >> 1 = 5`  |     |         |

---

### **5. Assignment Operators**

Used to assign values to variables:

| Operator | Description         | Example                |
| -------- | ------------------- | ---------------------- |
| `=`      | Assign              | `c = 10`               |
| `+=`     | Add and assign      | `c += 5` → `c = c + 5` |
| `-=`     | Subtract and assign | `c -= 3` → `c = c - 3` |
| `*=`     | Multiply and assign | `c *= 2` → `c = c * 2` |
| `/=`     | Divide and assign   | `c /= 4` → `c = c / 4` |
| `%=`     | Modulus and assign  | `c %= 3` → `c = c % 3` |

---

### **6. Unary Operators**

Operate on a single operand:

| Operator | Description | Example                   |
| -------- | ----------- | ------------------------- |
| `++`     | Increment   | `num++` → `num = num + 1` |
| `--`     | Decrement   | `num--` → `num = num - 1` |

---

## Running the Code

1. Save the file as `operators.go`.
2. Open a terminal and navigate to the file's location.
3. Run:

   ```sh
   go run operators.go
   ```

4. Expected output:

   ```
   Arithmetic Operators:
   Addition: 15
   Subtraction: 5
   Multiplication: 50
   Division: 2
   Modulus: 0

   Relational Operators:
   a == b: false
   a != b: true
   ...

   Logical Operators:
   x && y (AND): false
   x || y (OR): true
   !x (NOT): false
   ```
