package ltkgo_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "github.com/irryan/tcp"
    "github.com/irryan/tcp/middleware"

    . "ltkgo"

    "encoding/xml"
    "fmt"
    "log"
    "os"
    "reflect"
    "time"
)

var _ = Describe("LlrpConnection", func() {
    var (
        server tcp.TcpServer
        addr, port string
        conn *LlrpConnection
        em *middleware.ExpectationMiddleware
    )

    BeforeEach(func() {
        addr = "127.0.0.1"
        port = "8080"

        em = new(middleware.ExpectationMiddleware)
        server = tcp.NewTcpServer(log.New(os.Stdout, "", log.LstdFlags), port, em)
        Expect(server.Start()).To(Succeed())
    })

    AfterEach(func() {
        Expect(em.HasFailedExpectations()).To(BeFalse())
        Expect(em.HasRemainingExpectations()).To(BeFalse())
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

            em.AddExpectation(func(buff []byte) ([]byte, error) {
                defer GinkgoRecover()
                data, err := ParseLlrpFrame(buff)
                Expect(err).ToNot(HaveOccurred())

                var v LlrpMessage
                Expect(xml.Unmarshal(data, &v)).To(Succeed())
                Expect(v).To(Equal(m))

                return nil, nil
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
            expected := LlrpMessage{X:10}
            b, err := xml.Marshal(expected)
            Expect(err).ToNot(HaveOccurred())

            em.AddExpectation(func(buff []byte) ([]byte, error) {
                defer GinkgoRecover()
                data, err := ParseLlrpFrame(buff)
                Expect(err).ToNot(HaveOccurred())

                var v LlrpMessage
                Expect(xml.Unmarshal(data, &v)).To(Succeed())
                Expect(v).To(Equal(m))

                frame, _ := NewLlrpFrame(b)
                return frame, nil
            })

            resp, err := conn.TransactMessage(m, reflect.TypeOf((*LlrpMessage)(nil)).Elem())
            Expect(resp.(LlrpMessage)).To(Equal(expected))
            Expect(err).ToNot(HaveOccurred())
            time.Sleep(100*time.Millisecond)
        })
    })
})
