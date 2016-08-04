package ltkgo_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "github.com/irryan/tcp/tcp"

    . "ltkgo"

    "encoding/xml"
    "fmt"
    "reflect"
    "time"
)

var _ = Describe("LlrpConnection", func() {
    var (
        server *tcp.TestTcpServer
        addr, port string
        conn *LlrpConnection
    )

    BeforeEach(func() {
        addr = "127.0.0.1"
        port = "8080"

        server = tcp.NewTestTcpServer(port)
        Expect(server.Start()).To(Succeed())
    })

    AfterEach(func() {
        Expect(server.HasFailedExpectations()).To(BeFalse())
        Expect(server.HasRemainingExpectations()).To(BeFalse())
        server.Stop()
    })

    Describe("NewLlrpConnection", func() {
        It("Opens a connection", func() {
            conn, err := NewLlrpConnection(fmt.Sprintf("%s:%s", addr, port))
            Expect(conn).ToNot(BeNil())
            Expect(err).ToNot(HaveOccurred())
        })
    })

    manageConnection := func() {
        BeforeEach(func() {
            var err error
            conn, err = NewLlrpConnection(fmt.Sprintf("%s:%s", addr, port))
            Expect(err).ToNot(HaveOccurred())
        })

        AfterEach(func() {
            Expect(conn.Close()).To(Succeed())
        })
    }

    Describe("SendMessage", func() {
        manageConnection()

        It("Sends a message which is received", func() {
            m := LlrpMessage{X: 5}

            server.AddExpectation(func(buff []byte) error {
                defer GinkgoRecover()
                data, err := ParseLlrpFrame(buff)
                Expect(err).ToNot(HaveOccurred())

                var v LlrpMessage
                Expect(xml.Unmarshal(data, &v)).To(Succeed())
                Expect(v).To(Equal(m))
                return nil
            })

            Expect(conn.SendMessage(m)).To(Succeed())
            time.Sleep(100*time.Millisecond)
        })

        It("Does a thing", func() {})
    })

    Describe("TransactMessage", func() {
        manageConnection()

        It("Sends a message which is received and receives the expected response", func() {
            m := LlrpMessage{X: 5}

            server.AddExpectation(func(buff []byte) error {
                defer GinkgoRecover()
                return nil
            })

            Expect(conn.TransactMessage(m, reflect.TypeOf((*LlrpMessage)(nil)).Elem())).To(Succeed())
            time.Sleep(100*time.Millisecond)
        })
    })
})
