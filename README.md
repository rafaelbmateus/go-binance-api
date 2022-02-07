# go-binance-bot

Go Bincance Bot is a RESTful API.

| Verb | Route | Description |
|---|---|---|
| GET | /accounts/me | Get current account information |
| GET | /exchanges | Get exchanges info |
| GET | /symbols/:symbol/orders | Response list of orders successfully |
| POST | /symbols/:symbol/orders | New order created with successfully |
| GET | /symbols/:symbol/orders/:id | Response order successfully |
| DELETE | /symbols/:symbol/orders/:id | Order canceled with successfully |

For full handshake check the [OpenAPI Specification](docs/open-api.yaml).

## Setup Environment

Create a env file in root folder with your
[binance api key](https://www.binance.com/en/support/faq/360002502072)
like this:

```bash
BINANCE_API_KEY=<YOU_API_KEY_HERE>
BINANCE_API_SECRET=<YOU_API_SECRET_HERE>
```

## How to use?

To up the containers run this make command:

```bash
make up
```

Will be available in [localhost:3000](http://localhost:3000)

### Get 

## Binance API

https://binance-docs.github.io/apidocs/#change-log

https://github.com/edeng23/binance-trade-bot

https://github.com/chrisleekr/binance-trading-bot
