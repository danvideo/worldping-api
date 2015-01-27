package api

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"github.com/torkelo/grafana-pro/pkg/middleware"
	"github.com/torkelo/grafana-pro/pkg/setting"
	"github.com/torkelo/grafana-pro/pkg/util"
	"strconv"
)

func GraphiteProxy(c *middleware.Context) {
	proxyPath := c.Params("*")
	target, _ := url.Parse(setting.GraphiteUrl)

	director := func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.Header.Add("X-Account-Id", strconv.FormatInt(c.AccountId, 10))
		req.URL.Path = util.JoinUrlFragments(target.Path, proxyPath)
		
	}
	
	proxy := &httputil.ReverseProxy{Director: director}

	proxy.ServeHTTP(c.RW(), c.Req.Request)
}