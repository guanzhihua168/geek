package http

import (
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	sky "github.com/SkyAPM/go2sky/plugins/http"
	simplejson "github.com/bitly/go-simplejson"
	"github.com/sirupsen/logrus"
	"parent-api-go/global"
)

var (
	HttpDefault *HttpClient
)

type HttpClient struct {
	client       *http.Client
	headers      map[string]string
	responseType string
}

func init() {
	HttpDefault = &HttpClient{
		client: &http.Client{
			Transport: &http.Transport{
				Dial: func(netr, addr string) (net.Conn, error) {
					conn, e := net.DialTimeout(netr, addr, time.Second*6)
					if e != nil {
						return nil, e
					}
					conn.SetDeadline(time.Now().Add(time.Second * 6))
					return conn, nil
				},
				MaxIdleConns:          200,
				ResponseHeaderTimeout: time.Second * 6,
			},
			Timeout: 6 * time.Second,
		},
		headers: map[string]string{},
	}

	if global.SKWTrace != nil {
		if client, e := sky.NewClient(global.SKWTrace, sky.WithClient(HttpDefault.client)); e == nil {
			HttpDefault.client = client
			return
		} else {
			logrus.WithField("trace", global.SKWTrace).Error("new trace client fail:" + e.Error())
		}
	}
}

func NewHttpClient() *HttpClient {
	ht := &HttpClient{
		client: &http.Client{
			Transport: &http.Transport{
				Dial: func(netr, addr string) (net.Conn, error) {
					conn, e := net.DialTimeout(netr, addr, time.Second*6)
					if e != nil {
						return nil, e
					}
					conn.SetDeadline(time.Now().Add(time.Second * 6))
					return conn, nil
				},
				MaxIdleConns:          200,
				ResponseHeaderTimeout: time.Second * 6,
			},
			Timeout: 6 * time.Second,
		},
		headers: map[string]string{},
	}
	if global.SKWTrace != nil {
		if client, e := sky.NewClient(global.SKWTrace, sky.WithClient(ht.client)); e == nil {
			ht.client = client
			return ht
		} else {
			logrus.WithField("trace", global.SKWTrace).Error("new trace client fail:" + e.Error())
		}
	}
	return ht
}

// add headers
func (c *HttpClient) AddHeader(k, v string) *HttpClient {
	c.headers[k] = v
	return c
}

func (c *HttpClient) Get(uri string) ([]byte, error) {
	req, _ := http.NewRequest("GET", uri, nil)
	return c.doRequest(req, []byte(""))
}

func (c *HttpClient) Post(uri string, data *url.Values) ([]byte, error) {
	req, _ := http.NewRequest("POST", uri, strings.NewReader(data.Encode()))
	return c.doRequest(req, []byte(data.Encode()))
}

func (c *HttpClient) PostJson(uri string, data []byte) ([]byte, error) {
	req, _ := http.NewRequest("POST", uri, strings.NewReader(string(data)))
	return c.doRequest(req, data)
}

func (c *HttpClient) doRequest(req *http.Request, data []byte) ([]byte, error) {
	if len(c.headers) > 0 {
		for k, v := range c.headers {
			if strings.ToLower(k) != "host" {
				req.Header.Set(k, v)
			} else {
				req.Host = v
			}

		}
	}

	resp, err := c.client.Do(req)
	logFields := logrus.Fields{
		"url":          req.URL.String(),
		"method":       req.Method,
		"header":       req.Header,
		"request-body": string(data),
		"server-name":  "service-activity",
	}
	if err != nil {
		logFields["error"] = err.Error()
		logrus.WithFields(logFields).Warning("Http Request Done Fail")
		return []byte{}, errors.New("Request Connect Fail")
	}
	logFields["respone-status"] = resp.StatusCode
	b, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	logFields["response-body"] = string(b)
	if resp.StatusCode != 200 {
		logrus.WithFields(logFields).Warning("status-code !200")
	} else {
		logrus.WithFields(logFields).Info("http request")
	}
	return b, nil
}

func (c HttpClient) JsonFormatter(b []byte) *simplejson.Json {
	json, e := simplejson.NewJson(b)
	if e != nil {
		logrus.WithFields(logrus.Fields{
			"[]byte": string(b),
			"error":  "Json Formatter error:" + e.Error(),
		})
		return simplejson.New()
	}
	return json
}
