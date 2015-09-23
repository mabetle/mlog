package wlog

import (
	"fmt"
)

// Inspect
func (l *WrapLogger) Inspect() {
	InspectLogger(l)
}

// InspectLogger
func InspectLogger(logger *WrapLogger) {
	fmt.Printf("\nInspect Logger:\n")

	fmt.Printf("\nCatalog: %s\n", logger.Catalog)

	for _, a := range GetAppenders() {
		a.Inspect(logger.Catalog)
	}
}
