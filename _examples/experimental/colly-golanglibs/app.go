package main

import (
	"time"

	"github.com/asdine/storm"
	"github.com/gin-gonic/gin"

	// experimental addons
	// onion "github.com/sniperkit/colly/plugins/net/protocol/http/proxy/onion"
	// morty "github.com/sniperkit/colly/plugins/net/protocol/http/proxy/morty"

	// core
	colly "github.com/sniperkit/colly/pkg"
	proxy "github.com/sniperkit/colly/pkg/proxy/default"
)

/*
	Refs:
	- https://github.com/keyid/overlord/blob/master/overlord.go
	- https://github.com/keyid/overlord/blob/master/src/overlord/models/page.go
*/

// concurrent maps, datasets and databooks defaults
var (
	conf_file    string
	results_file string
	showVersion  bool = false
	verbose      bool = false
	version           = APP_VERSION
)

func init() {

	const (
		default_conf         = "example.toml"
		usage_conf           = "conf file path"
		usage_version        = "show version"
		default_showversion  = false
		usage_results_file   = "results file path"
		default_results_file = "example.results"
		usage_log_file       = "enable stdout to log"
		default_log_file     = true
		default_worker_qd    = 10000
		usage_worker_qd      = "queue depth for worker requests"

		usage_verbose   = "print debug logs"
		default_verbose = false

		usage_memprofile   = "write mem profile to file"
		default_memprofile = false

		usage_cpuprofile   = "write cpu profile to file"
		default_cpuprofile = false

		usage_enable_ui   = "enable terminal ui"
		default_enable_ui = true

		usage_dump_failures   = "enable 4xx status requests dump to file"
		defaule_dump_failures = false

		usage_dump_location   = "location of dump requests"
		default_dump_location = "."
	)

	flag.StringVar(&conf_file, "conf", default_conf, usage_conf)
	flag.StringVar(&conf_file, "c", default_conf, usage_conf+" (shorthand)")
	flag.StringVar(&results_file, "o", default_results_file, usage_results_file+" (shorthand)")
	flag.BoolVar(&showVersion, "version", default_showversion, usage_version)
	flag.BoolVar(&cpu_profile, "p", default_cpuprofile, usage_cpuprofile)
	flag.BoolVar(&mem_profile, "m", default_memprofile, usage_memprofile)
	flag.BoolVar(&enable_log, "d", default_log_file, usage_log_file)
	flag.BoolVar(&verbose, "v", default_verbose, usage_verbose)
	flag.IntVar(&worker_qd, "q", default_worker_qd, usage_worker_qd)
	flag.BoolVar(&enable_ui, "u", default_enable_ui, usage_enable_ui)
	flag.BoolVar(&dump_failures, "f", defaule_dump_failures, usage_dump_failures)
	flag.StringVar(&dump_location, "l", default_dump_location, usage_dump_location)

}

///// App ///////////////////////////////////////////////
type Application struct {
	Name string
	// Version   models.Version
	Database  *storm.DB
	Collector colly.Collector
	// Config    models.Config
}

func NewApp() (*Application, error) {

	db, _ := storm.Open("shared/storage/kvs/storm/servers.db")

	app := &Application{
		Name: "Overlord",
		// Version:  models.Version{0, 1, 0},
		Database: db,
		// Config:   models.Config{},
		Collector: *colly.NewCollector(
			//colly.Debugger(&debug.LogDebugger{}),
			colly.IgnoreRobotsTxt(),
			colly.MaxDepth(15),
			colly.CacheDir("./cache/"),
			//colly.DisallowedDomains("facebook.com", "twitter.com"),
			colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64; rv:52.0) Gecko/20100101 Firefox/52.0"),
			//colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
			colly.Async(false),
			//colly.URLFilters(
			//	regexp.MustCompile("(.+|^facebook|^twitter)$"),
			//),
		),
	}
	// app.PrintBanner()
	// defer app.Database.Close()
	// app.InitWebUI()
	// app.SetCollectorLimits(2, 4)

	return app, nil
}

func (self Application) SetCollectorLimits(parallelism int, delay time.Duration) {
	self.Collector.Limit(&colly.LimitRule{
		Parallelism: parallelism,
		RandomDelay: delay * time.Second,
	})
}

func (self Application) SetCollectorProxies(socksProxies string) (err error) {
	proxies, err := proxy.RoundRobinProxySwitcher(socksProxies)
	if err == nil {
		self.Collector.SetProxyFunc(proxies)
	}
	return err
}

///// Main ////////////////////////////////////////////
func (a Application) PrintBanner() {
	// fmt.Println(Bold(Magenta(fmt.Sprintf("%v: Network Detector (v%v)", a.Name, a.Version.String()))))
	// fmt.Println(Bold(Magenta(fmt.Sprintf("%v: Network Detector", a.Name))))
	// fmt.Println(Gray("=================================="))
}

func (a *Application) InitWebUI() {
	// TODO: Provide a web API that serves results of crawling
	router := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	// v1 := router.Group("/api/v1")
	// v1.GET("/pages", a.Pages)
	// v1.GET("/domains", a.Domains)
	// v1.GET("/servers", a.Servers)

	go router.Run(":8080")
}

/*
func LoadJSONConfig(path string) (config app.Config) {
	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}
	json.Unmarshal(fileData, &config)
	return config
}

func (a *Application) Servers(c *gin.Context) {
	var servers []models.Server
	err := app.Database.All(&servers)
	if err != nil {
		c.JSON(200, gin.H{"error": err})
	} else {
		c.IndentedJSON(200, servers)
	}
}

func (app *Application) Domains(c *gin.Context) {
	var domains []models.Domain
	err := app.Database.All(&domains)
	if err != nil {
		c.JSON(200, gin.H{"error": err})
	} else {
		c.IndentedJSON(200, domains)
	}
}

func (app *Application) Pages(c *gin.Context) {
	var pages []models.Page
	err := app.Database.All(&pages)
	if err != nil {
		c.JSON(200, gin.H{"error": err})
	} else {
		c.IndentedJSON(200, pages)
	}
}
*/
