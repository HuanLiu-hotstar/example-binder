package main

import (
	"context"
	"encoding/json"

	"github.com/hotstar/hs-core-ui-models-go/feature/image"
	"github.com/hotstar/hs-core-ui-models-go/widget"
)

type BinderPlugin struct{}

var Binder BinderPlugin

func (p BinderPlugin) Execute(ctx context.Context, input json.RawMessage) (interface{}, error) {
	data := make([]Collection, 0, 10)
	if err := json.Unmarshal(input, &data); err != nil {
		// log.Err(err).Msg("error in unmarshal tray widget")
		return nil, err
	}
	// images := []*image.Image{}
	items := []*widget.ScrollableTrayWidget_Item{
		{Widget: &widget.ScrollableTrayWidget_Item_VerticalContentPoster{
			VerticalContentPoster: &widget.VerticalContentPosterWidget{
				Data: &widget.VerticalContentPosterWidget_Data{
					Image: &image.Image{
						Src: "hello",
					},
				},
			},
		}},
	}
	for _, d := range data {
		im := &image.Image{
			Src: d.Images.VerticalImage,
		}
		it :=
			&widget.ScrollableTrayWidget_Item{
				Widget: &widget.ScrollableTrayWidget_Item_VerticalContentPoster{
					VerticalContentPoster: &widget.VerticalContentPosterWidget{
						Data: &widget.VerticalContentPosterWidget_Data{
							Image: im,
						},
					},
				},
			}

		items = append(items, it)
	}

	scrollableTrayWidget := &widget.ScrollableTrayWidget{
		Data: &widget.ScrollableTrayWidget_Data{
			Items: items,
		},
	}

	return scrollableTrayWidget, nil
}

type Collection struct {
	Archived    string `json:"archived"`
	AssetType   string `json:"assetType"`
	CategoryID  string `json:"categoryId"`
	ChannelName string `json:"channelName"`
	ContentID   string `json:"contentId"`
	ContentType string `json:"contentType"`
	Description string `json:"description"`
	EntityType  string `json:"entityType"`
	EpisodeCnt  int32  `json:"episodeCnt"`
	Images      struct {
		HorizontalImage string `json:"horizontalImage"`
		VerticalImage   string `json:"verticalImage"`
	} `json:"images"`
	IsSocialEnabled bool     `json:"isSocialEnabled"`
	Lang            []string `json:"lang"`
	LangObjs        []struct {
		Name string `json:"name"`
	} `json:"langObjs"`
	LoginNudgeStatus string   `json:"loginNudgeStatus"`
	PlaybackURI      string   `json:"playbackUri"`
	Premium          bool     `json:"premium"`
	StudioName       string   `json:"studioName"`
	Title            string   `json:"title"`
	Trailers         []string `json:"trailers"`
	URI              string   `json:"uri"`
	Vip              bool     `json:"vip"`
}
