package handle

import (
	"fmt"
	"strconv"
	"superboker2/model"
	"superboker2/query"
	"superboker2/serve"

	"github.com/gin-gonic/gin"
)

type HandleCategory struct {
	Sev serve.InSevCategory
}

func (h *HandleCategory) CategoryAddHandle(c *gin.Context) {
	entiy := query.NewEntity()
	m := new(model.Category)

	if err := c.BindJSON(m); err != nil {
		c.JSON(200, entiy)
		return
	}
	fmt.Println(m)
	if _, err := h.Sev.Edit(m); err != nil {
		c.JSON(200, entiy)
		return
	}
	if _, err := h.Sev.Add(m); err != nil {
		c.JSON(200, entiy)
		return
	}
	entiy.OK("true")
	c.JSON(200, entiy)
}
func (h *HandleCategory) CategoryInfoHandle(c *gin.Context) {
	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	entiy := query.NewEntity()
	m := h.Sev.ExistByCategoryID(id) //根据Id找记录
	if m == nil {                    //找不到
		c.JSON(200, entiy)
		return
	}
	entiy.OK(m)
	c.JSON(200, entiy)
}
func (h *HandleCategory) CategoryListHandle(c *gin.Context) {
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
func (h *HandleCategory) CategoryEditHandle(c *gin.Context) {
	entiy := query.NewEntity()
	m := new(model.Category)
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
func (h *HandleCategory) CategoryDelHandle(c *gin.Context) {
	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	e := query.NewEntity()
	if _, err := h.Sev.Delete(id); err != nil {
		c.JSON(200, e)
		return
	}
	e.OK("")
	c.JSON(200, e)
}
func (h *HandleCategory) CategoryListSonHandle(c *gin.Context) {
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
	total, err := h.Sev.ListSonTotal()
	fmt.Println(total, err)
	if err != nil {
		c.JSON(200, entiy)
		return
	}
	ms, err := h.Sev.ListSon(p) //获取到记录列表
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
