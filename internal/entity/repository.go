package entity

// Preload is a type for repository to associate different entity onto current repository.
type Preload []string

// ListFilter commonly used for repository as an argument for a filter in a function returning list of data.
type ListFilter struct {
	Column   string
	Operator string
	Value    any
}

// Sort commonly used for repository returning lists of data.
type Sort string

const (
	// SortAsc constant.
	SortAsc Sort = "ASC"
	// SortDesc constant.
	SortDesc Sort = "DESC"
)

// ListSort common repository entity.
type ListSort struct {
	Column string
	Order  Sort
}
