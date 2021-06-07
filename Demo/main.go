package main

import (
	"bufio"
	"fmt"
	cbot "github.com/LazarenkoA/ChatbotCommunication"
	"os"
)

func main() {
	bot := new(cbot.ChatBot).New(new(cbot.xu).New())
	//bot := new(cbot.ChatBot).New(new(cbot.botinplace))
	fmt.Println("Введите вопрос:")

	myscanner := bufio.NewScanner(os.Stdin)
	for myscanner.Scan() {
		question := myscanner.Text()
		if question == "exit" {
			break
		}

		if answer, err := bot.Send(question); err == nil {
			fmt.Println(answer)
		} else {
			fmt.Printf("Произошла ошибка: %v\n", err)
		}
	}
}
