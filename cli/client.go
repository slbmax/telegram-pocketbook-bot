package cli

import (
	"encoding/json"
	errors "github.com/slbmax/telegram-pocketbook-bot/lib/custom-errors"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

const (
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func New(host, token string) Client {
	return Client{
		host:     host,
		basePath: "bot" + token,
		client:   http.Client{},
	}
}

func (c *Client) Updates(limit, offset int) ([]Update, error) {
	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))
	query.Add("offset", strconv.Itoa(offset))

	rawData, err := c.sendRequest(getUpdatesMethod, query)
	if err != nil {
		return nil, err
	}

	var res UpdatesResponse

	if err := json.Unmarshal(rawData, &res); err != nil {
		return nil, err
	}

	return res.Result, nil
}

func (c *Client) SendMessage(chatID int, text string) error {
	query := url.Values{}
	query.Add("chat_id", strconv.Itoa(chatID))
	query.Add("text", text)

	_, err := c.sendRequest(sendMessageMethod, query)

	return errors.WrapIfErr("can`t send message", err)
}

func (c *Client) sendRequest(method string, query url.Values) (data []byte, err error) {
	defer func() { err = errors.WrapIfErr("can't send request", err) }()

	url := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	request, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	request.URL.RawQuery = query.Encode()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
