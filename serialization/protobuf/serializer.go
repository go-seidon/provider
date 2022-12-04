package protobuf

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

type protobufSerializer struct {
}

func (s *protobufSerializer) Marshal(i interface{}) ([]byte, error) {
	msg, ok := i.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("invalid message")
	}
	return proto.Marshal(msg)
}

func (s *protobufSerializer) Unmarshal(i []byte, o interface{}) error {
	msg, ok := o.(proto.Message)
	if !ok {
		return fmt.Errorf("invalid message")
	}
	return proto.Unmarshal(i, msg)
}

func NewSerializer() *protobufSerializer {
	return &protobufSerializer{}
}
