package housemd

import (
	"fmt"
	"sync"
)

type CacheUsers struct {
	Data map[int64]TUser
	Mu   sync.Mutex
}

func New() *CacheUsers {
	data := CacheUsers{Data: make(map[int64]TUser)}
	return &data
}

func (c *CacheUsers) addUser(user TUser) error {
	if user.Id == 0 {
		return fmt.Errorf("user (id: %d) is not identified", user.Id)
	}
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Data[user.Id] = user
	return nil
}
func (c *CacheUsers) NewMsg(user TUser) error {
	if val, ok := c.Data[user.Id]; ok {
		c.Mu.Lock()
		defer c.Mu.Unlock()
		val.MsgCount += 1
		c.Data[user.Id] = val
	} else {
		c.addUser(user)
		c.NewMsg(user)
		return nil
	}
	return nil
}
func (c *CacheUsers) GetUser(id int64) (TUser, bool) {
	if val, ok := c.Data[id]; ok {
		return c.Data[id], true
	} else {
		return val, false
	}
}

func (c *CacheUsers) UserTrust(id int64) int {
	return c.Data[id].MsgCount

}
