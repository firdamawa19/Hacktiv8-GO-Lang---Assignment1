package main
import (
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	name		string
	address		string
	job			string
	motivate	string
}

func printBiodata(student Student){
	fmt.Println("Nama\t\t\t: ", student.name)
	fmt.Println("Alamat\t\t\t: ", student.address)
	fmt.Println("Pekerjaan\t\t: ", student.job)
	fmt.Println("Alasan belajar Go\t: ", student.motivate)
}

func main () {

	person := map[string] Student {
		"1": {"Amanda", "New York", "Doctor", "Go"},
		"2": {"Johan", "Los Angeles", "Police", "Go"},
		"3": {"Kimmy", "New Jersey", "Consultant", "Go"},
		"4": {"John", "California", "Lecturer", "Go"},
		"5": {"Steve", "Washington", "musician", "Go"},
		"6": {"Christian", "California	", "Lecturer", "Go"},
	}


	if len(os.Args)>1 {
		var argumen=os.Args[1:][0]
		nomor,_:=strconv.Atoi(argumen)

		if (nomor < (len(person)+1)) && (nomor>0){
			printBiodata(person[argumen])

		} else {
			fmt.Println("Maaf, nomor absen tidak ditemukan")
		}

	} else {
		fmt.Println("Masukkan nomor urut absen dalam bentuk angka setelah perintah biodata.go")
	}
	
}