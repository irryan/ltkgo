package ltkgo

import (
    "encoding/xml"
    "net"
    "reflect"
    "time"
)

type LlrpConnection struct {
    close chan struct{}
    conn net.Conn

    messageMap map[reflect.Type][]interface{}
    messageQueue []interface{}
}

func NewLlrpConnection(address string) (*LlrpConnection, error) {
    conn, err := net.Dial("tcp", address)
    if err != nil {
        return nil, err
    }

    conn.SetDeadline(time.Now().Add(1e9))
    llrpConn := &LlrpConnection{
        close: make(chan struct{}),
        conn: conn,
        messageMap: make(map[reflect.Type][]interface{}),
        messageQueue: make([]interface{}, 0, 0),
    }

    // go func(quit chan struct{}) {
    //     for quit != nil {
    //         var b []byte

    //         n, err := llrpConn.conn.Read(b)
    //         if err != nil {
    //             quit = nil
    //         }

    //         var serializer BinarySerializer
    //         b, err = serializer.Deserialize(b)
    //         if err != nil {
                
    //         }

    //         binder := func (data []byte) interface{} {
                
    //         }

    //         v := binder(b)
    //         t := reflect.TypeOf(v)
    //         if _, ok := llrpConn.messageMap[t]; !ok {
    //             llrpConn.messageMap[t] = make([]*interface{}, 0, 0)
    //         }

    //         llrpConn.messageMap[t] = append(llrpConn.messageMap[t], &v)
    //         llrpConn.messageQueue = append(llrpConn.messageQueue, &v)
    //     }
    // }(llrpConn.close)

    return llrpConn, nil
}

func (c* LlrpConnection) Close() error {
    return c.conn.Close()
}

func (c *LlrpConnection) SendMessage(message LlrpMessage) error {
    b, err := xml.Marshal(message)
    if err != nil {
        return err
    }

    frame, err := NewLlrpFrame(b)
    if err != nil {
        return err
    }

    n, err := c.conn.Write(frame)
    if err != nil {
        return err
    }

    if n != len(frame) {

    }

    return nil
}

func (c *LlrpConnection) TransactMessage(message LlrpMessage, expectedType reflect.Type) (interface{}, error) {
    b, err := xml.Marshal(message)
    if err != nil {
        return nil, err
    }

    frame, err := NewLlrpFrame(b)
    if err != nil {
        return nil, err
    }

    n, err := c.conn.Write(frame)
    if err != nil {
        return nil, err
    }

    if n != len(frame) {

    }

    frame = make([]byte, 4096)
    _, err = c.conn.Read(frame)
    if err != nil {
        return nil, err
    }

    data, err := ParseLlrpFrame(frame)
    if err != nil {
        return nil, err
    }

    var m LlrpMessage
    err = xml.Unmarshal(data, &m)
    if err != nil {
        return nil, err
    }

    return m, nil
}
