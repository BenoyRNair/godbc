# Copyright (c) 2009 Benoy R Nair. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

TARG=godbc
CGOFILES=\
	godbc.go
CGO_LDFLAGS=cgodbc.o -lodbc

CLEANFILES+=cgodbc.o example

include $(GOROOT)/src/Make.pkg

godbc_godbc.so: cgodbc.o godbc.cgo4.o
	gcc $(_CGO_CLFAGS_$(GOARCH)) $(_CGO_LDFLAGS_$(GOOS)) -o $@ godbc.cgo4.o $(CGO_LDFLAGS)

example: install example.go
	$(GC) example.go
	$(LD) -o $@ example.$O

cgodbc.o: cgodbc.c cgodbc.h
	gcc -o cgodbc.o -c cgodbc.c
