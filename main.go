package main

import (
	"bufio"
	"fmt"
	ai "lab5-program-api/cohere_package"
	"os"
	"strconv"
	"time"

	cohere "github.com/cohere-ai/cohere-go/v2"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter topic of your nightmare(diplom):0 ")

	file, err := os.Create("sds_document.txt")
	check(err)

	defer file.Close()

	topic, err := reader.ReadString('\n')
	check(err)
	current_date := time.Now()
	current_date_str := fmt.Sprintf("%02d-%02d-%04d", current_date.Day(), int(current_date.Month()), current_date.Year())
	version := "1.0\n"

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("System Design Specification\nТема: " + topic + "\nДaтa: " + current_date_str + "\nBepciя: " + version + "\n3мicт\n")
	check(err)

	chapters_slice := []string{"Управління документами", "Огляд", "Інструменти розробки та стандарти", "Процеси системи",
		"Користувацький інтерфейс", "Безепека додатку", "Проектування бази даних", "Інтерфейси додатку", "Дані", "Впровадження"}

	for i := 0; i < len(chapters_slice); i++ {
		_, err = writer.WriteString(chapters_slice[i] + "\n")
		check(err)
	}

	var chat_history []*cohere.Message
	var response_text string
	for i := 0; i < len(chapters_slice); i++ {
		fmt.Println("Chapter " + strconv.Itoa(i+1) + " is generating ...")
		prompt := "У архітектурному документі System Design Specification з темою " + topic + "напиши пункт " + strconv.Itoa(i) + "(" + chapters_slice[i] + ")"
		response_text, chat_history = ai.GetAiResponse(prompt, chat_history)

		_, err = writer.WriteString(strconv.Itoa(i+1) + ". " + chapters_slice[i] + "\n\n" + response_text)
		check(err)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
