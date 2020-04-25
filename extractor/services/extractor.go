package services

import (
	"challenge/models"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

func ExtractParkingData() error {
	var parkings []*models.Parking

	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	c := colly.NewCollector(
		colly.Async(true),
	)
	c.WithTransport(t)

	url, err := getURL()
	if err != nil {
		return err
	}

	c.OnHTML("div.entry-content.clearfix", func(e *colly.HTMLElement) {
		parkingTime, err := getParkingTime(e.ChildText("code > a"))
		if err != nil {
			return
		}

		e.ForEach("tbody > tr", func(_ int, el *colly.HTMLElement) {
			name := el.ChildText("td:nth-child(1)")
			status := el.ChildText("td:nth-child(2)")
			freeSpaces, err := strconv.Atoi(el.ChildText("td:nth-child(3)"))
			if err != nil {
				return
			}

			parking := &models.Parking{
				Name:       name,
				Status:     status,
				FreeSpaces: freeSpaces,
				Time:       parkingTime,
			}
			parkings = append(parkings, parking)
		})
	})

	c.OnResponse(func(r *colly.Response) {
		err := createParkingHTML(r.Body)
		if err != nil {
			return
		}
	})

	err = c.Visit(url)
	if err != nil {
		return err
	}
	c.Wait()

	err = createParkingJSON(parkings)
	if err != nil {
		return err
	}

	fmt.Printf("Extraction completed: %s\n", url)
	return nil
}

func getURL() (string, error) {
	if os.Getenv("APP_ENV") == "offline" {
		dir, err := filepath.Abs("./")
		if err != nil {
			return "", err
		}
		return "file://" + dir + "/parking.html", nil
	}

	return "https://www.hipark.de/parkplatzbelegung/", nil
}

func getParkingTime(text string) (int64, error) {
	reg, err := regexp.Compile("[0-9]{2,4}")
	if err != nil {
		return 0, err
	}

	matches := reg.FindAllString(text, -1)
	formatted := fmt.Sprintf("%s-%s-%sT%s:%s:%s.000Z", matches[2], matches[1], matches[0], matches[3], matches[4], matches[5])
	t, err := time.Parse(time.RFC3339, formatted)
	if err != nil {
		return 0, err
	}

	unix := t.UnixNano() / int64(time.Millisecond)

	return unix, nil
}

func createParkingHTML(body []byte) error {
	err := ioutil.WriteFile("parking.html", body, 0777)
	if err != nil {
		return err
	}

	return nil
}

func createParkingJSON(parkings []*models.Parking) error {
	jsonBody, err := json.Marshal(parkings)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("parking.json", jsonBody, 0777)
	if err != nil {
		return err
	}

	return nil
}
