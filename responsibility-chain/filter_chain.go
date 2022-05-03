package responsibility_chain

type FilterChain interface {
	Filter(n chan int) chan int
}

type Filter struct {
	Plural  bool
	Gt100   bool
	Lt10000 bool
}

func NewFilter(plural, gt100, lt10000 bool) *Filter {
	return &Filter{
		Plural:  plural,
		Gt100:   gt100,
		Lt10000: lt10000,
	}
}

// Filter 筛选
func (f *Filter) Filter(n int) (res []int) {
	var filter []FilterChain
	if f.Plural {
		filter = append(filter, NewFilterA())
	}
	if f.Gt100 {
		filter = append(filter, NewFilterB())
	}
	if f.Lt10000 {
		filter = append(filter, NewFilterC())
	}
	var nChan = make(chan int)
	go func(nChan chan int) {
		for i := 0; i < n; i++ {
			nChan <- i
		}
		close(nChan)
	}(nChan)

	for _, f := range filter {
		nChan = f.Filter(nChan)
	}

	for v := range nChan {
		res = append(res, v)
	}
	return res
}

type FilterA struct {
	N chan int
}

func NewFilterA() *FilterA {
	return &FilterA{}
}

// Filter 筛选复数
func (f *FilterA) Filter(n chan int) chan int {
	res := make(chan int)
	go func() {
		for v := range n {
			if v%2 == 0 {
				res <- v
			}
		}
		close(res)
	}()
	return res
}

type FilterB struct {
}

func NewFilterB() *FilterB {
	return &FilterB{}
}

// Filter 大于100
func (f *FilterB) Filter(n chan int) chan int {
	res := make(chan int)
	go func() {
		for v := range n {
			if v > 100 {
				res <- v
			}
		}
		close(res)
	}()
	return res
}

type FilterC struct {
}

func NewFilterC() *FilterC {
	return &FilterC{}
}

// Filter 小于10000
func (f *FilterC) Filter(n chan int) chan int {
	res := make(chan int)
	go func() {
		for v := range n {
			if v < 10000 {
				res <- v
			}
		}
		close(res)
	}()
	return res
}
