package service

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"singo/model"
	"singo/serializer"
	"singo/util"

	"github.com/jiang4869/go-judger"
)

// JudgeService 评测服务
type JudgeService struct {
	ID    uint   `json:"item_id"`
	Input string `json:"input"`
}

// Judge 评测
func (service *JudgeService) Judge() serializer.Response {

    item, err := model.GetItem(service.ID)

    if err != nil{
        return serializer.ParamErr("获取商品失败", err)
    }

    ex, err := os.Executable()
    if err != nil {
        return serializer.ServiceErr("", err)
    }
    exPath := filepath.Dir(ex)

    dirName := path.Join(exPath, "tmp", util.RandStringRunes(10))

    os.MkdirAll(dirName, os.ModePerm)

    codeFile := path.Join(dirName, "main.cpp")
    binFile := path.Join(dirName, "main")
    inputFile := path.Join(dirName, "input.txt")
    outputFile := path.Join(dirName, "output.txt")
    errorFile := path.Join(dirName, "error.txt")
    judgeFile := path.Join(dirName, "judge.log")
    svgFile := path.Join(dirName, "img.svg")
    imgFile := path.Join(dirName, "img.png")

    util.WriteFile(codeFile, item.Code)
    util.WriteFile(inputFile, service.Input)

    judger.ExecShell(fmt.Sprintf("g++ -o %s '%s'", binFile, codeFile))
    
	r, err := judger.Run(item.Time, item.Time, item.Memory, item.Memory, -1, -1, 0, 0, 0, []string{}, []string{}, binFile, inputFile, outputFile, errorFile, judgeFile, "c_cpp_file_io")
	if err != nil {
		return serializer.JudgeErr("", err)
	}
	b, _ := json.Marshal(r)
	fmt.Println(string(b))

    answer, err := util.ReadFile(outputFile)
    if err != nil{
		return serializer.JudgeErr("", err)
    }
    
    judger.ExecShell(fmt.Sprintf("cairosvg %s -o %s", svgFile, imgFile))

    img, err := util.ReadFile(imgFile)
    if err != nil{
        fmt.Println(err.Error())
    }

	return serializer.BuildOutputResponse(answer, img)
}
