package notice

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func createBody(to string, content string) string {
	body := fmt.Sprintf(`{"to": "%s", "messages": [{"type": "text", "text": "%s"}]}`, to, content)
	return body
}

func Line(content string) error {
	// 環境変数の取得
	token := os.Getenv("LINE_TOKEN")
	to := os.Getenv("LINE_TO")
	if token == "" || to == "" {
		return fmt.Errorf("LINE_TOKEN or LINE_TO is empty")
	}

	// リクエストの作成
	endpoint := "https://api.line.me/v2/bot/message/push"
	body := createBody(to, content)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	// リクエストの送信
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// レスポンスの確認
	if resp.StatusCode >= 400 {
		return fmt.Errorf("status code is %d", resp.StatusCode)
	}

	return nil
}
