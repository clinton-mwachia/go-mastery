# Embed Static Files

In Go 1.16 and later, the `//go:embed` directive allows developers to embed static files and directories into the final binary at compile time. This feature is particularly useful for bundling assets like configuration files, templates, or static web content directly within your Go applications, eliminating the need to manage separate asset files during deployment.

## Using `//go:embed`

To utilize the `//go:embed` directive, follow these steps:

1. **Import the `embed` package**: Even if you don't directly reference any identifiers from the `embed` package, importing it is necessary to enable the embedding functionality. If no identifiers are used, perform a blank import.

2. **Apply the `//go:embed` directive**: Place this directive immediately above the variable declaration that will hold the embedded data. The variable must be of type `string`, `[]byte`, or `embed.FS`.

3. **Specify the file or directory to embed**: Provide a relative path to the file or directory you wish to embed.

### Initialise a Module

```go
go mod init embed/files
```

### Embedding a Single File into a String

If you have a text file named `config.yaml` and want to embed its contents into a string variable:

```go
package main

import (
	_ "embed"
	"fmt"
)

//go:embed config.yaml
var configData string

func main() {
	fmt.Println(configData)
}
```

**Build the project**

```go
go build
```

**Run the binary**

```go
./files.exe
```

In this example:

- The `config.yaml` file's contents are embedded into the `configData` string variable.

- The `embed` package is imported with a blank identifier to enable embedding.

- At runtime, `configData` contains the contents of `config.yaml`, which are printed to the console.

### Embedding a Single File into a Byte Slice

To embed the same `config.yaml` file into a byte slice:

```go
package main

import (
	_ "embed"
	"fmt"
)

//go:embed config.yaml
var configData []byte

func main() {
	fmt.Println(string(configData))
}
```

Here, `configData` is a byte slice containing the contents of `config.yaml`, which is then converted to a string for printing.

### Embedding Multiple Files or an Entire Directory

To embed multiple files or all files within a directory, use the `embed.FS` type. For instance, to embed all files in a `static` directory:

```go
package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed static
var staticFiles embed.FS

func main() {
	files, err := fs.ReadDir(staticFiles, "static")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
```

In this code:

- The `static` directory and its contents are embedded into the `staticFiles` variable of type `embed.FS`.

- The program reads the embedded directory and lists the names of the files contained within it.

### Accessing Embedded Files

When using `embed.FS`, you can access embedded files using methods provided by the `fs.FS` interface. For example, to read the contents of an embedded file:

```go
data, err := staticFiles.ReadFile("static/example.txt")
if err != nil {
	log.Fatal(err)
}
fmt.Println(string(data))
```

This snippet reads the contents of `example.txt` from the embedded `static` directory and prints it.

## Important Considerations

- **Importing `embed` Package**: Always import the `embed` package when using the `//go:embed` directive. If you don't reference any identifiers from `embed`, use a blank import (`_ "embed"`).

- **Variable Types**: The variable to hold embedded data must be of type `string`, `[]byte`, or `embed.FS`.

- **Directive Placement**: The `//go:embed` directive must immediately precede the variable declaration without any intervening blank lines or comments.

- **Path Patterns**: Paths are relative to the directory containing the Go source file. You can use patterns like `*.txt` to embed multiple files matching a specific pattern.

- **Exclusions**: By default, files starting with `.` or `_` are excluded from embedding. To include such files, use the `all:` prefix in your pattern (e.g., `//go:embed all:static`).

## Considerations

- Relative Paths: Ensure that the paths specified in the //go:embed directive match the actual directory structure. The paths are relative to the directory containing main.go.​

- Content Types: Setting the appropriate Content-Type headers ensures that browsers correctly interpret and render the files.​

- Error Handling: Proper error handling is implemented to manage cases where embedded files cannot be read.​

## Conclusion

By embedding your HTML and CSS files directly into your Go application, you create a self-contained binary that simplifies deployment and ensures all necessary assets are included. This approach is particularly useful for small to medium-sized web applications where managing external static files might be cumbersome.

## Reference

For more detailed information, refer to the official Go documentation on the `embed` package: [embed](https://pkg.go.dev/embed)
