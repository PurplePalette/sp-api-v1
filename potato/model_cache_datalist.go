package potato

// DataList object is a list of something data
type DataList []interface{}

// Len implements length function for using sort
func (l DataList) Len() int {
	return len(l)
}

// Swap implements swap function for using sort
func (l DataList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// ByName is decorator of DataList for sorting by name
type ByName struct{ DataList }

// Less is compare method for sorting by name
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

// ByCreatedTime is decorator of DataList for sorting by created time
type ByCreatedTime struct{ DataList }

// Less is compare method for sorting by created time
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

// ByUpdatedTime is decorator of DataList for sorting by updated time
type ByUpdatedTime struct{ DataList }

// Less is compare method for sorting by updated time
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

// ByDifficulty is compare method for sorting by difficulty
type ByDifficulty struct{ DataList }

// Less is compare method for sorting by difficulty
func (b ByDifficulty) Less(i, j int) bool {
	switch v := b.DataList[i].(type) {
	case Level:
		return v.Rating < b.DataList[j].(Level).Rating
	default:
		return i < j
	}
}

// ByNotes is compare method for sorting by notes
type ByNotes struct{ DataList }

// Less is compare method for sorting by notes
func (b ByNotes) Less(i, j int) bool {
	switch v := b.DataList[i].(type) {
	case Level:
		return v.Notes < b.DataList[j].(Level).Notes
	default:
		return i < j
	}
}
