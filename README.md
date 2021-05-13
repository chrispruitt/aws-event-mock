# aws-event-mock

A simple tool to quickly mockup aws test events like cloudwatch logs

## Usage

```bash
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
› aws-event-mock create --message --type cloudwatch-log '{"message":"i am a cloudwatch log event. send me to lambda!"}' | jq
{
  "awslogs": {
    "data": "H4sIAAAAAAAA/0SNT0+EMBBHv8vv3E1atrq7vZGIXPQEN0MMyoQ0obRpB40hfHfDH93jmzfzZob/HinCQGVn/fB4ud5UdobA4Psy+inAgCnxyx9upuJIrburgwXS9JE+ow1s/fhsB6aYYN62tR3RCDhKqe2p/gkEg6e8zt9fi6rKy2KvF1808no3w3bHE1pntlMQYOsocesCjNJaap3drhcp5X8YBqcTr/WlWX4BAAD//wEAAP//obAJk+EAAAA="
  }
}
```