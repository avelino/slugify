# slugify
A Go slugify application that handles string

[![Build Status](https://travis-ci.org/avelino/slugify.svg?branch=master)](https://travis-ci.org/avelino/slugify)

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

