package Graph

import  (
	"../bag"
	)

type Graph struct {
	nodes int
	edges []*bag.Bag
}

// Creates a new undirected graph.
func New(vertices int) *Graph {
	g := &Graph{
		nodes: vertices,
		edges: make([]*bag.Bag, vertices),
	}
	for i := 0; i < vertices; i++ {
		g.edges[i] = bag.New()
	}
	return g
}

func (g *Graph) Get() []*bag.Bag{
	return g.edges
}

func (g *Graph) FindCountFriend(i int,j int) int{
	test := make([]bool,len(g.edges))
	count := 0
	for k:=0;k<len(g.edges);k++{
		test[k] = false
	}
	for k:=0;k<g.edges[i].Size();k++{
		test[k] = true
	}
	for k:=0;k<g.edges[j].Size();k++{
		if test[k]{count++}
	}
	return count
}

func (g *Graph) BFS() bool{
	bl := make([]bool,g.nodes)
	for i:=0;i<g.nodes;i++{
		bl[i]=false
	}
	l := make([]int,1)
	l[0] = 1
	for len(l) > 0{
		v := l[len(l)-1]
		l = l[len(l)-1]
		if !bl[v]{
			bl[v] = true
			for i:=0;i<g.edges[v].Size();i++{
				if !bl[g.edges[v].Count(i)]{
					l = append(l,g.edges[i])
				}
			}
		}
	}
	for i:=0;i<g.nodes;i++{
		if !bl[i]{
			return false
		}
	}
	return true
}

// Returns the number of vertices in the graph.
func (g *Graph) Vertices() int {
	return g.nodes
}

// Connects two vertices of a graph (may be a loopback).
func (g *Graph) Connect(a, b int) {
	g.edges[a].Insert(b)
}

// Disconnects two vertices of a graph (may be a loopback).
func (g *Graph) Disconnect(a, b int) {
	g.edges[a].Remove(b)
	g.edges[b].Remove(a)
}


