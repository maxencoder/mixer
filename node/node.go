package node

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/maxencoder/log"
	"github.com/maxencoder/mixer/config"
	"github.com/maxencoder/mixer/db"
)

const (
	Master = "master"
	Slave  = "slave"

	KeepAliveInterval = 1
)

type Node struct {
	sync.Mutex

	name string

	user     string
	password string

	idleConns int

	master *db.DB
	slave  *db.DB

	downAfterNoAlive time.Duration

	lastMasterPing time.Time
	lastSlavePing  time.Time
}

func NewNode(name, user, password string, idleConns, downAfterNoAlive int, master string, slaves []string) (*Node, error) {
	n := new(Node)

	n.name = name

	n.user = user
	n.password = password

	n.idleConns = idleConns

	n.downAfterNoAlive = time.Duration(downAfterNoAlive) * time.Second

	if len(master) == 0 {
		return nil, fmt.Errorf("node must have master defined")
	}

	var err error
	if n.master, err = n.openDB(master); err != nil {
		return nil, err
	}

	if len(slaves) == 0 {
		return nil, fmt.Errorf("node must have at least one slave defined")
	}

	// random slave
	slave := slaves[rand.Intn(len(slaves))]
	if n.slave, err = n.openDB(slave); err != nil {
		log.Error(err.Error())
		n.slave = nil
	}

	go n.run()

	return n, nil
}

func NewNodeFromConfig(cfg config.NodeConfig) (*Node, error) {

	return NewNode(
		cfg.Name,
		cfg.User,
		cfg.Password,
		cfg.IdleConns,
		cfg.DownAfterNoAlive,
		cfg.Master,
		cfg.Slaves,
	)
}

func (n *Node) run() {
	//to do
	//1 check connection alive
	//2 check remove mysql server alive

	t := time.NewTicker(KeepAliveInterval * time.Second)
	defer t.Stop()

	n.lastMasterPing = time.Now()
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
	return n.name
}

func (n *Node) GetMasterConn() (*db.SqlConn, error) {
	n.Lock()
	db := n.master
	n.Unlock()

	if db == nil {
		return nil, fmt.Errorf("master is down")
	}

	return db.GetConn()
}

func (n *Node) GetSelectConn() (*db.SqlConn, error) {
	var db *db.DB

	n.Lock()
	db = n.slave
	n.Unlock()

	if db == nil {
		return nil, fmt.Errorf("no live mysql server")
	}

	return db.GetConn()
}

func (n *Node) checkMaster() {
	n.Lock()
	db := n.master
	n.Unlock()

	if db == nil {
		// log.Info("no master avaliable")
		return
	}

	if err := db.Ping(); err != nil {
		log.Error("%s ping master %s error %s", n, db.Addr(), err.Error())
	} else {
		n.lastMasterPing = time.Now()
		return
	}

	if n.downAfterNoAlive > 0 && time.Since(n.lastMasterPing) > n.downAfterNoAlive {
		log.Error("%s putting master db down %s", n, n.master.Addr())

		n.DownMaster()
	}
}

func (n *Node) checkSlave() {
	n.Lock()
	db := n.slave
	n.Unlock()

	if db == nil {
		return
	}

	if err := db.Ping(); err != nil {
		log.Error("%s ping slave %s error %s", n, db.Addr(), err.Error())
	} else {
		n.lastSlavePing = time.Now()
	}

	if n.downAfterNoAlive > 0 && time.Since(n.lastSlavePing) > n.downAfterNoAlive {
		log.Error("%s slave db %s not alive over %ds, down it",
			n, db.Addr(), int64(n.downAfterNoAlive/time.Second))

		n.DownSlave()
	}
}

func (n *Node) openDB(addr string) (*db.DB, error) {
	db, err := db.Open(addr, n.user, n.password, "")
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConnNum(n.idleConns)
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

func (n *Node) LastMasterPing() time.Time {
	return n.lastMasterPing
}

func (n *Node) LastSlavePing() time.Time {
	return n.lastSlavePing
}

func (n *Node) DownAfterNoAlive() time.Duration {
	return n.downAfterNoAlive
}
