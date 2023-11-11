package table

type Rows []Row

func (r *Rows) Remove(i int) {
	if len((*r)) > i {
		(*r) = append((*r)[:i], (*r)[i+1:]...)
	}
}
