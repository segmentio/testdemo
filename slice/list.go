package slice

type kv struct {
	k string
	v string
}

type Local []kv

func (l Local) Get(key string) (string, error) {
	for _, i := range l {
		if i.k == key {
			return i.v, nil
		}
	}
	return "", nil
}

func (l *Local) Set(key, value string) error {
	sl := make(Local, 0, len(*l))
	for _, i := range *l {
		if i.k != key {
			sl = append(sl, i)
		}
	}
	*l = append(sl, kv{k: key, v: value})
	return nil
}
