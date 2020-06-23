package amq

type Config struct {
	Addr             string `mapstructure:"addr"`
	Host             string `mapstructure:"host"`
	UserName         string `mapstructure:"username"`
	Password         string `mapstructure:"password"`
	DestinationName  string `mapstructure:"destination_name"`
	SubscriptionName string `mapstructure:"subscription_name"`
}
