package main

import(
	"os/exec"
	"os"
	"log"
	"fmt"
)

func main(){
	//contact information is a struct with a name and phone number in each slice
//	contactInformation := infoReader()
	name := "ben"
	number:="5136028241"
	output := terminalWriter(messageWriter(name),number)
	fmt.Println("output is:",output)
	cmd := exec.Command(output) // maybe there is something with this program
	    cmd.Stdout = os.Stdout
	    cmd.Stderr = os.Stderr
	    err := cmd.Run()
	    if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	    }
}
// am getting oascript sendMessage.applescipt 5136028241 \"hello\n\" when i run simple file
