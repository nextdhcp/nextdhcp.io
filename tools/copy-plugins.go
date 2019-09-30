package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var (
	nextdhcpPluginDir = "./nextdhcp/plugin"
	pluginDest        = "./content/plugins"
)

func main() {
	// read all plugin directories
	dirs, err := ioutil.ReadDir(nextdhcpPluginDir)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	// copy plugin readme and store path to it
	for _, d := range dirs {
		if !d.IsDir() {
			continue
		}

		path := filepath.Join(nextdhcpPluginDir, d.Name(), "README.md")
		_, err := os.Stat(path)
		if err != nil {
			log.Printf("failed to open README.md in %s, skipping: %s", path, err.Error())
			continue
		}

		readme, err := os.Open(path)
		if err != nil {
			log.Printf("failed to open README.md in %s, skipping: %s", path, err.Error())
			continue
		}
		defer readme.Close()

		destPath := filepath.Join(pluginDest, d.Name()+".md")
		dest, err := os.Create(destPath)
		if err != nil {
			log.Printf("failed to open %s for writing: %s", destPath, err.Error())
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			io.Copy(dest, readme)
		}()
	}

	wg.Wait()
}
