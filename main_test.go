package main

import (
	"encoding/json"
	handlers "github.com/belfazt/proxy-app/api/handlers"
	middleware "github.com/belfazt/proxy-app/api/middleware"
	server "github.com/belfazt/proxy-app/api/server"
	utils "github.com/belfazt/proxy-app/api/utils"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
)

type Response struct {
	Status       int            `json:"status,omitempty"`
	Response     string         `json:"result,omitempty"`
	ResponseText []ResponseText `json:"res,omitempty"`
}

type ResponseText struct {
	Domain string
}

func init() {
	var wg = &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		utils.LoadEnv()
		middleware.Init()
		var app = server.SetUp()
		handlers.HandleRedirection(app)
		wg.Done()
		server.RunServer(app)
	}(wg)
	wg.Wait()
}

func TestSorting(t *testing.T) {
	var cases = []struct {
		Domain string
		Output string
	}{
		{Domain: "alpha", Output: `["alpha"]`},
		{Domain: "omega", Output: `["alpha","omega"]`},
		{Domain: "alpha", Output: `["alpha","alpha","omega"]`},
		{Domain: "beta", Output: `["alpha","alpha","omega","beta"]`},
		{Domain: "", Output: "error"},
	}

	for _, caze := range cases {
		var valuesToCompare = &Response{}
		var client = http.Client{}
		var req, err = http.NewRequest("GET", "http://localhost:8080/", nil)

		assert.Nil(t, err)

		req.Header.Add("domain", caze.Domain)

		var response, err2 = client.Do(req)

		assert.Nil(t, err2)

		var bytes, err3 = ioutil.ReadAll(response.Body)

		assert.Nil(t, err3)

		json.Unmarshal(bytes, valuesToCompare)
		assert.NotNil(t, valuesToCompare.Response)
		assert.Equal(t, caze.Output, valuesToCompare.Response)
	}
}
