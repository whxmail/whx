package main

import (
	//	"fmt"
	//	"encoding/json"
	"log"

	//	"github.com/emersion/go-imap"
	//	"github.com/emersion/go-imap/client"
)

func main() {
	log.Println("Connecting to server...")

	//Start ipserver
	test("Starting IPServer...")
	is := ipServer{}
	is.set("127.0.1.1")
	is.startServer()

	/*
		buf := make([]byte, 512)
		n, addr, err := is.conn.ReadFromIP(buf)
		checkError(err)
		test(string(buf[:n]))

		u := User{}
		json.Unmarshal(buf[:n], &u)
		test("u.Username")
		fmt.Println(u.Username)

		fmt.Println("Receive from client", addr.String(), string(buf[:n]))

		json.Unmarshal(buf, &u)

		is.conn.WriteToIP([]byte("Welcome Client!"), addr)
	*/

	data := is.getData() //return type: map[string]interface{}

	//	usr := User{}
	if dt, ok := data["CMD"]; ok {
		test(dt)
		cmd := dt.([]interface{})
		test(cmd)
		for _, v := range cmd {
			out := v.(string)
			test(out)
		}
		/*		if !ok {
					fmt.Println("Data error!")
				}
		*/
		/*		switch cmd[0] {
				case "LOGIN":
					usr.Username = cmd[1]
					usr.Password = cmd[2]
				case "LOGOUT":
					usr.Flag = false
				}*/
	}
	/*
		mail := usr.Username
		pwd := usr.Password

		test(mail)
		// Connect to imapserver
		test("Connect to IMAPServer")
		c, err := client.DialTLS("imap.qq.com:993", nil)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Connected")

		// Don't forget to logout
		defer c.Logout()

		//Obtain information about user

		// Login

		if err := c.Login(mail, pwd); err != nil {
			log.Fatal(err)
		}
		log.Println("Logged in")

		// List mailboxes
		mailboxes := make(chan *imap.MailboxInfo)
		done := make(chan error, 1)
		go func() {
			done <- c.List("", "*", mailboxes)
		}()

		log.Println("Mailboxes:")
		for m := range mailboxes {
			log.Println("* " + m.Name)
		}

		if err := <-done; err != nil {
			log.Fatal(err)
		}

		// Select INBOX
		mbox, err := c.Select("INBOX", false)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Flags for INBOX:", mbox.Flags)

		// Get the last 4 messages
		seqset, _ := imap.NewSeqSet("")
		seqset.AddRange(mbox.Messages-3, mbox.Messages)

		messages := make(chan *imap.Message)
		done = make(chan error, 1)
		go func() {
			done <- c.Fetch(seqset, []string{imap.EnvelopeMsgAttr}, messages)
		}()

		for msg := range messages {
			log.Println(msg.Envelope.Subject)
		}

		if err := <-done; err != nil {
			log.Fatal(err)
		}

		log.Println("Done!")
	*/
}
