# TagFilter

A simple Go library to filter items by tags.

## Usage

```go
package main

import (
	"fmt"

	"github.com/vcokltfre/tagfilter"
)

func main() {
	tags := []string{"tag1", "tag2", "tag3"}
	tagsExtended := []string{"tag1", "tag2", "tag3", "tag4", "tag5", "tag6"}

	// Any of tags 'tag1', 'tag2', 'tag3'
	fmt.Println(tagfilter.Match(tags, "tag1 tag2 tag3"))

	// All of tags 'tag1', 'tag2', 'tag3'
	fmt.Println(tagfilter.Match(tags, "+tag1 +tag2 +tag3"))

	// Any of tags 'tag1', 'tag2', 'tag3' and none of tags 'tag4', 'tag5', 'tag6'
	fmt.Println(tagfilter.Match(tags, "tag1 tag2 tag3 -tag4 -tag5 -tag6"))         // Matches
	fmt.Println(tagfilter.Match(tagsExtended, "tag1 tag2 tag3 -tag4 -tag5 -tag6")) // Doesn't match

	// All of tags 'tag1', 'tag2' and none of tags 'tag4', 'tag5', 'tag6'
	fmt.Println(tagfilter.Match(tags, "+tag1 +tag2 -tag4 -tag5 -tag6"))         // Matches
	fmt.Println(tagfilter.Match(tagsExtended, "+tag1 +tag2 -tag4 -tag5 -tag6")) // Doesn't match
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
