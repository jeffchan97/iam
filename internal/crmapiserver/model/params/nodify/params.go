package nodify

// 映射对应erp客户的操作
// addCus 添加客户
// updCus 修改客户
// delCus 删除客户
// mergeCus 合并客户
// addLink 添加联系人
// updLink 修改联系人
// delLink 删除联系人
// mergeLink 合并联系人

var CrmTypeMapping map[string]NodifyParams = map[string]NodifyParams{
	"ec":   &EcNodifyParams{},
	"huke": &HukeNodifyParams{},
}

type NodifyParams interface {
	GetCrmType() string
	GetType() (string, bool)
	GetCusId() string
}
