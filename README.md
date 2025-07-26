# go-enum-generate

`go-enum-generate` is a simple yet powerful command-line tool designed to streamline the creation of Go enumerations (enums) from `YAML` or `JSON` definition files. It automatically generates idiomatic Go code, complete with methods for string conversion, parsing, validation, JSON/text marshalling, and more, helping you maintain type safety and consistency in your projects.

<div style="text-align: right">
<strong>Author:</strong> <em>Antoine de Barbarin</em>
</div>

-----

## Features

* **Flexible Input:** Define your enums using either `YAML` (`enums.yaml`) or `JSON` (`enums.json`) files.
* **Idiomatic Go Output:** Generates clean, `gofmt`-friendly Go code that adheres to common Go conventions for enums.
* **Comprehensive Enum Methods:**
	* `String()`: Converts enum values to their string representation.
	* `Parse()`: Parses strings back into enum values.
	* `Value()`: Returns the underlying `uint` value.
	* `MarshalText()` / `UnmarshalText()`: Enables seamless `JSON` and `text` (e.g., `YAML`) serialization/deserialization.
	* `IsValid()`: Checks if an enum value is one of the defined members.
	* `Args()`: Returns a slice of all enum string representations, useful for CLI arguments.
	* `Description()`: Provides a formatted string listing all available enum values and their descriptions.
	* `Cast()`: Safely converts a `uint` back to an enum type with validation.
* **Automatic Naming Conventions:** Handles conversion of enum names to appropriate Go naming conventions (e.g., `Color` type, `colors` internal struct, `Colors` variable).
* **Pluralization:** Automatically determines plural forms if not explicitly provided.
* **Robust Validation:** Includes checks for valid Go identifiers and unique enum keys/values to prevent conflicts.
* **`go generate` Integration:** Designed to work seamlessly with `go generate` for automated code generation.

-----

## Getting Started

### Installation

You can install `go-enum-generate` globally (recommended for general use) or integrate it directly into your project via `go generate`.

#### Global Installation (Recommended for general use)

To install the tool globally, run:

```bash
go install github.com/debarbarinantoine/go-enum-generate@latest
```

This will place the `go-enum-generate` executable in your `$GOBIN` path, making it available from any directory.

#### Local Integration with `go generate` (for project-specific generation)

If you prefer to manage the generator's version as a project dependency, add a `//go:generate` directive to any Go file in your project (e.g., `main.go` or a dedicated `generate.go` file):

```go
//go:generate go run github.com/debarbarinantoine/go-enum-generate
```

Then, run `go mod tidy` to add the generator as a dependency.

### Usage

1.  **Create your enum definition file:**
    In the root of your Go project, create an `enums.yaml` or `enums.json` file.

    **`enums.yaml` example:**

    ```yaml
    # This file defines a list of custom enumerations for your application.
    # Each entry in the list represents a single enumeration.

    # Required: The singular name of the enumeration (e.g., "Color")
    - name: Color

      # Optional: The plural name (omit if not needed)
      # Automatic plural rule: s, sh, x, z, ch or j ending -> add 'es', y ending -> switch 'y' with 'ies', otherwise -> add 's'
      plural: Colors

      # Required: A list of key-value pairs for the enum members
      values:

          # Required: The programmatic key (e.g., "RED")
        - key: red

          # Required: The associated value (e.g., a hex code, a string description)
          value: "#FF0000"
        - key: green
          value: "#00FF00"
        - key: blue
          value: "#0000FF"
    ```

    **`enums.json` example:**

    ```json
    [
      {
        "name": "Color",
        "plural": "Colors",
        "values": [
          {
            "key": "red",
            "value": "#FF0000"
          },
          {
            "key": "green",
            "value": "#00FF00"
          },
          {
            "key": "blue",
            "value": "#0000FF"
          }
        ]
      }
    ]
    ```

2.  **Generate your enums:**

	* If installed globally:
	  ```bash
	  go-enum-generate
	  ```
	* If using `go generate`:
	  ```bash
	  go generate ./...
	  ```

	This will create an `enum` directory in your project root (if it doesn't exist) and place the generated Go files (e.g., `enum/color.go`) inside it.

### Command-Line Flags

* `-f`, `--force`: Overwrite existing enum files without prompting. By default, the tool will return an error if a file already exists and `--force` is not used.
* `-h`, `--help`: Show help message and exit.

### Example Generated Code

Here's an example of the Go code generated for the `Color` enum defined above:

```go
package enum

import (
	"fmt"
	"slices"
	"strings"
)

// This file has been created automatically by `go-enum-generate`
// Do not modify nor edit this file directly.
// To modify this enum, edit the enums.json or enums.yaml definition file

type Color uint

const (
	red Color = iota
	green
	blue
)

func (e Color) String() string {
	switch e {
	case red:
		return "#FF0000"
	case green:
		return "#00FF00"
	case blue:
		return "#0000FF"
	default:
		return fmt.Sprintf("Unknown Color (%d)", e.Value())
	}
}

type colors struct {
	red Color
	green Color
	blue Color
}

var Colors = colors{
	red: red,
	green: green,
	blue: blue,
}

func (e colors) Values() []Color {
	return []Color{
		red,
		green,
		blue,
	}
}

func (e *Color) Parse(str string) error {
	
	str = strings.ToUpper(str)
	
	switch str {
	case Colors.red.String():
		*e = Colors.red
	case Colors.green.String():
		*e = Colors.green
	case Colors.blue.String():
		*e = Colors.blue
	default:
		return fmt.Errorf("invalid Color: %s", str)
	}
	
	return nil
}

func (e Color) Value() uint {
	return uint(e)
}

func (e Color) MarshalText() ([]byte, error) {
	return []byte(e.String()), nil
}

func (e *Color) UnmarshalText(text []byte) error {
	return e.Parse(string(text))
}

func (e Color) IsValid() bool {
	if !slices.Contains(Colors.Values(), e) {
		return false
	}
	return true
}

func (e colors) Args() []string {
	var args []string
	
	for _, enumVal := range e.Values() {
		args = append(args, enumVal.String())
	}
	return args
}

func (e colors) Description() string {
	var strBuilder strings.Builder
	strBuilder.WriteString("\tAvailable Colors:\n")
	for _, enumVal := range e.Values() {
		strBuilder.WriteString(fmt.Sprintf("=> %d -> %s\n", enumVal.Value(), enumVal.String()))
	}
	return strBuilder.String()
}

func (e colors) Cast(value uint) (Color, error) {
	if !slices.Contains(e.Values(), Color(value)) {
		return 0, fmt.Errorf("invalid cast Color: %d", value)
	}
	return Color(value), nil
}
```

-----

## Contributing

Contributions are welcome\! If you find a bug or have a feature request, please open an issue or submit a pull request on the GitHub repository.

-----

## License

This project is licensed under the MIT License.
