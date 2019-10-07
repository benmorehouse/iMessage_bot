package main

import(
	"fmt"
	"encoding/csv"
	"log"
	"os"
	"sync"
	"bufio"
	"strings"
	"errors"
	"os/exec"
	"flag"
)

type contact struct{
	Name string
	Number string
}

func iMessage(){
	csv_filename := flag.String("file","", "csv data file for contacts and their names")
	message := flag.String("message","","message that will be sent to all participants")
	flag.Parse()

	if *csv_filename == ""{
		log.Fatal("You have not passed in a csv file")
	}

	wg := sync.WaitGroup{}

	file, err := os.Open(*csv_filename)

	if err != nil{
		log.Fatal("Not able to open the inputted CSV file")
	}

	file_reader := csv.NewReader(file) // file is io.reader that reads file

	data , err := file_reader.Read()
	var contactList []contact
	for err == nil{
		wg.Add(1)
		var temp contact
		go func(data []string){
			temp.Init(&wg,data)
			if temp.Name == ""{
				return
			}else{
				contactList = append(contactList, temp) // initializes all the contacts
			}
		}(data)
		data , err = file_reader.Read()
	}
	wg.Wait()

	if *message == ""{
		fmt.Println("Please enter in a message to send to all users")
		fmt.Println("To enter in their name, type \"NAME\" with spaces on both sides")
		fmt.Println("Press ENTER when finished")
		reader := bufio.NewReader(os.Stdin)
		*message , err = reader.ReadString('\n')
	}

	wg.Add(len(contactList))
	for i:=0;i<len(contactList);i++{
		go func(string){
			*message , err = contactList[i].CreateMessage(*message)
			if err != nil{
				log.Fatal("Entered message is empty")
			}
			contactList[i].SendMessage(&wg,*message)
		}(*message)

		if err != nil{
			log.Fatal("Unable to send message to",contactList[i].Name)
		}
	}
	wg.Wait()
}

func (this *contact) Init(wg *sync.WaitGroup, input []string){
	defer wg.Done()
	if len(input) < 2{
		return
	}

	this.Name = input[0]
	this.Number = input[1]
}

func (this *contact) SendMessage(wg *sync.WaitGroup, message string){
//osascript sendMessage.applescript 1235551234 "Hello there!"
	defer wg.Done()
	output := "osascript sendMessage.applescipt "
	output += this.Name + string(" ")
	output += message

	cmd := exec.Command(output)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func (this contact) CreateMessage(input string)(string, error){
	if input == ""{
		err := errors.New("Error: message is empty")
		return input, err
	}

	messageField := strings.Fields(input)
	for i , val := range messageField{
		if val == "NAME"{
			messageField[i] = this.Name
		}
	}

	return strings.Join(messageField,string(" ")),nil
}

