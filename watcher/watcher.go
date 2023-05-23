package watcher

import (
	"log"

	fsnotify "github.com/fsnotify/fsnotify"
)

func StartWatcher(directory string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("Failed to create watcher:", err)
	}
	defer watcher.Close()

	// directory := "./.logs"

	// Add the directory to the watcher
	err = watcher.Add(directory)
	if err != nil {
		log.Fatal("Failed to add directory to watcher:", err)
	}

	log.Println("Watcher started. Monitoring directory:", directory)

	// Start watching for events
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Create == fsnotify.Create {
				log.Println("New file created:", event.Name)
				ReadFile(event.Name)
			}

			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Println("File modified:", event.Name)
				ReadFile(event.Name)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("Error occurred:", err)
		}
	}
}
