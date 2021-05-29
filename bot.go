package ChatbotCommunication

import (
	"bytes"
	"encoding/json"
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type ChatBot struct {
	UID        *uuid.UUID
	httpClient *http.Client
	cookieJar  *cookiejar.Jar
}

func (b *ChatBot) New() *ChatBot {
	b.UID, _ = uuid.NewV4()

	b.cookieJar, _ = cookiejar.New(nil)
	b.httpClient = &http.Client{
		Transport: &http.Transport{},
		//CheckRedirect: func(req *http.Request, via []*http.Request) error {
		//	return nil
		//},
		Jar:     b.cookieJar,
		Timeout: time.Minute * 5,
	}

	return b
}

func (b *ChatBot) Send(question string) (string, error) {
	body := map[string]interface{}{
		"bot":  "Добрый",
		"text": question,
		"uid":  b.UID.String(),
	}
	data, err := json.Marshal(&body)
	if err != nil {
		return "", fmt.Errorf("ошибка сериализации JSON: %w", err)
	}

	resp, err := b.httpClient.Post("https://xu.su/api/send", "application/json; charset=utf-8", bytes.NewBuffer(data))
	if err != nil {
		return "", fmt.Errorf("ошибка ошибка отправки запроса: %w", err)
	}
	defer resp.Body.Close()

	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	responce := map[string]interface{}{}
	if err := json.Unmarshal(rbody, &responce); err != nil {
		return "", fmt.Errorf("ошибка десериализации JSON: %w", err)
	}

	return responce["text"].(string), nil
}
