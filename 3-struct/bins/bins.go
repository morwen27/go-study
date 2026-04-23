package bins

import (
	"errors"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"creationDate"`
	Name      string    `json:"name"`
}

type BinList struct {
	Bins []*Bin `json:"bins"`
}

func ConstructBin(id string, name string, private bool) (*Bin, error) {
	if id == "" || name == "" {
		return nil, errors.New("Id or name can't be empty")
	}
	bin := Bin{
		Id:        id,
		Name:      name,
		CreatedAt: time.Now(),
		Private:   private,
	}

	return &bin, nil
}

func ConstructBinList() *BinList {
	return &BinList{}
}

func AddBin(bins []*Bin) *BinList {
	slice := make([]*Bin, 0)
	binList := BinList{
		Bins: append(slice, bins...),
	}

	return &binList
}
