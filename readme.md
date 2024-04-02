# Thoughtful AI technical challenge

## Description
A utility for sorting packages based on their weight, dimensions, and area.
Will return one of three status labels (STANDARD, SPECIAL, REJECTED) based on the provided values

Usage
```go
package main

import (
	"fmt"
	"github.com/weisbartb/thoughtful"
)

func main()  {
    fmt.Sprintf("package should go to %s ",thoughtful.Sort(100,100,50,14))
}
```

## Notes
- Numbers are stored as floats so that fractional units can be used,
if the requirement was to support unit conversion, a decimal type would be used with a conversion scale.
- More concrete types for numbers should be used past a prototype (weight should be a base in grams and multiplied a const for KG, similar to time package)
  - Wanted to respect time constraints so built to requirements
- Assumed packages must be valid 3 dimensional objects (no negative values)
- An error should return along with rejected for invalid values (for go), requirements said string return only.
- Coverage hits 100% of all logic
- Attempted to be as concise as possible while maintaining readability