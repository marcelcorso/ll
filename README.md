when i feel down i revert to simpler times when i used to get excited about pointers to pointers

this has a linked list implementation in go because i was bored


```go
package main

import (
        "fmt"

        "github.com/marcelcorso/ll"
)

l := ll.LL{}
l.Insert()
l.Insert(8)
l.Insert(-1)
l.Insert(54)
l.Insert(2)

l.Each(func(val int) {
	fmt.Print(val)
})
fmt.Println("")

```

prints

```go
-12854
```
