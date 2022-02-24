package main

import (
	"fmt"
	"os"
	"path"

	"github.com/FAdemoglu/patika/helper"
)

func main() {
	args := os.Args

	var bookSlice = []string{"The Idiot", "The Castle", "Native Son", "Sapiens", "Emotional Quality"}
	var wannaSearchBook string
	if len(args) == 1 {
		projectName := path.Base(args[0])
		fmt.Printf("%s uygulamasında kullanabileceğiniz komutlar : \n search=> arama işlemi için \n list=> listeleme işlemi için\n", projectName)
		return
	}

	fmt.Printf("Kullanıcı tarafından girilen ilk değer: %s \n", args[1])

	if len(args) > 3 {
		for i := 2; i < len(args)-1; i += 2 {
			wannaSearchBook += args[i] + " " + args[i+1]
		}
	}
	argument := helper.LowercaseString(args[1])
	if argument == "search" && len(args) < 3 {
		fmt.Printf("Aramak istediğiniz kitabı eklemelisiniz")
	} else {
		switch argument {
		case "search":
			search(bookSlice, wannaSearchBook)
		case "list":
			list(bookSlice)
		default:
			fmt.Printf("Girdiğiniz komut ile ilgili bir yönlendirme bulunmamaktadır.")
		}
	}
}

func search(books []string, search string) {
	var searchedBooks []string
	for i := range books {
		if books[i] == search {
			fmt.Println("Kitap Listede Var")
			searchedBooks = append(searchedBooks, books[i])
		}
	}
	if len(searchedBooks) == 0 {
		fmt.Println("Kitap listede yok.")
	} else {
		fmt.Printf("Aranan kitaplar %v", searchedBooks)
	}

}

func list(books []string) {
	for _, books := range books {
		fmt.Printf("%s\n", books)
	}
}
