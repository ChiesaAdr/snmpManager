package main

import (
	"fmt"

	// "github.com/ChiesaAdr/snmpManager/resources"
	"github.com/gosnmp/gosnmp"
)

//TODO: Test this with snmpMocker (https://git.intelbras.com.br/olt-software/snmpmocker) project
func TestTrapListener() {
	// file := os.Stdout
	conn := gosnmp.Default
	conn.Target = "10.100.34.66"
	tl := gosnmp.NewTrapListener()
	// tl.OnNewTrap = resources.MyTrapHandler(file)
	tl.Params = conn
	err := tl.Listen("0.0.0.0:162")
	if err != nil {
		fmt.Println(err)
		return
	}
}
func main() {
	TestTrapListener()
}
