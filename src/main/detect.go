package main

import (
	"net"
	"fmt"
	"strconv"
	"sync"
	"strings"
	"bytes"
	"os"
)


var w sync.WaitGroup

func detect_one(ip_str string,port_int int,ch chan  int){
	defer func() {
		<-ch
		w.Done()
	}()
	//fmt.Println("detecting : ",ip_str,strconv.Itoa(port_int))
	conn,err:=net.Dial("tcp",ip_str+":"+strconv.Itoa(port_int))
	if err!=nil{
		//fmt.Println("aaaaa",conn)
		return
	}else{
		fmt.Println("detected:",ip_str,":",port_int)
	}
	defer conn.Close()


}

func detect(ip_str string)  {
	//fmt.Println("detecting:",ip_str)
	for port:=21;port<=32;port++{
		ch:=make(chan int,50)
		ch <-1
		w.Add(1)
		go detect_one(ip_str,port,ch)
	}
	w.Wait()
}


func StringIpToInt(ipstring string) int {
	ipSegs := strings.Split(ipstring, ".")
	var ipInt int = 0
	var pos uint = 24
	for _, ipSeg := range ipSegs {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | tempInt
		pos -= 8
	}
	return ipInt
}

func IpIntToString(ipInt int) string {
	ipSegs := make([]string, 4)
	var len int = len(ipSegs)
	buffer := bytes.NewBufferString("")
	for i := 0; i < len; i++ {
		tempInt := ipInt & 0xFF
		ipSegs[len-i-1] = strconv.Itoa(tempInt)
		ipInt = ipInt >> 8
	}
	for i := 0; i < len; i++ {
		buffer.WriteString(ipSegs[i])
		if i < len-1 {
			buffer.WriteString(".")
		}
	}
	return buffer.String()
}


func main() {

	if (len(os.Args)!=3){
		return
	}



	//detect("10.10.3.227")

	for ip:=StringIpToInt(os.Args[1]);ip<=StringIpToInt(os.Args[2]);ip++{
		detect(IpIntToString(ip))
	}










}