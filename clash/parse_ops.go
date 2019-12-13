package clash

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"net/http"
	u "net/url"
)

type ClashFile struct {
	Proxy []interface{} `yaml:"Proxy"`
}

type ProxyProvider struct {
	Proxies []interface{} `yaml:"proxies"`
}

func Parse(c *gin.Context) {
	url := c.Query("source")
	if url == "" {
		c.JSON(http.StatusOK, nil)
		return
	}
	url, err := u.QueryUnescape(url)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, nil)
		return
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, nil)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, nil)
		return
	}
	var file ClashFile
	err = yaml.Unmarshal(body, &file)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, nil)
		return
	}
	var target ProxyProvider
	target.Proxies = file.Proxy
	res, err := yaml.Marshal(&target)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, nil)
		return
	}
	c.String(http.StatusOK, string(res))
}