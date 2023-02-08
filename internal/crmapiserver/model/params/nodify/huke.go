package nodify

import (
	"github.com/marmotedu/iam/internal/crmapiserver/model"
	"github.com/marmotedu/iam/internal/crmapiserver/store/crm"
	"strconv"
)

var hukeTypeMapping map[string]string = map[string]string{
	"1010": "addCus",
	"1001": "mergeCus",
	"1002": "addCus",
	"1005": "addLink",
	"1006": "updLink",
	"1007": "delLink",
}

type HukeNodifyParams struct {
	Type        int              `json:"type" binding:"required,number"`
	Cid         int              `json:"cid" binding:"required,number"`
	Eid         int              `json:"eid"`
	MergedCid   []int            `json:"mergedCid"`
	CidList     []int            `json:"cidList"`
	ContactList []map[string]int `json:"contactList"`
}

func (h *HukeNodifyParams) GetCrmType() string {
	return "huke"
}

func (h *HukeNodifyParams) GetType() (string, bool) {
	action := strconv.Itoa(h.Type)
	action, ok := hukeTypeMapping[action]
	return action, ok
}

func (h *HukeNodifyParams) GetCusId() string {
	return strconv.Itoa(h.Cid)
}

func (h *HukeNodifyParams) GetStore(applet *model.CrmApplet) crm.CrmStore {
	return &crm.HukeStore{
		&crm.Common{applet},
	}
}