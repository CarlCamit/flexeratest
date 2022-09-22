package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/carlcamit/flexeratest/user"
)

const (
	// CSV Index Columns
	ComputerID = iota
	UserID
	ApplicationID
	ComputerType
	Comment

	TargetApplicationID = "374"
	Laptop              = "laptop"
	Desktop             = "desktop"
)

func main() {
	csvFile := flag.String("file", "", "csv file to parse")
	flag.Parse()
	if *csvFile == "" {
		fmt.Println("please provide filepath to a CSV")
		return
	}

	f, err := os.Open(*csvFile)
	if err != nil {
		log.Fatalf("filepath %s not found, %v", *csvFile, err)
	}

	users := user.NewUsers()
	r := csv.NewReader(f)
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// Skip row if not the target application ID
		if row[ApplicationID] != TargetApplicationID {
			continue
		}

		// Check if the user exists in the user map
		// already otherwise create a new user
		u, ok := users[row[UserID]]
		if !ok {
			u = user.NewUser()
		}

		if strings.ToLower(row[ComputerType]) == Desktop {
			u.AddIfUniqueDesktop(row[ComputerID])
		} else {
			u.AddIfUniqueLaptop(row[ComputerID])
		}

		// Update user ID with updated user value
		users[row[UserID]] = u
	}

	fmt.Printf("total number of applications required is %v\r\n", users.TotalApplications())
}
