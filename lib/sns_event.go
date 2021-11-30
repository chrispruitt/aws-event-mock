package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/aws/aws-lambda-go/events"
)

func (event *Event) GetSNSEvent() (resp string, err error) {

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
	event.Message = strings.ReplaceAll(event.Message, "\n", "")

	result := events.SNSEvent{
		Records: []events.SNSEventRecord{
			{
				EventSource:          "aws.sns",
				EventVersion:         "1.0",
				EventSubscriptionArn: "arn:aws:sns:us-east-1:{{{accountId}}}:ExampleTopic",
				SNS: events.SNSEntity{
					Type:             "Notification",
					MessageID:        "95df01b4-ee98-5cb9-9903-4c221d41eb5e",
					TopicArn:         "arn:aws:sns:us-east-1:{{{accountId}}}:ExampleTopic",
					Subject:          "example subject",
					Timestamp:        time.Now().UTC(),
					SignatureVersion: "1",
					Signature:        "EXAMPLE",
					SigningCertURL:   "EXAMPLE",
					UnsubscribeURL:   "EXAMPLE",
					MessageAttributes: map[string]interface{}{
						"Test": struct {
							Type  string `json:"Type"`
							Value string `json:"Value"`
						}{"String", "teststringvalue"},
					},
					Message: event.Message,
				},
			},
		},
	}

	marshal, err := json.Marshal(result)
	if err != nil {
		return
	}

	resp = string(marshal)

	return
}
