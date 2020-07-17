package service

import (
	"demo_order/dao"
	"demo_order/model"
	"fmt"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type ReqAddParam struct {
	OrderId  string `json:"order_id"`
	UserName string `json:"user_name"`
	Amount   string `json:"amount"`
	Status   string `json:"status"`
	FileUrl  string `json:"file_url"`
}

type RespGetParam struct {
	ID       int64   `json:"id"`
	OrderId  string  `json:"order_id"`
	UserName string  `json:"user_name"`
	Amount   float64 `json:"amount"`
	Status   string  `json:"status"`
	FileUrl  string  `json:"file_url"`
}

type OrderService struct {
	dao dao.OrderDao
}

// 初始化
func NewOrderService(dao dao.OrderDao) *OrderService {
	return &OrderService{dao: dao}
}

// 新增数据
func (s *OrderService) AddOrder(req ReqAddParam) error {

	amount, _ := strconv.ParseFloat(req.Amount, 64)
	order := &model.Order{
		ID:        0,
		OrderId:   req.OrderId,
		UserName:  req.UserName,
		Amount:    amount,
		Status:    req.Status,
		FileUrl:   req.FileUrl,
		CreatedAt: time.Now().Unix(),
	}
	err := s.dao.CreateOrder(order)
	if err != nil {
		return err
	}
	return nil
}

// 查询数据
func (s *OrderService) GetOrder(id int64) (*RespGetParam, error) {

	order, err := s.dao.QueryOrder(id)
	if err != nil {
		return nil, err
	}
	res := &RespGetParam{
		ID:       order.ID,
		OrderId:  order.OrderId,
		UserName: order.UserName,
		Amount:   order.Amount,
		Status:   order.Status,
		FileUrl:  order.FileUrl,
	}
	return res, nil
}

// 查询数据列表
func (s *OrderService) GetOrders(username string, page, limit int) ([]*RespGetParam, error) {

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	username = "%" + username + "%"
	res, err := s.dao.QueryOrders(username, page, limit)

	if err != nil {
		return nil, err
	}

	var list []*RespGetParam
	for _, v := range res {
		resp := &RespGetParam{
			ID:       v.ID,
			OrderId:  v.OrderId,
			UserName: v.UserName,
			Amount:   v.Amount,
			Status:   v.Status,
			FileUrl:  v.FileUrl,
		}
		list = append(list, resp)
	}
	return list, nil
}

// 更新数据
func (s *OrderService) UpdateOne(param RespGetParam) error {

	order := model.Order{
		ID:        param.ID,
		OrderId:   param.OrderId,
		UserName:  param.UserName,
		Amount:    param.Amount,
		Status:    param.Status,
		FileUrl:   param.FileUrl,
		CreatedAt: time.Now().Unix(),
	}
	err := s.dao.UpdateOrder(&order)
	if err != nil {
		return err
	}
	return nil
}

// 更新文件路径
func (s *OrderService) UpdateFileUrl(id int64, url string) error {

	err := s.dao.UpdateUrl(id, url)
	if err != nil {
		return err
	}
	return nil
}

// 下载文件
func (s *OrderService) DownloadFile(id int64) (string, error) {

	order, err := s.dao.QueryOrder(id)
	if err != nil {
		return "", err
	}
	url := order.FileUrl

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return "", err
	}
	// 文件命名
	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("download-%d.jpg", timeStamp)
	// 存放文件路径
	filePath := fmt.Sprintf("download/%s", fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return "", nil
	}

	defer file.Close()
	_, err = file.Write(body)
	if err != nil {
		return "", err
	}
	return filePath, nil
}

// 文件导出
func (s *OrderService) ExportOrder() error {

	var (
		file  *xlsx.File
		sheet *xlsx.Sheet
		row   *xlsx.Row
		cell  *xlsx.Cell
	)

	res, err := s.dao.QueryAll()
	if err != nil {
		fmt.Println("导出读取数据错误")
	} else {
		fmt.Println(res)
	}

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("order")
	if err != nil {
		fmt.Printf(err.Error())
	}
	if sheet != nil {
		for _, v := range res {
			row = sheet.AddRow()
			cell = row.AddCell()
			cell.Value = strconv.Itoa(int(v.ID))
			cell = row.AddCell()
			cell.Value = v.OrderId
			cell = row.AddCell()
			cell.Value = v.UserName
			cell = row.AddCell()
			cell.Value = strconv.FormatFloat(v.Amount, 'f', -1, 64)
			cell = row.AddCell()
			cell.Value = v.Status
			cell = row.AddCell()
			cell.Value = v.FileUrl
		}
	}

	err = file.Save("demo_order.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
	return nil
}
