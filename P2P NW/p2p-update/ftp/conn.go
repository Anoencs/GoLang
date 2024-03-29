package ftp

import (
	"net"
)

// Conn represents a connection to the FTP server
type Conn struct {
	conn     net.Conn
	dataType dataType
	dataPort *dataPort
	rootDir  string //place public file
	workDir  string //current working dir
}

// NewConn returns a new FTP connection
func NewConn(conn net.Conn, rootDir string) *Conn {
	return &Conn{
		conn:    conn,
		rootDir: rootDir,
		workDir: "/",
	}
}
