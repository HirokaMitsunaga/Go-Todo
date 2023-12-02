package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings (logFile string){
	//OpenFileの引数でファイルの読み書き、ファイルの作成、ファイルの追記を指定している
	//0666はログファイルのパーミッション
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	//
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
} 