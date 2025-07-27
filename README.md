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

Then, run `go get github.com/debarbarinantoine/go-enum-generate` to add the generator as a dependency (or add `@latest` at the end of the line).

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
// DO NOT MODIFY NOR EDIT THIS FILE DIRECTLY.
// To modify this enum, edit the enums.json or enums.yaml definition file
// To know more about `go-enum-generate`, see go to `https://github.com/debarbarinantoine/go-enum-generate`
// Generated at: 2025-07-27 23:19:52

type Color uint

const (
	red Color = iota
	green
	blue
)

var colorKeys = make(map[Color]struct{}, 3)
var colorValues = make(map[string]Color, 3)
var colorKeysArray = make([]Color, 3)
var colorValuesArray = make([]string, 3)

func init() {
	colorKeys[red] = struct{}{}
	colorKeysArray[0] = red
	colorValues["#FF0000"] = red
	colorValuesArray[0] = "#FF0000"
	
	colorKeys[green] = struct{}{}
	colorKeysArray[1] = green
	colorValues["#00FF00"] = green
	colorValuesArray[1] = "#00FF00"
	
	colorKeys[blue] = struct{}{}
	colorKeysArray[2] = blue
	colorValues["#0000FF"] = blue
	colorValuesArray[2] = "#0000FF"
}

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

func (e *Color) Parse(str string) error {
	
	str = strings.TrimSpace(str)
	
	if val, ok := colorValues[str]; ok {
		*e = val
		return nil
	}
	return fmt.Errorf("invalid Color: %s", str)
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
	if _, ok := colorKeys[e]; !ok {
		return false
	}
	return true
}

type colors struct {
	Red Color
	Green Color
	Blue Color
}

var Colors = colors{
	Red: red,
	Green: green,
	Blue: blue,
}

func (e colors) Values() []Color {
	return slices.Clone(colorKeysArray)
}

func (e colors) Args() []string {
	return slices.Clone(colorValuesArray)
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
	if _, ok := colorKeys[Color(value)]; !ok {
		return 0, fmt.Errorf("invalid cast Color: %d", value)
	}
	return Color(value), nil
}
```

---

## Performance

This version of the go-enum-generate produces blazing fast enums! 🚀

Using a 50 values long enum, I executed a test code to see how the generated code would perform. The enum used corresponds to the one in `performance-test.yaml` file.

Here are the results on my desktop computer using the code following the results:

```
2025/07/27 23:38:47 --- Starting Enum Test Program ---
2025/07/27 23:38:47 Note: Time measurements for single operations (e.g., String()) may be extremely small (nanoseconds).
2025/07/27 23:38:47 For O(1) operations, multiple iterations are used to get measurable times.
2025/07/27 23:38:47 Initial Test Enum: #FF0000 (Value: 0)
2025/07/27 23:38:47 
--- Testing String() method ---
2025/07/27 23:38:47 Call: #FF0000.String() => "#FF0000". Took: 150ns
2025/07/27 23:38:47 
--- Testing Parse() method (case-sensitive, trimmed) ---
2025/07/27 23:38:47 Call: Parse("#FF0000") => #FF0000. Took: 401ns
2025/07/27 23:38:47 Call: Parse(" #FF0000 ") => #FF0000. Took: 361ns
2025/07/27 23:38:47 Call: Parse("NON_EXISTENT_COLOR") correctly returned error: invalid Color: NON_EXISTENT_COLOR. Took: 702ns
2025/07/27 23:38:47 Call: Parse("#ff0000" [lowercase]) correctly returned error (case-sensitive): invalid Color: #ff0000. Took: 1.683µs
2025/07/27 23:38:47 
--- Testing IsValid() method (O(1) loop) ---
2025/07/27 23:38:47 Performed 2000000 valid/invalid IsValid() checks. Took: 21.051371ms (Avg per check: 10ns)
2025/07/27 23:38:47 
--- Testing Value() method ---
2025/07/27 23:38:47 Call: #FF0000.Value() => 0. Took: 30ns
2025/07/27 23:38:47 
--- Testing MarshalText()/UnmarshalText() methods ---
2025/07/27 23:38:47 Call: MarshalText(#FF00FF) => "#FF00FF". Took: 130ns
2025/07/27 23:38:47 Call: UnmarshalText("#FF00FF") => #FF00FF. Took: 972ns
2025/07/27 23:38:47 
--- Testing Values() method ---
2025/07/27 23:38:47 Call: Colors.Values() returned 50 values. Took: 781ns
2025/07/27 23:38:47 First 5 values: [#FF0000 #00FF00 #0000FF #FFFF00 #00FFFF]
2025/07/27 23:38:47 
--- Testing Args() method ---
2025/07/27 23:38:47 Call: Colors.Args() returned 50 strings. Took: 912ns
2025/07/27 23:38:47 First 5 args: [#FF0000 #00FF00 #0000FF #FFFF00 #00FFFF]
2025/07/27 23:38:47 
--- Testing Description() method ---
2025/07/27 23:38:47 Call: Colors.Description() generated. Took: 30.187µs
2025/07/27 23:38:47     Available Colors:
=> 0 -> #FF0000
... # [skipping lines here]
=> 49 -> #DAA520

2025/07/27 23:38:47 
--- Testing Cast() method (O(1) loop) ---
2025/07/27 23:38:47 Performed 2000000 valid/invalid Cast() checks. Took: 390.72923ms (Avg per check: 195ns)
2025/07/27 23:38:47 
--- All Tests Completed Successfully ---
```

Here is the code used to test the performance of the code generated with *go-enum-generate* tool:

```go
package main

import (
	"log"
	"strings"
	"time"
	
	"enumTest/enum"
)

//go:generate go run github.com/debarbarinantoine/go-enum-generate@latest --force
func main() {
	log.Println("--- Starting Enum Test Program ---")
	log.Println("Note: Time measurements for single operations (e.g., String()) may be extremely small (nanoseconds).")
	log.Println("For O(1) operations, multiple iterations are used to get measurable times.")
	
	// --- General Setup ---
	// Pick a test color from your long list
	// Make sure this key/value exists in your enums.yaml
	testColorEnum := enum.Colors.Red // Use one of your generated enum constants
	testColorValueStr := "#FF0000"   // Use its exact string value from enums.yaml
	
	log.Printf("Initial Test Enum: %s (Value: %d)", testColorEnum.String(), testColorEnum.Value())
	
	// --- Test String() method ---
	log.Println("\n--- Testing String() method ---")
	start := time.Now()
	strVal := testColorEnum.String()
	duration := time.Since(start)
	log.Printf("Call: %s.String() => \"%s\". Took: %s", testColorEnum.String(), strVal, duration)
	
	// --- Test Parse() method ---
	log.Println("\n--- Testing Parse() method (case-sensitive, trimmed) ---")
	var parsedColor enum.Color
	
	// Test valid parse (exact match)
	start = time.Now()
	err := parsedColor.Parse(testColorValueStr)
	duration = time.Since(start)
	if err != nil {
		log.Fatalf("ERROR: Parsing \"%s\" failed: %v", testColorValueStr, err)
	}
	log.Printf("Call: Parse(\"%s\") => %s. Took: %s", testColorValueStr, parsedColor.String(), duration)
	
	// Test valid parse with whitespace (should work due to TrimSpace)
	start = time.Now()
	err = parsedColor.Parse(" " + testColorValueStr + " ")
	duration = time.Since(start)
	if err != nil {
		log.Fatalf("ERROR: Parsing with spaces \"%s\" failed: %v", " "+testColorValueStr+" ", err)
	}
	log.Printf("Call: Parse(\" %s \") => %s. Took: %s", testColorValueStr, parsedColor.String(), duration)
	
	// Test invalid parse (non-existent string)
	start = time.Now()
	err = parsedColor.Parse("NON_EXISTENT_COLOR")
	duration = time.Since(start)
	if err != nil {
		log.Printf("Call: Parse(\"NON_EXISTENT_COLOR\") correctly returned error: %v. Took: %s", err, duration)
	} else {
		log.Fatalf("ERROR: Parse(\"NON_EXISTENT_COLOR\") unexpectedly succeeded.")
	}
	
	// Test invalid parse (incorrect case - should fail as Parse is case-sensitive)
	start = time.Now()
	err = parsedColor.Parse(strings.ToLower(testColorValueStr)) // Assuming testColorValueStr is not already lowercase
	duration = time.Since(start)
	if err != nil {
		log.Printf("Call: Parse(\"%s\" [lowercase]) correctly returned error (case-sensitive): %v. Took: %s", strings.ToLower(testColorValueStr), err, duration)
	} else {
		log.Fatalf("ERROR: Parse(\"%s\" [lowercase]) unexpectedly succeeded (case-sensitive issue).", strings.ToLower(testColorValueStr))
	}
	
	// --- Test IsValid() method (Loop for performance) ---
	log.Println("\n--- Testing IsValid() method (O(1) loop) ---")
	numIterations := 1000000 // One Million iterations to get measurable time
	start = time.Now()
	for i := 0; i < numIterations; i++ {
		_ = enum.Colors.Blue.IsValid()  // Valid check
		_ = enum.Color(99999).IsValid() // Invalid check (use a value well outside your defined range)
	}
	duration = time.Since(start)
	log.Printf("Performed %d valid/invalid IsValid() checks. Took: %s (Avg per check: %s)",
		numIterations*2, duration, duration/time.Duration(numIterations*2))
	
	// --- Test Value() method ---
	log.Println("\n--- Testing Value() method ---")
	start = time.Now()
	valUint := testColorEnum.Value()
	duration = time.Since(start)
	log.Printf("Call: %s.Value() => %d. Took: %s", testColorEnum.String(), valUint, duration)
	
	// --- Test MarshalText() / UnmarshalText() methods ---
	log.Println("\n--- Testing MarshalText()/UnmarshalText() methods ---")
	initialForMarshal := enum.Colors.Magenta // Pick another color
	
	start = time.Now()
	marshaled, err := initialForMarshal.MarshalText()
	durationMarshal := time.Since(start)
	if err != nil {
		log.Fatalf("ERROR: Marshaling %s failed: %v", initialForMarshal.String(), err)
	}
	log.Printf("Call: MarshalText(%s) => \"%s\". Took: %s", initialForMarshal.String(), string(marshaled), durationMarshal)
	
	var unmarshaled enum.Color
	start = time.Now()
	err = unmarshaled.UnmarshalText(marshaled)
	durationUnmarshal := time.Since(start)
	if err != nil {
		log.Fatalf("ERROR: Unmarshaling \"%s\" failed: %v", string(marshaled), err)
	}
	log.Printf("Call: UnmarshalText(\"%s\") => %s. Took: %s", string(marshaled), unmarshaled.String(), durationUnmarshal)
	
	// --- Test Values() method ---
	log.Println("\n--- Testing Values() method ---")
	start = time.Now()
	allColors := enum.Colors.Values() // Access the global enum var
	duration = time.Since(start)
	log.Printf("Call: Colors.Values() returned %d values. Took: %s", len(allColors), duration)
	log.Printf("First 5 values: %v", allColors[:min(5, len(allColors))])
	
	// --- Test Args() method ---
	log.Println("\n--- Testing Args() method ---")
	start = time.Now()
	args := enum.Colors.Args() // Access the global enum var
	duration = time.Since(start)
	log.Printf("Call: Colors.Args() returned %d strings. Took: %s", len(args), duration)
	log.Printf("First 5 args: %v", args[:min(5, len(args))])
	
	// --- Test Description() method ---
	log.Println("\n--- Testing Description() method ---")
	start = time.Now()
	desc := enum.Colors.Description() // Access the global enum var
	duration = time.Since(start)
	log.Printf("Call: Colors.Description() generated. Took: %s", duration)
	// You can print the full description if you want to see it:
	log.Println(desc)
	
	// --- Test Cast() method (Loop for performance) ---
	log.Println("\n--- Testing Cast() method (O(1) loop) ---")
	numIterations = 1000000                   // One Million iterations
	validUintToCast := uint(enum.Colors.Blue) // A valid uint value
	invalidUintToCast := uint(99999)          // An invalid uint value
	
	start = time.Now()
	for i := 0; i < numIterations; i++ {
		_, _ = enum.Colors.Cast(validUintToCast)   // Valid cast
		_, _ = enum.Colors.Cast(invalidUintToCast) // Invalid cast
	}
	duration = time.Since(start)
	log.Printf("Performed %d valid/invalid Cast() checks. Took: %s (Avg per check: %s)",
		numIterations*2, duration, duration/time.Duration(numIterations*2))
	
	log.Println("\n--- All Tests Completed Successfully ---")
}
```

-----

## Contributing

Contributions are welcome\! If you find a bug or have a feature request, please open an issue or submit a pull request on the GitHub repository.

-----

## License

This project is licensed under the MIT License.
