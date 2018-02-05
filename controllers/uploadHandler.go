package controllers

import (
	"net/http"
	"fmt"
	"myGrpc/flashes"
	"mime/multipart"
	"io"
	"os"
	"math/rand"
)
const maxFileSize = 35 * 1024 * 1024
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type NewFile struct {
	FileName string
	FullPath string
}
var FileStorage = make(map[string]*NewFile)

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength > maxFileSize {
		flash := []byte("File too large.")
		flashes.SetFlash(w, "message", flash)

		http.Redirect(w, r, "/", 301)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, maxFileSize)

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	file, handle, err := r.FormFile("uploadImage")
	if err != nil {
		fmt.Fprintf(w, "Handle file err: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer file.Close()

	mimeType := handle.Header.Get("Content-Type")
	switch mimeType {
	case "image/png", "image/jpeg":
		save(w, r, file, handle)
	default:
		flash := []byte("The format file is not valid.")
		flashes.SetFlash(w, "message", flash)

		http.Redirect(w, r, "/", 301)
	}
}

func save(w http.ResponseWriter,  r *http.Request, file multipart.File, handle *multipart.FileHeader) {
	buf := make([]byte, 500 * 1024)
	dst, err := os.Create(RandStr(14, handle.Filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer dst.Close()

	for {
		// read a chunk
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("File read error: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := dst.Write(buf[:n]); err != nil {
			fmt.Println("File write error: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	flash := []byte("File saved")
	flashes.SetFlash(w, "message", flash)

	http.Redirect(w, r, "/show-image", 301)
}

func RandStr(n int, filename string) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
	}
	newFilename := string(b) + "_" +  filename
	path := "tmp/" + newFilename

	// the key will be a user ID or username
	FileStorage["asd"] = &NewFile{
		FileName: string(b) + "_" +  filename,
		FullPath: path,
	}

	return path
}