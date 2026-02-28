// Example: Prefix customisation
//
// Demonstrates setting a custom prefix, clearing the prefix, and restoring it.
//
// Run: go run ./examples/prefix/
package main

import (
	"fmt"

	"github.com/z0mbix/cliout"
)

func main() {
	// Default prefix is "»"
	fmt.Println("=== Default prefix ===")
	cliout.Info("using the default prefix")

	// Arrow prefix
	fmt.Println("\n=== Arrow prefix ===")
	cliout.SetPrefix("->")
	cliout.Info("using arrow prefix")
	cliout.Warn("warnings use it too")

	// Asterisk prefix
	fmt.Println("\n=== Asterisk prefix ===")
	cliout.SetPrefix("*")
	cliout.Info("using asterisk prefix")

	// App name as prefix
	fmt.Println("\n=== App name prefix ===")
	cliout.SetPrefix("[myapp]")
	cliout.Info("prefixed with app name")
	cliout.Error("errors too")

	// Unicode prefix
	fmt.Println("\n=== Unicode prefixes ===")
	cliout.SetPrefix("✓")
	cliout.Info("checkmark prefix")

	cliout.SetPrefix("●")
	cliout.Info("bullet prefix")

	cliout.SetPrefix("▸")
	cliout.Info("triangle prefix")

	// No prefix at all
	fmt.Println("\n=== No prefix ===")
	cliout.ClearPrefix()
	cliout.Info("no prefix, just the message")
	cliout.Warn("warnings also have no prefix")

	// Restore the prefix
	fmt.Println("\n=== Restored prefix ===")
	cliout.SetPrefix("»")
	cliout.Info("prefix is back to default")
}
