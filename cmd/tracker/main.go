package main

import (
	"context"
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

	c := client.NewClient(parsed)

	searchRes, err := c.Search(context.Background(), "0000175352124")
	if err != nil {
		return err
	}

	if len(searchRes) == 0 {
		return fmt.Errorf("nothing found")
	}

	tracked := map[string]*statusSet{}

	for _, r := range searchRes {
		trackRes, err := c.Track(context.Background(), r.TrackID)
		if err != nil {
			return err
		}

		if !trackRes.Result {
			continue
		}

		ss := &statusSet{
			history: make([]string, 0, len(trackRes.Statuses)),
		}
		for _, s := range trackRes.Statuses {
			// if current status
			// I fucking hate boxberry api so much
			if s.Status == r.Status {
				ss.current = s.Name
			}

			ss.history = append(ss.history, s.DateTime+": "+s.Name)
		}

		tracked[r.OrderID] = ss
	}

	for orderID, ss := range tracked {
		fmt.Printf("Информация по заказу: %s\n", orderID)
		fmt.Println(ss)
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
