package main

import (
	"context"
	"encoding/json"
	"log"
	"plugin"
	"testing"

	"github.com/hotstar/hs-core-page-compositor/binder/template/api"
)

// go test -v -run=TestPlugin
func TestPlugin(t *testing.T) {
	p, err := plugin.Open("playback.so")
	if err != nil {
		panic(err)
	}
	pb, err := p.Lookup("Playback")
	if err != nil {
		panic(err)
	}
	x, ok := pb.(api.Template)
	if !ok {
		panic("not implement template API")
	}
	ctx := context.TODO()
	bye := `{"playbackData":{"playbackSets":[{
		"playbackUrl":"hotstar.com"
	}]}}`
	res, err := x.Execute(ctx, json.RawMessage(bye))
	log.Printf("res:%+v, err:%s", res, err)

}
