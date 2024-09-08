package main

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/daniilty/boxberry-track/pkg/client"
)

type statusSet struct {
	current string
	history []string
}

func (s *statusSet) String() string {
	return fmt.Sprintf("Текущее состояние: %s\nИстория состояний: \n\t%s", s.current, strings.Join(s.history, "\n\t"))
}

func run() error {
	u := "https://boxberry.ru"
	parsed, _ := url.ParseRequestURI(u)

	if len(os.Args) < 2 {
		return errors.New("usage: boxberry-track orderID1 orderID2 orderID3")
	}

	c := client.NewClient(parsed)

	for i, arg := range os.Args[1:] {
		searchRes, err := c.Search(context.Background(), arg)
		if err != nil {
			return err
		}

		if len(searchRes.ParcelWithStatuses) == 0 {
			return fmt.Errorf("nothing found")
		}

		for j, p := range searchRes.ParcelWithStatuses {
			fmt.Printf("%d:%d. Информация по отправлению %s:\n"+
				"\tОжидаемая дата доставки: %s\n"+
				"\tТекущий статус: %s\n",
				i+1, j+1, p.OrderID, p.DeliveryDate, p.Status)

			if len(p.Statuses) != 0 {
				fmt.Println("\tИстория состояний:")
			}
			for n, s := range p.Statuses {
				fmt.Printf("\t\t%d. %s: %s\n", n+1, s.DateTime, s.Name)
			}
		}

	}

	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
