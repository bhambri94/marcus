package main

import (
	"github.com/bhambri94/marcus/config"
	"github.com/bhambri94/marcus/sheets"
	"github.com/bhambri94/marcus/voluum"
)

func main() {
	config.SetConfig()
	values, _, SheetName := voluum.GetDailyVoluumReport()
	sheets.ClearSheet(SheetName)
	writeRange := SheetName + "!A1"
	sheets.BatchWrite(writeRange, values)
}
