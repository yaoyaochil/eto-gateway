package config

type Server struct {
	System                 System    `mapstructure:"system" json:"system" yaml:"system"`
	GormMysqlDTaiwan       GormMysql `mapstructure:"mysql_d_taiwan" json:"mysql_d_taiwan" yaml:"mysql_d_taiwan"`
	GormMysqlTaiwanCain    GormMysql `mapstructure:"mysql_taiwan_cain" json:"mysql_taiwan_cain" yaml:"mysql_taiwan_cain"`
	GormMysqlTaiwanCain2ND GormMysql `mapstructure:"mysql_taiwan_cain_2nd" json:"mysql_taiwan_cain_2nd" yaml:"mysql_taiwan_cain_2nd"`
	GormMysqlSystem        GormMysql `mapstructure:"mysql_system" json:"mysql_system" yaml:"mysql_system"`
	GormMysqlTaiwanLogin   GormMysql `mapstructure:"mysql_taiwan_login" json:"mysql_taiwan_login" yaml:"mysql_taiwan_login"`
	GormMysqlTaiwanBilling GormMysql `mapstructure:"mysql_taiwan_billing" json:"mysql_taiwan_billing" yaml:"mysql_taiwan_billing"`
	Zap                    Zap       `mapstructure:"zap" json:"zap" yaml:"zap"`
	JWT                    JWT       `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
