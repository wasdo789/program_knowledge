package structure

import "container/list"

type Stu struct {
	Name string
}

type StuUnion struct {
	Parents map[*Stu]*Stu
}

func NewStuUnion(students []*Stu) *StuUnion {
	u := &StuUnion{}
	u.Parents = make(map[*Stu]*Stu)
	for _, stu := range students {
		u.Parents[stu] = stu
	}
	return u

}

func (u *StuUnion) FindFather(s *Stu) *Stu {
	if len(u.Parents) == 0 || s == nil {
		return nil
	}
	l := list.New()
	for {
		h, ok := u.Parents[s]
		if !ok {
			return nil
		} else if h != s {
			l.PushBack(s)
			s = h
		} else {
			for l.Len() > 0 {
				front := l.Remove(l.Front()).(*Stu)
				u.Parents[front] = h

			}
			return h
		}
	}

}

func (u *StuUnion) Merge(s1, s2 *Stu) bool {
	//
	h1 := u.FindFather(s1)
	if h1 == nil {
		return false
	}
	h2 := u.FindFather(s2)
	if h2 == nil {
		return false
	}
	//
	u.Parents[h1] = h2
	return true
}

func (u *StuUnion) IsSame(s1, s2 *Stu) bool {
	h1 := u.FindFather(s1)
	h2 := u.FindFather(s2)
	if h1 == nil || h2 == nil {
		return false
	}
	return h1 == h2
}
