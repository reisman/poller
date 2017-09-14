package voting

import (
	"github.com/pkg/errors"
	"utilities"
)

type VotingManagementService interface {
	Create(initiator *Initiator, name string, description *Description, answers []*Answer) (*Voting, error)
	Find(lookup string) ([]*Voting, error)
	Delete(voting *Voting) error
}

type service struct{}

func (srv *service) Create(initiator *Initiator, name string, description *Description, answers []*Answer) (*Voting, error) {
	if initiator == nil {
		return nil, errors.New("Initiator not provided")
	}

	if name == "" {
		return nil, errors.New("Name for the voting not provided")
	}

	if description == nil {
		return nil, errors.New("Description for the voting not provided")
	}

	if answers == nil {
		return nil, errors.New("Answers for the voting not provided")
	}

	identity := utilities.CreateIdentity()

	voting := Voting{
		Id:          identity,
		Initiator:   initiator,
		Description: description,
		Name:        name,
		Answers:     answers,
	}

	return &voting, nil
}

func (srv *service) Find(lookup string) ([]*Voting, error) {
	votings := make([]*Voting, 0)
	return votings, nil
}

func (srv *service) Delete(voting *Voting) error {
	return nil
}

func NewService() VotingManagementService {
	return &service{}
}
