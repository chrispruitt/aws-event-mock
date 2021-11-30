package lib

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestGetSNSEvent(t *testing.T) {
	testMessage := "{\"greet\":\"hello\",\n\"name\":\"bobby bob\"}"

	event := Event{
		Message: testMessage,
	}

	response, err := event.GetSNSEvent()
	if err != nil {
		t.Errorf("Error on GetSNSEvent, %s", err)
	}

	var snsEventResponse events.SNSEvent
	err = json.Unmarshal([]byte(response), &snsEventResponse)
	if err != nil {
		t.Errorf("Unable to unmarshal SNS Event string, %s", err)
	}

	expected := "{\"greet\":\"hello\",\"name\":\"bobby bob\"}"

	if !reflect.DeepEqual(snsEventResponse.Records[0].SNS.Message, expected) {
		fmt.Printf("Expected %v, got %v", expected, snsEventResponse.Records[0].SNS.Message)
		t.Error("SNS is not equal to expected response")
	}
}
