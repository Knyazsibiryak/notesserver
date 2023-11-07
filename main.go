package main

import (
	"bufio"
	"fmt"
	"os"
)

type Note struct {
	ID      int
	Name    string
	Surname string
	Text    string
}

var notes []Note

func main() {
	fmt.Println("Консольное приложение")
	menu()
}

func menu() {
	for {
		fmt.Println("\nМеню:")
		fmt.Println("1) Добавить заметку")
		fmt.Println("2) Посмотреть заметку")
		fmt.Println("3) Удалить заметку")
		fmt.Println("4) Завершить работу")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Выберите пункт меню: ")
		option, _ := reader.ReadString('\n')

		switch option {
		case "1\n":
			addNote()
		case "2\n":
			viewNote()
		case "3\n":
			deleteNote()
		case "4\n":
			fmt.Println("Работа приложения завершена.")
			return
		default:
			fmt.Println("Некорректный выбор. Попробуйте еще раз.")
		}
	}
}

func addNote() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите имя: ")
	name, _ := reader.ReadString('\n')
	name = name[:len(name)-1]

	fmt.Print("Введите фамилию: ")
	surname, _ := reader.ReadString('\n')
	surname = surname[:len(surname)-1]

	fmt.Print("Введите заметку: ")
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]

	id := len(notes) + 1
	note := Note{ID: id, Name: name, Surname: surname, Text: text}
	notes = append(notes, note)

	fmt.Println("Заметка добавлена. ID заметки:", id)
}

func viewNote() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите ID заметки: ")
	idStr, _ := reader.ReadString('\n')
	var id int
	fmt.Sscanf(idStr, "%d", &id)

	found := false

	for _, note := range notes {
		if note.ID == id {
			fmt.Println("Автор:")
			fmt.Println(note.Name, note.Surname)
			fmt.Println("Заметка:")
			fmt.Println(note.Text)

			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Заметка с ID %d не найдена.\n", id)
	}
}

func deleteNote() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите ID: ")
	idStr, _ := reader.ReadString('\n')
	var id int
	fmt.Sscanf(idStr, "%d", &id)

	found := false

	for i, note := range notes {
		if note.ID == id {
			notes = append(notes[:i], notes[i+1:]...)
			fmt.Printf("Заметка с ID %d удалена.\n", id)

			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Заметка с ID %d не найдена.\n", id)
	}
}
