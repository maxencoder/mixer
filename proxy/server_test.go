package proxy

import (
	"sync"
	"testing"
	"time"

	"github.com/maxencoder/mixer/config"
	"github.com/maxencoder/mixer/db"
	"github.com/maxencoder/mixer/node"
)

var testServerOnce sync.Once
var testServer *Server
var testDBOnce sync.Once
var testDB *db.DB

var testConfigData = []byte(`
addr : 127.0.0.1:4000
user : root
password : 

nodes :
- 
    name : node1 
    down_after_noalive : 300
    idle_conns : 16
    user: root
    password:
    master : 127.0.0.1:3306
    slave : 
- 
    name : node2
    down_after_noalive : 300
    idle_conns : 16
    user: root
    password:
    master : 127.0.0.1:3307

- 
    name : node3 
    down_after_noalive : 300
    idle_conns : 16
    user: root
    password:
    master : 127.0.0.1:3308
`)

func newTestServer(t *testing.T) *Server {
	f := func() {
		cfg, err := config.ParseConfigData(testConfigData)
		if err != nil {
			t.Fatal(err.Error())
		}

		if err := node.ParseNodes(cfg); err != nil {
			t.Fatal(err)
		}

		testServer, err = NewServer(cfg)
		if err != nil {
			t.Fatal(err)
		}

		go testServer.Run()

		time.Sleep(1 * time.Second)
	}

	testServerOnce.Do(f)

	return testServer
}

func newTestDB(t *testing.T) *db.DB {
	newTestServer(t)

	f := func() {
		var err error
		testDB, err = db.Open("127.0.0.1:4000", "root", "", "mixer")

		if err != nil {
			t.Fatal(err)
		}

		testDB.SetMaxIdleConnNum(4)
	}

	testDBOnce.Do(f)
	return testDB
}

func newTestDBConn(t *testing.T) *db.SqlConn {
	db := newTestDB(t)

	c, err := db.GetConn()

	if err != nil {
		t.Fatal(err)
	}

	return c
}

func TestServer(t *testing.T) {
	newTestServer(t)
}
