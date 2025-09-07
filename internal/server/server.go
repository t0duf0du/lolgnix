package server

import (
	"bufio"
	"fmt"
	"net"
	"slices"
	"strconv"

	requests "github.com/codecrafters-io/http-server-starter-go/internal/requests"

	"github.com/k0kubun/pp"
)

type Server struct {
	/*
	   Address or URI is the address of our restaurant.
	*/
	Host         string // The building our restaurant is in.
	Port         int    // The bloody door.
	URL          string
	AllowedPaths []string
}

func NewServer(
	host string,
	port int,
	paths []string,
) (*Server, error) {
	portStr := strconv.Itoa(port)
	url := host + ":" + portStr

	return &Server{
		Host:         host,
		Port:         port,
		URL:          url,
		AllowedPaths: paths,
	}, nil
}

func (s Server) getURL() string {
	return s.URL
}

func (s Server) Run() error {
	/*
	  See my gangstas, the whole purpose of this(and upcoming excercises)
	  is to build shit from scratch. But we got to respect the fucking
	  abstractions dawg. We jump between abstractions one project/learning
	  unit a time. Otherwise we would be fabricating our own bloody
	  wafers.

	  We are learning HTTP, which is an application layer protocol. So we
	  we will be using a Protocol that sits one layer below http, wiz TCP.
	*/

	/*
	   Listener is kinda like our main waiter:
	   1. Gives the customer a table(connection)
	   2. Takes order of the customer (The HTTP request string)
	*/

	listener, err := net.Listen("tcp", s.getURL())
	if err != nil {
		return fmt.Errorf("error in creating listner, %s", err.Error())
	}

	defer listener.Close() // I don't know why we do this shit.

	for {
		/*
		   The waiter(listener) accepted incoming customer and gave them
		   a table(connection) to sit on.
		*/
		conn, err := listener.Accept()
		if err != nil {
			return fmt.Errorf(
				"error in accecpting connection for the listener : %v. \n error: %v",
				listener,
				err,
			)
		}

		s.handleRequest(conn)
	}
}

func (s Server) handleRequest(conn net.Conn) error {
	reader := bufio.NewReader(conn)
	for {
		requestStr, err := readHttpHeaders(reader)
		if err != nil {
			return fmt.Errorf(
				"error in reading the string: %v",
				err,
			)
		}

		rm, err := requests.NewRequestManager(requestStr)

		if slices.Contains(s.AllowedPaths, rm.R.URLPath) {
			pp.Println("We cool here now. All chrome no static")
		} else {
			pp.Println("AYO WATCH WHERE YOU GOING BOI!!!!!!")
		}

		conn.Write([]byte(requestStr))
	}
}

func readHttpHeaders(reader *bufio.Reader) (string, error) {
	var request []byte

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			return "", fmt.Errorf("error in reading bytes:%v", err)
		}

		request = append(request, line...)
		if string(line) == "\r\n" {
			break
		}
	}
	return string(request), nil
}
