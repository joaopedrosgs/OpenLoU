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
	urlLogin := "http://localhost:8000/"
	randomLogin, _ := accountserver.GenerateRandomString(10)
	randomMail, _ := accountserver.GenerateRandomString(10)

	form := url.Values{
		"login":    {randomLogin},
		"email":    {randomMail},
		"password": {"senha12345"},
	}
	sendRequest(urlLogin, form, "register")

	loginres := sendRequest(urlLogin, form, "login")
	if loginres == nil {
		println("Erro ao logar")

		return
	}
	result := communication.Answer{}
	err := json.Unmarshal(loginres, &result)
	if err != nil {
		println(err.Error())
	}
	conn, err := net.Dial("tcp", "127.0.0.1:8080")

	key := result.Data.(string)

	begin := time.Now()
	request := communication.Request{}
	request.Key = key
	request.Type = 203

	data := make(map[string]uint)
	data["X"] = 11
	data["Y"] = 11
	data["CityID"] = 9
	data["Type"] = 1
	data["Range"] = 10
	data["Continent"] = 1
	request.Data = data
	sendWithData(key, conn, request)
	println(((time.Now().Sub(begin) * 1 / 1000000) * time.Millisecond).String())

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

func sendWithData(key string, conn net.Conn, request communication.Request) {

	buffer := make([]byte, 256)
	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)
	bytes, _ := json.Marshal(request)

	for i := 0; i < 10; i++ {
		_, err := writer.Write(bytes)

		if err != nil {
			println(err.Error())

			break
		}
		writer.Flush()

		_, err = reader.Read(buffer)
		if err != nil {
			println(err.Error())

			break
		}
		println(string(buffer))
	}
}
