package geom

type SliceOri2 []Ori2

func (s *SliceOri2) Len() int {
	return len(*s)
}

func (s *SliceOri2) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *SliceOri2) Delete(i int) {
	end := s.Len() - 1
	if i < end {
		s.Swap(i, end)
	}

	*s = (*s)[:end]
}

func (s *SliceOri2) Append(item interface{}) {
	i, _ := item.(Ori2)
	*s = append(*s, i)
}
