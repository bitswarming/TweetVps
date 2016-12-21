package cmd

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
	"net/http"
	"strings"
	"strconv"
	"time"
	"github.com/PuerkitoBio/goquery"
	"log"
	"os/exec"
)

var ExpireTime int64
var User string
var Pass string
var MyDomain string
var db *bolt.DB

func init() {
	// Read value back in a different read-only transaction.
	/*var err error
	db, err = bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
		if err := db.View(func(tx *bolt.Tx) error {
			// Create a bucket.
			b, err := tx.CreateBucketIfNotExists([]byte("DoneCommands"))
			if err != nil {
				fmt.Println("bbb")
				return err
			}
			_ = b
			return nil
		}); err != nil {
			log.Fatal(err)
		}*/
	fetchCmd.Flags().StringVarP(&User, "user", "", "", "Twitter user name")
	fetchCmd.Flags().Int64Var(&ExpireTime, "expire", 0, "Seconds after which command is expired, and ignored")
	fetchCmd.Flags().StringVarP(&Pass, "secret", "", "", "Password for decoding")
	fetchCmd.Flags().StringVarP(&MyDomain, "mydomain", "", "", "domain which listens for commands")
	RootCmd.AddCommand(fetchCmd)
}


var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetches twitter feed",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		db, err = bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
		if err != nil {
			fmt.Println(err)
		}

		fetchtwitter(User, ExpireTime, Pass, MyDomain)
		//	defer db.Close()
	},
}

func fetchtwitter(twitteruser string, expire int64, secret string, mydomain string) {
	tweettext := ""
	commandalreadydone := false
	resp, err := http.Get("https://twitter.com/" + twitteruser)
	if err != nil {
	}
	t := time.Now()
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".content").Each(func(i int, s *goquery.Selection) {
		tweettext = s.Find("p.tweet-text").Text()
		timems, _ := s.Find("span.js-short-timestamp").Attr("data-time")
		fmt.Println(tweettext)
		nowtime := t.UTC().Unix()
		at, _ := strconv.ParseInt(timems, 10, 64)
		actual := nowtime - at
		//fmt.Println(tweettext, timems, actual)
		dec, _ := decryptString(tweettext, secret)
		//                      fmt.Println(dec)
		splited := strings.Split(dec, ":")
		fmt.Println(dec)
		// Iterate over items in sorted key order.
		if err := db.Update(func(tx *bolt.Tx) error {
			//b := tx.Bucket([]byte("DoneCommands"))
			b, err := tx.CreateBucketIfNotExists([]byte("DoneCommands"))
			if err != nil {
				fmt.Println("aaa")
				return err
			}

			if err := b.ForEach(func(k, v []byte) error {
				if strings.EqualFold(string(k), tweettext) {
					fmt.Printf("exists %s decoded %s.\n", k, v)
					commandalreadydone = true
				}
				return nil
			}); err != nil {
				return err
			}

			return nil
		}); err != nil {
			log.Fatal(err)
		}
		//////
		if actual <= expire && strings.EqualFold(mydomain, splited[0]) && !commandalreadydone {

			fmt.Println(tweettext, timems, actual, dec, splited[1])
			// Start a write transaction.
			if err := db.Update(func(tx *bolt.Tx) error {
							b, err := tx.CreateBucketIfNotExists([]byte("DoneCommands"))
				if err != nil {
					return err
				}

				// Set the value "bar" for the key "foo".
				if err := b.Put([]byte(tweettext), []byte(dec)); err != nil {
					return err
				}
				return nil
			}); err != nil {
				log.Fatal(err)
			}
			fmt.Println("running command", splited[1])
			cmd := exec.Command("/usr/bin/sudo", "/bin/bash", "-c", splited[1])
			_, _ = cmd.CombinedOutput()
		}
	})
	//db.Close()

}
