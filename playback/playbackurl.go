package main

import (
	"context"
	"encoding/json"

	"github.com/hotstar/hs-core-page-compositor/binder/template/api"
	"github.com/hotstar/hs-core-ui-models-go/widget"
	"github.com/rs/zerolog/log"
)

type PlaybackUrl struct{}

var Playback PlaybackUrl

func NewPlaybackUrlBinder() api.Template {
	return PlaybackUrl{}
}

func (p PlaybackUrl) Execute(ctx context.Context, input json.RawMessage) (interface{}, error) {
	playbacks := GetPlaybackUrls{}
	if err := json.Unmarshal(input, &playbacks); err != nil {
		log.Err(err).Msg("error in unmarshal tray widget")
		return nil, err
	}
	// feed PlaybackUrl
	w := &widget.PlayerWidget{
		Data: &widget.PlayerWidget_Data{
			PlayerConfig: &widget.PlayerWidget_PlayerConfig{
				PlaybackUrls: []string{},
			},
		},
	}

	urls := make([]string, 0, len(playbacks.PlaybackData.PlaybackSets))
	for _, set := range playbacks.PlaybackData.PlaybackSets {
		urls = append(urls, set.PlaybackURL)
	}
	w.Data.PlayerConfig.PlaybackUrls = urls
	log.Debug().Msgf("playerWidget:%+v", w)
	return w, nil
}

type GetPlaybackUrls struct {
	Typename       string           `json:"__typename"`
	Message        string           `json:"message"`
	PlaybackData   PlaybackData     `json:"playbackData"`
	AdditionalInfo []AdditionalInfo `json:"additionalInfo"`
}
type PlaybackData struct {
	Match                bool           `json:"match"`
	ContentID            string         `json:"contentId"`
	ProfileRequestConfig string         `json:"profileRequestConfig"`
	PlaybackSets         []PlaybackSets `json:"playbackSets"`
}
type PlaybackSets struct {
	PlaybackCdnType string `json:"playbackCdnType"`
	PlaybackURL     string `json:"playbackUrl"`
	TagsCombination string `json:"tagsCombination"`
	TokenAlgorithm  string `json:"tokenAlgorithm"`
}
type AdditionalInfo struct {
	Key   string   `json:"key"`
	Value []string `json:"value"`
}
