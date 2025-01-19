package loudec

import (
	"github.com/quagmt/udecimal"
	"github.com/samber/lo"
)

// LoMax returns the maximum value in the collection.
func LoMax(collection []udecimal.Decimal) udecimal.Decimal {
	return lo.MaxBy(collection, func(a, b udecimal.Decimal) bool {
		return a.GreaterThan(b)
	})
}

// LoMin returns the minimum value in the collection.
func LoMin(collection []udecimal.Decimal) udecimal.Decimal {
	return lo.MinBy(collection, func(a, b udecimal.Decimal) bool {
		return a.LessThan(b)
	})
}

// LoSum returns the sum of all values in the collection.
func LoSum(collection []udecimal.Decimal) udecimal.Decimal {
	sum := udecimal.Decimal{}
	lo.ForEach(collection, func(item udecimal.Decimal, index int) {
		sum = sum.Add(item)
	})
	return sum
}

// LoMean returns the mean value of the collection.
func LoMean(collection []udecimal.Decimal) (udecimal.Decimal, error) {
	return LoSum(collection).Div64(uint64(len(collection)))
}
