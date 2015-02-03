TDD Go Logo
===========

This project is an idea for a TDD exercise in using Go.  It is pretty bare.  You need to implement the real code!

Idea:
-----

Create a basic [Logo](http://en.wikipedia.org/wiki/Logo_%28programming_language%29) interpreter in Go lang.  It can just be a small subset of the language like:

    forward <pixels> 
    right <deg>
    left <deg> 
    penup
    pendown

See [ACSLogo For Mac OS X](http://www.alancsmith.co.uk/logo/) as a real Logo implementation and to compare your output.  ACSLogo's help menu is a handy reference for Logo commands and usage.

Sample usage:
------------

    ./tdd_go_logo <my_logo_file>

It will output the results to a PNG file.

To keep things simple, its not interactive and just outputs to a file.

The [draw2d package](https://code.google.com/p/draw2d/) can be used as an library to draw to an image.

See [Samples](https://code.google.com/p/draw2d/wiki/Samples) for example code using the `draw2d` package.

Setup:
------

1. Install Go lang and set your GOPATH environment variable

2. Install the pre-requisite library `freetype-go`:

    go get code.google.com/p/freetype-go/freetype

4. Install the `draw2d` package:

    go get code.google.com/p/draw2d/draw2d

5. Run the example starting app:

    go build

    ./tdd_go_logo draw_square.txt

6. Check the resulting image:

    open Logo-output.png 

6. Start TDDing your tests for the real implementation!

You can use the built-in `testing` package or something like [goconvey](http://goconvey.co/).

Extensions:
-----------

* Display results in a window rather than a file
* Interactive use
* Support more Logo commands

