package exec

import (
	"github.com/shuvigoss/pcs/config"
	"time"
	"fmt"
	"bytes"
	"sync"
	"net"
	"golang.org/x/crypto/ssh"
)

const (
	Gray    = uint8(iota + 90)
	Red
	Green
	Yellow
	Blue
)

func RunCommands(config config.Config, command string) {
	var wg sync.WaitGroup
	for _, host := range config.Hosts {
		if host.Port == 0 {
			host.Port = config.Globalport
		}

		if host.Username == "" {
			host.Username = config.Globalname
		}

		if host.Password == "" {
			host.Password = config.Globalpwd
		}

		h := host
		wg.Add(1)
		go func() {
			var session *ssh.Session

			for i := 0; i < 3; i ++ {
				tmp, e := connect(h.Username, h.Password, h.Host, h.Port)
				if e != nil {
					time.Sleep(1 * time.Second)
					continue
				}
				session = tmp
				break
			}
			defer func() {
				if session != nil {
					session.Close()
				}
				wg.Done()
			}()

			if session == nil {
				fmt.Printf("connect to %s error", h.Host)
				return
			}

			var b bytes.Buffer
			session.Stdout = &b
			if err := session.Run(command); err != nil {
				fmt.Printf("Failed to run: %s with error %s", command, err)
				return
			}
			fmt.Printf("\033[%dm<< %s return >>\033[0m  \n %s \n\n", Blue, h.Host, b.String())

		}()
	}

	wg.Wait()

}

func connect(user, password, host string, port int) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	keyboardInteractiveChallenge := func(
		user,
		instruction string,
		questions []string,
		echos []bool,
	) (answers []string, err error) {
		fmt.Println(questions)
		if len(questions) == 0 {
			return []string{}, nil
		}
		return []string{password}, nil
	}

	auth = append(auth, ssh.KeyboardInteractive(keyboardInteractiveChallenge))
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}
