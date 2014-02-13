package mysql

import (
	"reflect"
	"testing"
)

func TestStmt_DropTable(t *testing.T) {
	str := `drop table if exists mixer_test_stmt`

	c := newTestConn()

	s, err := c.Prepare(str)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := s.Exec(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestStmt_CreateTable(t *testing.T) {
	str := `CREATE TABLE IF NOT EXISTS mixer_test_stmt (
          id BIGINT(64) UNSIGNED  NOT NULL,
          str VARCHAR(256),
          f DOUBLE,
          e enum("test1", "test2"),
          u tinyint unsigned,
          i tinyint,
          PRIMARY KEY (id)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8`

	c := newTestConn()
	defer c.Close()

	s, err := c.Prepare(str)

	if err != nil {
		t.Fatal(err)
	}

	if _, err = s.Exec(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestStmt_Delete(t *testing.T) {
	str := `delete from mixer_test_stmt`

	c := newTestConn()
	defer c.Close()

	s, err := c.Prepare(str)

	if err != nil {
		t.Fatal(err)
	}

	if _, err := s.Exec(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestStmt_Insert(t *testing.T) {
	str := `insert into mixer_test_stmt (id, str, f, e, u, i) values (?, ?, ?, ?, ?, ?)`

	c := newTestConn()
	defer c.Close()

	s, err := c.Prepare(str)

	if err != nil {
		t.Fatal(err)
	}

	if pkg, err := s.Exec(1, "a", 3.14, "test1", 255, -127); err != nil {
		t.Fatal(err)
	} else {
		if pkg.AffectedRows != 1 {
			t.Fatal(pkg.AffectedRows)
		}
	}

	s.Close()
}

func TestStmt_Select(t *testing.T) {
	str := `select str, f, e from mixer_test_stmt where id = ?`

	c := newTestConn()
	defer c.Close()

	s, err := c.Prepare(str)
	if err != nil {
		t.Fatal(err)
	}

	if result, err := s.Query(1); err != nil {
		t.Fatal(err)
	} else {
		if len(result.Data) != 1 {
			t.Fatal(len(result.Data))
		}

		if len(result.Fields) != 3 {
			t.Fatal(len(result.Fields))
		}

		if str, _ := result.GetString(0, 0); str != "a" {
			t.Fatal("invalid str", str)
		}

		if f, _ := result.GetFloat(0, 1); f != float64(3.14) {
			t.Fatal("invalid f", f)
		}

		if e, _ := result.GetString(0, 2); e != "test1" {
			t.Fatal("invalid e", e)
		}

		if str, _ := result.GetStringByName(0, "str"); str != "a" {
			t.Fatal("invalid str", str)
		}

		if f, _ := result.GetFloatByName(0, "f"); f != float64(3.14) {
			t.Fatal("invalid f", f)
		}

		if e, _ := result.GetStringByName(0, "e"); e != "test1" {
			t.Fatal("invalid e", e)
		}

	}

	s.Close()
}

func TestStmt_ResultsetDump(t *testing.T) {
	str := `select * from mixer_test_stmt where id = ?`

	c := newTestConn()
	defer c.Close()

	s, err := c.Prepare(str)
	if err != nil {
		t.Fatal(err)
	}

	if result, err := s.Query(1); err != nil {
		t.Fatal(err)
	} else {
		p := new(resultsetPacket)
		p.Status = result.Status

		for i := range result.Fields {
			p.Fields = append(p.Fields, fieldPacket(result.Fields[i].Dump()))
		}

		if rows, err := result.DumpRows(); err != nil {
			t.Fatal(err)
		} else {
			for i := range rows {
				p.Rows = append(p.Rows, rowPacket(rows[i]))
			}
		}

		if r, err := p.Parse(true); err != nil {
			t.Fatal(err)
		} else {
			if !reflect.DeepEqual(r, result) {
				t.Fatal("result set not equal")
			}
		}

	}
}

func TestStmt_NULL(t *testing.T) {
	str := `insert into mixer_test_stmt (id, str, f, e) values (?, ?, ?, ?)`

	c := newTestConn()
	defer c.Close()

	s, err := c.Prepare(str)

	if err != nil {
		t.Fatal(err)
	}

	if pkg, err := s.Exec(2, nil, 3.14, nil); err != nil {
		t.Fatal(err)
	} else {
		if pkg.AffectedRows != 1 {
			t.Fatal(pkg.AffectedRows)
		}
	}

	s.Close()

	str = `select * from mixer_test_stmt where id = ?`
	s, err = c.Prepare(str)

	if err != nil {
		t.Fatal(err)
	}

	if r, err := s.Query(2); err != nil {
		t.Fatal(err)
	} else {
		if b, err := r.IsNullByName(0, "id"); err != nil {
			t.Fatal(err)
		} else if b == true {
			t.Fatal(b)
		}

		if b, err := r.IsNullByName(0, "str"); err != nil {
			t.Fatal(err)
		} else if b == false {
			t.Fatal(b)
		}

		if b, err := r.IsNullByName(0, "f"); err != nil {
			t.Fatal(err)
		} else if b == true {
			t.Fatal(b)
		}

		if b, err := r.IsNullByName(0, "e"); err != nil {
			t.Fatal(err)
		} else if b == false {
			t.Fatal(b)
		}
	}

	s.Close()
}

func TestStmt_Unsigned(t *testing.T) {
	str := `insert into mixer_test_stmt (id, u) values (?, ?)`

	c := newTestConn()
	defer c.Close()

	s, err := c.Prepare(str)

	if err != nil {
		t.Fatal(err)
	}

	if pkg, err := s.Exec(3, uint8(255)); err != nil {
		t.Fatal(err)
	} else {
		if pkg.AffectedRows != 1 {
			t.Fatal(pkg.AffectedRows)
		}
	}

	s.Close()

	str = `select u from mixer_test_stmt where id = ?`

	s, err = c.Prepare(str)
	if err != nil {
		t.Fatal(err)
	}

	if r, err := s.Query(3); err != nil {
		t.Fatal(err)
	} else {
		if u, err := r.GetUint(0, 0); err != nil {
			t.Fatal(err)
		} else if u != uint64(255) {
			t.Fatal(u)
		}
	}

	s.Close()
}

func TestStmt_Signed(t *testing.T) {
	str := `insert into mixer_test_stmt (id, i) values (?, ?)`

	c := newTestConn()
	defer c.Close()

	s, err := c.Prepare(str)

	if err != nil {
		t.Fatal(err)
	}

	if _, err := s.Exec(3, 255); err == nil {
		t.Fatal(err)
	}

	if _, err := s.Exec(uint64(18446744073709551516), int8(-128)); err != nil {
		t.Fatal(err)
	}

	s.Close()

}

func TestStmt_Trans(t *testing.T) {
	c := newTestConn()
	defer c.Close()

	if _, err := c.Exec(`insert into mixer_test_stmt (id, str) values (1002, "abc")`); err != nil {
		t.Fatal(err)
	}

	if err := c.Begin(); err != nil {
		t.Fatal(err)
	}

	str := `select str from mixer_test_stmt where id = ?`

	s, err := c.Prepare(str)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := s.Query(1002); err != nil {
		t.Fatal(err)
	}

	if err := c.Commit(); err != nil {
		t.Fatal(err)
	}

	if r, err := s.Query(1002); err != nil {
		t.Fatal(err)
	} else {
		if str, _ := r.GetString(0, 0); str != `abc` {
			t.Fatal(str)
		}
	}

	if err := s.Close(); err != nil {
		t.Fatal(err)
	}
}
