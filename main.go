package main

import (
	"fmt"

	"github.com/bhambri94/marcus/marcusConfig"
	"github.com/bhambri94/marcus/sheets"
	"github.com/bhambri94/marcus/voluum"
)

func main() {
	marcusConfig.SetConfig()
	fmt.Println(marcusConfig.Configurations.SpreadsheetId)

	values, _, SheetName := voluum.GetDailyVoluumReport()
	sheets.ClearSheet(SheetName)
	writeRange := SheetName + "!A1"
	sheets.BatchWrite(writeRange, values)
}
