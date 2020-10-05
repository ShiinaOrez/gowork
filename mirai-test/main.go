package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"

	"github.com/Logiase/gomirai"
	"github.com/Logiase/gomirai/message"
)

func main() {
	var qq uint = 2485533524

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c := gomirai.NewClient("default", "http://127.0.0.1:8001", "12345678")
	c.Logger.Level = logrus.TraceLevel
	key, err := c.Auth()
	if err != nil {
		c.Logger.Fatal(err)
	}
	b, err := c.Verify(qq, key)
	if err != nil {
		c.Logger.Fatal(err)
	}

	go func() {
		err = b.FetchMessages()
		if err != nil {
			c.Logger.Fatal(err)
		}
	}()

	for {
		select {
		case e := <-b.Chan:
			switch e.Type {
			case message.EventReceiveGroupMessage:
				_, err = b.SendGroupMessage(e.Sender.Group.Id, 0, message.PlainMessage("Test"))
				if err != nil {
					fmt.Println(err)
				}
			}
		case <-interrupt:
			fmt.Println("######")
			fmt.Println("interrupt")
			fmt.Println("######")
			c.Release(qq)
			return
		}

	}
}