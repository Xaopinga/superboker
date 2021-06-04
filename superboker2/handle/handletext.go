package handle

import (
	"fmt"
	"strconv"
	"superboker2/model"
	"superboker2/query"
	"superboker2/serve"

	"github.com/gin-gonic/gin"
)

type HandleText struct {
	Sev serve.InSevText
}

func (h *HandleText) TextAddHandle(c *gin.Context) {
	e := query.NewEntity()
	m := new(model.Text)
	if err := c.BindJSON(m); err != nil {
		e.Msg = "参数不正确"
		c.JSON(500, e)
		return
	}
	if len(m.Textcontent) < 50 {
		c.String(200, "提交内容太少了")
		return
	}
	b, err := h.Sev.Add(m)
	if err != nil {
		c.JSON(500, e)
		return
	}
	if b {
		e.OK("ok")
		c.JSON(200, e)
		return
	}
	fmt.Println("添加文章数据发生未知错误err", err)
}

func (h *HandleText) TextDeletHandle(c *gin.Context) {
	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	entiy := query.NewEntity()
	b, err := h.Sev.Del(id)
	if err != nil {
		c.JSON(500, entiy)
		return
	}
	if b {
		entiy.OK("ok")
		c.JSON(200, entiy)
		return
	}
	fmt.Println("删除文章数据发生未知错误err", err)
}
func (h *HandleText) TextInfoHandle(c *gin.Context) {
	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	entiy := query.NewEntity()
	if m := h.Sev.ExistByTextID(id); m == nil {
		c.JSON(500, entiy)
	} else {
		entiy.OK(m)
		c.JSON(200, entiy)
	}
}
func (h *HandleText) TextEditHandle(c *gin.Context) {
	entiy := query.NewEntity()
	m := new(model.Text)
	if err := c.BindJSON(m); err != nil {
		c.JSON(200, entiy)
		return
	}
	fmt.Println(m)
	if _, err := h.Sev.Edit(m); err != nil {
		c.JSON(200, entiy)
		return
	}
	entiy.OK("Ok")
	c.JSON(200, entiy)
}
func (h *HandleText) TextListHandle(c *gin.Context) {
	p := new(query.Pagequery)
	entiy := query.NewEntity()
	err := c.BindJSON(p)
	fmt.Println(p)
	if err != nil {
		c.JSON(200, entiy)
		return
	}

	if p.Size > 100 || p.Size <= 0 {
		p.Size = 10
	}
	total, err := h.Sev.GetTotal()
	fmt.Println(total, err)
	if err != nil {
		c.JSON(200, entiy)
		return
	}
	ms, err := h.Sev.List(p) //获取到记录列表
	if err != nil {
		fmt.Print(err)
		c.JSON(200, entiy)
		return
	}
	pageTotal := 0
	if total%p.Size == 0 {
		pageTotal = int(total / p.Size)
	} else {
		pageTotal = int(total/p.Size) + 1
	}
	entiy.OK(ms)
	entiy.Total = total
	entiy.TotalPage = pageTotal
	c.JSON(200, entiy)
}
