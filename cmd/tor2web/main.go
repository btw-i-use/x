// Command tor2web is a simple caching proxy to tor onion sites.
package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/birkelund/boltdbcache"
	"github.com/gregjones/httpcache"
	"golang.org/x/net/proxy"
)

var (
	dbLoc        = flag.String("db-loc", "./cache.db", "cache location on disk (boltdb)")
	torSocksAddr = flag.String("tor-socks-addr", "127.0.0.1:9050", "tor socks address")
	httpPort     = flag.String("port", "8000", "HTTP port")
)

func main() {
	flag.Parse()

	// Create a socks5 dialer
	dialer, err := proxy.SOCKS5("tcp", *torSocksAddr, nil, proxy.Direct)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP transport
	tr := &http.Transport{
		Dial: dialer.Dial,
	}

	c, err := boltdbcache.New(*dbLoc, boltdbcache.WithBucketName("darkweb"))
	if err != nil {
		log.Fatalln(err)
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

	log.Fatalln(http.ListenAndServe(":"+*httpPort, rp))
}
