package main

import (
	"fmt"
	"strconv"
	"./DataStructure/Graph"
	"os"
	"encoding/csv"
	"bufio"
	"io"
	"log"
	"./Algorithm"
	"math"
	"strings"
)

var SIZE float64

func main(){
	SIZE = 0
	var path string
	for {
		fmt.Println("Enter Path")
		fmt.Scanln(&path)
		if path[:4]=="exit"{
			break
		}else{
			for{
				command,text := start()
				if len(command) == 3 && command[0]=="RUN"{
					if command[1]=="LinkedList" && command[2][:len(command[2])-1]=="Quick"{
						LinkedListQuick(path)
					}else if command[1]=="Matrix" && command[2][:len(command[2])-1]=="Quick"{
						MatrixQuick(path)
					}else if command[1]=="LinkedList" && command[2][:len(command[2])-1]=="Insertion"{
						LinkedListInsertion(path)
					}else if command[1]=="Matrix" && command[2][:len(command[2])-1]=="Insertion"{
						MatrixInsertion(path)
					}else if command[1]=="LinkedList" && command[2][:len(command[2])-1]=="Merge"{
						LinkedListMerge(path)
					}else if command[1]=="Matrix" && command[2][:len(command[2])-1]=="Merge"{
						MatrixMerge(path)
					}else if command[1]=="LinkedList" && command[2][:len(command[2])-1]=="Bubble"{
						LinkedListBubble(path)
					}else if command[1]=="Matrix" && command[2][:len(command[2])-1]=="Bubble"{
						MatrixBubble(path)
					}else{
						fmt.Println(len(command[2]),len(command[1]))
					}
				}else if len(command) == 5 &&  command[0]=="RUN"{
					if command[1]=="LinkedList" && command[2]=="Optimum" && command[3]=="Insertion"{
						num,err := strconv.Atoi(command[4][:len(command[4])-1])
						if err!=nil{
							fmt.Println("Please Send Number!,not Character")
						}else{
							LinkListOptimumInsertion(path,num)
						}
					}else if command[1]=="Matrix" && command[2]=="Optimum" && command[3]=="Insertion"{
						num,err := strconv.Atoi(command[4][:len(command[4])-1])
						if err!=nil{
							fmt.Println("Please Send Number!,not Character")
						}else{
							MatrixOptimumInsertion(path,num)
						}
					}else if command[1]=="LinkedList" && command[2]=="Optimum" && command[3]=="Bubble"{
						num,err := strconv.Atoi(command[4][:len(command[4])-1])
						if err!=nil{
							fmt.Println("Please Send Number!,not Character")
						}else{
							LinkListOptimumBubble(path,num)
						}
					}else if command[1]=="Matrix" && command[2]=="Optimum" && command[3]=="Bubble"{
						num,err := strconv.Atoi(command[4][:len(command[4])-1])
						if err!=nil{
							fmt.Println("Please Send Number!,not Character")
						}else{
							MatrixOptimumBubble(path,num)
						}
					}
				}else if len(command) == 1 && command[0][:len(command[0])-1]=="exit"{
					break
				}else{
					error(text)
				}
			}
		}
	}
}

func start() ([]string,string) {
	reader := bufio.NewReader(os.Stdin)
	//fmt.Println(reflect.TypeOf(reader))
	fmt.Print("root@ubuntu:~# ")
	text, _ := reader.ReadString('\n')
	text = string(text)
	com := strings.Split(text, " ")
	return com,text
}

func LinkedListQuick(path string){
	g,e := INIT(true,path)
	i:=0
	fmt.Println(e.Size())
	fmt.Println(22)
	for g.BFS(SIZE){
		fmt.Println(i)
		e.QuickSort()
		e.Clean(g)
		i++
	}
}

func MatrixQuick(path string){
	g,e := INIT(false,path)
	i:=0
	for g.BFS(SIZE){
		fmt.Println(i)
		e.QuickSort()
		e.Clean(g)
		i++
	}
}

func LinkedListInsertion(path string){
	g,e := INIT(true,path)
	i:=0
	fmt.Println(SIZE)
	for g.BFS(SIZE){
		fmt.Println(i)
		e.InsertionSort()
		e.Clean(g)
		i++
	}
}

func MatrixInsertion(path string){
	g,e := INIT(false,path)
	i:=0
	for g.BFS(SIZE){
		fmt.Println(i)
		e.InsertionSort()
		e.Clean(g)
		i++
	}
}

func LinkedListMerge(path string){
	g,e := INIT(true,path)
	i:=0
	for g.BFS(SIZE){
		fmt.Println(i)
		e.MergeSort()
		e.Clean(g)
		i++
	}
}

func MatrixMerge(path string){
	g,e := INIT(false,path)
	i:=0
	for g.BFS(SIZE){
		fmt.Println(i)
		e.MergeSort()
		e.Clean(g)
		i++
	}
}

func LinkedListBubble(path string){
	g,e := INIT(true,path)
	i:=0
	for g.BFS(SIZE){
		fmt.Println(i)
		e.BubbleSort()
		e.Clean(g)
		i++
	}
}

func MatrixBubble(path string){
	g,e := INIT(false,path)
	i:=0
	for g.BFS(SIZE){
		fmt.Println(i)
		e.BubbleSort()
		e.Clean(g)
		i++
	}
}

func LinkListOptimumInsertion(path string,N int){
	g,e := INIT(true,path)
	i:=0
	for g.BFS(SIZE){
		fmt.Println(i)
		e.OptimumInsertion(N)
		e.Clean(g)
		i++
	}
}

func MatrixOptimumInsertion(path string,N int){
	g,e := INIT(false,path)
	i:=0
	for g.BFS(SIZE){
		fmt.Println(i)
		e.OptimumInsertion(N)
		e.Clean(g)
		i++
	}
}

func LinkListOptimumBubble(path string,N int){
	g,e := INIT(true,path)
	i:=0
	for g.BFS(SIZE){
		fmt.Println(i)
		e.OptimumBubble(N)
		e.Clean(g)
		i++
	}
}

func MatrixOptimumBubble(path string,N int){
	g,e := INIT(false,path)
	i:=0
	for g.BFS(SIZE){
		fmt.Println(i)
		e.OptimumBubble(N)
		e.Clean(g)
		i++
	}
}

func INIT(b bool,path string) (*Graph.Graph,*Algorithm.Edges){
	if b{
		g,e := SaveListGraph(path)
		e = complete(g,e)
		return g,e
	}else{
		a,e := SaveMatrixGraph(path)
		g := SaveGraph(a)
		fmt.Println(e.Size())
		e = complete(g,e)
		return g,e
	}
}

func complete(graph *Graph.Graph,edges *Algorithm.Edges) *Algorithm.Edges{
	for i:=0;i<edges.Size();i++{
		x := edges.Get()[i].I
		y := edges.Get()[i].J
		z := graph.FindCountFriend(x,y)
		c := float64(z+1)/(math.Min(float64(graph.Get()[x].Size()-1),float64(graph.Get()[y].Size()-1)))
		edges.Get()[i].C = c
	}
	return edges
}

func SaveGraph(arr [3802930][3]int) *Graph.Graph{
	graph := Graph.New(150004)
	for i:=1;i<=arr[0][0];i++{
		graph.Connect(arr[i][0],arr[i][1])
	}
	return graph
}

func SaveMatrixGraph(path string) ([3802930][3]int,*Algorithm.Edges){
	edge := Algorithm.New()
	spars := [3802930][3]int{}
	csvFile, _ := os.Open(path)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	i := 1
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		in1, _ := strconv.Atoi(line[0])
		in2, _ := strconv.Atoi(line[1])
		SIZE = math.Max(float64(SIZE),float64(in1))
		spars[i][0] = in1
		spars[i][1] = in2
		edge.Add(in1,in2,float64(0))
		i++
	}
	spars[0][0] = i-1
	return spars,edge
}

func SaveListGraph(path string) (*Graph.Graph,*Algorithm.Edges){
	edge := Algorithm.New()
	graph := Graph.New(150004)
	csvFile, _ := os.Open(path)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		in1, _ := strconv.Atoi(line[0])
		in2, _ := strconv.Atoi(line[1])
		SIZE = math.Max(float64(SIZE),float64(in1))
		graph.Connect(in1,in2)
		edge.Add(in1,in2,float64(0))
	}
	return graph,edge
}

func error(text string){
	text = "'"+ text[:len(text)-1] + "' is not recognized as an internal or external command,\n operable program or batch file."

	fmt.Println(text)
}