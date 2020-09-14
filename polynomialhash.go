package polynomialhash

import "math"

const (
	//TableSize holds the size of hashtable.
	TableSize       int     = 50 //size of the hashtable.
	hashCoefficient float64 = 3  //it can be any prime number 11,26,31...
)

//GetPolynomialHashCode...Gets the hash code for string type argument returns hashcode..
func GetPolynomialHashCode(str string, noofHashSlots int) int {
	var hashcode int
	var prime float64 = 100003
	x := 34
	for i := len(str) - 1; i >= 0; i-- {
		hashcode = int(math.Mod(float64(hashcode*x+int(str[i])), prime))

		//polynomialhash function==> (hash*x + str[i]) % prime
	}
	hashcode = universalHash(hashcode, prime, noofHashSlots)
	return hashcode
}

func universalHash(number int, primenumber float64, noofslots int) int {
	var first int = 51
	var second int = 2
	//as universalhash function==>((a*number + b)) chance of collision is 1/noofslots
	return int(math.Mod(math.Mod(float64(first*number+second), primenumber), float64(noofslots)))
}
