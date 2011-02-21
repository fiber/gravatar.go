
include $(GOROOT)/src/Make.inc

TARG=github.com/fiber/gravatar

GOFILES=gravatar.go

include $(GOROOT)/src/Make.pkg

fmt:
	gofmt -w *.go
