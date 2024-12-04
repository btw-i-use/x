// Command tor2web is a simple caching proxy to tor onion sites.
package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
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

func newTransport(torSocksAddr string, dbLoc string) (http.RoundTripper, error) {
	dialer, err := proxy.SOCKS5("tcp", torSocksAddr, nil, proxy.Direct)
	if err != nil {
		return nil, err
	}

	tr := &http.Transport{
		Dial: dialer.Dial,
	}

	if dbLoc == "" {
		return tr, nil
	}

	c, err := boltdbcache.New(dbLoc, boltdbcache.WithBucketName("tor2web"))
	if err != nil {
		return nil, err
	}

	ttr := httpcache.NewTransport(c)
	ttr.Transport = tr
	return ttr, nil
}

func Run(args []string) error {
	fs := flag.NewFlagSet("tor2web", flag.ExitOnError)
	dbLoc := fs.String("db-loc", "", "cache location on disk (boltdb)")
	torSocksAddr := fs.String("tor-socks-addr", "127.0.0.1:9050", "tor socks address")
	httpPort := fs.String("port", "8000", "HTTP port")

	if err := fs.Parse(args); err != nil {
		return err
	}

	transport, err := newTransport(*torSocksAddr, *dbLoc)
	if err != nil {
		return err
	}

	rp := &httputil.ReverseProxy{
		Transport: transport,
		Rewrite: func(pr *httputil.ProxyRequest) {
			pr.SetURL(&url.URL{
				Scheme: "http",
				Host:   pr.Out.Host,
			})
		},
	}

	return http.ListenAndServe(":"+*httpPort, rp)
}
