// Copyright (C) 2023 Paul Gorman
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"crypto/tls"
	"flag"
	"io"
	"log"
	"net"
)

var optAddr string
var optCertFile string
var optKeyFile string
var optLogLevel int
var optPort string

func init() {
	flag.StringVar(&optAddr, "addr", "127.0.0.1", "IP address on which to serve the MUD")
	flag.StringVar(&optCertFile, "cert", "cert.pem", "TLS certificate file")
	flag.StringVar(&optKeyFile, "key", "key.pem", "TLS key file")
	flag.StringVar(&optPort, "port", "2323", "port on which to serve web interface")
	flag.Parse()
}

func main() {
	cert, err := tls.LoadX509KeyPair(optCertFile, optKeyFile)
	if err != nil {
		log.Fatal(err)
	}
	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
	listener, err := tls.Listen("tcp", optAddr+":"+optPort, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			// Echo all incoming data.
			io.Copy(c, c)
			// Shut down the connection.
			c.Close()
		}(conn)
	}

}
