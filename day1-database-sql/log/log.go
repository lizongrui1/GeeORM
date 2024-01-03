package log

import (
	"log"
	"os"

	"golang.org/x/tools/go/analysis/passes/defers"
)

var (
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m", log.LstdFlags|log.Lshortfile)   //使用 log.Lshortfile 支持显示文件名和代码行号
	infoLog  = log.New(os.Stdout, "\033[34m[info]\033[0m", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex
)

//log methods
var (
	Error = errorLog.Println
	Errorf = errorLog.Printf
	Info = infoLog.Println
	Infof = infoLog.Printf
)

//log levels
const(
	InfoLevel = iota    //第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1
	ErrorLevel
	Disabled
)

//SetLevel controls log level
func SetLevel (level int){
	mu.Lock
	defer mu.Unlock()

	for_,logger := range loggers{
		logger.SetOutput(os.Stdout)
	}

	
}

