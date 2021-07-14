package notify

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"time"

	"github.com/riete/requests"
)

func SendDingTalkMarkdown(title, message, robotUrl, secret string, isAtAll bool) string {
	var body = map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"title": title,
			"text":  fmt.Sprintf("### %s\n\n%s", title, message),
		},
		"at": map[string]interface{}{
			"isAtAll": isAtAll,
		},
	}
	return SendDingTalk(title, message, robotUrl, secret, body)
}

func SendDingTalkText(title, message, robotUrl, secret string, isAtAll bool) string {
	var body = map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": fmt.Sprintf("%s\n\n%s", title, message),
		},
		"at": map[string]interface{}{
			"isAtAll": isAtAll,
		},
	}
	return SendDingTalk(title, message, robotUrl, secret, body)
}

func SendDingTalk(title, message, robotUrl, secret string, body map[string]interface{}) string {
	timestamp := fmt.Sprintf("%d000", time.Now().Unix())
	sign := fmt.Sprintf("%s\n%s", timestamp, secret)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(sign))
	signB64 := base64.StdEncoding.EncodeToString([]byte(h.Sum(nil)))

	v := url.Values{}
	v.Add("sign", signB64)
	signUrlEncode := v.Encode()
	sendUrl := fmt.Sprintf("%s&timestamp=%s&%s", robotUrl, timestamp, signUrlEncode)

	r := requests.NewRequest()
	if err := r.PostJson(sendUrl, body); err != nil {
		return err.Error()
	} else {
		return r.Content
	}
}
