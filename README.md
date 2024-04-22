# Rod with errors

The idea behind this small, quick library is to extend rod pages so that they call OnError before timeout kills the page.

This is to save debug logs, take screenshots, etc.

It also allows you to call page.OnError() at any time to perform some error handling code that you choose.

If you need to access the original Page that this error-handling friendly page is built off of, simple access the `.Page` property


The flow of using this library looks like this

```go
    package main

import (
	"fmt"
	"time"

	"github.com/rob117/rod-with-errors/errorpage"
	"github.com/go-rod/rod"
)

func main() {
	browser := rod.New().MustConnect()
	page := errorpage.New(browser.MustPage("https://wikipedia.org"), OnPageError)
	// Close the page when we are done using it to prevent memory leaks
	defer page.Close()
	// Set a timeout for all page operations
	page.WithTimeout(10 * time.Second)
	// Do whatever you want here. If the timeout happens, your OnError method is called
	// and then the page closes
	page.CancelTimeout()
	// Chaining is also possible
	page.WithTimeout(10 * time.Second).MustScreenshot("Success.png").MustLoad("https://wikipedia.org")
	// when chain is done, cancel that timer and start a new one
	page.CancelTimeout().WithTimeout(35 * time.Second).MustScreenshot("SecondSuccess.png")
	// etc.
}

func OnPageError(p *errorpage.PageWithOnError) {
	fmt.Println("Error detected")
	p.MustScreenshot("Error.png")
}
```
