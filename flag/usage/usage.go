// Package usage provides (subjectively) nicer flag package formatting.
package usage

import (
	"flag"
	"fmt"
	"os"
	"reflect"
)

// Example of usage.
type Example struct {
	// Help description.
	Help string

	// Command example.
	Command string
}

// Config for output.
type Config struct {
	// Usage line (defaults to [options]).
	Usage string

	// Examples to output.
	Examples []Example
}

// Output usage.
func Output(config *Config) func() {
	return func() {

		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "  Usage:\n\n")
		fmt.Fprintf(os.Stderr, "    %s [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  Flags:\n\n")

		flag.CommandLine.VisitAll(func(f *flag.Flag) {
			u := fmt.Sprintf("    -%s", f.Name)
			name, usage := flag.UnquoteUsage(f)

			if len(name) > 0 {
				u += " " + name
			}

			s := fmt.Sprintf("%-25s %s", u, usage)
			if !isZeroValue(f, f.DefValue) {
				// if _, ok := f.Value.(*stringValue); ok {
				// 	// put quotes on the value
				// 	s += fmt.Sprintf(" (default %q)", f.DefValue)
				// } else {
				s += fmt.Sprintf(" (default %v)", f.DefValue)
				// }
			}
			fmt.Fprint(os.Stderr, s, "\n")
		})

		if len(config.Examples) > 0 {
			fmt.Fprintf(os.Stderr, "\n  Examples:\n")
			for _, example := range config.Examples {
				fmt.Fprintf(os.Stderr, "\n    %s\n", example.Help)
				fmt.Fprintf(os.Stderr, "    $ %s\n", example.Command)
			}
		}

		fmt.Fprint(os.Stderr, "\n")
	}
}

// isZeroValue guesses whether the string represents the zero
// value for a flag. It is not accurate but in practice works OK.
func isZeroValue(f *flag.Flag, value string) bool {
	// Build a zero value of the flag's Value type, and see if the
	// result of calling its String method equals the value passed in.
	// This works unless the Value type is itself an interface type.
	typ := reflect.TypeOf(f.Value)
	var z reflect.Value
	if typ.Kind() == reflect.Ptr {
		z = reflect.New(typ.Elem())
	} else {
		z = reflect.Zero(typ)
	}
	if value == z.Interface().(flag.Value).String() {
		return true
	}

	switch value {
	case "false":
		return true
	case "":
		return true
	case "0":
		return true
	}
	return false
}
