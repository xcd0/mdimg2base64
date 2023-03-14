# mdimg2base64

windows上でWindowsキー+shift+Sで撮ったスクリーンショットを
wsl上のvimで編集中のmarkdownにbase64で埋め込みたかった。

## install

```sh
go install github.com/xcd0/mdimg2base64@latest
```

[ghg]( https://github.com/Songmu/ghg) を使うと簡単。
```
ghg get xcd0/mdimg2base64
```

## usage
vimrcに
```
nnoremap sp :r!mdimg2base64<CR>
```
のように書く。spは何でもよい。

`windows + shift + s`などでクリップボードにスクリーンショットを保存する。  
この状態でmdimg2base64を呼ぶと![](data:image/png;base64,iVBO........)というテキストが標準出力に出力される。  
なのでvimでmarkdownを開き、ノーマルモードで`sp`と打てば画像の文字列が挿入される。  




## もっといい名前なかったのか

ネーミングセンスをください\_(:3 」∠ )\_


