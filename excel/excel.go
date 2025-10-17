package excel

import (
	"fmt"
	"time"

	"github.com/xuri/excelize/v2"
	"github.com/zollidan/sofascore/models"
)

// ========= EXCEL ==========

func SaveExcel(games []models.Game) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Заголовки
	headers := []interface{}{"ID", "Дата/Время", "Домашняя команда", "Гостевая команда", "Счет дома", "Счет гостей", "Статус", "Турнир", "Тур"}
	if err := f.SetSheetRow("Sheet1", "A1", &headers); err != nil {
		fmt.Printf("❌ Ошибка записи заголовков: %v\n", err)
		return
	}

	// Данные
	for idx, game := range games {
		matchTime := time.Unix(game.StartTimestamp, 0).Format("02.01.2006 15:04")

		row := []interface{}{
			game.ID,
			matchTime,
			game.HomeTeamName,
			game.AwayTeamName,
			game.HomeScore,
			game.AwayScore,
			game.StatusType,
			game.TournamentName,
			game.Round,
		}

		if err := f.SetSheetRow("Sheet1", fmt.Sprintf("A%d", idx+2), &row); err != nil {
			fmt.Printf("❌ Ошибка записи строки %d: %v\n", idx+2, err)
			return
		}
	}

	// Сохранение файла
	filename := "sofascore_games.xlsx"
	if err := f.SaveAs(filename); err != nil {
		fmt.Printf("❌ Ошибка сохранения файла: %v\n", err)
		return
	}

	fmt.Printf("✅ Данные сохранены в %s\n", filename)
}