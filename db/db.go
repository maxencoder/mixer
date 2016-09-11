package db

import (
	"container/list"
	"fmt"
	. "github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/client"
	"sync"
	"sync/atomic"
)

type DB struct {
	sync.Mutex

	addr         string
	user         string
	password     string
	db           string
	maxIdleConns int

	idleConns *list.List

	connNum int32
}

func Open(addr string, user string, password string, dbName string) (*DB, error) {
	db := new(DB)

	db.addr = addr
	db.user = user
	db.password = password
	db.db = dbName

	db.idleConns = list.New()
	db.connNum = 0

	return db, nil
}

func (db *DB) Addr() string {
	return db.addr
}

func (db *DB) String() string {
	return fmt.Sprintf("%s:%s@%s/%s?maxIdleConns=%v",
		db.user, db.password, db.addr, db.db, db.maxIdleConns)
}

func (db *DB) Close() error {
	db.Lock()

	for {
		if db.idleConns.Len() > 0 {
			v := db.idleConns.Back()
			co := v.Value.(*client.Conn)
			db.idleConns.Remove(v)

			co.Close()

		} else {
			break
		}
	}

	db.Unlock()

	return nil
}

func (db *DB) Ping() error {
	c, err := db.PopConn()
	if err != nil {
		return err
	}

	err = c.Ping()
	db.PushConn(c, err)
	return err
}

func (db *DB) SetMaxIdleConnNum(num int) {
	db.maxIdleConns = num
}

func (db *DB) GetIdleConnNum() int {
	return db.idleConns.Len()
}

func (db *DB) GetConnNum() int {
	return int(db.connNum)
}

func (db *DB) newConn() (*client.Conn, error) {
	co, err := client.Connect(db.addr, db.user, db.password, db.db)
	if err != nil {
		return nil, err
	}

	return co, nil
}

func (db *DB) tryReuse(co *client.Conn) error {
	if co.IsInTransaction() {
		//we can not reuse a connection in transaction status
		if err := co.Rollback(); err != nil {
			return err
		}
	}

	if !co.IsAutoCommit() {
		//we can not  reuse a connection not in autocomit
		if _, err := co.Execute("set autocommit = 1"); err != nil {
			return err
		}
	}

	//connection may be set names early
	//we must use default utf8
	if co.GetCharset() != DEFAULT_CHARSET {
		if err := co.SetCharset(DEFAULT_CHARSET); err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) PopConn() (co *client.Conn, err error) {
	db.Lock()
	if db.idleConns.Len() > 0 {
		v := db.idleConns.Front()
		co = v.Value.(*client.Conn)
		db.idleConns.Remove(v)
	}
	db.Unlock()

	if co != nil {
		if err := co.Ping(); err == nil {
			if err := db.tryReuse(co); err == nil {
				//connection may alive
				return co, nil
			}
		}
		co.Close()
	}

	co, err = db.newConn()
	if err == nil {
		atomic.AddInt32(&db.connNum, 1)
	}
	return
}

func (db *DB) PushConn(co *client.Conn, err error) {
	var closeConn *client.Conn = nil

	if err != nil {
		closeConn = co
	} else {
		if db.maxIdleConns > 0 {
			db.Lock()

			if db.idleConns.Len() >= db.maxIdleConns {
				v := db.idleConns.Front()
				closeConn = v.Value.(*client.Conn)
				db.idleConns.Remove(v)
			}

			db.idleConns.PushBack(co)

			db.Unlock()

		} else {
			closeConn = co
		}
	}

	if closeConn != nil {
		atomic.AddInt32(&db.connNum, -1)

		closeConn.Close()
	}
}

type SqlConn struct {
	*client.Conn

	db *DB
}

// XXX we should not always push err=nil; old driver used to track err
// state. PushConn decides wheather to close a conn based on it's err status.
// Change API to Close(err) or two discinct functions.
func (p *SqlConn) Close() {
	if p.Conn != nil {
		p.db.PushConn(p.Conn, nil)
		p.Conn = nil
	}
}

func (db *DB) GetConn() (*SqlConn, error) {
	c, err := db.PopConn()
	return &SqlConn{c, db}, err
}
