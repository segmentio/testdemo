package mapstring

type Remote map[string]string

func (r Remote) Get(key string) (string, error) {
	return r[key], nil
}

func (r Remote) Set(key, value string) error {
	r[key] = value
	return nil
}
