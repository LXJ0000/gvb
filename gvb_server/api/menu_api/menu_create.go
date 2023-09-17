package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/ctype"
)

type ImageSort struct {
	ImageID int `json:"image_id"`
	Sort    int `json:"sort"`
}

type Menu struct {
	MenuTitle     string      `json:"menu_title"`
	MenuTitleEn   string      `json:"menu_title_en"`
	Slogan        string      `json:"slogan"`
	Abstract      ctype.Array `json:"abstract"`
	AbstractTime  int         `json:"abstract_time"`
	BannerTime    int         `json:"banner_time"`
	Sort          int         `json:"sort"`
	ImageSortList []ImageSort `json:"image_sort_list"` // 具体图片的顺序
}

func (MenuApi) MenuCreateView(c *gin.Context) {

}
