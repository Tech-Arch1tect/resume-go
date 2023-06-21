package lib

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

func Watch(directory string, cmd *cobra.Command) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("Creating watcher failed: ", err)
	}
	defer watcher.Close()
	done := make(chan bool)
	// start a goroutine to watch for events
	go func() {
		defer close(done)
		for {
			select {
			case _, ok := <-watcher.Events:
				if !ok {
					log.Fatal("Watcher error: ", err)
				}
				// generate resume
				json := GetMapFromJson(cmd.Flag("resume").Value.String())
				GenerateResume(json, cmd.Flag("template").Value.String(), cmd.Flag("output").Value.String())
				log.Println("Resume generated")
			case err := <-watcher.Errors:
				log.Fatal("Watcher error: ", err)
			}
		}

	}()
	log.Println("Watching directory: ", directory)
	err = watcher.Add(directory + "/...")
	if err != nil {
		log.Fatal("Watch directory failed: ", err)
	}
	<-done
}
