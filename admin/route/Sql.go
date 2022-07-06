package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"turan/example-goWeb/admin/db"
)


type sqlRes struct {
	Id string `gorm:"column:s_id"`
	Name string `gorm:"column:s_name"`
	Birth string `gorm:"column:s_birth"`
	Sex string `gorm:"column:s_sex"`
	Chinese string `gorm:"column:chinese"`
	Math string `gorm:"column:math"`
}

//查询"01"课程比"02"课程成绩高的学生的信息及课程分数
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

//-- 2、查询"01"课程比"02"课程成绩低的学生的信息及课程分数
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


type scoreFull struct {
	Id string `gorm:"column:s_id"`
	Name string `gorm:"column:s_name"`
	Birth string `gorm:"column:s_birth"`
	Sex string `gorm:"column:s_sex"`
	Full string `gorm:"column:full"`
}

//# 返回所有学生的总分成绩和信息
func getScoreFull(c *gin.Context)  {

	offSet := c.DefaultQuery("offset","0")
	limit := c.DefaultQuery("limit","2")
	sql :=`
	select s4.*,s.s_score+s2.s_score+s3.s_score as full from Score s,Score s2 ,Score s3 ,Student s4 where s.c_id ='01' and  s2.c_id ='02' and
                                                        s3.c_id='03'and s.s_id = s4.s_id =s2.s_id =s3.s_id order by full limit ?,?`
	var scoreFull []*scoreFull
	err := db.GetDB().Raw(sql, offSet, limit).Scan(&scoreFull).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"err":err.Error()})
	}else {
		c.JSON(http.StatusOK,scoreFull)
	}
}



//查询平均成绩 大于等于60分的同学的学生编号和学生姓名和平均成绩
func sqlThree(c *gin.Context)  {
	type sqlThreeRes struct {
		Id string `gorm:"column:s_id"`
		Name string `gorm:"column:s_name"`
		Avg string `gorm:"column:avg"`
	}
	var sqlThree []*sqlThreeRes
	sql := `

		SELECT S2.s_id,S.s_name, ROUND(avg(S2.s_score),1) as avg FROM Student S ,Score S2 
		WHERE S2.s_id=S.s_id GROUP BY S2.s_id having avg >=60;
		`
	err := db.GetDB().Raw(sql).Scan(&sqlThree).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"err":err.Error()})
	}else {
		c.JSON(http.StatusOK,sqlThree)
	}
}

func sqlFour(c *gin.Context)  {
	type sqlFourRes struct {
		Id string `gorm:"column:s_id"`
		Name string `gorm:"column:s_name"`
		Avg string `gorm:"column:avg"`
	}
	var fourRes []*sqlFourRes

	sql :=`

select s2.s_id,s.s_name, round(avg(s2.s_score),2) as avg  from Student s , Score s2
where s.s_id =s2.s_id group by s2.s_id , s.s_name having avg<=60
union select s_id ,s_name,0 as avg from Student
where s_id not in (select distinct  s_id from Score  ) `

	err := db.GetDB().Raw(sql).Scan(&fourRes).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"err":err.Error()})
	}else {
		c.JSON(http.StatusOK,fourRes)
	}
}

func sqlFifth(c *gin.Context)  {
	type fifthRes struct {
		Id string `gorm:"column:s_id"`
		Name string `gorm:"column:s_name"`
		Course string `gorm:"column:course"`
		Total string `gorm:"column:total"`
	}
	var fifth []*fifthRes
	sql:=`
			select a.s_id,b.s_name, count(a.c_id) as course ,sum(a.s_score) as total
			from Score a join Student b on a.s_id=b.s_id group by a.s_id ,b.s_name
			`
	err := db.GetDB().Raw(sql).Scan(&fifth).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"err":err.Error()})
	}else {
		c.JSON(http.StatusOK,fifth)
	}
}

func sqlSeventh(c *gin.Context)  {
	type sqlSeventhRes struct {
		Id string `gorm:"column:s_id"`
		Name string `gorm:"column:s_name"`
		Birth string `gorm:"column:s_birth"`
		Sex string `gorm:"column:s_sex"`
	}
	var sqlSeventh []*sqlSeventhRes
	sql:=`
select * from Student
where s_id in  (select s_id from Score where  c_id =
( select c_id from  Course where t_id =(select t_id from Teacher where t_name = '张三')) group by s_id)`

	err := db.GetDB().Raw(sql).Scan(&sqlSeventh).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"err":err.Error()})
	}else {
		c.JSON(http.StatusOK,sqlSeventh)
	}

}