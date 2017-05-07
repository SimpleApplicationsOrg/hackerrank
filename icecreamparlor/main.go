package main
import ("fmt"
	"bufio"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	trips := 0
	fmt.Fscanf(in, "%d\n", &trips)

	for i := 0; i < trips; i++ {
		money := 0
		fmt.Fscanf(in, "%d\n", &money)
		numFlavors := 0
		fmt.Fscanf(in, "%d\n", &numFlavors)

		flavorsPrices := make([]int, numFlavors)
		for j := 0; j < numFlavors; j++ {
			fmt.Fscanf(in, "%d", &flavorsPrices[j])
		}
		fmt.Fscanf(in, "\n")

		fmt.Println(flavorsBought(money,flavorsPrices))
	}
}

func flavorsBought(money int, flavorsPrices []int) (int, int) {
	sortedPrices := make([]int, len(flavorsPrices))
	copy(sortedPrices,flavorsPrices)
	sort.Ints(sortedPrices)

	pricesUsed := make([]int, 2)
	for i := 0; i < len(sortedPrices); i++ {
		complement := money - sortedPrices[i];
		position := sort.SearchInts(sortedPrices, complement)
		if position >= 0 && position < len(sortedPrices) && sortedPrices[position] == complement {
			pricesUsed[0] = sortedPrices[i]
			pricesUsed[1] = sortedPrices[position]
			break
		}
	}

	flavorsBought := make([]int, 2)
	flavorsBought[0] = findFlavor(flavorsPrices, pricesUsed[0], -1)
	flavorsBought[1] = findFlavor(flavorsPrices, pricesUsed[1], flavorsBought[0])
	sort.Ints(flavorsBought)

	return flavorsBought[0] + 1, flavorsBought[1] + 1
}

func findFlavor(flavorsPrices []int, price int, excludeFlavor int) int {
	flavor := -1
	for i := 0; i < len(flavorsPrices); i++ {
		if flavorsPrices[i] == price && i != excludeFlavor {
			flavor = i
		}
	}
	return flavor
}