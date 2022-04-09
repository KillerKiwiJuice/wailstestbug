package main

import (
	"log"

	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/tidwall/gjson"
)

var (
	DATA_SOCKET_ADDR string
	DATA_SOCKET_PORT string
	DATA_SOCKET_TYPE string

	AUTH_SOCKET_ADDR string
	AUTH_SOCKET_PORT string
	AUTH_INIT_KEY    string
	AUTH_SOCKET_TYPE string
	POLY_URL         string
	POLY_MSG         string
	POLY_CLUSTER     string
	POLY_KEY         string
)

const INTERVAL_MULTIPLIER = 100
const MINIMUM_TRADE_SIZE = 0

var old_auth_key = uuid.UUID{}
var auth_key = uuid.New()

var quotes_map = cmap.New()
var knminfiuasodba = gjson.Result{}
var nmaksndkjasnkudas = time.Time{}
var lmakjsndjkasnjkdnas = websocket.CloseError{}

type TradeMessage struct {
	Sym string  `json:"sym"`
	X   int64   `json:"x"`
	I   int64   `json:"i"`
	Z   int64   `json:"z"`
	P   float64 `json:"p"`
	BP  float64 `json:"bp"`
	BS  int64   `json:"bs"`
	BX  string  `json:"bx"`
	AP  float64 `json:"ap"`
	AS  int64   `json:"as"`
	AX  string  `json:"ax"`
	C   string  `json:"c"`
	S   int64   `json:"s"`
	T   int64   `json:"t"`
	A   string  `json:"a"`
}

type QuoteMessage struct {
	Sym string  `json:"sym"`
	Bx  string  `json:"bx"`
	Bp  float64 `json:"bp"`
	Bs  int64   `json:"bs"`
	Ax  string  `json:"ax"`
	Ap  float64 `json:"ap"`
	As  int64   `json:"as"`
	T   int64   `json:"t"`
	Z   int64   `json:"z"`
}

type ApplicationLog struct {
	Timestamp   int64  `json:"time"`
	MessageType string `json:"msgType"`
	Message     string `json:"msg"`
}

func process_poly_json2(json_pipe *NonBlockingChan, master_queue chan TradeMessage) {
	for {

		json_arr, ok := <-json_pipe.Recv
		if ok {
			// value was received
			//json_arr := json_arr.(gjson.Result)
			log.Println(json_arr)
			// process it...
		} else {
			// channel was closed
		}
	}
}

func recv_poly_data2(json_queue *NonBlockingChan, app *App) {

	// COMMENT OUT BELOW FUNCTION CODE TO FIX SIGSEV
	log.Printf("connecting to %s", "wss://testnet-explorer.binance.org/ws/block")
	c, _, err := websocket.DefaultDialer.Dial("wss://testnet-explorer.binance.org/ws/block", nil)
	if err != nil {
		log.Fatal("Internal error dial:", err)
	}
	log.Println(c)
}

func process_poly_json(json_pipe *NonBlockingChan) {
	for {
		//log.Println("pjs2")
		// continue
		json_arr, ok := <-json_pipe.Recv
		if ok {
			log.Println(json_arr)
		}
	}
}

func recv_poly_data(json_queue *NonBlockingChan, app *App) {
	for {
		// This send is not blocking. Allows the websocket to read without any overhead.
		// Essentially, it's a buffered pipe. Reads will occur when the reader is ready to read.
		// The pipe can fill up to any size that the process allows. This means high liquidity market
		// times will likely fill the buffer up since the json processing is too slow to consume all the
		// websocket messages that are being read above.
		//log.Println("wtf", parsed_str)
		json_queue.Send <- "test"
	}
}

func AllDownhillFromHere(app *App) {
	log.Println("wtf")
	gjson_nonblocking_channel := NewNonBlockingChannel()

	// Uncomment below for simple working test
	// go recv_poly_data(gjson_nonblocking_channel, app)
	// process_poly_json(gjson_nonblocking_channel)

	// Uncomment below for (broken) websocket test
	master_queue := make(chan TradeMessage)

	go process_poly_json2(gjson_nonblocking_channel, master_queue)
	recv_poly_data2(gjson_nonblocking_channel, app)
}
