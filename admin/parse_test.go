// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlparser

import "testing"

func TestValid(t *testing.T) {
	checks := []struct {
		in  string
		out string
	}{{
		in: "add route r123 hash (type modulo, route btmdb-2005, btmdb-2007)",
	}, {
		in: "add route r123 range (0 to 100 route btmdb-2005)",
	}, {
		in: "add route r123 range (-100 to 100 route btmdb-2005, 100 to 200 route btmdb-2007)",
	}, {
		in: "add route r123 range (inf to 100 route btmdb-2005)",
	}, {
		in: "alter route r123 hash (type modulo, route btmdb-2005, btmdb-2007)",
	}, {
		in: "alter route r123 range (0 to 100 route btmdb-2005)",
	}, {
		in: "alter route r123 range (-100 to 100 route btmdb-2005, 100 to 200 route btmdb-2007)",
	}, {
		in: "alter route r123 range (inf to 100 route btmdb-2005)",
	}, {
		in: "delete route r123-withdash",
	}, {
		in:  "show routes",
		out: "show routes",
	}}

	for _, tcase := range checks {
		if tcase.out == "" {
			tcase.out = tcase.in
		}
		tree, err := Parse(tcase.in)
		if err != nil {
			t.Errorf("in: %s, err: %v", tcase.in, err)
			continue
		}
		out := String(tree)
		if out != tcase.out {
			t.Errorf("out: %s, want %s", out, tcase.out)
		}
	}
}

func TestInvalid(t *testing.T) {
	checks := []string{
		"add route r123 hash ()",
		"add route r123 range (1.334 to 232 route b11)",
		"add route range (-100 to 100 route btmdb-2005)",
		"add routes r123 range (infi to inf route btmdb-2005)",
	}

	for _, c := range checks {
		if _, err := Parse(c); err == nil {
			t.Errorf("expected parse error for: %s", c)
		} else {
			t.Log(err)
		}
	}
}
