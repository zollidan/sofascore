package main

import (
	"encoding/json"
	"fmt"

	"github.com/zollidan/sofascore/client"
	"github.com/zollidan/sofascore/excel"
	"github.com/zollidan/sofascore/utils"
)

// ============ MAIN ============

func main() {
	date, err := utils.GetUserDateChoice()
	if err != nil {
		fmt.Printf("❌ Ошибка: %v\n", err)
		return
	}

	client := client.NewAPIClient()
	games, err := client.FetchGames(date)
	if err != nil {
		fmt.Printf("❌ Ошибка загрузки: %v\n", err)
		return
	}

	jsonData, _ := json.Marshal(games)
	fmt.Printf("\nРазмер данных: %d байт (%.2f KB)\n", len(jsonData), float64(len(jsonData))/1024)

	excel.SaveExcel(games)
}
