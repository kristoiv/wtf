package internal

import (
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/kristoiv/wtf"
)

//go:generate protoc --go_out=. internal.proto

func MarshalItem(item *wtf.Item) ([]byte, error) {
	return proto.Marshal(&Item{
		Id:      string(item.ID),
		Title:   item.Title,
		Created: item.Created.UnixNano(),
		Changed: item.Changed.UnixNano(),
		Checked: item.Checked,
	})
}

func UnmarshalItem(data []byte, item *wtf.Item) error {
	pb := Item{}
	if err := proto.Unmarshal(data, &pb); err != nil {
		return err
	}

	item.ID = wtf.ItemID(pb.GetId())
	item.Title = pb.GetTitle()
	item.Created = time.Unix(0, pb.GetCreated()).UTC()
	item.Changed = time.Unix(0, pb.GetChanged()).UTC()
	item.Checked = pb.GetChecked()

	return nil
}
