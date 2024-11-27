package outputters

// Dependency Injection: injecting an interface (which the function is dependent
// upon to operate) as a parameter in order for us to have more control

import (
	"fmt"
	"io"
	"os"
)

// Fprintf requires an io.Writer interface. We use that dependency and inject it
// as a param into our Greet function.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Charlie")
}