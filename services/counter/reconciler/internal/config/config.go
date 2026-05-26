package config

type Config struct {
	Name  string
	Redis RedisConf
	Scan  ScanConf
}

type RedisConf struct {
	Host string
	Pass string `json:",optional"`
	Type string `json:",default=node"`
}

type ScanConf struct {
	IntervalHours     int     `json:",default=1"`
	BatchSize         int     `json:",default=256"`
	BatchIntervalMs   int     `json:",default=10"`
	ThresholdAbsolute int64   `json:",default=100"`
	ThresholdPercent  float64 `json:",default=1"`
}
