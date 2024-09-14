package streamdeck

import (
	"sync"
)

type ResponseChannel chan StreamDeckEvent

var (
	responseChannels = make(map[string]ResponseChannel)
	responseMutex    sync.Mutex
)

func registerResponseChannel(context string) ResponseChannel {
	responseMutex.Lock()
	defer responseMutex.Unlock()
	ch := make(ResponseChannel, 1)
	responseChannels[context] = ch
	return ch
}

func unregisterResponseChannel(context string) {
	responseMutex.Lock()
	defer responseMutex.Unlock()
	delete(responseChannels, context)
}

func sendResponse(context string, event StreamDeckEvent) {
	responseMutex.Lock()
	ch, ok := responseChannels[context]
	responseMutex.Unlock()
	if ok {
		select {
		case ch <- event:
		default:
		}
	}
}

// func waitForResponse(context string, timeout time.Duration) (StreamDeckEvent, error) {
// 	ch := registerResponseChannel(context)
// 	defer unregisterResponseChannel(context)

// 	select {
// 	case response := <-ch:
// 		return response, nil
// 	case <-time.After(timeout):
// 		return nil, fmt.Errorf("timeout waiting for response")
// 	}
// }
