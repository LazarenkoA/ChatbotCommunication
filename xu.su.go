package ChatbotCommunication

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type XU struct {
	parent *BotCreator
}

func (b *XU) New(parent *BotCreator) {
	b.parent = parent
}

func (b *XU) Send(question string) (string, error) {
	body := map[string]interface{}{
		"bot":  "Добрый",
		"text": question,
		"uid":  b.parent.UID.String(),
	}
	data, err := json.Marshal(&body)
	if err != nil {
		return "", fmt.Errorf("ошибка сериализации JSON: %w", err)
	}

	resp, err := b.parent.httpClient.Post("https://xu.su/api/send", "application/json; charset=utf-8", bytes.NewBuffer(data))
	if err != nil {
		return "", fmt.Errorf("ошибка ошибка отправки запроса: %w", err)
	}
	defer resp.Body.Close()

	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil || !(resp.StatusCode >= http.StatusOK && resp.StatusCode <= http.StatusIMUsed) {
		return "", err
	}

	responce := map[string]interface{}{}
	if err := json.Unmarshal(rbody, &responce); err != nil {
		return "", fmt.Errorf("ошибка десериализации JSON: %w", err)
	}

	return responce["text"].(string), nil
}
