package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type StundentParticipant struct {
	Participants []*Student `json:"participants"`
}

type Student struct {
	ID                string `json:"id"`
	StudentCode       string `json:"student_code"`
	StudentName       string `json:"student_name"`
	StudentAddress    string `json:"student_address"`
	StudentOccupation string `json:"student_occupation"`
	JoiningReason     string `json:"joining_reason"`
}

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Enter code argument")
		return
	}

	code := os.Args[1]

	var studentList StundentParticipant
	jsonFile, err := os.ReadFile("./data.json")

	if err != nil {
		fmt.Println("Error ketika membuka file:", err)
		return
	}

	err = json.Unmarshal(jsonFile, &studentList)
	if err != nil {
		fmt.Println("Error ketika membuka file 2:", err)
		return
	}

	var studentData *Student

	for _, student := range studentList.Participants {
		if code == student.StudentCode {
			studentData = student
		}
	}

	if studentData == nil {
		fmt.Println("Student tidak ditemukan")
		return
	}

	fmt.Printf("ID             : %s \n", studentData.ID)
	fmt.Printf("Name           : %s \n", studentData.StudentName)
	fmt.Printf("Address        : %s \n", studentData.StudentAddress)
	fmt.Printf("Job            : %s \n", studentData.StudentOccupation)
	fmt.Printf("Joining Reason : %s \n", studentData.JoiningReason)
}
