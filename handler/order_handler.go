package handler

import (
	"demo_order/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var OService *service.OrderService

type NewReq struct {
	OrderId  string `form:"order_id"`
	UserName string `form:"user_name"`
	Amount   string `form:"amount"`
	Status   string `form:"status"`
	FileUrl  string `form:"file_url"`
}

type GetReq struct {
	ID int64 `form:"id"`
}

type UpdateReq struct {
	ID      int64   `form:"id"`
	Amount  float64 `form:"amount"`
	Status  string  `form:"status"`
	FileUrl string  `form:"file_url"`
}

type ListReq struct {
	UserName string `form:"user_name"`
	Page     int    `form:"page"`
	Limit    int    `form:"limit"`
}

// 添加数据
func NewOne(c *gin.Context) {

	newReq := NewReq{}
	if err := c.ShouldBind(&newReq); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数不正确",
		})
		return
	}

	param := service.ReqAddParam{
		OrderId:  newReq.OrderId,
		UserName: newReq.UserName,
		Amount:   newReq.Amount,
		Status:   newReq.Status,
		FileUrl:  newReq.FileUrl,
	}

	err := OService.AddOrder(param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "添加失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "添加成功",
	})
}

// 查询单条数据
func GetOne(c *gin.Context) {

	getReq := GetReq{}
	if err := c.ShouldBind(&getReq); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数不正确",
		})
		return
	}

	result, err := OService.GetOrder(getReq.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "未找到相关信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}

// 更新数据
func UpdateOne(c *gin.Context) {

	updateReq := UpdateReq{}
	if err := c.ShouldBind(&updateReq); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数不正确",
		})
		return
	}

	param := service.RespGetParam{
		ID:      updateReq.ID,
		Amount:  updateReq.Amount,
		Status:  updateReq.Status,
		FileUrl: updateReq.FileUrl,
	}

	err := OService.UpdateOne(param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "更新失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "更新成功",
	})
}

// 查询列表数据
func GetList(c *gin.Context) {

	listReq := ListReq{}
	if err := c.ShouldBind(&listReq); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数不正确",
		})
		return
	}

	result, err := OService.GetOrders(listReq.UserName, listReq.Page, listReq.Limit)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "未找到相关信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}

// 文件上传
func Upload(c *gin.Context) {

	getReq := GetReq{}
	if err := c.ShouldBind(&getReq); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数不正确",
		})
		return
	}
	file, _ := c.FormFile("file")

	now := time.Now()
	// 文件夹路径
	fileDir := fmt.Sprintf("upload/%d/%d/%d", now.Year(), now.Month(), now.Day())
	// ModePerm是0777，这样拥有该文件夹路径的执行权限
	err := os.MkdirAll(fileDir, os.ModePerm)
	// 文件路径
	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, file.Filename)
	filePathStr := filepath.Join(fileDir, fileName)

	err = c.SaveUploadedFile(file, filePathStr)
	fileUrl := fmt.Sprintf("upload/%d/%d/%d/%s", now.Year(), now.Month(), now.Day(), fileName)
	uerr := OService.UpdateFileUrl(getReq.ID, fileUrl)
	if uerr != nil {
		fmt.Println(uerr)
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "上传文件失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": fmt.Sprintf("'%s' uploaded!", file.Filename),
	})
}

// 文件下载
// 通过传入id下载fileurl的文件
func Download(c *gin.Context) {

	getReq := GetReq{}
	if err := c.ShouldBind(&getReq); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数不正确",
		})
		return
	}
	res, err := OService.DownloadFile(getReq.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "下载文件失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": fmt.Sprintf("'%s' download!", res),
	})

}

// 文件导出
func Export(c *gin.Context) {

	err := OService.ExportOrder()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "导出文件失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": "导出文件成功",
	})
}
