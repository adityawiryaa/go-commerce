package port

type ProductOptions struct {
	Id string
}

type FindOptions struct {
	Id         string
	Name       string
	Ids        []string
	SortBy     string
	SortType   string
	Limit      int
	Offset     int
	IsPaginate bool
}
