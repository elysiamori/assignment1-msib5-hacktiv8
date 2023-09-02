package main

/*
VALDY RAMADHAN
GLNG-KS08-09
*/
import (
	"encoding/json"
	"fmt"
	"hacktiv8-go/go/assignment1/model"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// fungsi untuk argumen
	if len(os.Args) != 2 {
		fmt.Println("Masukkan student code sebagai argumen")
		return
	}

	studentCode := os.Args[1]

	// fungsi untuk membaca data json
	jsonPath := filepath.Join("data", "participants.json")

	participants, err := os.ReadFile(jsonPath)
	if err != nil {
		log.Fatal(err)
	}
	var student model.People

	if err := json.Unmarshal(participants, &student); err != nil {
		log.Fatal(err)
	}

	// fungsi untuk output data
	found := false
	for _, student8 := range student.People {
		if student8.StudentCode == studentCode {
			fmt.Println("=====================================")
			fmt.Printf(" ID : %s\n ", student8.ID)
			fmt.Printf("Student Code : %s\n ", student8.StudentCode)
			fmt.Printf("Name : %s\n ", student8.StudentName)
			fmt.Printf("Address : %s\n ", student8.StudentAddr)
			fmt.Printf("Occupation : %s\n ", student8.StudentOcc)
			fmt.Printf("Reason : %s\n ", student8.Reason)
			fmt.Println("=====================================")
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Mahasiswa dengan student code %s tidak ditemukan\n", studentCode)
	}
}
