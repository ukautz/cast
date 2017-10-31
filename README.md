[![Build Status](https://travis-ci.org/ukautz/cast.svg?branch=master)](https://travis-ci.org/ukautz/cast)
[![Coverage](http://gocover.io/_badge/github.com/ukautz/cast?v=0.1.3)](http://gocover.io/github.com/ukautz/cast)
[![GoDoc](https://godoc.org/github.com/ukautz/gpath?status.svg)](https://godoc.org/github.com/ukautz/gpath)

cast
====

cast is a library providing a sugar layer API for converting "unclean" (user) input into specific
low level Go types.

My primary use case: dealing with deep & complex YAML/JSON configuration files, which can
contain string("123") numbers or annoying float64(0.0) numbers where int is expected and so on.
See also github.com/ukautz/gpath to that end.

This library is not optimized for performance but for simplifying interaction with user input - and keeping
too often written, ugly, repetitive conditional type conversion code from my other projects.

See [godoc.org/github.com/ukautz/cast](https://godoc.org/github.com/ukautz/cast) 