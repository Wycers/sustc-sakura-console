package util

import (
	"net"
	"fmt"
	"io/ioutil"
)


func Exchange(str string) *[]byte {
	var remoteAddress, _ = net.ResolveTCPAddr("tcp4", "127.0.0.1:9090")
	session, err := net.DialTCP("tcp4", nil, remoteAddress)
	if err != nil {
		fmt.Println("Connect failed：", err)
		return nil
	}
	session.Write([]byte(str))
	bytes, err := ioutil.ReadAll(session)
	if err != nil {
		fmt.Println("Recive failed：", err)
		return nil
	}
	session.Close()
	return &bytes
}
