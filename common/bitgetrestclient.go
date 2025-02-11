package common

import (
	"github.com/MoyunRz/bitget-sdk/config"
	"github.com/MoyunRz/bitget-sdk/constants"
	"github.com/MoyunRz/bitget-sdk/utils"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type BitgetRestClient struct {
	ApiKey       string
	ApiSecretKey string
	Passphrase   string
	BaseUrl      string
	HttpClient   http.Client
	Signer       *Signer
	IsTestNet    bool
}

func (p *BitgetRestClient) Init() *BitgetRestClient {
	p.ApiKey = config.ApiKey
	p.ApiSecretKey = config.SecretKey
	p.BaseUrl = config.BaseUrl
	p.Passphrase = config.PASSPHRASE
	p.IsTestNet = config.IsTestNet
	p.Signer = new(Signer).Init(config.SecretKey)
	p.HttpClient = http.Client{
		Timeout: time.Duration(config.TimeoutSecond) * time.Second,
	}
	return p
}

func (p *BitgetRestClient) DoPost(uri string, params string) (string, error) {
	timesStamp := utils.TimesStamp()
	//body, _ := internal.BuildJsonParams(params)

	sign := p.Signer.Sign(constants.POST, uri, params, timesStamp)
	requestUrl := config.BaseUrl + uri

	buffer := strings.NewReader(params)
	request, err := http.NewRequest(constants.POST, requestUrl, buffer)

	utils.Headers(request, p.ApiKey, timesStamp, sign, p.Passphrase)
	if err != nil {
		return "", err
	}
	if p.IsTestNet {
		request.Header.Add("paptrading", "1")
	}
	response, err := p.HttpClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bodyStr, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	responseBodyString := string(bodyStr)
	return responseBodyString, err
}

func (p *BitgetRestClient) DoGet(uri string, params map[string]string) (string, error) {
	timesStamp := utils.TimesStamp()
	body := utils.BuildGetParams(params)

	sign := p.Signer.Sign(constants.GET, uri, body, timesStamp)

	requestUrl := p.BaseUrl + uri + body

	request, err := http.NewRequest(constants.GET, requestUrl, nil)
	if err != nil {
		return "", err
	}
	utils.Headers(request, p.ApiKey, timesStamp, sign, p.Passphrase)
	if p.IsTestNet {
		request.Header.Add("paptrading", "1")
	}
	response, err := p.HttpClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bodyStr, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	responseBodyString := string(bodyStr)
	return responseBodyString, err
}
