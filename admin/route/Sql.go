package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"turan/example-goWeb/admin/db"
)

//查询"01"课程比"02"课程成绩高的学生的信息及课程分数

func sqlOne(c *gin.Context)  {
	course := c.Query("course")

	if len(course) != 2 {
		c.JSON(http.StatusBadRequest,gin.H{"err":"参数错误"})

	}else {

		type sqlRes struct {
			Id string `gorm:"column:s_id"`
			Name string `gorm:"column:s_name"`
			Birth string `gorm:"column:s_birth"`
			Sex string `gorm:"column:s_sex"`
			Chinses string `gorm:"column:chinses"`
			Math string `gorm:"column:math"`
		}

		var res []*sqlRes

		sqlStr := `
		(select s.*,a.s_score as chinses,b.s_score as math  from Score a ,Score b ,Student s
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
