package ltkgo

import (
    "bytes"
    "encoding/binary"
)

type LlrpFrame []byte

func NewLlrpFrame(data []byte) (LlrpFrame, error) {
    buff := new(bytes.Buffer)
    err := binary.Write(buff, binary.BigEndian, int32(len(data)))
    if err != nil {
        return nil, err
    }

    err = binary.Write(buff, binary.BigEndian, data)
    if err != nil {
        return nil, err
    }

    return buff.Bytes(), nil
}

func ParseLlrpFrame(frame LlrpFrame) ([]byte, error) {
    var len int32
    err := binary.Read(bytes.NewBuffer(frame), binary.BigEndian, &len)
    if err != nil {
        return nil, err
    }

    return frame[4:len + 4], nil
}