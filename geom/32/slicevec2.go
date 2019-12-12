package geom

type SliceVec2 []Vec2

func (s *SliceVec2) Len() int {
	return len(*s)
}

func (s *SliceVec2) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *SliceVec2) Delete(i int) {
	end := s.Len() - 1
	if i < end {
		s.Swap(i, end)
	}

	*s = (*s)[:end]
}

func (s *SliceVec2) Append(item interface{}) {
	i, _ := item.(Vec2)
	*s = append(*s, i)
}
