package model

type Magnets []Magnet

type Magnet struct {
	Link     string  //磁链
	Size     float64 //大小
	SizeText string  //大小(网站显示)
}

func (p Magnets) Len() int { return len(p) }

func (p Magnets) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type SortBySize struct{ Magnets }

// 根据影片的大小降序排序
func (p SortBySize) Less(i, j int) bool {
	return p.Magnets[i].Size > p.Magnets[j].Size
}
