package data

type SliceInt []int

func (s *SliceInt) Len() int {
	return len(*s)
}

func (s *SliceInt) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *SliceInt) Delete(i int) {
	end := s.Len() - 1
	if i < end {
		s.Swap(i, end)
	}
	
	*s = (*s)[:end]
}

func (s *SliceInt) Append(t interface{}) {
	i, ok := t.(int)
	if !ok {
		panic("wrong type")
	}
	
	*s = append(*s, i)
}

