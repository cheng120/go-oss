package go_oss

import (
	"flag"
	"go-oss/log"
	"go-oss/router"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)
type Server struct {
	port     string
	logDir   string
	logLevel string
	dataDir  string
}
type FileInfo struct {
	CreateTime   time.Time
	Md5          string
	Bucket  	 string
	DownloadPath string `json:",omitempty"`
}
type UploadResponseInfo struct {
	ErrInfo
	File FileInfo
}
type FileServerInfo struct {
	ErrInfo
	ID         string
	FileNumber int
}
var svr = &Server{}

func main(){
	flag.StringVar(&svr.logDir, "logDir", "logs", "dir to save all logs")
	flag.StringVar(&svr.dataDir, "dataDir", "data", "data directory")
	flag.StringVar(&svr.port, "port", "50010", "web api port")
	flag.StringVar(&svr.logLevel, "logLevel", "MORE", `log level: NORMAL, MORE, MUCH`)
	flag.Parse()
	var show log.ShowLevel
	switch svr.logLevel {
	case "NORMAL":
		show = log.NORMAL
	case "MORE":
		show = log.MORE
	case "MUCH":
		show = log.MUCH
	default:
		log.Fatal("log_level only support: NORMAL, MORE, MUCH")
	}
	l := log.NewLoggerEx(1, show, "")
	if svr.logDir != "" {
		if err := os.MkdirAll(svr.logDir, os.ModePerm); err != nil {
			log.Fatal(err)
		}
		rw := log.NewRotateWriter(path.Join(svr.logDir, "repo.log"))
		defer rw.Close()
		l.SetOutput(rw)
	}
	log.SetStd(l)
	if loc, err := time.LoadLocation("Asia/Shanghai"); err == nil {
		time.Local = loc
	} else {
		log.Warn(err)
	}
	var err error
	svr.dataDir, err = filepath.Abs(svr.dataDir)
	if err != nil {
		log.Fatal(err)
	}
	//if _, err := InitDB(); err == nil {
	//	log.Println("DB init done")
	//	fileNum = getFileNumInDB()
	//} else {
	//	log.Fatal(err)
	//	return
	//}
	//defer db.Close()
	//go deleteExpiredFile()
	mainRouter := router.NewRouter()
	log.Infof("run server on: %s", svr.port)
	err := http.ListenAndServe(":"+svr.port, mainRouter)
}