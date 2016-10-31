package main

import (
    "log"

    "github.com/emersion/go-imap/client"
    "github.com/emersion/go-imap"
)

func main() {
    log.Println("Connecting to server...")

    // Connect to server
    c, err := client.DialTLS("mail.example.org:993", nil)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Connected")

    // Don't forget to logout
    defer c.Logout()

    // Login
    if err := c.Login("username", "password"); err != nil {
        log.Fatal(err)
    }
    log.Println("Logged in")

    // List mailboxes
    mailboxes := make(chan *imap.MailboxInfo)
    done := make(chan error, 1)
    go func () {
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
    seqset.AddRange(mbox.Messages - 3, mbox.Messages)

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
