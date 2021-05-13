test:
	go run main.go create --message --type cloudwatch-log '{"message":"i am a cloudwatch log event. send me to lambda!"}'

generate-enums: 
	go-enum --names --file lib/event_type_enum.go

unit:
	# Execute test recursivly
	go test -v ./...