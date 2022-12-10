package service

import (
	"bytes"
	"context"
	"github.com/cornelk/hashmap"
	"github.com/unsafe-risk/simple-db/query/executor"
	"log"
	"net"
)

type Service struct {
	executors *hashmap.Map[string, *executor.Executor]
	ctx       context.Context
	cancel    context.CancelFunc
}

func New() *Service {
	return &Service{
		executors: hashmap.New[string, *executor.Executor](),
	}
}

func (s *Service) ListenAndServe(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	go func() {
		defer func(lis net.Listener) {
			err := lis.Close()
			if err != nil {
				log.Println(err)
			}
		}(lis)

		for {
			conn, err := lis.Accept()
			if err != nil {
				log.Println(err)
				continue
			}
			go s.handler(conn)
		}
	}()
	return nil
}

func (s *Service) handler(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)
	tempBuffer := [4096]byte{}
	buffer := bytes.NewBuffer(nil)
	for {
		n, err := conn.Read(tempBuffer[:])
		if err != nil {
			log.Println(err)
			return
		}
		buffer.Write(tempBuffer[:n])
		if n < 4096 {
			func() {
				input := buffer.Bytes()
				t := bytes.SplitN(input, []byte(" "), 2)
				if len(t) != 2 {
					if _, err := conn.Write([]byte("ERR: invalid input")); err != nil {
						log.Println(err)
					}
					return
				}

				exe, ok := s.executors.Get(string(t[0]))
				if !ok {
					if _, err := conn.Write([]byte("ERR: executor not found")); err != nil {
						log.Println(err)
					}
					return
				}

				result, err := exe.Run(t[1])
				if err != nil {
					if _, err := conn.Write([]byte("ERR: " + err.Error())); err != nil {
						log.Println(err)
					}
					return
				}

				if _, err := conn.Write(result); err != nil {
					log.Println(err)
				}
			}()
			buffer.Reset()
		}
	}
}

func (s *Service) RegisterTable(name string, columns []string, types []int) error {
	exe, err := executor.New()
	if err != nil {
		return err
	}

	if err := exe.SetTable(name, columns, types); err != nil {
		return err
	}

	s.executors.Set(name, exe)

	return nil
}

func (s *Service) Close() error {
	s.executors.Range(func(key string, value *executor.Executor) bool {
		if err := value.Close(); err != nil {
			log.Println(err)
		}
		return true
	})
	return nil
}
