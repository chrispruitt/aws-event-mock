# aws-event-mock

A simple tool to quickly mockup aws test events like cloudwatch logs or sns events

## Usage

```text
Usage:
  aws-event-mock [command]

Available Commands:
  create      returns a valid aws event in json format
  help        Help about any command

Flags:
  -h, --help      help for aws-event-mock
      --version   version for aws-event-mock
```

## Examples

```bash
aws-event-mock create --message --type cloudwatch-log --message '{"message":"i am a cloudwatch log event. send me to lambda!"}' | jq
{
  "awslogs": {
    "data": "H4sIAAAAAAAA/0SNT0+EMBBHv8vv3E1atrq7vZGIXPQEN0MMyoQ0obRpB40hfHfDH93jmzfzZob/HinCQGVn/fB4ud5UdobA4Psy+inAgCnxyx9upuJIrburgwXS9JE+ow1s/fhsB6aYYN62tR3RCDhKqe2p/gkEg6e8zt9fi6rKy2KvF1808no3w3bHE1pntlMQYOsocesCjNJaap3drhcp5X8YBqcTr/WlWX4BAAD//wEAAP//obAJk+EAAAA="
  }
}
```

```bash
# from file use '@' before file path (supports relative or absolute)
aws-event-mock create --type cloudwatch-log --message "@/some/filepath"
{
  "awslogs": {
    "data": "H4sIAAAAAAAA/0SNT0+EMBBHv8vv3E1atrq7vZGIXPQEN0MMyoQ0obRpB40hfHfDH93jmzfzZob/HinCQGVn/fB4ud5UdobA4Psy+inAgCnxyx9upuJIrburgwXS9JE+ow1s/fhsB6aYYN62tR3RCDhKqe2p/gkEg6e8zt9fi6rKy2KvF1808no3w3bHE1pntlMQYOsocesCjNJaap3drhcp5X8YBqcTr/WlWX4BAAD//wEAAP//obAJk+EAAAA="
  }
}

aws-event-mock create --type sns --message '{"message":"i am an sns message. send me to lambda!"}' | jq
{
  "Records": [
    {
      "EventVersion": "1.0",
      "EventSubscriptionArn": "arn:aws:sns:us-east-1:{{{accountId}}}:ExampleTopic",
      "EventSource": "aws.sns",
      "Sns": {
        "Signature": "EXAMPLE",
        "MessageId": "95df01b4-ee98-5cb9-9903-4c221d41eb5e",
        "Type": "Notification",
        "TopicArn": "arn:aws:sns:us-east-1:{{{accountId}}}:ExampleTopic",
        "MessageAttributes": {
          "Test": {
            "Type": "String",
            "Value": "teststringvalue"
          }
        },
        "SignatureVersion": "1",
        "Timestamp": "2021-11-30T22:08:23.379513Z",
        "SigningCertUrl": "EXAMPLE",
        "Message": "{\"message\":\"i am an sns message. send me to lambda!\"}",
        "UnsubscribeUrl": "EXAMPLE",
        "Subject": "example subject"
      }
    }
  ]
}
```

### Roadmap

- [ ] Add more event types
  - [x] Cloudwatch Log
  - [ ] Cloudwatch
  - [x] SNS
  - [ ] S3
  - [ ] Dynamo
  - [ ] API Gateway