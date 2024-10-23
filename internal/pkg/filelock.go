package pkg

import "sync"

var filelocks = make(map[string]*sync.Mutex)

func GetLock(filepath string) *sync.Mutex {
	_, ok := filelocks[filepath]
	if !ok {
		filelocks[filepath] = &sync.Mutex{}
	}

	return filelocks[filepath]
}