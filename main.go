package main

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"net/url"
	"os"
	"regexp"
	"strconv"
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

	// Uncomment after OK
	//hotelID, err := extractHotelID(BookingLink)
	//if err != nil {
	//	stdErr(err)
	//}
	//
	//hotelName, err := getPlaceName(page)
	//if err != nil {
	//	stdErr(err)
	//	return
	//}
	//
	//hotelDesc, err := getPlaceDescription(page)
	//if err != nil {
	//	stdErr(err)
	//	return
	//}
	//
	//bookingData := &booking.Data{
	//	RoomDetail: booking.RoomDetail{
	//		Property: booking.Property{
	//			AccommodationType:        booking.AccommodationType{},
	//			HasDesignatedSmokingArea: false,
	//			RoomsDetails:             nil,
	//			HighFloorStartsAt:        0,
	//			Name:                     hotelName,
	//			ID:                       hotelID,
	//			TypeName:                 "",
	//		},
	//	},
	//}
	//res, err := json.MarshalIndent(bookingData, "", "\t")
	//if err != nil {
	//	stdErr(err)
	//}
	//
	//if err := os.WriteFile("./scraped-booking.json", res, 0644); err != nil {
	//	stdErr(err)
	//}
	//
	//fmt.Printf("Hotel ID: %d", hotelID)
	//fmt.Printf("Hotel name: %s", hotelName)
	//fmt.Printf("Hotel description: %s", hotelDesc)

	time.Sleep(time.Hour * 1)
}

func getPlaceDescription(page *rod.Page) (string, error) {
	desc, err := page.Element("div.hp-description")
	if err != nil {
		return "", fmt.Errorf("failed to get hotel description: %w", err)
	}

	descText, err := desc.Text()
	if err != nil {
		return "", fmt.Errorf("failed to get hotel description text: %w", err)
	}

	return descText, nil
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

var labelIDRegex = regexp.MustCompile(`hotel-(\d+)`)

func extractHotelID(raw string) (int, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return 0, fmt.Errorf("parse url: %w", err)
	}

	label := u.Query().Get("label")
	m := labelIDRegex.FindStringSubmatch(label)
	if len(m) != 2 {
		return 0, fmt.Errorf("id not found in label: %s", label)
	}

	idInt, err := strconv.Atoi(m[1])
	if err != nil {
		return 0, fmt.Errorf("id could not be parsed to int: %w", err)
	}

	return idInt, nil
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
