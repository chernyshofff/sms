package main

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {
	// setup watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	// use goroutine to start the watcher
	go func() {
		for {
			select {
			// provide the list of events to monitor
			case event := <-watcher.Events:
				if event.Op&fsnotify.Create == fsnotify.Create {
					fmt.Println("File created:", event.Name)
				}
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					fmt.Println("File removed:", event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("Error:", err)
			}
		}
	}()

	// provide the directory to monitor
	err = watcher.Add("C:/users/IT/go/src/tmp")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
