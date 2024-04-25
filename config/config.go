package config

type Server struct {
	System           System    `mapstructure:"system" json:"system" yaml:"system"`
	GormMysqlDTaiwan GormMysql `mapstructure:"mysql_d_taiwan" json:"mysql_d_taiwan" yaml:"mysql_d_taiwan"`
	GormMysqlSystem  GormMysql `mapstructure:"mysql_system" json:"mysql_system" yaml:"mysql_system"`
	Zap              Zap       `mapstructure:"zap" json:"zap" yaml:"zap"`
	JWT              JWT       `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
