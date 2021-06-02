package potato

type DataList []interface{}

func (l DataList) Len() int {
	return len(l)
}
func (l DataList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

type ByName struct{ DataList }

func (b ByName) Less(i, j int) bool {
	switch v := b.DataList[i].(type) {
	case Background:
		return v.Name < b.DataList[j].(Background).Name
	case Effect:
		return v.Name < b.DataList[j].(Effect).Name
	case Engine:
		return v.Name < b.DataList[j].(Engine).Name
	case Particle:
		return v.Name < b.DataList[j].(Particle).Name
	case Skin:
		return v.Name < b.DataList[j].(Skin).Name
	case Level:
		return v.Name < b.DataList[j].(Level).Name
	default:
		return i < j
	}
}

type ByCreatedTime struct{ DataList }

func (b ByCreatedTime) Less(i, j int) bool {
	switch v := b.DataList[i].(type) {
	case Background:
		return v.CreatedTime < b.DataList[j].(Background).CreatedTime
	case Effect:
		return v.CreatedTime < b.DataList[j].(Effect).CreatedTime
	case Engine:
		return v.CreatedTime < b.DataList[j].(Engine).CreatedTime
	case Particle:
		return v.CreatedTime < b.DataList[j].(Particle).CreatedTime
	case Skin:
		return v.CreatedTime < b.DataList[j].(Skin).CreatedTime
	case Level:
		return v.CreatedTime < b.DataList[j].(Level).CreatedTime
	default:
		return i < j
	}
}

type ByUpdatedTime struct{ DataList }

func (b ByUpdatedTime) Less(i, j int) bool {
	switch v := b.DataList[i].(type) {
	case Background:
		return v.UpdatedTime < b.DataList[j].(Background).UpdatedTime
	case Effect:
		return v.UpdatedTime < b.DataList[j].(Effect).UpdatedTime
	case Engine:
		return v.UpdatedTime < b.DataList[j].(Engine).UpdatedTime
	case Particle:
		return v.UpdatedTime < b.DataList[j].(Particle).UpdatedTime
	case Skin:
		return v.UpdatedTime < b.DataList[j].(Skin).UpdatedTime
	case Level:
		return v.UpdatedTime < b.DataList[j].(Level).UpdatedTime
	default:
		return i < j
	}
}

type ByDifficulty struct{ DataList }

func (b ByDifficulty) Less(i, j int) bool {
	switch v := b.DataList[i].(type) {
	case Level:
		return v.Rating < b.DataList[j].(Level).Rating
	default:
		return i < j
	}
}

type ByNotes struct{ DataList }

func (b ByNotes) Less(i, j int) bool {
	switch v := b.DataList[i].(type) {
	case Level:
		return v.Notes < b.DataList[j].(Level).Notes
	default:
		return i < j
	}
}
