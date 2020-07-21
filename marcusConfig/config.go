package marcusConfig

type struct Configs {
	SpreadsheetId string
	TrafficSourcesShortlisted []string
	TrafficSourceFilteringEnabled bool
	IncludeTrafficSources string
	VoluumAccessId string
	VoluumAccessKey string
}

var Configurations Configs

func SetConfig(){
	input, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	Configurations = json.Unmarshal(input, &Configurations)
	fmt.Println(Configurations)
}