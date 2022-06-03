package serializer

import "singo/model"

// item 商品序列化器
type Item struct {
	ID        uint     `json:"id"`
	Name      string   `json:"name"`
	Brief     string   `json:"brief"`
	Picture   string   `json:"picture"`
	Tag       []string `json:"tag"`
	Introduce string   `json:"introduce,omitempty"`
	Algorithm string   `json:"algorithm,omitempty"`
}

// BuildItem 序列化商品
func BuildItem(item model.Item, tags []model.Tag) Item {
	var tag []string = []string{}
	for _, i := range tags {
		tag = append(tag, i.TagName)
	}
	return Item{
		ID:        item.ID,
		Name:      item.Name,
		Brief:     item.Brief,
		Picture:   item.Picture,
		Tag:       tag,
		Introduce: item.Introduce,
		Algorithm: item.Algorithm,
	}
}

// BuildItemResponse 序列化商品响应
func BuildItemResponse(item model.Item, tags []model.Tag) Response {
	return Response{
		Data: BuildItem(item, tags),
	}
}

// BuildItemBrief 序列化商品简易信息
func BuildItemBrief(item model.Item, tags []model.Tag) Item {
	var tag []string = []string{}
	for _, i := range tags {
		tag = append(tag, i.TagName)
	}
	return Item{
		ID:      item.ID,
		Name:    item.Name,
		Brief:   item.Brief,
		Picture: item.Picture,
		Tag:     tag,
	}
}

// BuildItemList 序列化商品列表
func BuildItemList(items []model.Item, tags [][]model.Tag) []Item {
	var itemList []Item = []Item{}
	for index, data := range items {
		itemList = append(itemList, BuildItemBrief(data, tags[index]))
	}
	return itemList
}

// BuildItemListResponse 序列化商品列表响应
func BuildItemListResponse(item []model.Item, tags [][]model.Tag) Response {
	return Response{
		Data: BuildItemList(item, tags),
	}
}
