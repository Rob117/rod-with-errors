# Rod with errors

The idea behind this small, quick library is to extend rod pages so that they call OnError before timeout kills the page.

This is to save debug logs, take screenshots, etc.

It also allows you to call page.OnError() at any time to perform some error handling code that you choose.

The pages are completely interoperable with existing rod pages.
