package serializer

import "singo/model"

// item 商品序列化器
type Item struct {
	ID        uint     `json:"id"`
	Name      string   `json:"name"`
	Brief     string   `json:"brief"`
	Picture   string   `json:"picture"`
	Tag       []string `json:"tag"`
	Input     []string `json:"input,omitempty"`
	Introduce string   `json:"introduce,omitempty"`
	Algorithm string   `json:"algorithm,omitempty"`
	Code      string   `json:"code,omitempty"`
	Time      int     `json:"time,omitempty"`
	Memory    int     `json:"memory,omitempty"`
	OutputImg bool     `json:"output_img,omitempty"`
}

// BuildItem 序列化商品
func BuildItem(item model.Item, tags []model.Tag, inputs []model.Input) Item {
	var tag []string = []string{}
	var input []string = []string{}
	for _, i := range tags {
		tag = append(tag, i.TagName)
	}

	for _, i := range inputs {
		input = append(tag, i.Input)
	}
	return Item{
		ID:        item.ID,
		Name:      item.Name,
		Brief:     item.Brief,
		Picture:   item.Picture,
		Tag:       tag,
		Introduce: item.Introduce,
		Algorithm: item.Algorithm,
		Input:     input,
	}
}

// BuildItemResponse 序列化商品响应
func BuildItemResponse(item model.Item, tags []model.Tag, inputs []model.Input) Response {
	return Response{
		Data: BuildItem(item, tags, inputs),
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

// BuildItemFull 序列化商品完整信息
func BuildItemFull(item model.Item, tags []model.Tag, inputs []model.Input) Item {
	result := BuildItem(item, tags, inputs)
	result.Code = item.Code
	result.Time = item.Time
	result.Memory = item.Memory
	result.OutputImg = item.OutputImg
	return result
}

// BuildItemFullResponse 序列化商品完整信息响应
func BuildItemFullResponse(item model.Item, tags []model.Tag, inputs []model.Input) Response {
	return Response{
		Data: BuildItemFull(item, tags, inputs),
	}
}
