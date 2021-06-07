package ChatbotCommunication

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type Botinplace struct {
	parent *BotCreator
}

func (b *Botinplace) New(parent *BotCreator) {
	b.parent = parent
}

func (b *Botinplace) Send(question string) (result string, e error) {
	defer func() {
		if err := recover(); err != nil {
			result = ""
			e, _ = err.(error)
		}
	}()

	params := []string{
		fmt.Sprintf("msg=%s", question),
		fmt.Sprintf("mykey=%s", b.getmyKey()),
	}

	data := bytes.NewBufferString(strings.Join(params, "&"))
	// resp, err := b.parent.httpClient.Post(fmt.Sprintf("https://botinplace.ru/brain/index.php?%d", time.Now().Unix()), "application/x-www-form-urlencoded; charset=UTF-8", data)

	req, err := http.NewRequest("POST", fmt.Sprintf("https://botinplace.ru/brain/index.php?%d", time.Now().Unix()), data)
	if err != nil {
		return "", err
	}
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("x-authorization", "token")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	//req.Header.Set("accept-encoding", "gzip, deflate, br")
	req.Header.Set("origin", "https://botinplace.ru")
	req.Header.Set("referer", "https://botinplace.ru/")
	req.Header.Set("accept", "*/*")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-gpc", "1")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36")
	req.ContentLength = int64(data.Len())

	resp, err := b.parent.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("ошибка отправки запроса: %w", err)
	}
	defer resp.Body.Close()

	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil || !(resp.StatusCode >= http.StatusOK && resp.StatusCode <= http.StatusIMUsed) {
		return "", err
	}

	//reader, err := gzip.NewReader(resp.Body)
	//body, _ := ioutil.ReadAll(reader)
	//reader.Close()

	responce := map[string]interface{}{}
	if err := json.Unmarshal(rbody, &responce); err != nil {
		return "", fmt.Errorf("ошибка десериализации JSON: %w", err)
	}

	return responce["info"].(map[string]interface{})["msg"].(string), nil
}

func (b *Botinplace) getmyKey() string {
	resp, err := b.parent.httpClient.Get("https://botinplace.ru")
	if err != nil {
		log.Println("ошибка получения myKey")
		return ""
	}
	defer resp.Body.Close()

	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	var re = regexp.MustCompile(`(?m)mykey:[\s]*\'([^']+)`)
	if match := re.FindAllStringSubmatch(string(rbody), -1); len(match) > 0 {
		return match[0][1]
	}

	return ""
}
