all:target


target:
	go build

plugin:
	go build -o playback.so -buildmode=plugin  binder.go main.go 
clean:
	rm -f playback 
