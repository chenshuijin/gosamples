package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type GraceSvr struct {
	ch         chan bool
	waitGroup  *sync.WaitGroup
	l          *net.TCPListener
	remoteHost string
	timeOut    int
}

func New(netName, laddr, remoteHost string) (*GraceSvr, error) {
	if os.Getenv("grace") == "true" {
		return newFromFD(3, remoteHost)
	} else {
		return new(netName, laddr, remoteHost)
	}
}

func new(netName, laddr, remoteHost string) (*GraceSvr, error) {
	log.Println("start from listen...")
	addr, err := net.ResolveTCPAddr(netName, laddr)
	if err != nil {
		log.Fatal("an error occured : ", err)
		return nil, err
	}
	ln, err := net.ListenTCP(netName, addr)
	if err != nil {
		log.Fatal("an error occured : ", err)
		return nil, err
	}
	log.Println("listening on :", laddr)
	return &GraceSvr{
		ch:         make(chan bool),
		waitGroup:  &sync.WaitGroup{},
		l:          ln,
		remoteHost: remoteHost,
	}, nil
}

func newFromFD(fd uintptr, remoteHost string) (*GraceSvr, error) {
	log.Println("start from fd...")
	file := os.NewFile(fd, "")
	ln, err := net.FileListener(file)
	if err != nil {
		log.Fatalln("an error occured : ", err)
		return nil, err
	}

	log.Println("listening on flie:", file)
	return &GraceSvr{
		ch:         make(chan bool),
		waitGroup:  &sync.WaitGroup{},
		l:          ln.(*net.TCPListener),
		remoteHost: remoteHost,
	}, nil
}

func (self *GraceSvr) Start() error {
	log.Println("start ...")
	go self.Accept()
	sigs := make(chan os.Signal)
	log.Println("waiting for signal...")
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	log.Println("signal")
	for sig := range sigs {
		switch sig {
		case syscall.SIGTERM:
			log.Println("sigterm...")
			self.Stop()
			self.l.SetDeadline(time.Now().Add(time.Duration(self.timeOut) * time.Second))
			os.Exit(-127)
		case syscall.SIGHUP, syscall.SIGINT:
			log.Println("begin to restart...")
			self.Restart()
		default:
			log.Println("exit...")

			os.Exit(0)
		}
	}
	log.Println("start exit")
	return nil
}

func (self *GraceSvr) Stop() error {
	self.waitGroup.Wait()
	return nil
}

func (self *GraceSvr) Restart() {
	file, err := self.l.File()
	if err != nil {
		log.Fatalln("get socket file err: ", err)
	}
	lnfd := file.Fd()
	os.Setenv("grace", "true")
	log.Println("restart fd:", lnfd)
	execSpec := &syscall.ProcAttr{
		Env: os.Environ(),
		Files: []uintptr{
			os.Stdin.Fd(),
			os.Stdout.Fd(),
			os.Stderr.Fd(),
			lnfd,
		},
	}
	fork, err := syscall.ForkExec(os.Args[0], os.Args, execSpec)
	if err != nil {
		log.Fatal("start new process err:", err)
	}
	log.Println("SIGHUP received: fork-exec to", fork)
	self.waitGroup.Wait()
	os.Exit(0)
}

func (self *GraceSvr) Accept() {
	log.Println("begin to accept...")
	for {
		conn, err := self.l.Accept()
		if err != nil {
			log.Fatal("accept err:", err)
			continue
		}
		log.Println("new conn:", conn)
		go self.handleConnection(conn, self.remoteHost)
	}
}

func (self *GraceSvr) handleConnection(conn net.Conn, remoteHost string) {
	self.waitGroup.Add(1)
	defer self.waitGroup.Done()
	defer conn.Close()

	addr, err := net.ResolveTCPAddr("tcp", remoteHost)
	if err != nil {
		log.Println("err:", err)
		return
	}
	remoteConn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Println("err:", err)
	}
	defer remoteConn.Close()

	for {
		data := make([]byte, 2048)
		count, err := conn.Read(data)
		if err != nil {
			log.Println("err:", err)
			return
		}
		log.Println("recv count:", count)
		log.Printf("recv data:\n%s\n", data)
		count, err = remoteConn.Write(data)
		if err != nil {
			log.Println("err:", err)
			return
		}

		log.Println("send count:", count)

		count, err = remoteConn.Read(data)
		if err != nil {
			log.Println("err:", err)
			return
		}
		log.Println("read from remote count:", count)
		log.Printf("read from remote data:\n%s\n", data)
		count, err = conn.Write(data)
		if err != nil {
			log.Println("err:", err)
			return
		}
		log.Println("write to local count:", count)
		log.Printf("write to local data:\n%s\n", data)
	}

	return
}
