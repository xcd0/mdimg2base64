package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	cb "github.com/atotto/clipboard"
	ci "github.com/xcd0/clipboard-image"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile) // ログの出力書式を設定する

	// wslの時はwindowsのクリップボードを参照したい
	// そういう処理はclipboard-imageライブラリのほうを修正した
	r, err := ci.Read()
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	ct := http.DetectContentType(b)
	//log.Printf("ct : %s", ct)
	if strings.HasPrefix(ct, "image/") {
		// 画像だったらbase64にして出力する
		//_, format, err := image.DecodeConfig(bytes.NewReader(b))
		//if err != nil {
		//	log.Fatal(err)
		//}
		//log.Printf("format : %s", format)
		str := base64.StdEncoding.EncodeToString(b)
		b64 := fmt.Sprintf("data:image/%s;base64,%s", ct[6:], str)
		if cb.WriteAll(b64) != nil {
			log.Fatal(err)
		}
	}
}
