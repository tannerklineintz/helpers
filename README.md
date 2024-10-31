# Helpers

Things that make my life easier.

## Deep Copy
returns a deep copy of any type. example:
```go
package main

import (
	"fmt"

	"github.com/tannerklineintz/helpers"
)

func main() {
	mySlice := [][][]string{
		{
			{"Apple", "Banana"},
			{"Cherry", "Date"},
		},
		{
			{"Elderberry", "Fig"},
			{"Grape", "Honeydew"},
		},
	}
	test := helpers.DeepCopy(mySlice)
	fmt.Println(test.([][][]string))
}
```