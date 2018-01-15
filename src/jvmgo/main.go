package main
import "fmt"
import "strings"
import "jvmgo/classpath"

func main()  {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd)  {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("Xjre:%s classpath:%s class:%s args:%v\n", cmd.XjreOption, cmd.cpOption, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	fmt.Printf("class name:%v\n", className)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Counld not find or load main class %s\n error:%s\n", cmd.class, err)
		return 
	}
	fmt.Printf("class data:%v\n", classData)
}