package ltkgo

import (
    "encoding/xml"
)

type LlrpConnection struct {}

func NewLlrpConnection(address string) (*LlrpConnection, error) {
    return nil, nil
}

func (c *LlrpConnection) TransactMessage(message LlrpMessage) (LlrpMessage, error) {
    b, err := xml.Marshal(message)
    if err != nil {
        return nil, err
    }

    return nil, nil
}
