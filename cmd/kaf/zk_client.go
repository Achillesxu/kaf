package main

import (
	"errors"
	"fmt"
	"github.com/go-zookeeper/zk"
	"log"
	"path"
	"strings"
	"time"
)

type nilLogger struct{}

func (nilLogger) Printf(format string, a ...interface{}) {
	// do nothing
}

type Auth struct {
	Scheme  string
	Payload []byte
}

type ZkClient struct {
	Servers        []string
	SessionTimeout time.Duration
	KafkaRootPath  string
	Auth           *Auth
	SilentLog      bool
	ZkConn         *zk.Conn
	ZkEvent        <-chan zk.Event
}

// NewZkClient create a zookeeper client to the Zookeeper ensemble where an Apache Kafka cluster maintains
// metadata, and reads some information from the /brokers/topics tree
func NewZkConfig(Servers []string, KafkaRootPath string, SessionTimeout time.Duration, SilentLog bool) *ZkClient {
	return &ZkClient{
		Servers:        Servers,
		KafkaRootPath:  KafkaRootPath,
		SessionTimeout: SessionTimeout,
		SilentLog:      SilentLog,
	}
}

// Connect connect to Zookeeper ensemble, Up to three attempts to connect
func (zc *ZkClient) Connect() (err error) {
	defaultLog := log.New(errWriter, "[go-zookeeper] ", log.Lshortfile|log.LstdFlags)
	logger := zk.WithLogger(defaultLog)
	if !zc.SilentLog {
		logger = zk.WithLogger(nilLogger{})
	}

	zc.ZkConn, zc.ZkEvent, err = zk.Connect(zc.Servers, zc.SessionTimeout, logger)
	if err != nil {
		return
	}
	if zc.Auth != nil {
		auth := zc.Auth
		err = zc.ZkConn.AddAuth(auth.Scheme, auth.Payload)
		if err != nil {
			return
		}
	}
	n := 0
	failed := false

loop:
	for {
		select {
		case event, ok := <-zc.ZkEvent:
			n += 1
			if ok && event.State == zk.StateConnected {
				break loop
			} else if n > 3 {
				failed = true
				break loop
			}
		}
	}
	if failed {
		err = errors.New(
			fmt.Sprintf("Failed to connect to %s!", strings.Join(zc.Servers, ",")),
		)
	}
	return
}

func (zc *ZkClient) LsDir(target string) (children []string, err error) {
	err = zc.IsConnected()
	if err != nil {
		return
	}
	cleanPath := path.Clean(target)

	children, _, err = zc.ZkConn.Children(cleanPath)
	if err != nil {
		return
	}
	return
}

func (zc *ZkClient) GetData(target string) (Data []byte, err error) {
	err = zc.IsConnected()
	if err != nil {
		return
	}
	cleanPath := path.Clean(target)

	Data, _, err = zc.ZkConn.Get(cleanPath)
	if err != nil {
		return
	}
	return
}

func (zc *ZkClient) IsConnected() (err error) {
	state := zc.ZkConn.State()
	if state == zk.StateConnected || state == zk.StateHasSession {
		err = nil
	} else {
		err = errors.New("connection is disconnected")
	}
	return
}

func (zc *ZkClient) Close() {
	zc.ZkConn.Close()
}
