package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Image struct {
	ImgBase64 string `json:"imageData"`
}

type allImages []Image

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func getProfileImages(w http.ResponseWriter, r *http.Request) {
	var images allImages
	vars := mux.Vars(r)
	key := vars["id"]
	files, _ := ioutil.ReadDir("./source")
	for _, f := range files {
		imgFile, err := os.Open("./source/" + f.Name())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer imgFile.Close()
		fInfo, _ := imgFile.Stat()
		size := fInfo.Size()
		buffer := make([]byte, size)

		fReader := bufio.NewReader(imgFile)
		fReader.Read(buffer)

		invStr := reverse(f.Name())
		invStr = invStr[:6]
		nStr := reverse(invStr)
		fns := f.Name()
		fNameStr := fns[:len(fns)-len(nStr)]

		if key == fNameStr {
			var image Image
			imgBase64Str := base64.StdEncoding.EncodeToString(buffer)
			image.ImgBase64 = imgBase64Str

			images = append(images, image)
		}
	}
	json.NewEncoder(w).Encode(images)
}

func getImages(w http.ResponseWriter, r *http.Request) {
	var images allImages
	files, _ := ioutil.ReadDir("./source")
	for _, f := range files {
		imgFile, err := os.Open("./source/" + f.Name())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer imgFile.Close()
		fInfo, _ := imgFile.Stat()
		size := fInfo.Size()
		buffer := make([]byte, size)

		fReader := bufio.NewReader(imgFile)
		fReader.Read(buffer)

		var image Image
		imgBase64Str := base64.StdEncoding.EncodeToString(buffer)
		image.ImgBase64 = imgBase64Str

		images = append(images, image)
	}
	json.NewEncoder(w).Encode(images)
}

func getOneImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	_, err := os.Stat("./source/" + key + ".jpg")
	if os.IsNotExist(err) {
		os.Exit(1)
	}
	imgFile, err := os.Open("./source/" + key + ".jpg")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer imgFile.Close()
	fInfo, _ := imgFile.Stat()
	size := fInfo.Size()
	buffer := make([]byte, size)

	fReader := bufio.NewReader(imgFile)
	fReader.Read(buffer)

	imgBase64Str := base64.StdEncoding.EncodeToString(buffer)

	var img Image
	img.ImgBase64 = imgBase64Str

	json.NewEncoder(w).Encode(img)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/api/msi/images", getImages)
	myRouter.HandleFunc("/api/msi/image/{id}", getOneImage)
	myRouter.HandleFunc("/api/msi/images/{id}", getProfileImages)
	log.Fatal(http.ListenAndServe(":8888", myRouter))
}

func main() {
	handleRequests()
}
