package ChatbotCommunication

type botinplace struct {
	parent *ChatBot
}

func (b *botinplace) New(parent *ChatBot) {
	b.parent = parent
}

func (b *botinplace) Send(question string) (string, error) {

	return "", nil
}
