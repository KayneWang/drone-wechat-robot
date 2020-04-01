package main

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type (
	Repo struct {
		Owner string `json:"owner"`
		Name  string `json:"name"`
	}

	Build struct {
		Tag        string `json:"tag"`
		Event      string `json:"event"`
		Number     int    `json:"number"`
		Commit     string `json:"commit"`
		CommitLink string `json:"commit_link"`
		Ref        string `json:"ref"`
		Branch     string `json:"branch"`
		Author     string `json:"author"`
		Message    string `json:"message"`
		Status     string `json:"status"`
		Link       string `json:"link"`
		Started    int64  `json:"started"`
		Created    int64  `json:"created"`
	}

	Config struct {
		Key                 string   `json:"key"`
		MsgType             string   `json:"msgtype"`
		Content             string   `json:"content"`
		MentionedList       []string `json:"mentioned_list"`
		MentionedMobileList []string `json:"mentioned_mobile_list"`
		Base64              string   `json:"base64"`
		Md5                 string   `json:"md5"`
	}

	Job struct {
		Started int64 `json:"started"`
	}

	Plugin struct {
		Repo   Repo
		Build  Build
		Config Config
		Job    Job
	}
)

type (
	TextContent struct {
		Content             string   `json:"content"`
		MentionedList       []string `json:"mentioned_list"`
		MentionedMobileList []string `json:"mentioned_mobile_list"`
	}
	MarkdownContent struct {
		Content string `json:"content"`
	}
	ImageContent struct {
		Base64 base64 `json:"base64"`
		Md5    md5    `json:"md5"`
	}
)

func (p Plugin) Exec() error {
	url := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + p.Config.Key

	var data interface{}

	switch p.Config.MsgType {
	case "text":
		text := TextContent{
			Content:             p.Config.Content,
			MentionedList:       p.Config.MentionedList,
			MentionedMobileList: p.Config.MentionedMobileList,
		}

		data = struct {
			MsgType string      `json:"msgtype"`
			Text    TextContent `json:"text"`
		}{p.Config.MsgType, text}

	case "markdown":
		markdown := MarkdownContent{
			Content: p.Config.Content,
		}
		data = struct {
			MsgType  string          `json:"msgtype"`
			Markdown MarkdownContent `json:"markdown"`
		}{p.Config.MsgType, markdown}

	case "image":
		b64, err := base64.StdEncoding.DecodeString(p.Config.Base64)
		m5 := md5.Sum([]byte(p.Config.Md5))
		if err != nil {
			fmt.Printf("Error: failed to encoding base64")
			return err
		}
		image := ImageContent{
			Base64: b64,
			Md5:    m5,
		}
		data = struct {
			MsgType string       `json:"msgtype"`
			Image   ImageContent `json:"image"`
		}{p.Config.MsgType, image}

	default:
		return errors.New("Error: wrong msgtype, you should use either text, markdown, image, news")
	}

	b, _ := json.Marshal(data)
	fmt.Printf("Request data: %s\n", string(b))

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		fmt.Printf("Error: failed to create the request. %s\n", err)
	}
	request.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("Error: failed to excute the HTTP request. %s\n", err)
		return err
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	responseBody, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(responseBody))

	return nil
}
