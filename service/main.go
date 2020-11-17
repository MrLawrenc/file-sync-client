package main

import (
	. "file-sync-client/service/files"
	"file-sync-client/service/utils"
	"github.com/fsnotify/fsnotify"
	"log"
)

func init() {
	// log config 显示行数
	log.SetFlags(log.Ldate | log.Lshortfile)
}

func main() {
	var parentDir = "E://"

	allFiles := ListFiles(parentDir)
	println(allFiles)

	//watch file
	watcher, err := fsnotify.NewWatcher()
/*	if err != nil {
		log.Fatal("NewWatcher failed: ", err)
	}*/
	utils.PrintError(err,false)

	defer watcher.Close()

	done := make(chan bool)
	go func() {
		defer close(done)
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fileName := event.Name
				op := event.Op
				log.Printf("%s %s\n", fileName, op)

				switch op.String() {
				case "CREATE":
					log.Println("CREATE")
				case "WRITE":
					log.Println("WRITE")
				case "RENAME":
					log.Println("RENAME")
				case "REMOVE":
					log.Println("REMOVE")
				default:
					log.Println("Chmod")
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(parentDir)
	if err != nil {
		log.Fatal("Add failed:", err)
	}
	<-done

}
