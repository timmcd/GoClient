all: goclient

goclient: goclient.8
	8l -o GoClient goclient.8

goclient.8:
	8g goclient.go

clean:
	rm GoClient
	rm goclient.8