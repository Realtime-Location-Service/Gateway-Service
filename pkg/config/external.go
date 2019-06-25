package config

// LoadExternalCfg ...
func LoadExternalCfg() {
	LoadPingCfg()
	LoadAuthCfg()
	LoadMetaCfg()
	LoadHistoryCfg()
}
