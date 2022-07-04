package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"turan/example-goWeb/admin/db"
)

//查询"01"课程比"02"课程成绩高的学生的信息及课程分数
type sqlRes struct {
	Id string `gorm:"column:s_id"`
	Name string `gorm:"column:s_name"`
	Birth string `gorm:"column:s_birth"`
	Sex string `gorm:"column:s_sex"`
	Chinese string `gorm:"column:chinese"`
	Math string `gorm:"column:math"`
}


func sqlOne(c *gin.Context)  {
	course := c.Query("course")

	if len(course) != 2 {
		c.JSON(http.StatusBadRequest,gin.H{"err":"参数错误"})

	}else {



		var res []*sqlRes

		sqlStr := `
		(select s.*,a.s_score as chinese,b.s_score as math  from Score a ,Score b ,Student s
		where a.c_id=? and b.c_id ='02' and a.s_id = b.s_id and a.s_score >b.s_score
		and a.s_id = s.s_id)`

		err := db.GetDB().Raw(sqlStr,course).Scan(&res).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError,"err")
		}else {
			c.JSON(http.StatusOK,gin.H{"data":res})
		}



	}

}

func sqlTwo(c *gin.Context)  {

	sql := `
		select S4.*, s.s_score as chinese,s2.s_score as math  
	from 
		(select * from Score union all select '07' as s_id, '01'as c_id, '0' as s_score union all  select  '06' as s_id,'02'as c_id,'0' as s_score) s ,
		(select * from Score union all select '07' as s_id, '01' as c_id, '0' as s_score)  s2 , Student S4
	where s.c_id ='01' and  s2.c_id ='02' and s.s_id =s2.s_id and s.s_score < s2.s_score and s.s_id = S4.s_id`
	var res []*sqlRes
	err := db.GetDB().Raw(sql).Scan(&res).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"err":err.Error()})
	}else {
		c.JSON(http.StatusOK,res)
	}
}