/*
	DHT.go - A simplified implementation of a Distributed Hash Table which uses concurrency mechanisms 
	to mimic a P2P network allowing execution in absence of any external connection or communication.

	Author: Justin Underhay
	Date of last modification: Nov 21, 2018
*/

package main

import "fmt"
import "strconv"
import "time"
import "strings"

var nodeChans []chan string
var numNodes int


func node(ID int) {

	var storage = make(map[string]string)
	nodeChans[ID] = make(chan string)
	var keyULimit = 33 + ID*33
	var keyLLimit = ID*33
	var hash []string

	for {
		cmd := <- nodeChans[ID]

		switch cmd[:3] {	
		case "set":
			hash = strings.Split(cmd, " ")
			key, _ := strconv.Atoi(hash[1])

			if key >= keyLLimit && key <= keyULimit {
			//	fmt.Println("Node",ID,"storing",hash[2])
				storage[hash[1]] = hash[2]
			} else {
			//	fmt.Println("Node",ID,"forwarding set cmd")
				nodeChans[(ID+1)%numNodes] <- cmd
			}

		case "get":
			hash = strings.Split(cmd, " ")
			path, _ := strconv.Atoi(hash[2])
			
			if path == numNodes {
				fmt.Println("The value for this key was not found.")
			} else {

				ret, prs := storage[hash[1]]
				
				if prs {
					fmt.Println("The value",ret,"for this key was found in node",ID)
				} else {
					nodeChans[(ID+1)%numNodes] <- "get " + hash[1] + " " + strconv.Itoa(path+1) 
				}
			}

		}	
	
	}

		

}


func getHash(s string) int {

	var val int = 0
	for i := 0; i < len(s); i++ {
		val += int(s[i])
	}
	return val % 100
}


func main() {

	nodeChans = make([]chan string, 3)
	numNodes = 3
	
	for i := 0; i<3; i++ {
		go node(i)
	}

	fmt.Println("MiniDHT [Ver. 1.112]")
	fmt.Println("Author: J. Underhay  -  Date of last modification: 2018-11-21")

	var cmd string
	var one string
	var two string
	var quit bool = false
	

	for {
		fmt.Printf("==>")
		_, err := fmt.Scanln(&cmd, &one, &two)

		if err != nil && one != "" && two != "" {
			fmt.Println(err)
			fmt.Println("Unexpected error, quitting...")
		}


		switch cmd {
		case "quit":
			quit = true
		case "set":
			if one == "" || two == "" {
				fmt.Println("Invalid format of set command. Usage: set <key> <value>")
				break
			}

			hVal := getHash(one)
			msg := "set " + strconv.Itoa(hVal) + " " + two
			nodeChans[(hVal/numNodes) % numNodes] <- msg

		case "get":
			if one == "" {
				fmt.Println("Invalid format of get command. Usage: get <key>")
				break
			}

			hVal := getHash(one)
			msg := "get " + strconv.Itoa(hVal) + " 0"
			nodeChans[(hVal/numNodes) % numNodes] <- msg

		case "cmds":
			fmt.Println()
			fmt.Println("Recognzed commands:")
			fmt.Printf("\tset <key> <value>   Hashes <key> to a storage value and stores <value> under that hash\n")
			fmt.Printf("\tget <key>           Retrieves the value associated with the hashed value of <key> if it exists\n")
			fmt.Printf("\tquit                Exits the program\n")
		case "":
		default:
			fmt.Println("",cmd,"is not a recognized command. Type cmds for a list of valid commands.")
		}

		one = ""
		two = ""
		cmd = ""

		if quit {
			break
		}

		time.Sleep(time.Millisecond*100)
	}

    
}
