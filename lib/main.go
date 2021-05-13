package lib

type Event struct {
	Message string
}

func GetEvent(event *Event, eventType EventTypeEnum) (resp string, err error) {
	switch eventType {
	case EventTypeEnumCloudwatchLog:
		resp, err = event.GetCloudwatchLogEvent()
	}

	return
}
