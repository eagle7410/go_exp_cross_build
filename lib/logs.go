package lib

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const logfile = "development.log"

var logFileName string

func OpenLogFile() {
	t := time.Now()
	oldFileName := logFileName

	logFileName = t.Format("2006-01-02") + "_" + logfile

	if FileExists("./logs") == false {
		err := os.Mkdir("./logs", 0777)

		if err != nil {
			log.Fatal("Mkdir logs error %s", err)
		}
	}

	if len(oldFileName) == 0 || oldFileName != logFileName {
		fmt.Printf("Logging to file %v\n", logFileName)

		lf, err := os.OpenFile("./logs/"+logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)

		if err != nil {
			log.Fatal("OpenLogfile: os.OpenFile: %s", err)
		}

		log.SetOutput(lf)
	}
}

func LogRequest(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		OpenLogFile()

		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

		handler.ServeHTTP(w, r)
	})
}

func LogAppRun(port string) {
	mask := "\n=== App run At (" + time.Now().Format("2006-01-02T15:04:05") + ") in http://localhost:" + port + " ===\n"
	fmt.Printf(mask)
	log.Printf(mask)
}

func LogFatalf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
	log.Fatalf(format, v...)
}

func Logf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
	log.Printf(format, v...)
}

func LogEF(format string, v ...interface{}) {
	Logf("[0;31m"+format+"[39m\n", v...)
}
