package main

import (
	"fmt"
	"log"
	"os"

	"github.com/soarqin/go-webview2"
)

func main() {
	os.Setenv("WEBVIEW2_ADDITIONAL_BROWSER_ARGUMENTS", "--enable-features=msWebView2EnableDraggableRegions")
	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:     true,
		AutoFocus: true,
		WindowOptions: webview2.WindowOptions{
			Title:           "Minimal webview example",
			Width:           800,
			Height:          600,
			IconId:          2, // icon resource id
			Center:          true,
			Borderless:      true,
			BackgroundColor: &webview2.Color{R: 0, G: 0, B: 0, A: 0},
		},
	})
	if w == nil {
		log.Fatalln("Failed to load webview.")
	}
	defer w.Destroy()
	w.Bind("div_loaded", func(left int, top int, width int, height int) {
		fmt.Println("left:", left, " top:", top, " width:", width, " height:", height)
		w.SetSize(left+width, top+height, webview2.HintFixed)
	})
	w.SetSize(800, 600, webview2.HintFixed)
	w.Navigate("http://localhost:11451")
	w.Eval(`document.addEventListener('DOMContentLoaded', function() {
		const rect = document.getElementById('move').getBoundingClientRect();
		window.div_loaded(rect.left, rect.top, rect.width, rect.height);
	});`)
	w.Run()
}
