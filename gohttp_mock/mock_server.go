package gohtt_mock

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"sync"

	"github.com/gopheramit/go-httpclient/core"
)

var (
	MockUpServer = mockServer{
		mocks:      make(map[string]*Mock),
		httpClient: &httpClientMock{},
	}
)

type mockServer struct {
	enabled     bool
	serverMutex sync.Mutex
	httpClient  core.HttpClient
	mocks       map[string]*Mock
}

func (m *mockServer) getMockKey(mehtod, url, body string) string {
	hasher := md5.New()
	hasher.Write([]byte(mehtod + url + m.cleanBody(body)))
	key := hex.EncodeToString(hasher.Sum(nil))
	return key
}

func (m *mockServer) cleanBody(body string) string {
	body = strings.TrimSpace(body)
	if body == "" {
		return ""
	}
	body = strings.ReplaceAll(body, "\n", "")
	body = strings.ReplaceAll(body, "\t", "")
	return body

}
