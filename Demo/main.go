package main

import (
	"bufio"
	"fmt"
	cbot "github.com/LazarenkoA/ChatbotCommunication"
	"os"
)

func main() {
	//bot := new(cbot.BotCreator).New(new(cbot.XU))
	bot := new(cbot.BotCreator).New(new(cbot.Botinplace))
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
