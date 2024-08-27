package HaruiBot

// DefaultSets when application work, read config here.
type DefaultSets struct {
	DataFolder         string
	DisableChannelUser bool // in some situation,
	// we may not want a channel (ignore / unknown user) to use our bot,
	//so this function can enable to stop work for channel user yet.
	DataBaseLocation string
}

var ApplicationSets DefaultSets

func init() {
	// sets used in .env, if config is none, them use default sets.
	ApplicationSets = DefaultSets{
		DataFolder:         "/data/",
		DisableChannelUser: false,
		DataBaseLocation:   "/data/database",
	}
}
