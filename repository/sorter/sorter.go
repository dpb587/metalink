package sorter

import (
	"sort"

	"github.com/dpb587/metalink/repository"
)

func Sort(results []repository.File, via Sorter) {
	ps := &sorter{
		results: results,
		sorter:  via,
	}
	sort.Sort(ps)
}

type sorter struct {
	results []repository.File
	sorter  Sorter
}

func (s *sorter) Len() int {
	return len(s.results)
}

func (s *sorter) Swap(i, j int) {
	s.results[i], s.results[j] = s.results[j], s.results[i]
}

func (s *sorter) Less(i, j int) bool {
	return s.sorter.Less(s.results[i], s.results[j])
}
