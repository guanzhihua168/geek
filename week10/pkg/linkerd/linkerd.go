package linkerd

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	pkg_errors "github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type LinkerdResult struct {
	Code    int
	Message string
	Data    json.RawMessage
}
type LinkerdOptions func(c *Linkerd)
type WithReqstContext func(c *http.Request)

var LinkerdRequestClient *http.Client

func WithJson() LinkerdOptions {
	return func(c *Linkerd) {
		c.header["Content-Type"] = "application/json"
	}
}
func WithContext(f WithReqstContext) LinkerdOptions {
	return func(c *Linkerd) {
		c.reqContext = f
	}
}

func WithLogger(log *logrus.Entry) LinkerdOptions {
	return func(c *Linkerd) {
		if log == nil {
			return
		}
		c.log = log
	}
}
func NewLinkerd(host, appId, appToken, serverName string, options ...LinkerdOptions) *Linkerd {
	l := &Linkerd{
		host:     host,
		appId:    appId,
		appToken: appToken,
	}

	l.header = map[string]string{
		"Service-Token": appToken,
		"App-Id":        appId,
		"host":          serverName,
		"X-Trace":       "linkerd-proxy",
		"Content-Type":  "application/x-www-form-urlencoded",
	}

	for _, o := range options {
		o(l)
	}

	if l.log == nil {
		l.log = logrus.WithField("App-Id", appId)
	}
	return l
}

type Linkerd struct {
	host       string
	appId      string
	appToken   string
	header     map[string]string
	trace      string
	reqContext WithReqstContext
	log        *logrus.Entry
}

func (c *Linkerd) fullUrl(path string, query ...map[string]string) string {
	host := strings.TrimRight(c.host, "/") + "/" + strings.TrimLeft(path, "/")
	q := url.Values{}
	if len(query) > 0 {
		for k, v := range query[0] {
			q.Add(k, v)
		}
	}

	if len(q) > 0 {
		host = "?" + q.Encode()
	}
	return host
}
func (c *Linkerd) Get(path string, query ...map[string]string) (*LinkerdResult, error) {
	req, e := http.NewRequest("GET", c.fullUrl(path, query...), nil)
	c.log = c.log.WithField("request_url", c.fullUrl(path, query...))
	if e != nil {
		logrus.Warn("new request fail:" + e.Error())
		return nil, e
	}

	return c.do(req)
}

func (c *Linkerd) Post(path string, body []byte) (*LinkerdResult, error) {
	fpath := c.fullUrl(path)
	req, e := http.NewRequest("POST", fpath, bytes.NewReader(body))
	c.log = c.log.WithFields(logrus.Fields{
		"request_url":  fpath,
		"request_body": string(body),
	})
	if e != nil {
		c.log.Error("new request fail:" + e.Error())
		return nil, pkg_errors.Wrap(e, "new request fail")
	}
	return c.do(req)
}

func (c *Linkerd) do(req *http.Request) (lr *LinkerdResult, err error) {
	lr = &LinkerdResult{}
	for k, v := range c.header {
		if strings.ToLower(k) != "host" {
			req.Header.Set(k, v)
		} else {
			req.Host = v
		}
	}

	if c.reqContext != nil {
		c.reqContext(req)
	}

	resp, e := LinkerdRequestClient.Do(req)
	c.log = c.log.WithFields(logrus.Fields{
		"url":    req.URL.String(),
		"method": req.Method,
		"header": req.Header,
		"host":   req.Host,
	})

	if e != nil {
		c.log.Error(e.Error())
		err = e
		return
	}

	if resp.StatusCode != 200 {
		c.log.Error(resp.StatusCode)
		err = pkg_errors.Wrap(errors.New("请求非200"), "")
		return
	}
	rb, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	c.log = c.log.WithField("result", string(rb))
	c.log.Info("ok")

	if e := json.Unmarshal(rb, &lr); e != nil {
		err = e
		c.log.Error(e.Error())
		return
	}

	if lr.Code != 200 {
		c.log.Error("code !200")
	}
	return lr, nil
}
