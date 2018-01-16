package ShardReduce

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type X struct {
	a int
	b int
}

func TestReduce(t *testing.T) {
	var y []*interface{} = FakeMap()

	r := NewShardReduce(y)

	dones := r.Map(func(input X) X {
		fmt.Printf("%#v\n", input)
		y := input
		return y
	}).Filter(func(input *interface{}) bool {
		rand.Seed(time.Now().UnixNano())
		rnd := rand.Intn(100)
		fmt.Printf("\n---\n")
		if rnd > 50 {
			return true
		} else {
			return false
		}
	}).Filter(func(input *interface{}) bool {
		rand.Seed(time.Now().UnixNano())
		rnd := rand.Intn(100)
		fmt.Printf("\n---\n")
		if rnd > 50 {
			return true
		} else {
			return false
		}
	}).Map(func(input X) X {
		y := input
		fmt.Printf("%#v\n", input)
		return y
	}).Reduce(func(start X, n X) X {
		var brandNew X = X{start.a + n.a, start.b + n.b}
		return brandNew
		/*if n == nil {
			return start
		}
		var st X = (*start).(X)
		var ne X = (*n).(X)
		var brandNew interface{} = X{st.a + ne.a, st.b + ne.b}
		return &brandNew*/

	})

	fmt.Printf("---\nresult: %#v", *dones)

}

func FakeMap() []*interface{} {
	x := make([]*interface{}, 0)
	for y := 0; y < 25; y++ {
		var q interface{}
		q = X{y, y + 1}
		x = append(x, &q)
	}

	return x
}
