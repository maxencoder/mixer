# Copyright 2012, Google Inc. All rights reserved.
# Use of this source code is governed by a BSD-style license that can
# be found in the LICENSE file.

MAKEFLAGS = -s

admin.go: admin.y
	go tool yacc -o admin.go admin.y
	gofmt -w admin.go

clean:
	rm -f y.output admin.go
