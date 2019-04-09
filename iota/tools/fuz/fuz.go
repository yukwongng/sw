package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/pensando/sw/iota/tools/fuz/fuze"
	"github.com/pkg/errors"
)

const (
	helloString = "HELLO FUZ\n"
	logfie      = "fuz.log"
)

var jsonOut bool

func printMsg(msg string) {
	if !jsonOut {
		fmt.Print(msg)
	} else {
		log.Println(msg)
	}
}

func main() {
	var proto string
	var conns, rate int
	var timeDuration time.Duration
	var talk bool
	var port int
	var cps int
	var jsonInput string
	var input fuze.Input

	f, err := os.Create(logfie)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	flag.BoolVar(&talk, "talk", false, "talk starts concurrent clients, default is to run as a server")
	flag.BoolVar(&jsonOut, "jsonOut", false, "Print Json output")
	flag.StringVar(&jsonInput, "jsonInput", "", "Json Input")
	flag.IntVar(&conns, "conns", 1, "number of concurrent clients")
	flag.IntVar(&cps, "cps", 10, "number of connections per second")
	flag.IntVar(&port, "port", 12000, "Default port to listen or connect to ")
	flag.StringVar(&proto, "proto", "tcp", "protocol: [tcp | udp]")
	flag.IntVar(&rate, "rate", 100, "Bytes Per Second")
	defaultDuration, _ := time.ParseDuration("10s")
	flag.DurationVar(&timeDuration, "duration", defaultDuration, "Duation for which client should keep the connection open")

	flag.Parse()

	connDatas := []*fuze.ConnectionData{}
	duration := (int)(timeDuration.Seconds())
	if jsonInput != "" {

		// Open our jsonFile
		jsonFile, err := os.Open(jsonInput)
		// if we os.Open returns an error then handle it
		if err != nil {
			printMsg(err.Error())
		}
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()
		// read our opened xmlFile as a byte array.
		byteValue, _ := ioutil.ReadAll(jsonFile)
		if err := json.Unmarshal(byteValue, &input); err != nil {
			printMsg(err.Error())
		}

		for _, conn := range input.Connections {
			connDatas = append(connDatas, fuze.NewConnData(conn.ServerIPPort, conn.Proto, timeDuration, !talk))
		}

	} else {
		if duration < 1 {
			printMsg("error - duration must be > 0\n")
			return
		}
		if proto != "tcp" && proto != "udp" {
			printMsg("protocol must be either udp or tcp\n")
			return
		}
		if conns > 10000 || conns < 1 {
			printMsg("error specifying the number of connections, must be between 1-10000\n")
			return
		}

		ipPorts := flag.Args()
		/*if !talk && len(ipPorts) > 1 {
			fmt.Printf("server can only listen on one ip:port\n")
			return
		}*/
		if len(ipPorts) == 0 {
			ipPorts = []string{fmt.Sprintf("0.0.0.0:%d", port)}
		}
		for idx, ipPort := range ipPorts {
			strs := strings.Split(ipPort, ":")
			if len(strs) == 0 {
				ipPorts[idx] = ipPorts[idx] + fmt.Sprintf(":%d", port)
			}
			if strings.Contains(ipPort, ".") && net.ParseIP(strs[0]) == nil {
				printMsg(fmt.Sprintf("error parsing IP %s\n", ipPort))
				return
			}
		}
		printMsg(fmt.Sprintf("IP Ports: %+v\n", ipPorts))

		for _, ipport := range ipPorts {
			connDatas = append(connDatas, fuze.NewConnData(ipport, proto, timeDuration, !talk))
		}

	}
	// start server or client

	debug.SetMaxThreads(30000)
	done := make(chan error)
	if talk {
		go runClient(proto, connDatas, conns, rate, cps, duration, done)
	} else {
		go runServer(proto, connDatas, done)
	}

	exitFunc := func() {
		errMsg := ""
		fuzD := &fuze.Output{Connections: connDatas, ErrorMsg: errMsg}
		for _, connData := range connDatas {
			if connData.ErrorMsg != "" {
				fuzD.ErrorMsg = connData.ErrorMsg
			} else {
				//fuzD.SuccessConnections++
			}
			fuzD.FailedConnections += connData.Failed
			fuzD.SuccessConnections += connData.Success
		}

		if jsonOut {
			fuzJSON, _ := json.Marshal(fuzD)
			os.Stdout.Write(fuzJSON)
		} else {
			fmt.Fprintf(os.Stdout, "Fuz talk success")
		}

		if fuzD.FailedConnections != 0 {
			os.Exit(1)
		}
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigs:
		exitFunc()

	case <-done:
		exitFunc()
	}
}

func worker(id int, jobs <-chan net.Conn) {
	workingConns := []net.Conn{}

	bytesReceived := 0
	bytesSent := 0
	for true {
		for true {
			breakup := false
			select {
			case c := <-jobs:
				printMsg(fmt.Sprintf("client %s connected\n", c.RemoteAddr().String()))
				workingConns = append(workingConns, c)
			case <-time.After(20 * time.Millisecond):
				breakup = true
			}
			if breakup {
				break
			}
		}

		if len(workingConns) == 0 {
			time.Sleep(10 * time.Millisecond)
			continue
		}
		newWorkingConns := []net.Conn{}
		for _, wc := range workingConns {
			wc.SetReadDeadline(time.Now().Add(1 * time.Millisecond))
			in, err := bufio.NewReader(wc).ReadString('\n')
			if err != nil {
				if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
					//socket timeout, add it for later
				} else {
					if err != io.EOF {
						printMsg(err.Error())
					}
					//Connection done, no need to add it to working connection
					wc.Close()
				}
			} else {
				bytesReceived += len(in)
				len, _ := wc.Write([]byte(in))
				bytesSent += len
			}

			if err == nil || (err != io.EOF && (err.(net.Error).Timeout())) {
				newWorkingConns = append(newWorkingConns, wc)
			}
		}
		workingConns = newWorkingConns

	}
}

func runServer(proto string, connDatas []*fuze.ConnectionData,
	done chan<- error) {
	waitCh := make(chan error)
	workers := runtime.GOMAXPROCS(0) - 2
	if workers <= 0 {
		workers = 1
	}
	jobs := make(chan net.Conn, 16384)
	for w := 1; w <= workers; w++ {
		go worker(w, jobs)
	}
	for _, connData := range connDatas {
		connData := connData
		go func() {
			printMsg(fmt.Sprintf("running %s server on %s\n", connData.Proto, connData.ServerIPPort))
			l, err := net.Listen(connData.Proto, connData.ServerIPPort)
			if err != nil {
				printMsg(err.Error())
				waitCh <- fmt.Errorf("error listening %s, protocol %s, '%s'\n", connData.ServerIPPort, connData.Proto, err)
				return
			}
			defer l.Close()

			for {
				c, err := l.Accept()
				if err != nil {
					printMsg(err.Error())
					continue
				}
				connData.ClientIPPort = c.RemoteAddr().String()
				printMsg(fmt.Sprintf("client %s accepted\n", c.RemoteAddr().String()))
				jobs <- c
			}
		}()
	}

	for ii := 0; ii < len(connDatas); ii++ {
		err := <-waitCh
		if err != nil {
			printMsg(err.Error())
			done <- errors.Wrapf(err, "Error starting server")
			return
		}
	}

	done <- nil

}

var mutex sync.Mutex

func runClient(proto string, connDatas []*fuze.ConnectionData, conns, rate, cps, duration int,
	err chan<- error) {
	waitCh := make(chan error, len(connDatas)*conns)
	randomBytes := RandomBytes(rate)

	numConnsSpawned := 0
	totalSpawned := 0
	for _, connData := range connDatas {
		for ii := 0; ii < conns; ii++ {
			//connData := connData
			numConnsSpawned++
			totalSpawned++
			if numConnsSpawned > cps {
				time.Sleep(time.Second)
				numConnsSpawned = 0
			}
			go func(connData *fuze.ConnectionData) {
				start := time.Now()
				var msg error
				bytesReceived := 0
				bytesSent := 0
				updateStats := func() {
					printMsg(fmt.Sprintf("Initiating completion  to : %s\n", connData.ServerIPPort))
					mutex.Lock()
					connData.ConnDurtion = time.Since(start)
					connData.DataReceived += bytesReceived
					connData.DataSent += bytesSent
					if msg != nil {
						connData.ErrorMsg = msg.Error()
						atomic.AddInt32(&connData.Failed, 1)
					} else {
						atomic.AddInt32(&connData.Success, 1)
					}
					mutex.Unlock()
					printMsg(fmt.Sprintf("Completed connection to : %s Total : %v, Spawned : %v )\n", connData.ServerIPPort, len(connDatas)*1, totalSpawned))
					waitCh <- msg
				}
				conn, err := net.Dial(connData.Proto, connData.ServerIPPort)
				printMsg(fmt.Sprintf("Initiating connection to : %s\n", connData.ServerIPPort))
				if err != nil {
					msg = fmt.Errorf("error dialing %s, protocol %s, '%s'\n", connData.ServerIPPort, connData.Proto, err)
					updateStats()
					return
				}
				defer conn.Close()

				connData.ClientIPPort = conn.LocalAddr().String()
				printMsg(fmt.Sprintf("Established connection to : %s\n", connData.ServerIPPort))
				for ticks := duration; ticks > 0; ticks-- {
					written, err := fmt.Fprintf(conn, randomBytes)
					if err != nil {
						msg = fmt.Errorf("error '%s' writin to socket", err)
						updateStats()
						return
					}

					ch := make(chan error)
					go func() { // this goroutime still exist even when timeout
						bytesSent += written
						readData, err := bufio.NewReader(conn).ReadString('\n')
						if err != nil {
							msg = fmt.Errorf("error '%s' reading from connection", err)
							ch <- msg
							return
						}
						if readData != randomBytes {
							msg = fmt.Errorf("unexpected read value from client '%s'(%v)", readData, len(readData))
							ch <- msg
							return
						}
						bytesReceived += len(randomBytes)
						ch <- nil
					}()
					select {
					case <-ch:
						if msg != nil {
							updateStats()
							return
						}
					case <-time.After(time.Duration(90) * time.Second):
						msg = fmt.Errorf("Timeout on reading from server")
						updateStats()
						return
					}
					time.Sleep(time.Second)
				}

				msg = nil
				updateStats()
			}(connData)
		}
	}

	for ii := 0; ii < len(connDatas)*conns; ii++ {
		err := <-waitCh
		if err != nil {
			printMsg(err.Error())
		}
	}

	err <- nil
}

func RandomBytes(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b) + "\n"
}
