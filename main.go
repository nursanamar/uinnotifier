package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
	"github.com/nursanamar/uinnotifier/src/notifier"
	"github.com/nursanamar/uinnotifier/src/recorder"
)

type Announcement struct {
	Title       string
	Timestamp   int64
	Description string
	Url         string
}

func getRecentAnnouncements() ([]Announcement, error) {

	var announcements []Announcement
	const dateFormat = "02 Jan 2006"

	res, err := http.Get("https://uin-alauddin.ac.id/pengumuman")
	if err != nil {
		log.Print(err)
		return announcements, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Print("Bad responses", res)
		return announcements, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Print("Failed to parse")
		return announcements, err
	}

	doc.Find(".post-item").Each(func(i int, s *goquery.Selection) {
		var announcement Announcement

		postTime, err := time.Parse(dateFormat, s.Find(".post-meta-date").Text())
		if err == nil {
			announcement.Timestamp = postTime.Unix()
		}
		announcement.Title = s.Find(".text-primary > a").Text()
		announcement.Url = s.Find(".text-primary > a").AttrOr("href", "")
		announcement.Description = s.Find("p").Text()

		announcements = append(announcements, announcement)
	})

	return announcements, nil
}

func isSomethinNew(recent Announcement) bool {
	recorder := recorder.MakeRecoder("File")

	last, err := recorder.GetRecord()
	if err != nil {
		last = "0"
	}

	lastTimestamp, err := strconv.ParseInt(last, 10, 64)
	if err != nil {
		lastTimestamp = 0
	}

	if recent.Timestamp > lastTimestamp {
		recorder.Record(strconv.Itoa(int(recent.Timestamp)))
		return true
	}

	return false
}

func sendNotification(recent Announcement) error {
	notifier := notifier.MakeNotifier("mail")

	notifier.SetMessage("UIN AlAUDDIN : " + recent.Title + "\n" + recent.Url)
	err := notifier.SendMessage("nursan011@gmail.com")
	if err != nil {
		return err
	}

	return nil
}

func check(c chan Announcement, quit chan int) {
	uptimeTicker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-uptimeTicker.C:
			log.Print("Check Announcmecmnt")
			list, err := getRecentAnnouncements()
			if err != nil {
				log.Print(err)
			}

			if len(list) < 1 {
				log.Print("Empty Annoucmenct")
				continue
			}

			recent := list[0]
			if isSomethinNew(recent) {
				c <- recent
			}
		case <-quit:
			return
		}
	}
}

func main() {

	log.Print("Init....")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	recent := make(chan Announcement)
	quit := make(chan int)

	log.Print("Running check routine")
	go check(recent, quit)

	for {
		select {
		case r := <-recent:
			log.Print("New Announcement")
			log.Print("Sending Message")
			err := sendNotification(r)
			if err != nil {
				log.Print("Failed to send notification")
				log.Print(err)
			}
		}
	}
}
