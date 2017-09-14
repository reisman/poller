package voting

type Voting struct {
	Id          string
	Name        string
	Initiator   *Initiator
	Description *Description
	Answers     []*Answer
}

type Description struct {
	Id          string
	Question    string
	Description string
}

type Answer struct {
	Id          string
	Text        string
	Description string
}

type Initiator struct {
	Id   string
	Name string
}

type SelectedAnswer struct {
	Voter  *Voter
	Voting *Voting
	Answer *Answer
}

type Voter struct {
	Id   string
	Name string
}
