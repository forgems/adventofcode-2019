package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Recipe for creating an amount of product
type Recipe struct {
	Name     string
	Produces int64
	Requires map[string]int64
}

// ReadInput reads input data
func ReadInput(r io.Reader) map[string]Recipe {
	out := map[string]Recipe{}
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		r := Recipe{Requires: map[string]int64{}}
		line := scanner.Text()
		values := strings.Split(line, " => ")
		fmt.Sscanf(values[1], "%d %s", &r.Produces, &r.Name)
		ingridients := strings.Split(values[0], ", ")
		for i := range ingridients {
			var amount int64
			var name string
			fmt.Sscanf(ingridients[i], "%d %s", &amount, &name)
			r.Requires[name] = amount
		}
		out[r.Name] = r
	}

	return out
}

// Refinery type
type Refinery map[string]int64

// Produce amount of product, returns how much ore was used
func (r Refinery) Produce(product string, amount int64, recipes map[string]Recipe) (sum int64) {
	rec, ok := recipes[product]
	if !ok {
		return amount
	}
	producet := r[product]
	cnt := int64(0)
	for producet < amount {
		cnt++
		producet += rec.Produces
	}
	for ingridient, a := range rec.Requires {
		sum += r.Produce(ingridient, a*cnt, recipes)
	}
	r[product] = producet - amount
	return sum
}

func main() {
	data := ReadInput(os.Stdin)
	//fuel := int64(0)
	r := Refinery{}
	required := int64(1000000000000)
	//prev := r["ORE"]
	min, max := 1000000, 10000000
	mid := (min + max) / 2
	for max-min > 1 {
		mid := (max + min) / 2
		ore := r.Produce("FUEL", int64(mid), data)
		fmt.Println(mid, ore)
		if ore < required {
			min = mid
		} else if ore > required {
			max = mid
		} else if ore == required {
			break
		}
	}
	fmt.Println(mid)
}
