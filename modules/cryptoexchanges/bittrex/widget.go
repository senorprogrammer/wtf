package bittrex

import (
	"encoding/json"
	"fmt"
	"time"

	"net/http"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

var ok = true
var errorText = ""

const baseURL = "https://bittrex.com/api/v1.1/public/getmarketsummary"

// Widget define wtf widget to register widget later
type Widget struct {
	wtf.TextWidget

	app      *tview.Application
	settings *Settings
	summaryList
}

// NewWidget Make new instance of widget
func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, pages, settings.common, false),

		app:         app,
		settings:    settings,
		summaryList: summaryList{},
	}

	widget.SetRefreshFunction(widget.Refresh)

	ok = true
	errorText = ""

	widget.setSummaryList()

	return &widget
}

func (widget *Widget) setSummaryList() {
	for symbol, currency := range widget.settings.summary.currencies {
		mCurrencyList := widget.makeSummaryMarketList(symbol, currency.market)
		widget.summaryList.addSummaryItem(symbol, currency.displayName, mCurrencyList)
	}
}

func (widget *Widget) makeSummaryMarketList(currencySymbol string, market []interface{}) []*mCurrency {
	mCurrencyList := []*mCurrency{}

	for _, marketSymbol := range market {
		mCurrencyList = append(mCurrencyList, makeMarketCurrency(marketSymbol.(string)))
	}

	return mCurrencyList
}

func makeMarketCurrency(name string) *mCurrency {
	return &mCurrency{
		name: name,
		summaryInfo: summaryInfo{
			High:           "",
			Low:            "",
			Volume:         "",
			Last:           "",
			OpenBuyOrders:  "",
			OpenSellOrders: "",
		},
	}
}

/* -------------------- Exported Functions -------------------- */

// Refresh & update after interval time
func (widget *Widget) Refresh() {
	widget.updateSummary()

	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) updateSummary() {
	// In case if anything bad happened!
	defer func() {
		recover()
	}()

	client := &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	for _, baseCurrency := range widget.summaryList.items {
		for _, mCurrency := range baseCurrency.markets {
			request := makeRequest(baseCurrency.name, mCurrency.name)
			response, err := client.Do(request)

			if err != nil {
				ok = false
				errorText = "Please Check Your Internet Connection!"
				break
			} else {
				ok = true
				errorText = ""
			}

			if response.StatusCode != http.StatusOK {
				errorText = response.Status
				ok = false
				break
			} else {
				ok = true
				errorText = ""
			}

			defer response.Body.Close()
			jsonResponse := summaryResponse{}
			decoder := json.NewDecoder(response.Body)
			decoder.Decode(&jsonResponse)

			if !jsonResponse.Success {
				ok = false
				errorText = fmt.Sprintf("%s-%s: %s", baseCurrency.name, mCurrency.name, jsonResponse.Message)
				break
			}
			ok = true
			errorText = ""

			mCurrency.Last = fmt.Sprintf("%f", jsonResponse.Result[0].Last)
			mCurrency.High = fmt.Sprintf("%f", jsonResponse.Result[0].High)
			mCurrency.Low = fmt.Sprintf("%f", jsonResponse.Result[0].Low)
			mCurrency.Volume = fmt.Sprintf("%f", jsonResponse.Result[0].Volume)
			mCurrency.OpenBuyOrders = fmt.Sprintf("%d", jsonResponse.Result[0].OpenBuyOrders)
			mCurrency.OpenSellOrders = fmt.Sprintf("%d", jsonResponse.Result[0].OpenSellOrders)
		}
	}

	widget.display()
}

func makeRequest(baseName, marketName string) *http.Request {
	url := fmt.Sprintf("%s?market=%s-%s", baseURL, baseName, marketName)
	request, _ := http.NewRequest("GET", url, nil)

	return request
}
