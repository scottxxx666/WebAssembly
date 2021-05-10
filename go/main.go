package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"syscall/js"
	"time"
)

func main() {
	// document := js.Global().Get("document")
	// canvas := document.Call("getElementById", "canvas")
	// ctx := canvas.Call("getContext", "2d")
	// canvas.Set("width", js.ValueOf(500))
	// canvas.Set("height", js.ValueOf(500))
	//
	// image := js.Global().Call("eval", "new Image()")
	// image.Set("src", "data:image/png;base64,"+loadImage("image.png"))
	//
	// js.Global().Call("setInterval", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	// 	ctx.Call("clearRect", 0, 0, 500, 500)
	// 	ctx.Call("drawImage", image, 0, 0)
	//
	// 	return nil
	// }), js.ValueOf(50))

	select {}
}

func loadImage(path string) string {
	href := js.Global().Get("location").Get("href")
	u, err := url.Parse(href.String())
	if err != nil {
		log.Fatal(err)
	}
	u.Path = path
	u.RawQuery = fmt.Sprint(time.Now().UnixNano())

	log.Println("loading image file: " + u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}
