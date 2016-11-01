package node

import (
	"fmt"
	"sync"
	"time"

	"github.com/maxencoder/mixer/config"
	"github.com/maxencoder/mixer/db"
	"github.com/maxencoder/log"
)

const (
	Master = "master"
	Slave  = "slave"
)

type Node struct {
	sync.Mutex

	cfg config.NodeConfig

	//running master db
	db *db.DB

	master *db.DB
	slave  *db.DB

	downAfterNoAlive time.Duration

	lastMasterPing int64
	lastSlavePing  int64
}

func NewNode(cfg config.NodeConfig) (*Node, error) {
	n := new(Node)
	n.cfg = cfg

	n.downAfterNoAlive = time.Duration(cfg.DownAfterNoAlive) * time.Second

	if len(cfg.Master) == 0 {
		return nil, fmt.Errorf("must setting master MySQL node.")
	}

	var err error
	if n.master, err = n.openDB(cfg.Master); err != nil {
		return nil, err
	}

	n.db = n.master

	if len(cfg.Slave) > 0 {
		if n.slave, err = n.openDB(cfg.Slave); err != nil {
			log.Error(err.Error())
			n.slave = nil
		}
	}

	go n.run()

	return n, nil
}

func (n *Node) run() {
	//to do
	//1 check connection alive
	//2 check remove mysql server alive

	t := time.NewTicker(3000 * time.Second)
	defer t.Stop()

	n.lastMasterPing = time.Now().Unix()
	n.lastSlavePing = n.lastMasterPing
	for {
		select {
		case <-t.C:
			n.checkMaster()
			n.checkSlave()
		}
	}
}

func (n *Node) String() string {
	return n.cfg.Name
}

func (n *Node) GetMasterConn() (*db.SqlConn, error) {
	n.Lock()
	db := n.db
	n.Unlock()

	if db == nil {
		return nil, fmt.Errorf("master is down")
	}

	return db.GetConn()
}

func (n *Node) GetSelectConn() (*db.SqlConn, error) {
	var db *db.DB

	n.Lock()
	if n.cfg.RWSplit && n.slave != nil {
		db = n.slave
	} else {
		db = n.db
	}
	n.Unlock()

	if db == nil {
		return nil, fmt.Errorf("no alive mysql server")
	}

	return db.GetConn()
}

func (n *Node) checkMaster() {
	n.Lock()
	db := n.db
	n.Unlock()

	if db == nil {
		log.Info("no master avaliable")
		return
	}

	if err := db.Ping(); err != nil {
		log.Error("%s ping master %s error %s", n, db.Addr(), err.Error())
	} else {
		n.lastMasterPing = time.Now().Unix()
		return
	}

	if int64(n.downAfterNoAlive) > 0 && time.Now().Unix()-n.lastMasterPing > int64(n.downAfterNoAlive) {
		log.Error("%s down master db %s", n, n.master.Addr())

		n.DownMaster()
	}
}

func (n *Node) checkSlave() {
	if n.slave == nil {
		return
	}

	db := n.slave
	if err := db.Ping(); err != nil {
		log.Error("%s ping slave %s error %s", n, db.Addr(), err.Error())
	} else {
		n.lastSlavePing = time.Now().Unix()
	}

	if int64(n.downAfterNoAlive) > 0 && time.Now().Unix()-n.lastSlavePing > int64(n.downAfterNoAlive) {
		log.Error("%s slave db %s not alive over %ds, down it",
			n, db.Addr(), int64(n.downAfterNoAlive/time.Second))

		n.DownSlave()
	}
}

func (n *Node) openDB(addr string) (*db.DB, error) {
	db, err := db.Open(addr, n.cfg.User, n.cfg.Password, "")
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConnNum(n.cfg.IdleConns)
	return db, nil
}

func (n *Node) checkUpDB(addr string) (*db.DB, error) {
	db, err := n.openDB(addr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func (n *Node) UpMaster(addr string) error {
	n.Lock()
	if n.master != nil {
		n.Unlock()
		return fmt.Errorf("%s master must be down first", n)
	}
	n.Unlock()

	db, err := n.checkUpDB(addr)
	if err != nil {
		return err
	}

	n.Lock()
	n.master = db
	n.db = db
	n.Unlock()

	return nil
}

func (n *Node) UpSlave(addr string) error {
	n.Lock()
	if n.slave != nil {
		n.Unlock()
		return fmt.Errorf("%s, slave must be down first", n)
	}
	n.Unlock()

	db, err := n.checkUpDB(addr)
	if err != nil {
		return err
	}

	n.Lock()
	n.slave = db
	n.Unlock()

	return nil
}

func (n *Node) DownMaster() error {
	n.Lock()
	n.db = nil
	n.master = nil
	n.Unlock()
	return nil
}

func (n *Node) DownSlave() error {
	n.Lock()
	db := n.slave
	n.slave = nil
	n.Unlock()

	if db != nil {
		db.Close()
	}

	return nil
}

func (n *Node) Master() *db.DB {
	return n.master
}

func (n *Node) Slave() *db.DB {
	return n.slave
}

func (n *Node) LastMasterPing() int64 {
	return n.lastMasterPing
}

func (n *Node) LastSlavePing() int64 {
	return n.lastSlavePing
}

func (n *Node) DownAfterNoAlive() time.Duration {
	return n.downAfterNoAlive
}
