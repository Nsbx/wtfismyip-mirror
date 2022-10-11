package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/caddyserver/certmagic"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/oschwald/geoip2-golang"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	middlewarestd "github.com/slok/go-http-metrics/middleware/std"
)

var xffMode bool
var cityReader *geoip2.Reader
var orgReader *geoip2.Reader
var templateHTML *template.Template
var templateJSON *template.Template
var templateYAML *template.Template
var templateXML *template.Template
var templateClean *template.Template

type geoText struct {
	org         string
	details     string
	countryCode string
	city        string
	country     string
	state       string
}

type wtfResponse struct {
	IPv6        bool
	Address     string
	Hostname    string
	Geo         string
	ISP         string
	CountryCode string
	Tor         bool
	Myipwtf     bool
}

var ctx2 = context.TODO()

var ctx = context.Background()

var rdb *redis.Client

func main() {

	xffMode = false

	if len(os.Args) == 2 {
		if os.Args[1] == "--xff" {
			xffMode = true
		}
	}
	var err error

	rdb = redis.NewClient(&redis.Options{
		Addr:     "172.19.1.70:6379",
		Password: "",
		DB:       0})

	cityReader, err = geoip2.Open("/usr/local/wtf/GeoIP/GeoIP2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}

	orgReader, err = geoip2.Open("/usr/local/wtf/GeoIP/GeoIP2-ISP.mmdb")
	if err != nil {
		log.Fatal(err)
	}

	defer cityReader.Close()
	defer orgReader.Close()

	templateHTML, err = template.ParseFiles("/usr/local/wtf/static/html.template")
	if err != nil {
		log.Fatal(err)
	}
	templateJSON, err = template.ParseFiles("/usr/local/wtf/static/json.template")
	if err != nil {
		log.Fatal(err)
	}
	templateYAML, err = template.ParseFiles("/usr/local/wtf/static/yaml.template")
	if err != nil {
		log.Fatal(err)
	}
	templateXML, err = template.ParseFiles("/usr/local/wtf/static/xml.template")
	if err != nil {
		log.Fatal(err)
	}
	templateClean, err = template.ParseFiles("/usr/local/wtf/static/clean.template")
	if err != nil {
		log.Fatal(err)
	}

	mdlw := middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})

	r := mux.NewRouter()
	h := middlewarestd.Handler("", mdlw, r)

	r.Host("ipv5.wtfismyip.com").HandlerFunc(ipv5Handler)
	r.Host("ipv7.wtfismyip.com").HandlerFunc(ipv5Handler)
	r.Host("text.wtfismyip.com").HandlerFunc(text)
	r.Host("text.myip.wtf").HandlerFunc(text)
	r.Host("ipv4.text.wtfismyip.com").HandlerFunc(text)
	r.Host("ipv4.text.myip.wtf").HandlerFunc(text)
	r.Host("text.ipv4.wtfismyip.com").HandlerFunc(text)
	r.Host("text.ipv4.myip.wtf").HandlerFunc(text)
	r.Host("ipv6.text.wtfismyip.com").HandlerFunc(text)
	r.Host("ipv6.text.myip.wtf").HandlerFunc(text)
	r.Host("text.ipv6.wtfismyip.com").HandlerFunc(text)
	r.Host("text.ipv6.myip.wtf").HandlerFunc(text)
	r.Host("json.wtfismyip.com").HandlerFunc(json)
	r.Host("json.myip.wtf").HandlerFunc(json)
	r.Host("ipv4.json.wtfismyip.com").HandlerFunc(json)
	r.Host("ipv4.json.myip.wtf").HandlerFunc(json)
	r.Host("json.ipv4.wtfismyip.com").HandlerFunc(json)
	r.Host("json.ipv4.myip.wtf").HandlerFunc(json)
	r.Host("ipv6.json.wtfismyip.com").HandlerFunc(json)
	r.Host("ipv6.json.myip.wtf").HandlerFunc(json)
	r.Host("json.ipv6.wtfismyip.com").HandlerFunc(json)
	r.Host("json.ipv6.myip.wtf").HandlerFunc(json)
	r.Host("xml.wtfismyip.com").HandlerFunc(xml)
	r.Host("xml.myip.wtf").HandlerFunc(xml)
	r.Host("ipv4.xml.wtfismyip.com").HandlerFunc(xml)
	r.Host("ipv4.xml.myip.wtf").HandlerFunc(xml)
	r.Host("xml.ipv4.wtfismyip.com").HandlerFunc(xml)
	r.Host("xml.ipv4.myip.wtf").HandlerFunc(xml)
	r.Host("ipv6.xml.wtfismyip.com").HandlerFunc(xml)
	r.Host("ipv6.xml.myip.wtf").HandlerFunc(xml)
	r.Host("xml.ipv6.wtfismyip.com").HandlerFunc(xml)
	r.Host("xml.ipv6.myip.wtf").HandlerFunc(xml)
	r.Host("clean.wtfismyip.com").HandlerFunc(cleanHandle)
	r.HandleFunc("/clean", cleanHandle)
	r.HandleFunc("/headers", headers)
	r.HandleFunc("/test", test)
	r.HandleFunc("/json", json)
	r.HandleFunc("/yaml", yaml)
	r.HandleFunc("/xml", xml)
	r.HandleFunc("/text", text)
	r.HandleFunc("/text/isp", textisp)
	r.HandleFunc("/text/geo", textgeo)
	r.HandleFunc("/text/city", textcity)
	r.HandleFunc("/text/country", textcountry)
	r.HandleFunc("/text/ip", text)
	r.HandleFunc("/js", jsHandle)
	r.HandleFunc("/jsclean", jscleanHandle)
	r.HandleFunc("/js2", js2Handle)
	r.HandleFunc("/js2clean", js2cleanHandle)
	r.HandleFunc("/clean", cleanHandle)
	r.HandleFunc("/church", cleanHandle)
	r.HandleFunc("/traffic", trafficHandle)
	r.HandleFunc("/omgwtfbbq.png", trafficPngHandle)
	r.HandleFunc("/.git/config", gitconfigHandle)
	r.HandleFunc("/_ignition/health-check/}", healthHandle)
	r.HandleFunc("/public/_ignition/health-check/}", healthHandle)
	r.HandleFunc("/", wtfHandle).Methods("GET")
	r.HandleFunc("/", miscHandle).Methods("POST")
	r.HandleFunc("/", miscHandle).Methods("PUT")
	r.HandleFunc("/", miscHandle).Methods("DELETE")
	r.HandleFunc("/", miscHandle).Methods("TRACE")
	r.HandleFunc("/admin", adminHandle)
	r.HandleFunc("/administrator", adminHandle)
	r.HandleFunc("/metrics", metricsHandle)
	r.HandleFunc("/{foo:.*log$}", miscHandle)
	r.HandleFunc("/{foo:.*bak$}", miscHandle)
	r.HandleFunc("/{foo:.*swp$}", miscHandle)
	r.HandleFunc("/{foo:.*~$}", miscHandle)
	r.HandleFunc("/{foo:.*sql$}", sqlHandle)
	r.HandleFunc("/{foo:.*zip$}", zipHandle)
	r.HandleFunc("/{foo:.*gz$}", gzHandle)
	r.HandleFunc("/{foo:.*ini$}", iniHandle)
	r.HandleFunc("/{foo:.*php$}", trollHandle)
	r.HandleFunc("/{foo:.*asp$}", trollHandle)
	r.HandleFunc("/{foo:.*aspx$}", trollHandle)
	r.NotFoundHandler = http.HandlerFunc(custom404)

	config := certmagic.NewDefault()
	tags := []string{}
	config.CacheUnmanagedCertificatePEMFile(ctx2, "/docker/certs/wtf.cert.pem", "/docker/certs/wtf.key.pem", tags)
	tlsConfig := config.TLSConfig()

	srvHTTPS := &http.Server{
		ReadTimeout:  8 * time.Second,
		WriteTimeout: 16 * time.Second,
		IdleTimeout:  16 * time.Second,
		Addr:         ":10443",
		Handler:      h,
		TLSConfig:    tlsConfig,
	}

	srvHTTP := &http.Server{
		ReadTimeout:  8 * time.Second,
		WriteTimeout: 16 * time.Second,
		IdleTimeout:  16 * time.Second,
		Handler:      h,
		Addr:         ":10080",
	}

	go srvHTTP.ListenAndServe()
	srvHTTPS.ListenAndServeTLS("", "")
}

func geoData(ip string) (location geoText) {
	var isCityPresent bool

	address := net.ParseIP(ip)
	isp, err := orgReader.ISP(address)
	if err != nil {
		log.Println(err)
	}

	record, err := cityReader.City(address)
	if err != nil {
		log.Println(err)
	}

	if len(record.Subdivisions) > 0 {
		location.state = record.Subdivisions[0].IsoCode
	}

	location.city, isCityPresent = record.City.Names["en"]
	location.country, _ = record.Country.Names["en"]
	location.country = strings.Replace(location.country, "Palestinian Territory", "Occupied Palestinian Territory", 1)
	location.countryCode = record.Country.IsoCode

	if isCityPresent {
		if len(location.state) > 0 {
			location.details = location.city + ", " + location.state + ", " + location.country
		} else {
			location.details = location.city + ", " + location.country
		}
	} else {
		if len(location.country) > 0 {
			location.details = location.country
		} else {
			location.details = "Unknown"
		}
	}

	if len(location.countryCode) == 0 {
		location.countryCode = "Unknown"
	}

	location.org = isp.ISP
	return location
}

func reverseDNS(ip string) (response string) {
	omfg := make(chan string, 1)
	go func() {
		dnsName, err := net.LookupAddr(ip)
		if err != nil {
			omfg <- ip
		}
		if len(dnsName) == 0 {
			omfg <- ip
		} else {
			hostname := dnsName[0]
			omfg <- hostname[0 : len(hostname)-1]
		}
	}()

	select {
	case response = <-omfg:
		return response
	case <-time.After(2 * time.Second):
		return ip
	}
}

func custom404(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("/usr/local/wtf/static", filepath.Clean(r.URL.Path))
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No such fucking page!")
	}
	w.Write(contents)
}

func js2Handle(w http.ResponseWriter, r *http.Request) {
	contents, err := ioutil.ReadFile("/usr/local/wtf/static/js2")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No such fucking page!")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(contents)
}

func js2cleanHandle(w http.ResponseWriter, r *http.Request) {
	contents, err := ioutil.ReadFile("/usr/local/wtf/static/js2clean")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No such fucking page!")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(contents)
}

func miscHandle(w http.ResponseWriter, r *http.Request) {
	contents, err := ioutil.ReadFile("/usr/local/wtf/static/evil.log")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No such fucking page!")
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Write(contents)
}

func sqlHandle(w http.ResponseWriter, r *http.Request) {
	contents, err := ioutil.ReadFile("/usr/local/wtf/static/evil.sql")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No such fucking page!")
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Write(contents)
}

func zipHandle(w http.ResponseWriter, r *http.Request) {
	contents, err := ioutil.ReadFile("/usr/local/wtf/static/evil.zip")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No such fucking page!")
	}
	w.Header().Set("Content-Type", "application/zip")
	w.Write(contents)
}

func gzHandle(w http.ResponseWriter, r *http.Request) {
	contents, err := ioutil.ReadFile("/usr/local/wtf/static/evil.tar.gz")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No such fucking page!")
	}
	w.Header().Set("Content-Type", "application/gzip")
	w.Write(contents)
}

func iniHandle(w http.ResponseWriter, r *http.Request) {
	contents, err := ioutil.ReadFile("/usr/local/wtf/static/evil.ini")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No such fucking page!")
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Write(contents)
}

func adminHandle(w http.ResponseWriter, r *http.Request) {
	contents, err := ioutil.ReadFile("/usr/local/wtf/static/admin.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No such fucking page!")
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write(contents)
}

func trollHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<html><head><meta http-equiv=\"Refresh\" content=\"0; url=https://www.youtube.com/watch?v=sTSA_sWGM44\" /></head><body><p>TROLOLOLOL!</p></body></html>")
}

// lets add some really rudimentary and shitty IP allowlisting to block access to explicit metrics
func metricsHandle(w http.ResponseWriter, r *http.Request) {
	allowlistAddr, _ := rdb.Get(ctx, "allowlistAddr").Result()
	add := getAddress(r)
	if add == allowlistAddr {
		promhttp.Handler().ServeHTTP(w, r)
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("sorry dude"))
	}
}

func json(w http.ResponseWriter, r *http.Request) {
	add := getAddress(r)
	hostname := reverseDNS(add)
	geo := geoData(add)
	isIPv6 := strings.Contains(add, ":")
	isTor := isTorExit(add)
	resp := wtfResponse{isIPv6, add, hostname, geo.details, geo.org, geo.countryCode, isTor, false}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	templateJSON.Execute(w, resp)
}

func yaml(w http.ResponseWriter, r *http.Request) {
	add := getAddress(r)
	hostname := reverseDNS(add)
	geo := geoData(add)
	isIPv6 := strings.Contains(add, ":")
	isTor := isTorExit(add)
	resp := wtfResponse{isIPv6, add, hostname, geo.details, geo.org, geo.countryCode, isTor, false}
	w.Header().Set("Content-Type", "text/yaml")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	templateYAML.Execute(w, resp)
}

func text(w http.ResponseWriter, r *http.Request) {
	add := getAddress(r)
	response := add + "\n"
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	fmt.Fprintf(w, response)
}

func textisp(w http.ResponseWriter, r *http.Request) {
	add := getAddress(r)
	response := geoData(add).org + "\n"
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	fmt.Fprintf(w, response)
}

func textgeo(w http.ResponseWriter, r *http.Request) {
	add := getAddress(r)
	response := geoData(add).details + "\n"
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	fmt.Fprintf(w, response)
}

func textcountry(w http.ResponseWriter, r *http.Request) {
	add := getAddress(r)
	response := geoData(add).countryCode + "\n"
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	fmt.Fprintf(w, response)
}

func textcity(w http.ResponseWriter, r *http.Request) {
	add := getAddress(r)
	response := geoData(add).city + "\n"
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	fmt.Fprintf(w, response)
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Yes, the website is fucking running\n")
}

func jsHandle(w http.ResponseWriter, r *http.Request) {
	add := getAddress(r)
	hostname := reverseDNS(add)
	geo := geoData(add)
	isIPv6 := strings.Contains(add, ":")
	if isIPv6 && r.Host == "ipv4.wtfismyip.com" {
		w.WriteHeader(http.StatusMisdirectedRequest)
		w.Write([]byte("Fucking protocol error"))
	} else {
		response := "ip='" + add + "';\n"
		response += "hostname='" + hostname + "';\n"
		response += "geolocation='" + geo.details + "';\n"
		response += "document.write('<center><p><h2>Your fucking IPv4 address is:</h2></center>');document.write('<center><p>' + ip + '</center>');document.write('<center><p><h2>Your fucking IPv4 hostname is:</h2></center>');document.write('<center><p>' + hostname + '</center>');document.write('<center><p><h2>Geographic location of your fucking IPv4 address:</h2></center>');document.write('<center><p>' + geolocation + '</center>');"
		fmt.Fprintf(w, response)
	}
}

func jscleanHandle(w http.ResponseWriter, r *http.Request) {
	add := getAddress(r)
	hostname := reverseDNS(add)
	geo := geoData(add)
	isIPv6 := strings.Contains(add, ":")
	if isIPv6 && r.Host == "ipv4.wtfismyip.com" {
		w.WriteHeader(http.StatusMisdirectedRequest)
		w.Write([]byte("Fucking protocol error"))
	} else {
		response := "ip='" + add + "';\n"
		response += "hostname='" + hostname + "';\n"
		response += "geolocation='" + geo.details + "';\n"
		response += "document.write('<center><p><h2>Your IPv4 address is:</h2></center>');document.write('<center><p>' + ip + '</center>');document.write('<center><p><h2>Your IPv4 hostname is:</h2></center>');document.write('<center><p>' + hostname + '</center>');document.write('<center><p><h2>Geographic location of your IPv4 address:</h2></center>');document.write('<center><p>' + geolocation + '</center>');"
		fmt.Fprintf(w, response)
	}
}

func xml(w http.ResponseWriter, r *http.Request) {
	add := getAddress(r)
	hostname := reverseDNS(add)
	geo := geoData(add)
	isIPv6 := strings.Contains(add, ":")
	isTor := isTorExit(add)
	resp := wtfResponse{isIPv6, add, hostname, geo.details, geo.org, geo.countryCode, isTor, false}
	w.Header().Set("Content-Type", "application/xml")
	templateXML.Execute(w, resp)
}

func cleanHandle(w http.ResponseWriter, r *http.Request) {
	add := getAddress(r)
	isIPv6 := strings.Contains(add, ":")
	hostname := reverseDNS(add)
	geo := geoData(add)
	isTor := isTorExit(add)
	resp := wtfResponse{isIPv6, add, hostname, geo.details, geo.org, geo.countryCode, isTor, false}
	templateClean.Execute(w, resp)
}

func wtfHandle(w http.ResponseWriter, r *http.Request) {
	add := getAddress(r)
	isIPv6 := strings.Contains(add, ":")
	hostname := reverseDNS(add)
	geo := geoData(add)
	isTor := isTorExit(add)
	myipwtf := false
	if r.Host == "myip.wtf" {
		myipwtf = true
	}
	resp := wtfResponse{isIPv6, add, hostname, geo.details, geo.org, geo.countryCode, isTor, myipwtf}
	if r.TLS == nil {
		if myipwtf {
			http.Redirect(w, r, "https://myip.wtf/", 301)
		} else {
			http.Redirect(w, r, "https://wtfismyip.com/", 301)
		}
	} else {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("X-Hire-Me", "clint@wtfismyip.com")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		w.Header().Set("X-OMGWTF", "BBQ")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Content-Security-Policy", "default-src 'none'; img-src wtfismyip.com myip.wtf; script-src ipv4.wtfismyip.com wtfismyip.com myip.wtf ipv4.myip.wtf; style-src 'unsafe-inline'")
		w.Header().Set("X-DNS-Prefetch-Control", "off")
		w.Header().Set("Referrer-Policy", "no-referrer")
		w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
		w.Header().Set("X-Did-I-Set-Too-Many-Fucking-Headers", "Yes. I just wanted a fucking A from securityheaders.io.")
		templateHTML.Execute(w, resp)
	}
}

func headers(w http.ResponseWriter, r *http.Request) {
	var response string
	for name, val := range r.Header {
		response += name + ": " + val[0] + "\n"
	}
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, response)
}

func ipv5Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "No such fucking protocol")
}

func healthHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "I am healthy, motherfucker!")
}

func trafficHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><head><img src=\"/omgwtfbbq.png\" style=\"width: 100%; object-fit: contain\"></head></html>")
}

func gitconfigHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "[user]\n\tname = wtfismyip\n\temail = wtfismyip@nsa.gov\n\n[github]\n\tuser = wtfismyip\n\ttoken = lmfaotrolololo")
}

func trafficPngHandle(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("/usr/local/tmp/omgwtfbbq.png")
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "Fucking Error: %s", err)
	}
	w.Header().Set("Content-Type", "image/png")
	w.Write([]byte(content))
}

func getAddress(r *http.Request) (ip string) {
	if xffMode {
		ip = r.Header.Get("X-Forwarded-For")
		return
	}
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "0.0.0.0"
	}
	return
}

func isTorExit(ip string) bool {
	val, _ := rdb.Get(ctx, ip).Result()

	if val == "exit" {
		return true
	}
	return false
}
