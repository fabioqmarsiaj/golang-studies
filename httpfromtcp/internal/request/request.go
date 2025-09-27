package request

import (
	"errors"
	"io"
	"log"
	"strings"
	"unicode"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	content, err := io.ReadAll(reader)
	if err != nil {
		log.Fatalf("Error reading content from Reader: %s\n", err.Error())
	}
	fullContent := string(content)

	requestLine, err := parseRequestLine(fullContent)
	if err != nil {
		return nil, errors.New("fail creating RequestLine struct")
	}

	log.Printf("RequestLine created: %+v\n", requestLine)

	return &Request{
		RequestLine: *requestLine,
	}, nil
}

func parseRequestLine(fullContent string) (*RequestLine, error) {
	contentArr := strings.Split(fullContent, "\r\n")
	requestLine := contentArr[0]
	reqLinesArr := strings.Split(requestLine, " ")

	if len(reqLinesArr) < 3 {
		return nil, errors.New("request lines are incomplete")
	}

	err := requestLineValidations(reqLinesArr)

	if err != nil {
		return nil, errors.New("fail on RequestLine validations")
	}

	method := reqLinesArr[0]
	requestTarget := reqLinesArr[1]
	httpVersion := strings.Split(reqLinesArr[2], "/")

	return &RequestLine{
		HttpVersion:   httpVersion[1],
		RequestTarget: requestTarget,
		Method:        method,
	}, nil
}

func requestLineValidations(reqLinesArr []string) error {
	method := strings.TrimSpace(reqLinesArr[0])
	httpVersion := reqLinesArr[2]

	for _, r := range method {
		if !unicode.IsUpper(r) {
			return errors.New("method contains lowercase letters")
		}
	}
	log.Printf("Method is valid: %s\n", method)

	httpSplit := strings.Split(httpVersion, "/")
	if httpSplit[1] != "1.1" {
		return errors.New("HTPP Version is not 1.1")
	}
	log.Printf("HTTP Version is valid: %s\n", httpSplit[1])

	return nil
}
