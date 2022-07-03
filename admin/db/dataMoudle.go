package db

type  Course struct {
	Cid string `gorm:"column:c_id"`
	Name string	`gorm:"column:c_name"`
	Tid string `gorm:"column:t_id"`
}

type Score struct {
	Sid string `gorm:"column:s_id"`
	Cid string	`gorm:"column:c_id"`
	Sscore int	`gorm:"column:s_score"`
}

type Student struct {
	Sid string `gorm:"column:s_id"`
	Sname string `gorm:"column:s_name"`
	Sbirth string `gorm:"column:s_birth"`
	Ssex string`gorm:"column:s_sex"`
}

type Teacher struct {
	Tid string `gorm:"column:t_id"`
	Tname string	`gorm:"column:t_name"`

}