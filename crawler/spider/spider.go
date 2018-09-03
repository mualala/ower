package spider

import (
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/henrylee2cn/pholcus/app/downloader/request"
	"github.com/henrylee2cn/pholcus/app/pipeline/collector/data"
	"yhm.com/crawler/runtime/status"
)

type (
	//蜘蛛集合
	Spider struct {
		Name         string // 蜘蛛名称（应保证唯一性）
		Desc         string // 蜘蛛描述
		EnableCookie bool   // 所有请求是否使用cookie记录
		RuleTree     *RuleTree
		status       int // 执行状态
	}
	//采集规则树
	RuleTree struct {
		Root  func(*Context)   // 根节点(执行入口)
		Trunk map[string]*Rule // 节点散列表(执行采集过程)
	}
	// 采集规则节点
	Rule struct {
		ItemFields []string                                           // 结果字段列表(选填，写上可保证字段顺序)
		ParseFunc  func(*Context)                                     // 内容解析函数
		AidFunc    func(*Context, map[string]interface{}) interface{} // 通用辅助函数
	}

	Context struct {
		spider   *Spider           // 规则
		Request  *request.Request  // 原始请求
		Response *http.Response    // 响应流，其中URL拷贝自*request.Request
		text     []byte            // 下载内容Body的字节流格式
		dom      *goquery.Document // 下载内容Body为html时，可转换为Dom的对象
		items    []data.DataCell   // 存放以文本形式输出的结果数据
		files    []data.FileCell   // 存放欲直接输出的文件("Name": string; "Body": io.ReadCloser)
		err      error             // 错误标记
		sync.Mutex
	}
)

// 添加自身到蜘蛛菜单
func (s Spider) Register() *Spider {
	s.status = status.STOPPED
	return Species.Add(&s)
}
