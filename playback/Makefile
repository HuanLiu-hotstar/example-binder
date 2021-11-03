all:target


target:
	go build

plugin:
	go build -o playback.so -buildmode=plugin   playbackurl.go main.go 
clean:
	rm -f playback 
