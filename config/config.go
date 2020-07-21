package config

// Enter Spreadsheet id from Google Sheet Url
var SpreadsheetId = "1IWEtAXU6m083he9PIcVp9nVegEfO2aG1yslWOacb09E"
var trafficSourcesShortlisted []string

// This toogle will skip traffic source based filtering
var TrafficSourceFilteringEnabled = false

//This toggle will help to fetch only ACTIVE/ALL/ACHIVED traffic sources
var IncludeTrafficSources = "ACTIVE"

//This toggle is Traffic Sources Config which should be pushed to Google Sheets
func GetTrafficSourcesShortlisted() []string {
	trafficSourcesShortlisted = []string{"Advertizer", "Facebook", "Zeropark"}
	return trafficSourcesShortlisted
}
