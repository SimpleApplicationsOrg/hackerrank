package main
import ("fmt"
	"bufio"
	"os"
	"strings"
	"strconv")

const EDGE_DISTANCE = 6

type graph struct {
	Nodes []node
}

type node struct {
	Neighbors []int
}

func (g *graph) getNode(id int) *node {
	if g.Nodes[id].Neighbors == nil {
		g.Nodes[id] = node{make([]int,0)}
	}
	return &g.Nodes[id]
}

func (g *graph) addEdge(first int, second int) {
	g.getNode(first).Neighbors = append(g.getNode(first).Neighbors, second)
	g.getNode(second).Neighbors = append(g.getNode(second).Neighbors, first)
}

func (g *graph) shortestReach(startId int) []int {
	queue := make([]int,0)
	queue = append(queue,startId)

	distances := make([]int,len(g.Nodes))
	for indDist,_ := range distances { distances[indDist] = -1 }
	distances[startId] = 0

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if g.Nodes[node].Neighbors != nil {
			for _,neighbor := range g.Nodes[node].Neighbors {
				if distances[neighbor] == -1 {
					distances[neighbor] = distances[node] + EDGE_DISTANCE
					queue = append(queue, neighbor)
				}
			}
		}
	}
	return distances
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	nQueries,_ := strconv.Atoi(scanner.Text())
	for i := 0; i < nQueries; i++ {
		scanner.Scan()
		graphParams := strings.Split(scanner.Text()," ")
		nNodes,_ := strconv.Atoi(graphParams[0])
		g := &graph{make([]node,nNodes)}

		m,_ := strconv.Atoi(graphParams[1])
		for j := 0; j < m; j++ {
			scanner.Scan()
			edgeParams := strings.Split(scanner.Text()," ")
			u,_ := strconv.Atoi(edgeParams[0])
			v,_ := strconv.Atoi(edgeParams[1])

			g.addEdge(u-1,v-1)
		}

		scanner.Scan()
		s := strings.Trim(scanner.Text()," ")
		startId,_ := strconv.Atoi(s)
		startId -= 1

		distances := g.shortestReach(startId)
		for j := 0; j < len(distances); j++ {
			if j != startId {
				fmt.Print(distances[j])
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}