package main

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"net/url"
	"os"
	"time"
)

const (
	BookingLink = "https://www.booking.com/hotel/fr/hotelfrancais.fr.html?aid=311984&label=hotel-50298-fr-FiL6EYFft2PxTc2lRRg2LwS162155423944%3Apl%3Ata%3Ap1%3Ap2%3Aac%3Aap%3Aneg%3Afi%3Atikwd-351816081071%3Alp9207452%3Ali%3Adec%3Adm%3Appccp%3DUmFuZG9tSVYkc2RlIyh9YTQUGSsRwx9_3qo3uPTHyoo&sid=b945786d9ef83f9832fba8a61eee38d7&dest_id=-1456928&dest_type=city&dist=0&group_adults=2&group_children=0&hapos=1&hpos=1&no_rooms=1&req_adults=2&req_children=0&room1=A%2CA&sb_price_type=total&sr_order=popularity&srepoch=1771626391&srpvid=8fa59dc7ccaf01a4&type=total&ucfs=1&"

	// UserDataDir Create a new dir for rod to store user data in there.
	userDataDir = "/Users/malik/Documents/rod-chrome-data"
	profileDir  = "Profile 1"
)

func main() {
	u := launcher.NewUserMode().
		UserDataDir(userDataDir).
		ProfileDir(profileDir).
		Headless(false)

	wsURL := u.MustLaunch()

	browser := rod.New().
		ControlURL(wsURL).
		MustConnect()

	defer browser.MustClose()

	// It's often safer to use browser.MustPage("") then navigate
	page := browser.MustPage("")
	page.MustSetViewport(1920, 1080, 1, false)

	pageUrl, err := preparePageURL(BookingLink, "fr", "fr")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error parsing the page: %v", err)
	}

	page.MustNavigate(pageUrl)

	hotelName, err := getPlaceName(page)
	if err != nil {
		stdErr(err)
		return
	}

	fmt.Printf("Hotel name: %s", hotelName)
	//fmt.Printf("Hotel description: %s", hotelDesc)

	time.Sleep(time.Hour * 1)
}

func getPlaceName(page *rod.Page) (string, error) {
	h1, err := page.Element("div#hp_hotel_name")
	if err != nil {
		return "", fmt.Errorf("failed to get hotel name: %w", err)
	}

	h1Text, err := h1.Text()
	if err != nil {
		return "", fmt.Errorf("failed to get hotel name text: %w", err)
	}

	return h1Text, nil
}

func stdErr(err error) {
	_, _ = fmt.Fprintf(os.Stderr, err.Error())
}

func preparePageURL(val, language, location string) (string, error) {
	u, err := url.Parse(val)
	if err != nil {
		return "", fmt.Errorf("failed parsing the url: %v", err)
	}

	q := u.Query()
	q.Set("hl", language)
	q.Set("gl", location)

	u.RawQuery = q.Encode()

	return u.String(), nil
}
