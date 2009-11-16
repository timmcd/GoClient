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
	fmt.Print("\x1B[0;58r");
	fmt.Print("\x1B[59;1H");
	fmt.Print("---------------------------------------------------------------------------------");
	raddr, _ := net.ResolveTCPAddr(flag.Arg(0));
	conn, _ := net.DialTCP("tcp", nil, raddr);
	bytec := make(chan int);
	go readLoop(bytec, conn);
	go writeLoop(conn);
	<-bytec;
	fmt.Print("\x1B[r");
}

func readLoop(bytec chan int, conn *net.TCPConn) {
	buf := make([]byte, 1);
	var out string;
	for {
		_, err := conn.Read(buf);
		if err != nil {
			if err == os.EOF {
				fmt.Print("Disconnected from the MUD!");
				break;
			}
		}
		out += string(buf[0]);
		if string(buf[0]) == "\n" {
			fmt.Print("\x1B[58;1H");
			fmt.Print(out);
			fmt.Print("\x1B[60;1H");
			out = "";
		}
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