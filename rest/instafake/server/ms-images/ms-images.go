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
	ImgBase64 string `json:"image"`
}

type Images struct {
	ImagesArray []string `json:"images"`
}

func getImages(w http.ResponseWriter, r *http.Request) {
	allImages := make([]string, 0)
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

		imgBase64Str := base64.StdEncoding.EncodeToString(buffer)

		allImages = append(allImages, imgBase64Str)
	}
	var images Images
	images.ImagesArray = allImages
	json.NewEncoder(w).Encode(images)
}

func getOneImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	_, err := os.Stat("./source/" + key + ".png")
	if os.IsNotExist(err) {
		os.Exit(1)
	}
	imgFile, err := os.Open("./source/" + key + ".png")
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
	myRouter.HandleFunc("/api/msi/images/{id}", getOneImage)
	log.Fatal(http.ListenAndServe(":8888", myRouter))
}

func main() {
	handleRequests()
}
