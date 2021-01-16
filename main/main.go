package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"time"

	"github.com/gookit/ini/v2"
	"github.com/gorilla/feeds"
	"github.com/jessevdk/go-flags"
	"github.com/lggomez/go-zimrss/main/zim"
)

var commandlineFlags struct {
	NotebookPath string `short:"n" long:"notebook" description:"Zim notebook path" required:"true"`
	Title        string `short:"t" long:"title" description:"Feed title" required:"true"`
	Link         string `short:"l" long:"link" description:"Feed site link" required:"true"`
	Description  string `short:"d" long:"description" description:"Feed site description" required:"false"`
	AuthorName   string `short:"a" long:"authorname" description:"Author name" required:"true"`
	AuthorEmail  string `short:"e" long:"authoremail" description:"Author email" required:"false"`
}

func main() {
	// Validate commandline args and get notebook path
	args := os.Args
	args, err := flags.ParseArgs(&commandlineFlags, args)
	if err != nil {
		log.Fatal(err)
	}

	rootPath := commandlineFlags.NotebookPath

	// Parse zim notebook object from file
	notebookPath := path.Join(rootPath, "notebook.zim")
	err = ini.LoadExists(notebookPath)
	if err != nil {
		log.Fatal(err)
	}

	zimNotebook := zim.NotebookFromMap(ini.StringMap("Notebook"))
	fmt.Printf("%+v\r\n", zimNotebook)

	// traverse notebook pages from root notebook directory and sort them by creation date
	pages := traverseAndParsePageMetadata(rootPath, notebookPath)

	fmt.Printf("pages (len %v): %+v\r\n", len(pages), pages)
	sort.Stable(pages)
	fmt.Printf("pages (len %v): %+v (post sort)\r\n", len(pages), pages)

	if len(pages) == 0 {
		fmt.Println("This notebook has no pages. Finishing")
		return
	}

	rss, err := createRSSFeed(pages, rootPath, zimNotebook)
	if err != nil {
		log.Fatal(err)
	}

	// Write RSS feed file
	f, err := os.Create("rss.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString(rss)
	fmt.Println("RSS feed write DONE")
}

func createRSSFeed(pages zim.PageMetadataByCreationDate, rootPath string, zimNotebook zim.Notebook) (string, error) {
	// Create main feed element
	feed := &feeds.Feed{
		Title:       commandlineFlags.Title,
		Link:        &feeds.Link{Href: commandlineFlags.Link},
		Description: commandlineFlags.Description,
		Author:      &feeds.Author{Name: commandlineFlags.AuthorName, Email: commandlineFlags.AuthorEmail},
		Created:     pages[0].CreationDate.Add(-(60 * time.Minute)), // Use 1 hour prior to the first page creation as an arbitrary creation date
	}

	items := make([]*feeds.Item, 0, len(pages))
	for _, p := range pages {
		pageLink := p.PathToURL(rootPath, commandlineFlags.Link, zimNotebook.DefaultFileExtension)
		items = append(items, &feeds.Item{
			Title: p.Title,
			Link:  &feeds.Link{Href: pageLink},
			//Description: ,
			Author:  &feeds.Author{Name: commandlineFlags.AuthorName, Email: commandlineFlags.AuthorEmail},
			Created: p.CreationDate,
		})
	}
	feed.Items = items

	rss, err := feed.ToRss()
	if err != nil {
		log.Fatal(err)
	}
	return rss, nil
}

func traverseAndParsePageMetadata(rootPath string, notebookPath string) zim.PageMetadataByCreationDate {
	pages := zim.PageMetadataByCreationDate{}
	walkErr := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if path == notebookPath {
			return nil // ignore base notebook file
		}
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}

		if !info.IsDir() {
			// TODO: Detect and ignore empty pages (with empty content starting from line 7)
			pageMetadata, err := zim.ParsePage(path, info)
			if err != nil {
				return err
			}

			pages = append(pages, pageMetadata)
		}

		return nil
	})
	if walkErr != nil {
		panic(walkErr)
	}
	return pages
}
