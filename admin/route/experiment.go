package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"turan/example-goWeb/admin/db"
)

type queryExperimentRequest struct {
	UserId int `gorm:"user_id" json:"userId" form:"userId" `
	ProjectId int `gorm:"project_id" json:"projectId" form:"projectId" `
	Limit int `json:"limit" form:"limit"`
	OffSet  int `json:"offset" form:"offset"`
}
type queryExperimentReply struct {
	Id int `gorm:"id" json:"id"`
	Pid int `gorm:"pid" json:"pid"`
	UserId int `gorm:"user_id" json:"userId"`
	ProjectId int `gorm:"project_id" json:"projectId"`
	Name string `gorm:"name" json:"name"`
	ExperimentType string `gorm:"experiment_type" json:"experimentType"`
	Status int `gorm:"status" json:"status"`
	ErrMsg string `gorm:"err_msg" json:"errMsg"`
	StartAt int64 `gorm:"start_at" json:"startAt"`
	DoneAt int64 `gorm:"done_at" json:"doneAt"`
	CreateAt int64 `gorm:"create_at" json:"createAt"`
	UpdateAt int64 `gorm:"update_at" json:"updateAt"`
	DeleteAt int64 `gorm:"delete_at"`
}

func queryExperiment(ctx *gin.Context)  {
	var req queryExperimentRequest
	err := ctx.ShouldBindJSON(&req)
	if err!=nil {
		ctx.JSON(http.StatusOK,gin.H{"bindErr":err.Error()})
		return
	}

	sql := `select id,pid,user_id,project_id,name,experiment_type,status,err_msg,start_at,done_at,create_at,update_at,delete_at from experiment where 
			user_id =? and project_id = ? limit ?,?
			`
	var reply []*queryExperimentReply
	err = db.GetTdb().Raw(sql, req.UserId, req.ProjectId, req.Limit, req.OffSet).Scan(&reply).Error
	if err != nil {
		ctx.JSON(http.StatusOK,gin.H{"db":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,reply)
}