package pprint

import (
	"reflect"
	"fmt"
	"text/tabwriter"
	"os"
	"bytes"
)

var buf bytes.Buffer
var values []interface{}
var depth int

func push(v interface{}) {
	values = append(values, v)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 递归计算struct里内嵌的struct深度， 返回最深深度
func depthOfStruct(dep int, val reflect.Value) int {
	max := dep
	for i := 0; i < val.NumField(); i++ {
		v := val.Field(i)
		switch v.Kind() {
		case reflect.Ptr:
			v = reflect.Indirect(v)
			fallthrough
		case reflect.Struct:
			max = Max(max, depthOfStruct(dep+1, v))
		}
	}
	return max
}

// count 决定了填充几次 %s\t
func fill(dep int) {
	for i := 0; i < dep; i++ {
		buf.WriteString("%s\t%s\t")
		push("")
		push("")
	}
}

func w(v reflect.Value, dep int) {
	switch v.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fill(dep)
		buf.WriteString("%d")
		push(v)
	case reflect.Int64:
		fill(dep)
		if v.Type().String() == "time.Duration" {
			buf.WriteString("%v")
		} else {
			buf.WriteString("%d")
		}
		push(v)
	case reflect.String:
		fill(dep)
		buf.WriteString("%s")
		push(v)
	case reflect.Float32, reflect.Float64:
		fill(dep)
		buf.WriteString("%f")
		push(v)
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if i == 0 {
				buf.WriteString("--\t%s\t")
			} else {
				fill(depth - dep)
				buf.WriteString("%s\t|--\t%s\t")
				push("")
			}
			push(v.Type().Field(i).Name)
			w(v.Field(i), dep-1)
		}
		return
	default:
		fill(dep)
		buf.WriteString("%v")
		push(v)
	}
	buf.WriteString("\n")
}

func Format(v interface{}) {
	buf.Reset()
	ind := reflect.Indirect(reflect.ValueOf(v))
	tw := tabwriter.NewWriter(os.Stdout, 0, 2, 2, ' ', 0)
	dep := depthOfStruct(0, ind)
	depth = dep
	for i := 0; i < ind.NumField(); i++ {
		v := ind.Field(i)
		if v.Kind() == reflect.Ptr {
			v = reflect.Indirect(v)
		}
		buf.WriteString("%s\t")
		push(ind.Type().Field(i).Name)
		w(v, dep)

	}
	fmt.Fprintf(tw, buf.String(), values...)
	tw.Flush()
}
