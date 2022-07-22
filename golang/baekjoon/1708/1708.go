package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"sync"
)

type Interface interface {
	Take(i int) (x, y float64)
	Len() int
	Swap(i, j int)
	Slice(i, j int) Interface
}

type coordinates [][2]float64

func (c coordinates) Take(i int) (x, y float64) {
	return c[i][0], c[i][1]
}

func (c coordinates) Len() int {
	return len(c)
}

func (c coordinates) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c coordinates) Slice(i, j int) Interface {
	return c[i:j]
}

type Options struct {
	Pool *sync.Pool
}
type poolElStruct struct {
	lowerIndexes, upperIndexes []int
}

func New(points Interface) Interface {
	sort.Sort(pointSorter{i: points})
	return NewFromSortedArray(points)
}

func NewFromSortedArray(points Interface) Interface {
	o := Options{}
	return NewFromSortedArrayWithOptions(points, o)
}
func NewWithOptions(points Interface, o Options) Interface {
	sort.Sort(pointSorter{i: points})
	return NewFromSortedArrayWithOptions(points, o)
}

func NewFromSortedArrayWithOptions(points Interface, o Options) Interface {
	n := points.Len()
	if n < 3 {
		return points
	}
	var w sync.WaitGroup
	var lowerIndexes []int
	var upperIndexes []int
	var isPooledMemReceived bool
	if o.Pool != nil {
		poolElCandidate := o.Pool.Get()
		if poolElCandidate != nil {
			isPooledMemReceived = true
			poolEl := poolElCandidate.(*poolElStruct)
			lowerIndexes = poolEl.lowerIndexes[0:0]
			upperIndexes = poolEl.upperIndexes[0:0]
		}
	}
	if !isPooledMemReceived {
		lowerIndexes = make([]int, 0, 5)
		upperIndexes = make([]int, 0, 5)
	}
	lowerIndexes = append(lowerIndexes, 0, 1)
	upperIndexes = append(upperIndexes, n-1, n-2)
	w.Add(2)
	func() {
		for i := 2; i < n; i++ {
			x, y := points.Take(i)
			m := len(lowerIndexes)
			for m > 1 {
				x2, y2 := points.Take(lowerIndexes[m-2])
				x3, y3 := points.Take(lowerIndexes[m-1])
				if isOrientationPositive(x2, y2, x3, y3, x, y) {
					break
				}
				lowerIndexes = lowerIndexes[:m-1]
				m -= 1
			}
			lowerIndexes = append(lowerIndexes, i)
		}

		w.Done()
	}()
	func() {
		for i := n - 3; i >= 0; i-- {
			x, y := points.Take(i)
			m := len(upperIndexes)
			for m > 1 {
				x2, y2 := points.Take(upperIndexes[m-2])
				x3, y3 := points.Take(upperIndexes[m-1])
				if isOrientationPositive(x2, y2, x3, y3, x, y) {
					break
				}
				upperIndexes = upperIndexes[:m-1]
				m -= 1
			}
			upperIndexes = append(upperIndexes, i)

		}
		w.Done()
	}()
	w.Wait()

	upperIndexes = upperIndexes[:len(upperIndexes)-1]
	lowerIndexes = lowerIndexes[:len(lowerIndexes)-1]
	allIndexes := append(lowerIndexes, upperIndexes...)

	result := sortByIndexes(points, allIndexes)
	if o.Pool != nil {
		o.Pool.Put(&poolElStruct{lowerIndexes: allIndexes, upperIndexes: upperIndexes})
	}
	return result
}

func isOrientationPositive(x1, y1, x2, y2, x3, y3 float64) (isPositive bool) {
	return (x1-x3)*(y2-y3)-(y1-y3)*(x2-x3) > 0
}

type pointSorter struct {
	i Interface
}

func (s pointSorter) Less(i, j int) bool {
	x1, y1 := s.i.Take(i)
	x2, y2 := s.i.Take(j)
	return x1 < x2 || (x1 == x2 && y1 < y2)
}

func (s pointSorter) Swap(i, j int) {
	s.i.Swap(i, j)
}

func (s pointSorter) Len() int {
	return s.i.Len()
}

func sortByIndexes(points Interface, indices []int) Interface {
	s := indexSorter{indices: indices, points: points}
	sort.Sort(s)
	for i, index := range indices {
		points.Swap(i, index)
	}
	return points.Slice(0, len(indices))
}

type indexSorter struct {
	indices []int
	points  Interface
}

func (s indexSorter) Less(i, j int) bool {
	return s.indices[i] < s.indices[j]
}

func (s indexSorter) Swap(i, j int) {
	s.indices[i], s.indices[j] = s.indices[j], s.indices[i]
	s.points.Swap(s.indices[i], s.indices[j])
}

func (s indexSorter) Len() int {
	return len(s.indices)
}
func main() {
	var N int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscanf(r, "%d", &N)
	points := make([][2]float64, N+1)
	for i := 0; i <= N; i++ {
		fmt.Fscanln(r, &points[i][0], &points[i][1])
	}
	convexHull := New(coordinates(points))
	fmt.Println(convexHull.Len())
}
