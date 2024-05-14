package cache

import (
	"fmt"
	"sync"
)

type CacheUsers struct {
	Data map[int64]bool
	Mu   sync.Mutex
}

func New() *CacheUsers {
	data := CacheUsers{Data: make(map[int64]bool)}
	return &data
}

func (c *CacheUsers) addUser(userID int64) error {
	if userID == 0 {
		return fmt.Errorf("user (id: %d) is not identified", userID)
	}
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Data[userID] = true
	return nil
}
func (c *CacheUsers) NewMsg(userID int64) error {
	if c.Data[userID] {
		return nil
	}
	c.addUser(userID)
	return nil

}
func (c *CacheUsers) GetUser(userID int64) bool {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	return c.Data[userID]
}
