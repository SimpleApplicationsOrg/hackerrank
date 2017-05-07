package main

import ("fmt"
	"bufio"
	"os")

type node struct {
	data rune
	next *node
}

type stack struct {
	head *node
}

func (s *stack) push(data rune) {
	newNode := &node{data:data}
	if s.head == nil {
		s.head = newNode
	} else {
		newNode.next = s.head
		s.head = newNode
	}
}

func (s *stack) pop() *rune {
	if s.head == nil { return nil }

	popNode := s.head
	s.head = s.head.next
	return &popNode.data
}

func (s *stack) peek() *rune {
	if s.head == nil { return nil }

	return &s.head.data
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanf(in, "%d\n", &n)

	for i := 0; i < n; i++ {
		var exp string
		fmt.Fscanf(in, "%s\n", &exp)
		if balanced(exp) { fmt.Println("YES") } else { fmt.Println("NO") }
	}

}

var (
	openning = map[rune]int{'{':1, '[':2, '(':3}
	closing = map[rune]int{'}':1, ']':2, ')':3}
)

func balanced(exp string) bool {
	stackRunes := stack{}

	for _, v := range exp {

		if  _, openContains := openning[v]; openContains {
			stackRunes.push(v)
		} else {
			if stackRunes.head == nil || closing[v] != openning[*stackRunes.peek()] {
				return false
			}
			stackRunes.pop()
		}

	}

	if stackRunes.head != nil { return false }

	return true
}