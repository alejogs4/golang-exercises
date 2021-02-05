package main

import (
	"fmt"
	"log"
	"strings"

	"example.com/alejogs4/learning/linkparser"
)

const htmlSample string = `
<html>
<body>
  <h1>Hello!</h1>
	<a href="/other-page">A link to another page</a>
	<a href="/other-go">A link to another page</a>
	<ul>
		<li></li>
		<li></li>
		<li></li>
		<li></li>
		<li>
			<a href="/nested-link">
			Nested one
			<span>strong 1</span>
			<span>strong 2</span>
			</a>
			<!--[if lt IE 7]> <html class="ie ie6 lt-ie9 lt-ie8 lt-ie7" lang="en"> <![endif]-->
		</li>
	</ul>
</body>
</html>
`

func main() {
	r := strings.NewReader(htmlSample)

	links, error := linkparser.Parse(r)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Printf("%+v\n", links)

}
