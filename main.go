package ChatbotCommunication

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	bot := new(ChatBot).New()
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
