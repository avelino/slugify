# slugify
A Go slugify application that handles string

Example:

    package main
    
    import (
    	"fmt"
    	"github.com/avelino/slugify"
    )
    
    func main() {
    	text := "Example slugify"
    	fmt.Printf(text + ": " + slugify.Slugify(text))
    }

