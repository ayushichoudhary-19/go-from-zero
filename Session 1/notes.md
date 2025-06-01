# 🧠 Golang Fundamentals

## 📦 Packages and `main` Function

```go
package main
```

* Every Go file **must** begin with a `package` declaration.
* `main` is special: any package with `package main` can be compiled as an executable program.
* Other packages (like `package client`) are **library-style packages** that can be imported into other files.

🔁 **Analogy:**

* Go’s `main` is like Python’s `if __name__ == "__main__":`, but Go enforces this at compile time.
* You can’t run a file directly unless it’s in `package main` and contains a `main()` function.

---

## 🔧 The `main()` Function

```go
func main() {
    fmt.Println("Hello GoBlogr")
}
```

* Go programs **start** with the `main()` function.
* Must be in a file under `package main`.
* Functions are declared using the `func` keyword.

---

## 📤 `fmt.Println` vs `fmt.Print`

* `fmt.Println()` adds a newline.
* `fmt.Print()` doesn’t.
* `fmt.Printf()` allows formatted output like in C (`%s`, `%d`, etc.).

---

## 📚 Package Import Rules

```go
import "fmt"
```

* Imports must be **used**, or Go throws an unused import error.
* You can import multiple packages using parentheses:

```go
import (
    "fmt"
    "os"
    "strings"
)
```

---

## 🧾 Functions & Export Rules

```go
package client

func Func1() {
    fmt.Println("This is func1 from client package")
}
```

* Functions starting with **uppercase** are **exported** (public).
* Functions starting with **lowercase** are **private to the package**.
* So, in go we do not have export keyword, we just type that function's name starting with an Uppercase.
* In Go, exported identifiers start with capital letters (unlike Python's underscore conventions).
* Go does not have `public`, `private`, or `protected` keywords like Java — it uses capitalization instead!

---

## 💡 `byte` and `rune`

```go
var b byte   // byte is alias for uint8
var r rune   // rune is alias for int32
```

* `byte` is often used for binary data or characters in ASCII.
* `rune` is used for Unicode characters (UTF-8).

```go
char := 'A'        // this is a rune
fmt.Printf("%T\n", char)  // Output: int32
```

---

## 📦 Data Structures: Slices, Maps (but no stack/linked list)

### ❌ No built-in `stack` or `linked list`

* These can be implemented via **slices** and **structs**.
* Go prefers **composability** over pre-baked data types.
* Go prefers simplicity and performance over built-in containers like Python or Java.
---

### ✅ Maps

```go
m := map[string]string{
    "language": "Go",
    "type":     "compiled",
}
```

* `map[KeyType]ValueType`
* Only types that are **comparable** (i.e., support `==`) can be map keys.
* Cannot use slices, maps, or functions as keys.

---

## 🔢 Reading Input: `fmt.Scanln` vs `bufio.NewReader`

### 🔹 `fmt.Scanln(&input)`:

```go
var input string
fmt.Scanln(&input)
```

* Reads **only one word** (until the first space).
* Good for basic single token input.

### 🔸 Problem

```sh
Enter a sentence: Ayushi is a nice dev
Words: [Ayushi]
```

* The rest (`is a nice dev`) is ignored or passed to shell.

---

### ✅ `bufio.NewReader` to the Rescue

```go
reader := bufio.NewReader(os.Stdin)
input, _ := reader.ReadString('\n')
input = strings.TrimSpace(input)
```

* Reads the **entire line**, including spaces.
* Use `strings.Fields(input)` to split into words.

🔁 Like Python’s:

```python
input("Enter: ")          # for entire line
input.split()             # to tokenize
```

---

## 📖 `strings.Fields()` Explanation

```go
input := "Ayushi is a nice dev"
words := strings.Fields(input)
```

* Splits string on any whitespace (spaces, tabs, etc.).
* Returns a `[]string` (slice of strings).
* Empty strings are filtered out automatically.

---

## 🧮 Integers, Overflows, and Type Aliases

### 🔸 Unsigned Overflow

```go
var maxUint32 uint32 = 4294967295
fmt.Println(maxUint32 + 1) // Output: 0 — wraps around!
```

* Go does **not** throw runtime errors on overflow.
* `uint32` → 0 to 2³²−1 (4294967295)
* Overflow silently wraps back to zero — **use with caution**.

---

### ✅ Use Larger Types to Avoid It

```go
var bigger uint64 = 4294967296
```

---

## 🔍 Summary Table

| Concept             | Notes                              |
| ------------------- | ---------------------------------- |
| `package main`      | Required to run program            |
| `func main()`       | Entry point                        |
| `fmt.Println()`     | Prints with newline                |
| `Scanln`            | Reads only one word                |
| `bufio.Reader`      | Reads full line                    |
| `strings.Fields`    | Splits by whitespace               |
| `map[string]string` | Dictionary equivalent              |
| `byte`              | Alias for `uint8`                  |
| `rune`              | Alias for `int32`, handles Unicode |
| Integer overflow    | No runtime error, wraps silently   |

---

## ⚠️ Gotchas to Watch Out For

* **No semicolons** needed at line ends — Go automatically inserts them during parsing.
* **Must use** all declared variables (Go is strict on unused vars).
* Naming matters for visibility (`Func1` is exported, `func1` is not).
* `:=` is for short variable declaration and assignment.

---