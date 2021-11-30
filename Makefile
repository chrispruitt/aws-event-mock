test:
	go run main.go create --type cloudwatch-log --message '{"message":"i am a cloudwatch log event. send me to lambda!"}'
	go run main.go create --type sns --message '{"message":"i am an sns message. send me to lambda!"}'

generate-enums: 
	go-enum --names --file lib/event_type_enum.go

unit:
	# Execute test recursivly
	go test -v ./...

release:
	goreleaser --rm-dist