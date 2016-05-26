package main

import (
    "testing"
)

func Test_DoConnect(t *testing.T) {
	// 接続のみ実施
	//func DoConnect(url string, origin string) (ws *websocket.Conn, err error)

	// 接続出来る
	_, e := DoConnect("ws://echo.websocket.org", "http://echo.websocket.org/")
	if e != nil {
		t.Error("FAILURE [CONNECT FAILED]")
	}
	_, e = DoConnect("ws://echo.websocket.org", "http://aaaa.echo.websocket.org/")
	if e == nil {
		t.Error("FAILURE [CONNECT SUCCESS]")
	}
}

func Test_DoRequest(t *testing.T) {
	// メッセージのやり取りを実施
	//func DoRequest(statuses bool, requestBody string, ws *websocket.Conn) *checkers.Checker
}
