package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	port int
	help   bool

)

type result struct {
	host string
	name string
	finger string
}

var (
	socketChan chan bool
	resultChan chan *result
)


func main(){
	logo :=
"    __                __ __  _ _____                \n" +
"   / /   ____  ____ _/ // / (_) ___/_________ _____ \n  " +
  "/ /   / __ \\/ __ `/ // /_/ /\\__ \\/ ___/ __ `/ __ \\\n " +
 "/ /___/ /_/ / /_/ /__  __/ /___/ / /__/ /_/ / / / /\n" +
"/_____/\\____/\\__, /  /_/_/ //____/\\___/\\__,_/_/ /_/ \n " +
	          "/____/    /___/                         \n" +
	"    coded by 天下大木头"
	fmt.Println(logo)
	parserInput()
	now := time.Now()
	socketChan = make(chan bool)
	resultChan = make(chan *result)
	go TcpStart(port)
	for{
		select {
		case <-socketChan:
			fmt.Printf("[+] Log4j2Vuln Detected\n")
			res := <-resultChan
			fmt.Println((*res).host,(*res).name)
			content := fmt.Sprintf("Host: %s is vulnerable !\n",(*res).host)
			writeFile(fmt.Sprintf("log4j2ScanRes-%d.txt",now.Unix()),content)
		}
	}

}

func writeFile(filename string,content string) {
	outputFile, outputError := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if outputError != nil {
		fmt.Println("[!] An error Occured with file opening or creation\n")
		return
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	outputWriter.WriteString(content)
	outputWriter.Flush()
}

func parserInput() {
	flag.IntVar(&port, "p", 8001, "detect port")
	flag.BoolVar(&help, "help", false, "help info")
	flag.Parse()
	if help {
		flag.PrintDefaults()
		return
	}
}


