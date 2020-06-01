package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

/*
func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)
}
*/

func main() {
	//https://golang.org/pkg/log/
	//golangにはエラーのログを出すものがない
	/*
		logging.debug("")
		logging.info("")
		logging.warn("")
		logging.error("")
		logging.exception("")
	*/
	_, err := os.Open("alsdkjf")
	if err != nil {
		log.Fatalln("Exit", err)
	}

	log.Println("Logging!")
	log.Printf("%T, %v", "test", "test")

	log.Fatalf("%T, %v", "Test", "Test") //Fatalシリーズは処理終了！
	log.Fatalln("error!!")               //Faralln()を処理がここで終了

	fmt.Println("OK!")

}
