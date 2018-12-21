package biz

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/nagexiucai/gosh/config"
	"log"
	"github.com/nagexiucai/gosh/model"
	"net/http"
	"github.com/nagexiucai/gosh/handler"
	"fmt"
)

// router实例
// db实例

type Biz struct {
	Router *mux.Router
	DB     *gorm.DB
}

// 配置初始化
func (b *Biz) Initialize(config *config.Config) {
	// 更多参数：https://github.com/go-sql-driver/mysql#parameters
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local&autocommit=true",
		config.DB.Username,
		config.DB.Password,
		config.DB.IP,
		config.DB.Port,
		config.DB.Name,
		config.DB.Charset)
	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatalln("Could not connect database!", err.Error())
	}

	b.DB = model.DBMigrate(db)
	b.Router = mux.NewRouter()
	b.setRouters()
}

// 包装router的GET方法
func (b *Biz) Get(path string, f func(writer http.ResponseWriter, request *http.Request)) {
	b.Router.HandleFunc(path, f).Methods("GET")
}

// 包装router的POST方法
func (b *Biz) Post(path string, f func(writer http.ResponseWriter, request *http.Request)) {
	b.Router.HandleFunc(path, f).Methods("POST")
}

// 包装router的PUT方法
func (b *Biz) Put(path string, f func(writer http.ResponseWriter, request *http.Request)) {
	b.Router.HandleFunc(path, f).Methods("PUT")
}

// 包装router的DELETE方法
func (b *Biz) Delete(path string, f func(writer http.ResponseWriter, request *http.Request)) {
	b.Router.HandleFunc(path, f).Methods("DELETE")
}

// 处理Employee相关的请求
func (b *Biz) GetAllEmployees(writer http.ResponseWriter, request *http.Request) {
	handler.GetAllEmployees(b.DB, writer, request)
}
func (b *Biz) CreateEmployee(writer http.ResponseWriter, request *http.Request) {
	handler.CreateEmployee(b.DB, writer, request)
}
func (b *Biz) GetEmployee(writer http.ResponseWriter, request *http.Request) {
	handler.GetEmployee(b.DB, writer, request)
}
func (b *Biz) UpdateEmployee(writer http.ResponseWriter, request *http.Request) {
	handler.UpdateEmployee(b.DB, writer, request)
}
func (b *Biz) DeleteEmployee(writer http.ResponseWriter, request *http.Request) {
	handler.DeleteEmployee(b.DB, writer, request)
}
func (b *Biz) DisableEmployee(writer http.ResponseWriter, request *http.Request) {
	handler.DisableEmployee(b.DB, writer, request)
}
func (b *Biz) EnableEmployee(writer http.ResponseWriter, request *http.Request) {
	handler.EnableEmployee(b.DB, writer, request)
}

// 设置所有必须的路由
func (b *Biz) setRouters() {
	// 项目处理相关的路由
	b.Get("/employees", b.GetAllEmployees)
	b.Post("/employees", b.CreateEmployee)
	b.Get("/employees/{title}", b.GetEmployee)
	b.Put("/employees/{title}", b.UpdateEmployee)
	b.Delete("/employees/{title}", b.DeleteEmployee)
	b.Put("/employees/{title}/disable", b.DisableEmployee)
	b.Put("/employees/{title}/enable", b.EnableEmployee)
}

// 照所配置的路由运行
func (b *Biz) Run(host string) {
	err := http.ListenAndServe(host, b.Router)
	log.Fatal(err)
}
