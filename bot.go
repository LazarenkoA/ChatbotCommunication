package ChatbotCommunication

import (
	uuid "github.com/nu7hatch/gouuid"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type Iprovader interface {
	Send(string) (string, error)
	New(*BotCreator)
}

type BotCreator struct {
	UID        *uuid.UUID
	httpClient *http.Client
	cookieJar  *cookiejar.Jar
}

func (b *BotCreator) New(provader Iprovader) Iprovader {
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

	provader.New(b)
	return provader
}
