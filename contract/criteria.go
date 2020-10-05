package contract

type RepoCondition struct {
	Field        string
	Operation    string
	Subquery     bool
	Value        interface{}
	Conditions   []*RepoCondition
	OrConditions []*RepoCondition
}

type RepoJoin struct {
	Join string
}

type Preload struct {
	Relation  string
	Criterias *RepoCriterias
}

type RepoCriterias struct {
	Conditions   []*RepoCondition
	OrConditions []*RepoCondition
	Joins        []*RepoJoin
	GroupBy      []string
	Preloads     []*Preload
	Order        []string
}
