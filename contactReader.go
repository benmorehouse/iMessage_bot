package main

import (
	"fmt"
	"bufio"
	"os"
	"flag"
	"strings"
)

type orientee struct{
	name string
	number string
}

func scan_Phone(input string)string{
	var output string
	for _,val := range input{
		if val == '(' || val == ')' || val =='-'{
			continue
		}else{
			output += string(val)
		}
	}
	return output
}

func infoReader()[]orientee{
	/*
	This program will take in the structure of student's first names and their numbers and send them an automated message
	*/
	var contacts []orientee
	fmt.Print()
	// this array is gonna be appended with orientation contact information
	fileptr := flag.String("ContactFileScanner","information.txt","File scanner for the contact information")
	flag.Parse()
	file,err := os.Open(*fileptr)
	if err != nil{
		os.Exit(1)
	}
	s := bufio.NewScanner(file)
	var phone_number string
	// go through and get the phone number and the first name
	for s.Scan(){
		temp := strings.Fields(s.Text())
		marker := false // when marker is true it is time to read phone number and break the loop
		for _, val :=range temp{
			if marker == true{
				phone_number = scan_Phone(val)
				break
			}
			for i:=0;i<len(val)-2;i++{
				// inner loop to loop through each character to find .edu
				if val[i]=='e'{
					if val[i+1]=='d' && val[i+2]=='u'{
						marker = true
					}
				}
			}
		}
		contactsTemp := orientee{
			number: phone_number,
			name: temp[2],
		}
		contacts = append(contacts,contactsTemp)
	}
	return contacts
}

