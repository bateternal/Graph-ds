package Algorithm

import (
	"../DataStructure/Graph"
	"math"
)

type Edge struct {
	I int
	J int
	C float64
}

type Edges struct {
	data []*Edge
}

func (e *Edge) Get() (int,int,float64){
	return e.I,e.J,e.C
}

func New() *Edges{
	e := &Edges{
		data : make([]*Edge,0),
	}
	return e
}

func (e *Edges)Get() []*Edge{
	return e.data
}

func (e *Edges) Size() int{
	return len(e.data)
}

func (e *Edges) Add(i int,j int, C float64){
	edge := &Edge{i,j,C}
	e.data = append(e.data,edge)
}

func (e *Edges) Clean(graph *Graph.Graph){
	i,j,c := e.data[0].Get()
	graph.Disconnect(i,j)
	e.Remove(i,j,c)
	s1 := make([]int,0)
	s2 := make([]int,0)
	s3 := make([]float64,0)
	b := make([]bool,1500000)

	for k:=0;k<graph.Get()[i].Size();k++ {
		z := graph.FindCountFriend(graph.Get()[i].Get()[k], i)
		c := float64(z+1) / (math.Min(float64(graph.Get()[graph.Get()[i].Get()[k]].Size()-1), float64(graph.Get()[i].Size()-1)))
		s1 = append(s1,graph.Get()[i].Get()[k])
		s2 = append(s2,i)
		s3 =append(s3,c)
	}
	for k:=0;k<graph.Get()[j].Size();k++ {
		z := graph.FindCountFriend(graph.Get()[j].Get()[k], i)
		c := float64(z+1) / (math.Min(float64(graph.Get()[graph.Get()[j].Get()[k]].Size()-1), float64(graph.Get()[j].Size()-1)))
		for f := 0; f < e.Size(); f++ {
			if j == e.Get()[i].I && graph.Get()[j].Get()[k] == e.Get()[j].J || j == e.Get()[i].J && graph.Get()[j].Get()[k] == e.Get()[j].I{
				e.Get()[i].C = c
			}

		}
	}
}

func (e *Edges) Remove(i int,j int,C float64){
	for k:=0;k<len(e.data);k++{
		if e.data[k].I==i && e.data[k].J==j && e.data[k].C==C{
			e.Delete(k)
		}
		if e.data[k].I==j && e.data[k].J==i && e.data[k].C==C{
			e.Delete(k)
		}
	}
}

func (e *Edges) Delete(index int) {
	e.data[index] = e.data[len(e.data)-1]
	e.data = e.data[:len(e.data)-1]
}

func (e *Edges) InsertionSort(){
	for i:=1;i<len(e.data);i++{
		j := i-1
		key := e.data[i]
		for j>=0 && e.data[j].C > key.C{
			e.data[j+1] = e.data[j]
			j -= 1
		}
		e.data[j+1] = key
	}
}

func (e *Edges) BubbleSort(){
	changed := true
	for changed{
		changed = false
		for i:=0;i<len(e.data)-1;i++{
			if e.data[i].C > e.data[i+1].C{
				e.data[i],e.data[i+1] = e.data[i+1],e.data[i]
				changed = true
			}
		}
	}
}

func (e *Edges) MergeSort(){
	if len(e.data)>1{
		mid := int(len(e.data)/2)
		leftHalf := e.data[:mid]
		rightHalf := e.data[mid:]
		eLeft := new(Edges)
		eLeft.data = leftHalf
		eRight := new(Edges)
		eRight.data = rightHalf
		eLeft.MergeSort()
		eRight.MergeSort()
		e.data = append(eLeft.data,eRight.data...)

		i,j,k := 0,0,0

		for i<len(leftHalf) && j<len(rightHalf){
			if leftHalf[i].C < rightHalf[j].C{
				e.data[k] = leftHalf[i]
				i++
			}else{
				e.data[k] = rightHalf[j]
				j++
			}
			k++
		}
		for i<len(leftHalf){
			e.data[k]=leftHalf[i]
			i++
			k++
		}
		for j<len(rightHalf){
			e.data[k]=rightHalf[j]
			j++
			k++
		}
	}
}

func (e *Edges) QuickSort(){
	less := make([]*Edge,0)
	pivotlist := make([]*Edge,0)
	more := make([]*Edge,0)
	if len(e.data)<=1{return}else{
		pivot := e.data[0]
		for i:=0;i<len(e.data);i++{
			if e.data[i].C < pivot.C{
				less = append(less,e.data[i])
			}else if e.data[i].C > pivot.C{
				more = append(more,e.data[i])
			}else{
				pivotlist = append(pivotlist,e.data[i])
			}
		}
		eLess := new(Edges)
		eLess.data = less
		eLess.QuickSort()
		eMore := new(Edges)
		eMore.data = more
		eMore.QuickSort()
		e.data = append(eLess.data,pivotlist...)
		e.data = append(e.data,eMore.data...)
	}
}

func (e *Edges) OptimumInsertion(N int){
	less := make([]*Edge,0)
	pivotList := make([]*Edge,0)
	more := make([]*Edge,0)
	if len(e.data)<=1{return}else{
		pivot := e.data[0]
		for i:=0;i<len(e.data);i++{
			if e.data[i].C < pivot.C{
				less = append(less,e.data[i])
			}else if e.data[i].C > pivot.C{
				more = append(more,e.data[i])
			}else{
				pivotList = append(pivotList,e.data[i])
			}
		}
		eLess := new(Edges)
		eLess.data = less
		if len(less)>N {eLess.OptimumInsertion(N)}else{eLess.InsertionSort()}
		eMore := new(Edges)
		eMore.data = more
		if len(more)>N {eMore.OptimumInsertion(N)}else{eMore.InsertionSort()}
		e.data = append(eLess.data, pivotList...)
		e.data = append(e.data,eMore.data...)
	}
}

func (e *Edges) OptimumBubble(N int){
	less := make([]*Edge,0)
	pivotlist := make([]*Edge,0)
	more := make([]*Edge,0)
	if len(e.data)<=1{return}else{
		pivot := e.data[0]
		for i:=0;i<len(e.data);i++{
			if e.data[i].C < pivot.C{
				less = append(less,e.data[i])
			}else if e.data[i].C > pivot.C{
				more = append(more,e.data[i])
			}else{
				pivotlist = append(pivotlist,e.data[i])
			}
		}
		eLess := new(Edges)
		eLess.data = less
		if len(less)>N {eLess.OptimumInsertion(N)}else{eLess.BubbleSort()}
		eMore := new(Edges)
		eMore.data = more
		if len(more)>N {eMore.OptimumInsertion(N)}else{eMore.BubbleSort()}
		e.data = append(eLess.data,pivotlist...)
		e.data = append(e.data,eMore.data...)
	}
}