// Command tor2web is a simple caching proxy to tor onion sites.
package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/birkelund/boltdbcache"
	"github.com/gregjones/httpcache"
	"golang.org/x/net/proxy"
)

func main() {
	if err := Run(os.Args[1:]); err != nil {
		log.Fatalln(err)
	}
}

func Run(args []string) error {
	fs := flag.NewFlagSet("tor2web", flag.ExitOnError)
	dbLoc := fs.String("db-loc", "./cache.db", "cache location on disk (boltdb)")
	torSocksAddr := fs.String("tor-socks-addr", "127.0.0.1:9050", "tor socks address")
	httpPort := fs.String("port", "8000", "HTTP port")

	if err := fs.Parse(args); err != nil {
		return err
	}

	dialer, err := proxy.SOCKS5("tcp", *torSocksAddr, nil, proxy.Direct)
	if err != nil {
		return err
	}

	tr := &http.Transport{
		Dial: dialer.Dial,
	}

	c, err := boltdbcache.New(*dbLoc, boltdbcache.WithBucketName("darkweb"))
	if err != nil {
		return err
	}

	ttr := httpcache.NewTransport(c)
	ttr.Transport = tr

	rp := &httputil.ReverseProxy{
		Transport: ttr,
		Director: func(r *http.Request) {
			r.URL.Scheme = "http"
			r.URL.Host = r.Host
		},
	}

	return http.ListenAndServe(":"+*httpPort, rp)
}
