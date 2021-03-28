package tag

type Nothing struct{}
type Set map[string]Nothing

func BuildSet(tags ...string) Set {
	set := make(map[string]Nothing, len(tags))

	for _, tag := range tags {
		set[tag] = Nothing{}
	}

	return set
}

func BuildSetOfSize(size int) Set {
	return make(map[string]Nothing, size)
}

func (set Set) Add(tag string) {
	set[tag] = Nothing{}
}

func (set Set) Include(tag string) bool {
	if set == nil {
		return false
	}

	_, ok := set[tag]

	return ok
}

func (set Set) Exclude(tag string) bool {
	if set == nil {
		return true
	}

	_, ok := set[tag]

	return !ok
}
