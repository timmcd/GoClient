package main
import (
	"os";
	"flag";
	"net";
	"fmt";
	"bufio";
	"strings";
	)

func main() {
	raddr, _ := net.ResolveTCPAddr(flag.Arg(0));
	conn, _ := net.DialTCP("tcp", nil, raddr);
	bytec := make(chan int);
	go readLoop(bytec, conn);
	go writeLoop(conn);
	<-bytec;
}

func readLoop(bytec chan int, conn *net.TCPConn) {
	buf := make([]byte, 1);
	for {
		_, err := conn.Read(buf);
		if err != nil {
			if err == os.EOF {
				fmt.Print("Disconnected from the MUD!");
				break;
			}
		}
		fmt.Print(string(buf[0]));
	}
	bytec <- 1;
}

func writeLoop(conn *net.TCPConn) {
	in := bufio.NewReader(os.Stdin);
	for {
		line, _ := in.ReadString('\n');
		conn.Write(strings.Bytes(line))
	}
}