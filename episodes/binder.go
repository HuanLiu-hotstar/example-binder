package main

import (
	"context"
	"encoding/json"
	"path/filepath"

	"github.com/hotstar/hs-core-ui-models-go/feature/image"
	"github.com/hotstar/hs-core-ui-models-go/widget"
)

const androidPrefix string = "https://img1.hotstarext.com/image/upload/f_auto,t_web_hl_3x"

type BinderPlugin struct{}

var Binder BinderPlugin

func (p BinderPlugin) Execute(ctx context.Context, input json.RawMessage) (interface{}, error) {
	data := make([]Collection, 0, 10)
	if err := json.Unmarshal(input, &data); err != nil {
		return nil, err
	}

	items := []*widget.ScrollableTrayWidget_Item{}
	for _, d := range data {
		im := &image.Image{
			Src: getUrl(ctx, d.Images.HorizontalImage),
		}
		it :=
			&widget.ScrollableTrayWidget_Item{
				Widget: &widget.ScrollableTrayWidget_Item_DownloadableContent{
					DownloadableContent: &widget.DownloadableContentWidget{
						Data: &widget.DownloadableContentWidget_Data{
							Poster:      im,
							Title:       d.Title,
							Description: d.Description,
						},
					},
				},
			}

		items = append(items, it)
	}
	episodes := &widget.ScrollableTrayWidget{
		Data: &widget.ScrollableTrayWidget_Data{
			Items: items,
		},
	}
	return episodes, nil
}
func getUrl(ctx context.Context, source string) string {
	return filepath.Join(androidPrefix, source)
}

type Collection struct {
	Archived    bool   `json:"archived"`
	AssetType   string `json:"assetType"`
	CategoryID  int32  `json:"categoryId"`
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
