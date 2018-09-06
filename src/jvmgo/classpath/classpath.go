package classpath

import "os"
import "path/filepath"
import "fmt"

type ClassPath struct {
	bootClassPath Entry
	extClassPath Entry
	userClassPath Entry
}

func Parse(jreOption, cpOption string) * ClassPath {
	cp := &ClassPath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *ClassPath) parseBootAndExtClasspath(jreOption string)  {
	jreDir := getJreDir(jreOption)
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClassPath = newWildCardEntry(jreLibPath)
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClassPath = newWildCardEntry(jreExtPath)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		fmt.Printf("jreOption is " + jreOption)
		return jreOption
	}

	if exists("./jre") {
		return "./jre"
	}

	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	panic("Can not find jre folder")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

func (self *ClassPath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClassPath = newEntry(cpOption)
}


func (self *ClassPath) ReadClass(className string) ([]byte, Entry, error) {
	fmt.Printf("className:%s\n", className)
	className = className + ".class"
	data, entry, err := self.bootClassPath.readClass(className);
	if err == nil {
		return data, entry, err
	} else {
		fmt.Printf("self.bootClassPath.readClass err : %s\n", err)
		fmt.Printf("self.bootClassPath : %s\n", self.bootClassPath.String())
	}

	if data, entry, err := self.extClassPath.readClass(className); err == nil {
		return data, entry, err
	}

	return self.userClassPath.readClass(className)
}

func (self *ClassPath) String() string {
	return self.userClassPath.String()
}



