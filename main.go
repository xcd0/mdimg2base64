package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"golang.design/x/clipboard"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile) // ログの出力書式を設定する

	// wslの時はwindowsのクリップボードを参照したい

	if err := clipboard.Init(); err != nil {
		panic(err)
	}

	// クリップボードから画像を読み込む。
	var imgdata []byte = clipboard.Read(clipboard.FmtImage)

	ct := http.DetectContentType(imgdata)
	if !strings.HasPrefix(ct, "image/") {
		log.Printf("クリップボードに画像がありませんでした。")
		log.Printf("クリップボードの中身 :\n%v", string(imgdata))
		return
	}

	// 画像だったらbase64にして出力する
	str := base64.StdEncoding.EncodeToString(imgdata)
	b64 := fmt.Sprintf("![](data:image/%s;base64,%s)", ct[6:], str)

	if len(os.Args) == 2 && os.Args[1] == "-" {
		// 引数に - が指定されているとき標準出力に出力する。
		fmt.Println(string(b64))
	} else {
		changed := clipboard.Write(clipboard.FmtText, []byte(b64))
		//log.Printf("クリップボードに書き込んでいます。")
		select {
		case <-changed:
			//log.Printf("クリップボードへの書き込みが完了しました。")
		}
	}
}
