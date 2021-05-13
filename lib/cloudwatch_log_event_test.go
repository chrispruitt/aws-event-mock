package lib

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestGetCloudwatchLogEvent(t *testing.T) {
	// example of encoded raw data "H4sIAAAAAAAAAHWPwQqCQBCGX0Xm7EFtK+smZBEUgXoLCdMhFtKV3akI8d0bLYmibvPPN3wz00CJxmQnTO41whwWQRIctmEcB6sQbFC3CjW3XW8kxpOpP+OC22d1Wml1qZkQGtoMsScxaczKN3plG8zlaHIta5KqWsozoTYw3/djzwhpLwivWFGHGpAFe7DL68JlBUk+l7KSN7tCOEJ4M3/qOI49vMHj+zCKdlFqLaU2ZHV2a4Ct/an0/ivdX8oYc1UVX860fQDQiMdxRQEAAA=="
	// example of decoded raw data "{\"messageType\":\"DATA_MESSAGE\",\"owner\":\"123456789123\",\"logGroup\":\"testLogGroup\",\"logStream\":\"testLogStream\",\"subscriptionFilters\":[\"testFilter\"],\"logEvents\":[{\"id\":\"eventId1\",\"timestamp\":1440442987000,\"message\":\"[ERROR] First test message\"},{\"id\":\"eventId2\",\"timestamp\":1440442987001,\"message\":\"[ERROR] Second test message\"}]}"
	testMessage := "testmessage"

	event := Event{
		Message: testMessage,
	}

	response, err := event.GetCloudwatchLogEvent()
	if err != nil {
		t.Errorf("Error on zipAndBase64(), %s", err)
	}

	var cloudwatchLogsEvent events.CloudwatchLogsEvent
	err = json.Unmarshal([]byte(response), &cloudwatchLogsEvent)
	if err != nil {
		t.Errorf("Unable to unmarshal CloudwatchLogsData string, %s", err)
	}

	parsedResponse, err := cloudwatchLogsEvent.AWSLogs.Parse()
	if err != nil {
		t.Errorf("Unable to parse CloudWatch Logs data, %s", err)
	}

	expected := events.CloudwatchLogsData{
		Owner:               "123456789123",
		LogGroup:            "testLogGroup",
		LogStream:           "testLogStream",
		SubscriptionFilters: []string{"testFilter"},
		MessageType:         "DATA_MESSAGE",
		LogEvents: []events.CloudwatchLogsLogEvent{
			{
				ID:        "testeventid1",
				Timestamp: 1440442987000,
				Message:   testMessage,
			},
		},
	}

	if !reflect.DeepEqual(parsedResponse, expected) {
		t.Error("CloudwatchLogsData is not equal to expected response")
	}
}
