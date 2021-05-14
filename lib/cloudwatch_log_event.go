package lib

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"unicode/utf8"

	"github.com/aws/aws-lambda-go/events"
)

// Cloudwatch log event example
// {
// 	"awslogs": {
// 	  "data": "H4sIAAAAAAAAAHWPwQqCQBCGX0Xm7EFtK+smZBEUgXoLCdMhFtKV3akI8d0bLYmibvPPN3wz00CJxmQnTO41whwWQRIctmEcB6sQbFC3CjW3XW8kxpOpP+OC22d1Wml1qZkQGtoMsScxaczKN3plG8zlaHIta5KqWsozoTYw3/djzwhpLwivWFGHGpAFe7DL68JlBUk+l7KSN7tCOEJ4M3/qOI49vMHj+zCKdlFqLaU2ZHV2a4Ct/an0/ivdX8oYc1UVX860fQDQiMdxRQEAAA=="
// 	}
//   }

// type CloudWatchLogEvent struct {
// }

func (event *Event) GetCloudwatchLogEvent() (resp string, err error) {

	// read from file if first char is @
	if string(event.Message[0]) == "@" {
		_, i := utf8.DecodeRuneInString(event.Message)
		var b []byte
		b, err = ioutil.ReadFile(event.Message[i:])
		if err != nil {
			err = fmt.Errorf("Error: unable to parse file: %s", err)
			return
		}
		event.Message = string(b)
	}

	// Remove whitespaces
	event.Message = strings.Join(strings.Fields(event.Message), "")

	logData := events.CloudwatchLogsData{
		Owner:               "123456789123",
		LogGroup:            "testLogGroup",
		LogStream:           "testLogStream",
		SubscriptionFilters: []string{"testFilter"},
		MessageType:         "DATA_MESSAGE",
		LogEvents: []events.CloudwatchLogsLogEvent{
			{
				ID:        "testeventid1",
				Timestamp: 1440442987000,
				Message:   event.Message,
			},
		},
	}

	logDataBytes, err := json.Marshal(logData)
	if err != nil {
		err = fmt.Errorf("Error marshaling CloudwatchLogsData: %s", err.Error())
		return
	}

	encodedCloudwatchLogEvent, err := zipAndBase64(string(logDataBytes))
	if err != nil {
		return
	}

	result := events.CloudwatchLogsEvent{
		AWSLogs: events.CloudwatchLogsRawData{
			Data: encodedCloudwatchLogEvent,
		},
	}

	marshal, err := json.Marshal(result)
	if err != nil {
		return
	}

	resp = string(marshal)

	return
}

func zipAndBase64(str string) (resp string, err error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)

	if _, err = gz.Write([]byte(str)); err != nil {
		return
	}
	if err = gz.Flush(); err != nil {
		return
	}
	if err = gz.Close(); err != nil {
		return
	}

	resp = base64.StdEncoding.EncodeToString(b.Bytes())

	return
}
