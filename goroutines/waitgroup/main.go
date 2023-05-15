package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Notifier interface {
	Notify()
}

type EmailNotifier struct{}

func (n *EmailNotifier) Notify() {
	log.Println("Enviando Email...")
	time.Sleep(time.Second * 1)
}

type SMSNotifier struct{}

func (n *SMSNotifier) Notify() {
	log.Println("Enviando SMS...")
	time.Sleep(time.Second * 1)
}

type TelegramNotifier struct{}

func (n *TelegramNotifier) Notify() {
	log.Println("Enviando mensagem no Telegram...")
	time.Sleep(time.Second * 1)
}

func track(name string) func() {
	start := time.Now()
	return func() {
		log.Printf("%s, execution time %s\n", name, time.Since(start))
	}
}

func blockingExecution() {
	defer track("Blocking Execution")()

	notifiers := []Notifier{
		&EmailNotifier{},
		&SMSNotifier{},
		&TelegramNotifier{},
	}
	for _, n := range notifiers {
		n.Notify()
	}
}

func nonBlockingExecution() {
	defer track("Non-Blocking Execution")()

	notifiers := []Notifier{
		&EmailNotifier{},
		&SMSNotifier{},
		&TelegramNotifier{},
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(notifiers))
	for _, n := range notifiers {
		go func(wg *sync.WaitGroup, n Notifier) {
			n.Notify()
			wg.Done()
		}(wg, n)
	}
	wg.Wait()
}

func main() {
	blockingExecution()
	nonBlockingExecution()
	fmt.Println("Processo principal finalizado")
}
