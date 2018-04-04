package service

import (
	"sync"
	"github.com/wycers/sustc-sakura-console/model"
)

var Log = &logService{
	mutex: &sync.Mutex{},
}

type logService struct {
	mutex *sync.Mutex
}
func (srv *logService) CreateLog(log *model.Log) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	if err := db.Create(log).Error; err != nil {
		return err
	}
	return nil
}