package lib

import (
	"io"
	"log"
	"net/http"
	"os"
)

func InitResume() {
	// if resume.json already exists then exit
	if _, err := os.Stat("resume.json"); err == nil {
		log.Fatal("resume.json already exists")
	}
	// download https://raw.githubusercontent.com/jsonresume/resume-schema/master/sample.resume.json
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://raw.githubusercontent.com/jsonresume/resume-schema/master/sample.resume.json", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// create resume.json
	out, err := os.Create("resume.json")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// copy response body to resume.json
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// print success message
	log.Println("resume.json created")
}
