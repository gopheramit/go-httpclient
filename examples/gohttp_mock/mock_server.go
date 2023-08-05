package gohtt_mock

import "sync"

var (
	MockUpServer = make(map[string]*Mock)
)

type MockServer struct {
	enabled     bool
	serverMutex sync.Mutex
	mocks       map[string]*Mock
}
