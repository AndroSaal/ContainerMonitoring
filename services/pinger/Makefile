#CONFIG
CONFIGPATH = ./config
CONFIGNAME = local.yaml

run : clean build
	./.bin/pinger  --config_path=$(CONFIGPATH) --config_file=$(CONFIGNAME)
	
build :
	go build -o ./.bin/pinger ./internal/cmd/main.go

clean :
	rm -rf ./.bin