package main

import (
	"fmt"
	"github.com/jvmgo/ch03/classpath"
	"strings"
)

func main() {
	var cmd = parseCmd()
	if cmd.versionFlag {
		fmt.Printf("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, e := cp.ReadClass(className)
	if e != nil {
		fmt.Printf("Could not find or load class %s\n", cmd.class)
		return
	}
	fmt.Printf("class data:%v \n", classData)

}
