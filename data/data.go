package data

import (
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/kristoiv/wtf"
)

//go:generate protoc --go_out=plugins=grpc:. data.proto

func MarshalItem(item *wtf.Item) (*Item, error) {
	created, err := item.Created.MarshalBinary()
	if err != nil {
		return nil, err
	}

	changed, err := item.Changed.MarshalBinary()
	if err != nil {
		return nil, err
	}

	return &Item{
		Id:      string(item.ID),
		Title:   item.Title,
		Created: created,
		Changed: changed,
		Checked: item.Checked,
	}, nil
}

func UnmarshalItem(item *Item, out *wtf.Item) error {
	out.ID = wtf.ItemID(item.GetId())
	out.Title = item.GetTitle()
	out.Checked = item.GetChecked()

	created := &time.Time{}
	if err := created.UnmarshalBinary(item.GetCreated()); err != nil {
		return err
	}
	out.Created = *created

	changed := &time.Time{}
	if err := changed.UnmarshalBinary(item.GetChanged()); err != nil {
		return err
	}
	out.Changed = *changed

	return nil
}

func MarshalItemBytes(item *wtf.Item) ([]byte, error) {
	created, err := item.Created.MarshalBinary()
	if err != nil {
		return nil, err
	}

	changed, err := item.Changed.MarshalBinary()
	if err != nil {
		return nil, err
	}

	return proto.Marshal(&Item{
		Id:      string(item.ID),
		Title:   item.Title,
		Created: created,
		Changed: changed,
		Checked: item.Checked,
	})
}

func UnmarshalItemBytes(data []byte, item *wtf.Item) error {
	pb := Item{}
	if err := proto.Unmarshal(data, &pb); err != nil {
		return err
	}

	item.ID = wtf.ItemID(pb.GetId())
	item.Title = pb.GetTitle()
	item.Checked = pb.GetChecked()

	created := &time.Time{}
	if err := created.UnmarshalBinary(pb.GetCreated()); err != nil {
		return err
	}
	item.Created = created.UTC()

	changed := &time.Time{}
	if err := changed.UnmarshalBinary(pb.GetChanged()); err != nil {
		return err
	}
	item.Changed = changed.UTC()

	return nil
}