/*
 * 簡易ログライブラリについて
 *   - ログレベルによる制御は可能
 *   - ログオブジェクトの変更可能
 */
package main

import (
    "log"
    "os"
)

/****** log library START ******/

type DebugLevel int

const (
    DEBUG DebugLevel = iota // == 0
    INFO  = iota // == 1
    WARN  = iota // == 2
    ERROR = iota // == 3
    FATAL = iota // == 4
)

type Logger interface {
    Debugf(string, ...interface{})
    Infof(string, ...interface{})
    Warnf(string, ...interface{})
    Errorf(string, ...interface{})
    Fatalf(string, ...interface{})
    SetLogLevel(c DebugLevel)
    SetLogger(l Logger)
}

var logger Logger = &BasicLogger{
    Logger: log.New(os.Stderr, "", log.LstdFlags),
}

func (bl *BasicLogger) SetLogger(l Logger) {
    logger = l
}

// デフォルトログレベルはINFO(ライブラリ内スコープ変数なので小文字)
var loglevel DebugLevel = INFO

type BasicLogger struct {
    Logger *log.Logger
}

func (bl *BasicLogger) SetLogLevel(c DebugLevel) {
    loglevel = c
}

func (bl *BasicLogger) Debugf(format string, v ...interface{}) {
	if loglevel > DEBUG {
		return
	}
    bl.Logger.Printf("[DEBUG] " + format, v...)
}

func (bl *BasicLogger) Infof(format string, v ...interface{}) {
	if loglevel > INFO {
		return
	}
    bl.Logger.Printf("[INFO] " + format, v...)
}

func (bl *BasicLogger) Warnf(format string, v ...interface{}) {
	if loglevel > WARN {
		return
	}
    bl.Logger.Printf("[WARN] " + format, v...)
}

func (bl *BasicLogger) Errorf(format string, v ...interface{}) {
	if loglevel > ERROR {
		return
	}
    bl.Logger.Printf("[ERROR] " + format, v...)
}

func (bl *BasicLogger) Fatalf(format string, v ...interface{}) {
	if loglevel > FATAL {
		return
	}
    bl.Logger.Printf("[FATAL] " + format, v...)
}

/****** log library END ******/

func main(){

	// 標準出力へと送られる
	logger.SetLogLevel(DEBUG)     // 最低ログレベルを指定する
	logger.Debugf("this is debug")
	logger.Infof("%s", "this is info")
	logger.Warnf("%d %s %s",100, "this is ", "warn")
	logger.Errorf("this is debug")
	logger.Fatalf("this is debug")

	// アプリケーション側でLogger設定を上書きすることができる。出力先やフラグを変更できる
	var applogger Logger = &BasicLogger{
		Logger: log.New(os.Stderr, "", log.LstdFlags | log.Lmicroseconds | log.Lshortfile),
	}
	logger.SetLogger(applogger)
	logger.Debugf("this is debug")
	logger.Infof("%s", "this is info")
	logger.Warnf("%d %s %s",100, "this is ", "warn")
	logger.Errorf("this is debug")
	logger.Fatalf("this is debug")
}
