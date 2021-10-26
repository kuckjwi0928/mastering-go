package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	pid, _, _ := syscall.Syscall(39, 0, 0, 0)
	uid, _, _ := syscall.Syscall(24, 0, 0, 0)
	fmt.Println(pid)
	fmt.Println(uid)
	fd := 1
	message := []byte{'H', 'e', 'l', 'l', 'o'}
	syscall.Write(fd, message)
	command := "/bin/ls"
	env := os.Environ()
	syscall.Exec(command, []string{"ls", "-a", "-x"}, env)
}
