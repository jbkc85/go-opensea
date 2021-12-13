package opensea

var (
	API_ENDPOINT = "https://api.opensea.io/api"
)

type Client struct {
	APIKey   string
	LogLevel string
	URL      string
}

func New(APIKey string, version string, logLevel string) *Client {
	return &Client{
		APIKey:   APIKey,
		LogLevel: logLevel,
		URL:      API_ENDPOINT + "/" + version,
	}
}
