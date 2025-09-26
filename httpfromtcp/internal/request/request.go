package request

import (
	"io"
	"log"
	"strings"
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
	parseRequestLine(fullContent)

}

func parseRequestLine(fullContent string) (*RequestLine, error) {
	contentArr := strings.Split(fullContent, "\r\n")
	requestLine := contentArr[0]
	reqjuestLineBrokenDown := strings.Split()

}
