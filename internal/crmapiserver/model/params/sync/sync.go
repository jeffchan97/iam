package sync

type SyncParams struct {
	CusType string `json:"cusType" binding:"required"`
	CrmType string `json:"crmType" binding:"required"`
	Action  string `json:"action" binding:"required"`
}