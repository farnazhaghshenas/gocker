package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"
	"syscall"
)



func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	name :=RandStringRunes(6)
	err := os.Mkdir(name, 0777)
	if err != nil{
		log.Fatal(err)
	}
	err = os.Chdir(name)
	if err != nil{
		log.Fatal(err)
	}
	wd, err := os.Getwd()
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(wd)
 	origin, err := os.Open("../"+os.Args[1])
 	if err != nil{
 		log.Fatal(err)
	}
	destination, err := os.OpenFile(os.Args[1], os.O_CREATE|os.O_RDWR, 0777)
	if err != nil{
		log.Fatal(err)
	}
	_, err =io.Copy(destination, origin)
	if err != nil{
		log.Fatal(err)
	}
	destination.Close()
	origin.Close()

	cmd := exec.Command("bash", os.Args[1])
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:  syscall.CLONE_NEWUTS,
	}

	err = cmd.Start()
	if err != nil{
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\ndone\n")

}

