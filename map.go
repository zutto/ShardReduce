package ShardReduce

import (
	"reflect"
)

type ShardReduce struct {
	input []*interface{}
}

func NewShardReduce(input []*interface{}) *ShardReduce {
	x := ShardReduce{
		input: input,
	}

	return &x
}

//input func(*interface)
func (sr *ShardReduce) Filter(fFunc interface{}) *ShardReduce {
	tempStock := make([]*interface{}, 0)
	for _, v := range sr.input {
		r := reflect.ValueOf(fFunc).Call([]reflect.Value{reflect.ValueOf(v)}) // fFunc(v)

		if r[0].Bool() == true {
			tempStock = append(tempStock, v)
		}
	}
	sr.input = tempStock
	return sr
}

//input func(interface{}) *interface{}
func (sr *ShardReduce) Map(mapFunc interface{}) *ShardReduce {
	tempStock := make([]*interface{}, 0)
	for _, v := range sr.input {
		//r := mapFunc(v)
		r := reflect.ValueOf(mapFunc).Call([]reflect.Value{reflect.ValueOf(*v)})[0].Interface()
		if r != nil {
			tempStock = append(tempStock, &r)
		}
	}
	sr.input = tempStock
	return sr
}

//input *interface{}, *interface{} -> output *interface{}
func (sr *ShardReduce) Reduce(reduceFunc interface{}) *interface{} {
	var last interface{}

	for _, v := range sr.input {
		if last != nil {
			last = reflect.ValueOf(reduceFunc).Call([]reflect.Value{reflect.ValueOf(last), reflect.ValueOf(*v)})[0].Interface()
		} else {
			last = *v
			continue
		}
	}
	return &last
}
