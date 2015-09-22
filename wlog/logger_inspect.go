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
	fmt.Printf("\tCatalog: %s\n", logger.Catalog)

	fmt.Printf("\nAppenders:\n")
	for k, _ := range Appenders {
		fmt.Printf("\t%s\n", k)
	}

	// show catalog level maps.
	fmt.Printf("\nCatalogs Level Map:\n")
	for k, v := range levelMap {
		fmt.Printf("\t%s:%s\n", k, GetLevelLabel(v))
	}

	fmt.Printf("\nAppender Level Map:\n")
	for k, a := range Appenders {
		fmt.Printf("Appender : %s\n", k)
		fmt.Printf("Level Map:\n")
		alm := a.GetLevelMap()
		for k, v := range alm {
			fmt.Printf("\t%s:%s\n", k, GetLevelLabel(v))
		}
		fmt.Printf("\n")
	}
}
