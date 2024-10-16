package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func SendSms(code int) {

	url := "https://kqq4g8.api.infobip.com/sms/2/text/advanced"
	method := "POST"

	payload := map[string]any{
		"messages": []map[string]any{
			{
				"destinations": []map[string]any{
					{"to": "41793026727"},
				},
				"from": "InfoSMS",
				"text": fmt.Sprintf("Your verification code is: %v", code),
			},
		},
	}

	jsonPayload, _ := json.Marshal(payload)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(jsonPayload))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "your api key")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println(string(body))
	fmt.Println(string(body))
}
