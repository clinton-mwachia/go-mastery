## Web Development in Go

### Introduction

Covered in this tutorial:

- Creating a data structure with load and save methods
- Using the net/http package to build web applications
- Using the html/template package to process HTML templates
- Using the regexp package to validate user input
- Using closures

Assumed knowledge:

- Programming experience
- Understanding of basic web technologies (HTTP, HTML)
- Some UNIX/DOS command-line knowledge

## SetUp

### Create Module

```go
go mod init web/dev
```

**Create a `main.go` file** and add the code:-

```go
package main

import (
	"fmt"
	"os"
)

// defining data structures
type Page struct {
	Title string
	Body  []byte
}

// The Page struct describes how page data will be stored in memory
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

// load save pages + error handling
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
```

**Run the code**

```go
go run main.go
```

**EXPECTED OUTPUT**

```go
This is a sample Page.
```

### Reference

---

[Reference](https://go.dev/doc/articles/wiki/)
