package test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/joaopedrosgs/OpenLoU/accountserver"
	"github.com/joaopedrosgs/OpenLoU/communication"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

func RunTest() {
	url_login := "http://localhost:8000/"
	random_login, _ := accountserver.GenerateRandomString(10)
	random_mail, _ := accountserver.GenerateRandomString(10)

	form := url.Values{
		"login":    {random_login},
		"email":    {random_mail},
		"password": {"senha12345"},
	}
	sendRequest(url_login, form, "register")
	loginres := sendRequest(url_login, form, "login")
	if loginres == nil {
		return
	}
	result := communication.Answer{}
	err := json.Unmarshal(loginres, &result)
	if err != nil {
		println(err.Error())
	}
	key := result.Data.(string)
	request := communication.Request{}
	request.Key = key
	request.Type = 102
	request.Data = make(map[string]string)
	request.Data["X"] = "10"
	request.Data["Y"] = "10"
	request.Data["Continent"] = "1"
	request.Data["Range"] = "10"
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		return
	}
	buffer := make([]byte, 1024)
	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)
	bytes, _ := json.Marshal(request)
	inicio := time.Now()
	for i := 0; i < 10; i++ {
		_, err := writer.Write(bytes)
		if err != nil {
			break
		}
		writer.Flush()
		_, err = reader.Read(buffer)
		if err != nil {
			break
		}

	}
	println(time.Now().Sub(inicio).String())
	conn.Close()

}

func sendRequest(url string, form url.Values, sub string) []byte {
	body := bytes.NewBufferString(form.Encode())
	rsp, err := http.Post(url+sub, "application/x-www-form-urlencoded", body)
	if err != nil {
		return nil
	}
	defer rsp.Body.Close()
	body_byte, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}
	return body_byte
}
