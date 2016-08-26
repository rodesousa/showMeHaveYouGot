package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
	//	"reflect"
)

type Config struct {
	Server string
	Port   string
	File   map[string]string
	//Dir  []string
}

var conf Config

type HandlerService struct {
	Filename string
}

var controler []HandlerService
var controlerDir []HandlerService

var directories string = ""

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test")
}

func handlerInfos(w http.ResponseWriter, r *http.Request) {
	printing := "<h1> ShowMeHaveYouGot v0.1</h1> <br/> <br/>"
	for key := range conf.File {
		printing = printing + "- <a href=\"http://" + conf.Server + ":" + conf.Port + "/" + key + "\"> " + conf.File[key] + "</a> <br/> <br/>"
	}
	fmt.Fprintf(w, printing)
}

func (hs HandlerService) handler(w http.ResponseWriter, r *http.Request) {
	file, error := os.Open(hs.Filename)
	if error != nil {
		fmt.Fprintf(w, "no files")
	} else {
		info, _ := file.Stat()
		if info.Size() > 2000000 {
			buf := make([]byte, 50000)
			start := info.Size() - 50000
			file.ReadAt(buf, start)
			print := fmt.Sprintf("%s\n", buf)
			file.Close()
			fmt.Fprintf(w, print)
		} else {
			data, _ := ioutil.ReadFile(hs.Filename)
			file.Close()
			fmt.Fprintf(w, string(data))
		}
	}
}

func main() {
	data, er := ioutil.ReadFile("config.yaml")
	if er != nil {
		fmt.Println("erreur à la lecture du fichier de config.yaml")
	}

	err := yaml.Unmarshal([]byte(data), &conf)

	//fmt.Println(reflect.TypeOf(conf.Dir))
	//fmt.Println(conf.Dir)

	if err != nil {
		fmt.Println("erreur au parsing du fichier de config.yaml")
	}

	port := ":" + conf.Port
	controler = make([]HandlerService, len(conf.File))

	i := 0
	for key := range conf.File {
		controler[i].Filename = conf.File[key]
		ws := "/" + key
		http.HandleFunc(ws, controler[i].handler)
		i = i + 1
	}

	//	//for ele := range conf.Dir {
	//	//var a = conf.Dir[0]
	//	files, _ := ioutil.ReadDir("/home/rodesousa/gitpath/luz-lantern/github.com/rodesousa/printfiles")
	//	controlerDir = make([]HandlerService, len(files))
	//
	//	j := 0
	//	//path := conf.Dir[0][1:]
	//	for _, f := range files {
	//		ws := "/" + f.Name()
	//		controlerDir[j].Filename = conf.Dir[0] + "/" + f.Name()
	//		//fmt.Println(controlerDir[j].Filename)
	//		//fmt.Println(ws)
	//		//http.HandleFunc(ws, controlerDir[j].handler)
	//		http.HandleFunc(ws, handler)
	//		directories = directories + "<a href=\"localhost:" + conf.Port + ws + "\"> " + f.Name() + "</a> <br/>"
	//		j = j + 1
	//	}
	//
	//	http.HandleFunc("/dir", func(w http.ResponseWriter, r *http.Request) {
	//		fmt.Fprintf(w, directories)
	//	})
	//

	http.HandleFunc("/infos", handlerInfos)
	http.ListenAndServe(port, nil)
}
