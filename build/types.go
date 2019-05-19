package build

type Build struct {
	Data   Data
	Result []byte
}

type Data struct {
	Plugins   []string
	Telemetry bool
}
