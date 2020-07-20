package handler

import (
	"demo_order/model"
	"demo_order/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"time"
)

var OService *service.OrderService

// 添加数据
func AddOrder(c *gin.Context) {

	req := model.AddOrderReq{}
	if err := c.ShouldBind(&req); err != nil {
		response := model.ResponseData{
			Code:    -1,
			Message: "参数不正确",
		}
		c.JSON(200, response)
		return
	}

	err := OService.AddOrder(req)
	if err != nil {
		response := model.ResponseData{
			Code:    -1,
			Message: "添加失败",
		}
		c.JSON(200, response)
		return
	}

	response := model.ResponseData{
		Code:    1,
		Message: "添加成功",
	}
	c.JSON(200, response)
}

// 查询单条数据
func OrderDetail(c *gin.Context) {

	getReq := model.GetOrderReq{}
	if err := c.ShouldBind(&getReq); err != nil {
		response := model.ResponseData{
			Code:    -1,
			Message: "参数不正确",
		}
		c.JSON(200, response)
		return
	}

	result, err := OService.OrderDetail(getReq.ID)
	if err != nil {
		response := model.ResponseData{
			Code:    -1,
			Message: "未找到相关信息",
		}
		c.JSON(200, response)
		return
	}
	response := model.ResponseData{
		Code:    1,
		Message: result,
	}
	c.JSON(200, response)
}

// 更新数据
func UpdateOrder(c *gin.Context) {

	updateReq := model.Order{}
	if err := c.ShouldBind(&updateReq); err != nil {
		response := model.ResponseData{
			Code:    -1,
			Message: "参数不正确",
		}
		c.JSON(200, response)
		return
	}

	err := OService.UpdateOrder(updateReq)
	if err != nil {
		response := model.ResponseData{
			Code:    -1,
			Message: "更新失败",
		}
		c.JSON(200, response)
		return
	}
	response := model.ResponseData{
		Code:    1,
		Message: "更新成功",
	}
	c.JSON(200, response)
}

// 查询列表数据
func OrderList(c *gin.Context) {

	listReq := model.OrderListReq{}
	if err := c.ShouldBind(&listReq); err != nil {
		response := model.ResponseData{
			Code:    -1,
			Message: "参数不正确",
		}
		c.JSON(200, response)
		return
	}

	result, err := OService.OrderList(listReq.UserName, listReq.Page, listReq.Limit)
	if err != nil {
		response := model.ResponseData{
			Code:    -1,
			Message: "未找到相关信息",
		}
		c.JSON(200, response)
		return
	}
	response := model.ResponseData{
		Code:    1,
		Message: result,
	}
	c.JSON(200, response)
}

// 文件上传
func Upload(c *gin.Context) {

	getReq := model.GetOrderReq{}
	if err := c.ShouldBind(&getReq); err != nil {
		response := model.ResponseData{
			Code:    -1,
			Message: "参数不正确",
		}
		c.JSON(200, response)
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
	fileUrl := fmt.Sprintf("utils/upload/%d/%d/%d/%s", now.Year(), now.Month(), now.Day(), fileName)
	uerr := OService.UpdateFileUrl(getReq.ID, fileUrl)
	if uerr != nil {
		fmt.Println(uerr)
	}
	if err != nil {
		response := model.ResponseData{
			Code:    -1,
			Message: "文件上传失败",
		}
		c.JSON(200, response)
		return
	}

	response := model.ResponseData{
		Code:    1,
		Message: fmt.Sprintf("'%s' uploaded!", file.Filename),
	}
	c.JSON(200, response)
}

// 文件下载
// 通过传入id下载fileurl的文件
func Download(c *gin.Context) {

	getReq := model.GetOrderReq{}
	if err := c.ShouldBind(&getReq); err != nil {
		response := model.ResponseData{
			Code:    -1,
			Message: "参数不正确",
		}
		c.JSON(200, response)
		return
	}
	res, err := OService.DownloadFile(getReq.ID)
	if err != nil {
		response := model.ResponseData{
			Code:    -1,
			Message: "下载文件失败",
		}
		c.JSON(200, response)
		return
	}
	response := model.ResponseData{
		Code:    1,
		Message: fmt.Sprintf("'%s' download!", res),
	}
	c.JSON(200, response)
}

// 文件导出
func Export(c *gin.Context) {

	err := OService.ExportOrder()
	if err != nil {
		response := model.ResponseData{
			Code:    -1,
			Message: "导出文件失败",
		}
		c.JSON(200, response)
		return
	}
	response := model.ResponseData{
		Code:    1,
		Message: "导出文件成功",
	}
	c.JSON(200, response)
}
