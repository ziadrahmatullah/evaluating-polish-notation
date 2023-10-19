package main
import (
    "fmt"
    "bufio"
    "log"
    "strings"
    "os"
    "strconv"
)

type node struct{
    val int
    next *node
}

type stack struct{
    top *node
}

func (S *stack) push(item int){
	newNode := &node{val: item, next: nil}
	if S.top == nil {
		S.top = newNode
	} else {
		newNode.next = S.top
		S.top = newNode
	}
}

func (S *stack) pop() int {
	if S.top != nil {
		popped := S.top.val
		S.top = S.top.next
		return popped
	}
	return 0
}

func (S *stack) calculatePN(input []string) int{
    input = reverseArray(input)
    for _, val := range input{
        if number, err := strconv.Atoi(val); err == nil{
            S.push(number)
        }else{
            switch val{
            case "+":
                S.push(S.pop() + S.pop())
            case "-":
                S.push(S.pop() - S.pop())
            case "/":
                S.push(S.pop() / S.pop())
            case "*":
                S.push(S.pop() * S.pop())
            }
        }
    }
    return S.pop()
}

func reverseArray(arr []string) []string{
    for i, j := 0, len(arr)-1; i<j; i, j = i+1, j-1 {
       arr[i], arr[j] = arr[j], arr[i]
    }
    return arr
 }

func readInput() []string {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.Fields(line)
}

func main(){
    S := stack{}

    fmt.Print("Input: ")
    polishNotation := readInput()
    fmt.Println("Output: ", S.calculatePN(polishNotation))
}