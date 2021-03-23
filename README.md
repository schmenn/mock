# Mock
a go package for generating mock text

e.g. `lorem ipsum dolor sit amet` &rarr; `lOrEm iPsUm dOlOr sIt aMeT`

## Usage

Get it with:

`go get github.com/Schmenn/mock`


```go
package main

import (
	"fmt"

	"github.com/Schmenn/mock"
)

func main() {
	text := mock.Mock("lorem ipsum dolor sit amet")
	fmt.Println(text) // lOrEm iPsUm dOlOr sIt aMeT
}
```

for a command-line program click [here](https://github.com/Schmenn/mock/cmd/mock)
