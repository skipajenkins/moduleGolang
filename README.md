
# 🐹 Go Modules & Packages — Greetings Example

This lesson demonstrates how to create a Go module and import it into another Go application.
You’ll learn how Go handles modules, packages, and local replacements, which is a core concept in real-world Go development.

The project is split into two parts:

* greetings/ → a reusable Go module

* hello/ → an application that imports and uses the module

---

## 🎯 Learning Objectives

By completing this lesson, you will learn how to:

* Create a Go module using go mod init

* Export functions from a package

* Import a local module into another Go application

* Use go mod edit -replace for local development

* Understand how Go resolves dependencies

---

📁 Project Structure
```bash
module/
├── greetings/
│   ├── go.mod
│   └── greetings.go
└── hello/
    ├── go.mod
    └── hello.go
```

* greetings is the module we create
  
* hello is the application that consumes the module

---

## ⚙️ Environment Setup
###✅ Step 1: Install Go

Check if Go is installed:
```bash
go version
```

If Go is not installed, download it from:

👉 https://go.dev/dl/

### 📦 Creating the Greetings Module
```greetings/go.mo
module example.com/greetings

go 1.25.6
```

This defines a module path, which is how other Go programs will reference it.

```greetings/greetings.go
package greetings

import (
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
```

### 🔍 Key Points

* The function name Hello is capitalized → it is exported

* package greetings matches the folder name

* The function returns a string, making it reusable

---

## 🚀 Using the Module in an Application
```hello/hello.go
package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
```

---

### 🔁 Linking the Local Module (Important Step)

Because example.com/greetings is a local module, we must tell Go where to find it.

From inside the hello/ directory, run:
```bash
go mod edit -replace example.com/greetings=../greetings
```

Then tidy dependencies:
```bash
go mod tidy
```

This tells Go:

“When you see example.com/greetings, use the local folder instead.”

### ▶️ Running the Application

From the hello/ directory:
```bash
go run .
```

### 🧾 Example Output
```bash
Hi, Gladys. Welcome!
```

---

## 📚 Key Concepts
| Concept              | Explanation                                      |
|----------------------|--------------------------------------------------|
| Go Modules           | Dependency management using `go.mod`             |
| Packages             | Groups of related Go files                       |
| Exported Functions   | Capitalized names (e.g., `Hello`)                |
| Local Replace        | `go mod edit -replace` for local development     |
| Imports              | Use module path, not file paths                  |
| Entry Point          | `func main()`                                    |


---

## 🧠 How It Works (Conceptual Overview)

* greetings defines a reusable module

* The module exports a function (Hello)

* hello imports the module using its module path

* go mod edit -replace links the local folder

* Go automatically resolves and builds dependencies

This mirrors real production workflows, where modules may live in separate repositories.

---

## 🧩 Why This Matters

This pattern is foundational for:

* Microservices

* Shared internal libraries

* Versioned APIs

* Large Go codebases

Once you understand this, you understand how Go scales.

---

## 🛠 Built With

* Go

* Go Modules

---


## 📄 License

This project is open-source and available under the MIT License.

---
