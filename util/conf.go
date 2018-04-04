package util

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/wycers/sustc-sakura-console/log"
	"github.com/wycers/sustc-sakura-console/model"
	"github.com/jinzhu/gorm"
)


type Configuration struct {
	Server                string // server scheme, host and port
	StaticServer          string // static resources server scheme, host and port
	StaticResourceVersion string // version of static resources
	Loglevel              string // logging level: trace/debug/info/warn/error/fatal
	SessionSecret         string // HTTP session secret
	SessionMaxAge         int    // HTTP session max age (in seciond)
	RuntimeMode           string // runtime mode (dev/prod)
	SQLite                string // SQLite database file path
	MySQL                 string // MySQL connection URL
	StaticRoot            string // static resources file root path
	Port                  string // listen port
	AxiosBaseURL          string // axio base URL
	MockServer            string // mock server
}

// Zero push time.
var ZeroPushTime, _ = time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")

var logger = log.NewLogger(os.Stdout)
var Config *Configuration
var Models = []interface{}{
	&model.Report{}, &model.Log{},
}

const Version = "1.0.0"
const tablePrefix = "Sakura_"

func LoadConfig() {
	version := flag.Bool("version", false, "Check out current version.")
	confFilePath := flag.String("conf", "config.json", "Set path of Config file.")
	confLoglevel := flag.String("loglevel", "", "this will override Config.Loglevel if specified")
	confServer := flag.String("server", "", "this will override Conf.Server if specified")
	confStaticServer := flag.String("static_server", "", "this will override Conf.StaticServer if specified")
	confStaticResourceVer := flag.String("static_resource_ver", "", "this will override Conf.StaticResourceVersion if specified")
	confRuntimeMode := flag.String("runtime_mode", "", "this will override Conf.RuntimeMode if specified")
	confSQLite := flag.String("sqlite", "", "this will override Conf.SQLite if specified")
	confStaticRoot := flag.String("static_root", "", "this will override Conf.StaticRoot if specified")
	confPort := flag.String("port", "", "this will override Conf.Port if specified")

	if *version {
		fmt.Println(Version)
		os.Exit(0)
	}

	bytes, err := ioutil.ReadFile(*confFilePath)
	if err != nil {
		logger.Fatalf("loads Configuration file[%s] failed: %s.", *confFilePath, err.Error())
	}
	Config = &Configuration{}
	if err = json.Unmarshal(bytes, Config); err != nil {
		logger.Fatalf("parses Config file failed:%s", err.Error())
	}

	log.SetLevel(Config.Loglevel)
	Override(&Config.Loglevel, *confLoglevel)
	if *confLoglevel != "" {
		log.SetLevel(*confLoglevel)
	}

	home, err := UserHome()
	if err != nil {
		logger.Fatalf("can't find user home directory: ", err.Error())
	}
	logger.Debugf("${home} [%s]", home)


	Override(&Config.RuntimeMode, *confRuntimeMode)

	Override(&Config.Server, *confServer)

	Config.StaticServer = Config.Server
	Override(&Config.StaticServer, *confStaticServer)

	time := strconv.FormatInt(time.Now().UnixNano(), 10)
	logger.Debugf("${time} [%s]", time)
	Config.StaticResourceVersion = strings.Replace(Config.StaticResourceVersion, "${time}", time, 1)
	Override(&Config.StaticResourceVersion, *confStaticResourceVer)

	Config.SQLite = strings.Replace(Config.SQLite, "${home}", home, 1)
	if "" != *confSQLite {
		Config.SQLite = *confSQLite
	}
	Config.StaticRoot = ""
	if *confStaticRoot != "" {
		Config.StaticRoot = *confStaticRoot
		Config.StaticRoot = filepath.Dir(Config.StaticRoot)
	}

	Override(&Config.Port, *confPort)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	logger.Debugf("Configurations [%#v]", Config)
}

func Override(item *string, str string) {
	if str != "" {
		*item = str
	}
}