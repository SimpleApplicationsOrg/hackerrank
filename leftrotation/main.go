package main
import ("fmt"
	"bufio"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	n := 0;
	k := 0;

	fmt.Fscanf(in, "%d %d\n", &n , &k)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(in, "%d", &arr[i])
	}

	rotated := rotateLeft(arr, k)

	for i, v := range rotated {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}

}

func rotateLeft(arr []int, k int) []int {

	rotated := make([]int, len(arr))

	if k < len(arr) {

		copy(rotated[0:len(arr)-k], arr[k:])
		copy(rotated[len(arr)-k:], arr[0:k])

	} else {
		copy(rotated, arr)
	}

	return rotated

}