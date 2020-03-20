package model

type Movie struct {
	Gid      string   //id
	Fanhao   string   //番号
	Link     string   //链接
	Duration string   //时长
	Img      string   //封面图
	Uc       string   //uc码
	Lang     string   //语言
	Title    string   //片名
	Actress  []string //演员表
	Date     string   //发售时间
	Series   string   //系列
	Category []string //类别
	Snapshot []string //截图
	Magnets  []Magnet //种子信息
}

type MovieJson struct {
	Fanhao   string   //番号
	Link     string   //链接
	Duration string   //时长
	Title    string   //片名
	Actress  []string //演员表
	Date     string   //发售时间
	Series   string   //系列
	Category []string //类别
	Snapshot []string //截图
	Magnets  []Magnet //种子信息
}