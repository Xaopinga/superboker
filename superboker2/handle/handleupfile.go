package handle

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HandleUpfile struct {
}

func (h *HandleUpfile) Upfile(c *gin.Context) {
	// r.ParseMultipartForm(1024)
	// file := r.MultipartForm.File["file"][0]
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
		c.String(500, "解析文件失败")
		return
	}
	name := file.Filename
	arr := strings.Split(name, `.`)
	switch arr[len(arr)-1] {
	case "jpg":
	case "img":
	case "png":
	case "ico":
	default:
		c.String(500, "格式不正确")
		return
	}
	fmt.Println(file.Filename)

	body, err := file.Open()
	if err == nil {
		byt, err := ioutil.ReadAll(body)
		if err == nil {
			//写入文件
			path := `data/img/` + uuid.New().String() + "." + arr[len(arr)-1]
			os.MkdirAll(`data/img/`, 0766)
			err = ioutil.WriteFile(path, byt, 0666)
			if err == nil {
				c.String(200, path)
			} else {
				c.String(500, "失败")
			}

		}
	}
}
