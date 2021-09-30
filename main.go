package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/valyala/fastjson"
)

func main() {
	exists, _ := exists("downloads")
	if !exists {
		os.Mkdir("downloads", 0755)
	}

	domain := os.Args[1]
	amount := os.Args[2]

	tags := os.Args[3:]
	tagsAll := strings.Join(tags[:], "+")

	fmt.Println("Domain: " + domain)
	fmt.Println("Limit: " + amount)
	fmt.Println("Tags: " + tagsAll)
	fmt.Println("Beginning to scrape.")

	// begin shit

	client := &http.Client{}
	var data = strings.NewReader(``)
	var url = "https://api.r34.app/booru/gelbooru/posts?baseEndpoint=" + domain + "&limit=" + amount + "&tags=" + tagsAll
	fmt.Println(url)

	req, _ := http.NewRequest("GET", url, data)
	res, _ := client.Do(req)

	var p fastjson.Parser

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	newStr := buf.String()

	parsed, _ := p.Parse(newStr)

	for _, element := range parsed.GetArray() {
		r34Url := element.GetObject("high_res_file").Get("url").GetStringBytes()
		splitStr := strings.Split(string(r34Url), ".")
		r34ID := element.GetInt64("id")
		path := "downloads/" + fmt.Sprint(r34ID) + "." + splitStr[3]

		err := DownloadFile(path, string(r34Url))
		if err != nil {
			fmt.Println(err)
		}
	}
}

// https://golangcode.com/download-a-file-from-a-url/
func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)

	size, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
	s2 := strconv.Itoa(size)

	fmt.Println("Downloaded url: " + url + " with length of " + s2 + " bytes")
	return err
}

// https://stackoverflow.com/questions/10510691/how-to-check-whether-a-file-or-directory-exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
