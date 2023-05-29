package pagination

import (
	"math"
)


type Pagination struct{
    TotalItems float64
    PageSize float64
    CurrentPage float64
    BeforeAndAfterPage float64

}

func (p *Pagination)Offset() float64{
  return (p.CurrentPage -1) * p.PageSize
}

func (p *Pagination)TotalPages() int{
    return int(math.Ceil(p.TotalItems/p.PageSize))
}

func (p *Pagination)GeneratePageNumbers()[]int{

    var num []int
 
    totalPage :=  p.TotalPages()
 
    startPage := int(math.Max(2, p.CurrentPage - p.BeforeAndAfterPage)) //3 1-2
  
	endPage := int(math.Min(float64(totalPage), p.CurrentPage + p.BeforeAndAfterPage))
 
    for i := startPage; i <= endPage; i++ {
        if i !=  totalPage {
            num = append(num, i)
        }
             
     }

    return num
}