package main

import (
	//	"net"
	//"fmt"
	"log"

	"github.com/whxmail/whx/types"

	//	"github.com/emersion/go-imap"
	//	"github.com/emersion/go-imap/client"
)

func init() {

}

func main() {
	log.Println("Connecting to server...")
	server := types.NewTCPServer()

	//Register router
	server.MuxRegister("LOGIN", login)
	server.ListenAndServer("127.0.1.1:9999", myHandleFunc)

	//	cmd := data.GetCMD()

	/*
		if cmd[0] == "LOGIN" {
			mail := cmd[1]
			pwd := cmd[2]

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
		}
	*/
}
