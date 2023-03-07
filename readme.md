# mdimg2base64

windows上でWindowsキー+shift+Sで撮ったスクリーンショットを
wsl上のvimで編集中のmarkdownにbase64で埋め込みたかった。

## usage

```sh
go install github.com/xcd0/mdimg2base64@latest
```

とした後に、vimrcに
```
nnoremap sp :r!mdimg2base64<CR>
```
とかでどうでしょうか

## もっといい名前なかったのか

\_(:3 」∠ )\_


