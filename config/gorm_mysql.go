package config

type GormMysql struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (m *GormMysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

func (m *GormMysql) GetLogMode() string {
	return m.LogMode
}
