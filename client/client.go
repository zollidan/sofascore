package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zollidan/sofascore/config"
	"github.com/zollidan/sofascore/models"
	"github.com/zollidan/sofascore/parser"
)

// ============ API CLIENT ============

type APIClient struct {
	httpClient *http.Client
	baseURL    string
}

func NewAPIClient() *APIClient {
	return &APIClient{
		httpClient: &http.Client{Timeout: config.RequestTimeout},
		baseURL:    config.BaseAPIURL,
	}
}

func (c *APIClient) FetchGames(date string) ([]models.Game, error) {
	url := c.baseURL + date

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания запроса: %w", err)
	}

	c.setHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer resp.Body.Close()

	fmt.Printf("📡 Статус ответа: %d\n", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("неожиданный код ответа: %d", resp.StatusCode)
	}

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("ошибка парсинга JSON: %w", err)
	}

	events, ok := data["events"].([]interface{})
	if !ok || len(events) == 0 {
		return []models.Game{}, nil
	}

	fmt.Printf("✅ Получено событий: %d\n", len(events))

	return parser.ParseGames(events), nil
}

func (c *APIClient) setHeaders(req *http.Request) {
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "https://www.sofascore.com/")
}
