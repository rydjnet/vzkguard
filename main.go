package main

import (
	"fmt"
	"sync"
	"vzkguard/tbot"
)

type TUser struct {
	Id         int64
	UserName   string
	Role       bool
	MsgCount   int
	TrustScore int
}

type CacheUsers struct {
	Data map[int64]TUser
	Mu   sync.Mutex
}

func New() *CacheUsers {
	data := CacheUsers{Data: make(map[int64]TUser)}
	return &data
}

func (c *CacheUsers) AddUser(user TUser) error {
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
		return c.AddUser(user)
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

func main() {

	tbot.Start()

}
