package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"
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
	stderr, err := cmd.StderrPipe()
	stdout, err := cmd.StdoutPipe()
	if err != nil{
		log.Fatal(err)
	}
	err = cmd.Start()
	if err != nil{
		log.Fatal(err)
	}

	stderrRead, err := ioutil.ReadAll(stderr)
	if err != nil{
		log.Fatal(err)
	}
	stdoutRead, err := ioutil.ReadAll(stdout)
	if err != nil{
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%s\n",stderrRead)
	fmt.Printf("\n%s\n",stdoutRead)

}

