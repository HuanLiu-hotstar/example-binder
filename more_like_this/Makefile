all:target


target:
	go build

plugin:
	go build -o more_like_this.so -buildmode=plugin   binder.go main.go 
clean:
	rm -f more_like_this
