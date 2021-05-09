package tag

type Nothing struct{}
type Set map[string]Nothing

func BuildSet(tags []string) Set {
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

func (set Set) IncludeAll(tags []string) bool {
	if set == nil {
		return false
	}

	for _, tag := range tags {
		if _, ok := set[tag]; !ok {
			return false
		}
	}

	return true
}

func (set Set) IncludeAny(tags []string) bool {
	if set == nil {
		return true
	}
	if len(set) == 0 {
		return true
	}

	for _, tag := range tags {
		if _, ok := set[tag]; ok {
			return true
		}
	}

	return false
}

func (set Set) SubsetOf(tags []string) bool {
	if set == nil {
		return true
	}
	if len(tags) == 0 {
		return true
	}

	tagsSet := BuildSet(tags)

	for subTag := range set {
		if !tagsSet.Include(subTag) {
			return false
		}
	}

	return true
}

func (set Set) Exclude(tag string) bool {
	if set == nil {
		return true
	}

	_, ok := set[tag]

	return !ok
}
