package controllers

import (
	"log"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

type response_vendor struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Record      interface{} `json:"record"`
	Perpage     int         `json:"perpage"`
	Totalrecord int         `json:"totalrecord"`
	Time        string      `json:"time"`
}

func Vendorhome(c *fiber.Ctx) error {
	type payload_vendorhome struct {
		Vendor_search string `json:"vendor_search"`
		Vendor_page   int    `json:"vendor_page"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_vendorhome)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(response_vendor{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"vendor_search":   client.Vendor_search,
			"vendor_page":     client.Vendor_page,
		}).
		Post(PATH + "api/vendor")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*response_vendor)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":      result.Status,
			"message":     result.Message,
			"record":      result.Record,
			"perpage":     result.Perpage,
			"totalrecord": result.Totalrecord,
			"time":        time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func VendorSave(c *fiber.Ctx) error {
	type payload_vendorsave struct {
		Page          string `json:"page"`
		Sdata         string `json:"sdata" `
		Vendor_search string `json:"vendor_search" `
		Vendor_page   int    `json:"vendor_page" `
		Vendor_id     string `json:"vendor_id" `
		Vendor_name   string `json:"vendor_name" `
		Vendor_alamat string `json:"vendor_alamat" `
		Vendor_phone1 string `json:"vendor_phone1" `
		Vendor_phone2 string `json:"vendor_phone2" `
		Vendor_status string `json:"vendor_status" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_vendorsave)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"page":            client.Page,
			"sdata":           client.Sdata,
			"vendor_search":   client.Vendor_search,
			"vendor_page":     client.Vendor_page,
			"vendor_id":       client.Vendor_id,
			"vendor_name":     client.Vendor_name,
			"vendor_alamat":   client.Vendor_alamat,
			"vendor_phone1":   client.Vendor_phone1,
			"vendor_phone2":   client.Vendor_phone2,
			"vendor_status":   client.Vendor_status,
		}).
		Post(PATH + "api/vendorsave")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
