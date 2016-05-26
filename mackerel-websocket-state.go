package main

import (
	"os"
	"golang.org/x/net/websocket"

	"github.com/jessevdk/go-flags"
	"github.com/mackerelio/checkers"
)

type Options struct {
    Origin  string  `short:"o" long:"origin" value-name:"N" description:"origin url"`
    Url     string  `short:"u" long:"url" value-name:"N" description:"websocket endpoint url"`
    Body    string  `short:"b" long:"body" value-name:"N" description:"request body string"`
    Ignore  bool    `short:"i" long:"ignore" value-name:"N" description:"no confirming response message"`
    Status  bool    `short:"s" long:"status" value-name:"N" description:"confirming a communication status"`
}

var opts Options

func main() {
	stat := run(os.Args[1:])
	stat.Name = "WebSocketStat"
	stat.Exit()
}

func run(args[]string) *checkers.Checker {
	// 引数処理
	_, err := flags.ParseArgs(&opts, args)
	if err != nil {
		os.Exit(1)
	}

	// 接続
	ws, err := DoConnect(opts.Url, opts.Origin)

	// 接続エラーチェック
	if err != nil {
		return checkers.NewChecker(checkers.CRITICAL, "can not connected. " + err.Error())
	}

	// 接続確認のみの場合、この時点でsuccessを返す
	if opts.Ignore {
		return checkers.NewChecker(checkers.OK, "Connected.")
	}

	return DoRequest(opts.Status, opts.Body, ws)
}

// 接続のみ実施
func DoConnect(url string, origin string) (ws *websocket.Conn, err error) {
	ws, err = websocket.Dial(url, "", origin);
	return
}

// メッセージのやり取りを実施
func DoRequest(statuses bool, requestBody string, ws *websocket.Conn) *checkers.Checker {
	msg := ""

	// 送信
	err := websocket.Message.Send(ws, requestBody)
	if err != nil {
		// エラー
		return checkers.NewChecker(checkers.CRITICAL, "communication error: " + err.Error())
	}
	// 受信
	err  = websocket.Message.Receive(ws, &msg)
	if statuses && err != nil {
		// エラー
		return checkers.NewChecker(checkers.CRITICAL, "response:" + msg + ", communication error: " + err.Error())
	}

	return checkers.NewChecker(checkers.OK, "response=" + msg)
}
