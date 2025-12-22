package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"os/user"
	"strings"
	"sync"

	"github.com/gedex/go-instagram/instagram"
)

type App struct {
	client    *instagram.Client
	fileIndex int
	mu        sync.Mutex
}

func NewApp(c *instagram.Client) *App {
	return &App{
		client: c,
	}
}

func (a *App) nextFileIndex() int {
	a.mu.Lock()
	defer a.mu.Unlock()

	val := a.fileIndex
	a.fileIndex++
	return val
}

func (a *App) StartDownload(username string) {
	usr, _ := user.Current()
	baseDir := fmt.Sprintf("%v/Downloads", usr.HomeDir)

	users, _, err := a.client.Users.Search(username, nil)
	if err != nil {
		log.Println("search error:", err)
		return
	}

	var userID string
	for _, u := range users {
		if u.Username == username {
			userID = u.ID
			break
		}
	}

	if userID == "" {
		log.Println("user not found:", username)
		return
	}

	folder := fmt.Sprintf("[%s]%s", userID, username)
	log.Println("Downloading:", folder)

	a.FindPhotos(folder, user, baseDir)
}

func (a *App) FindPhotos(ownerFolder, userID, baseDir string) {
	dir := fmt.Sprintf("%v/%v", baseDir, ownerFolder)

	os.MkdirAll(dir, 0755)

	linkChan := make(chan string)
	defer close(linkChan)
	wg := new(sync.WaitGroup)

	//workers
	workerCount := 2
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go a.DownloadWorker(dir, linkChan, wg)
	}

	var next *instagram.ResponsePagination

	for {
		maxID := ""
		if next != nil {
			maxID = next.NextMaxID
		}
		params := &instagram.Parameters{
			Count: 5,
			MaxID: maxID,
		}

		media, pagination, error := a.client.Users.RecentMedia(userID, params)
		if err != nil {
			log.Println("Media error:", err)
			break
		}

		next = pagination

		for _, m := range media {
			linkChan <- m.Images.StandardResolution.URL
		}

		if len(media) == 0 || next.NextMaxID == "" {
			break
		}
		wg.Wait()
	}
}

func (a *App) DownloadWorker(destDir string, links chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for url := range links {
		ext := ".jpg"
		if strings.Contains(url, ".png") {
			ext = ".png"
		}

		resp, err := http.Get(url)
		if err != nil {
			log.Println("http error:", err)
			continue
		}

		img, _, err := image.Decode(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Println("decode error:", err)
			continue
		}

		b := img.Bounds()
		if b.Dx() <= 300 || b.Dy() <= 300 {
			continue
		}

		index := a.nextFileIndex()
		name := fmt.Sprintf("pic%04d%s", index, ext)

		out, err := os.Create(destDir + "/" + name)
		if err != nil {
			log.Println("file error:", err)
			continue
		}

		if ext == ".png" {
			png.Encode(out, img)
		} else {
			jpeg.Encode(out, img, nil)
		}

		out.Close()

		if index%30 == 0 {
			log.Println(index, "photos downloaded")
		}
	}
}
