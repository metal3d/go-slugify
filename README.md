# Slugify

This is a simple package that handle Slugify() function. The slugification is usefull for URL build from accentuated strings.

It replaces accentuated chars to non-accentuated and spaces by minus sign. All other chars (non-alphanumeric) are removed.

# Installation

    go get github.com/metal3d/go-slugify

# Usage

See http://godoc.org/github.com/metal3d/go-slugify

Example:

```go
package main

import (
    "fmt"
    "github.com/metal3d/go-slugify"
)

func main(){
    txt := "Être en été est à mi-chemin de noël"
    fmt.Println(slug)
    // print: etre-en-ete-est-a-mi-chemin-de-noel
}
```

# Misc

A big thanks to the author of the Javascript function that is the base of this package:
http://irz.fr/slugme-permalien-javascript-slug/

