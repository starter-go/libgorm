package libgorm

// Configuration 数据源配置信息
type Configuration struct {
	Name     string // 数据源的名称, 通常用于属性中,例如: 'datasource.{name}.driver'
	Driver   string
	User     string
	Password string
	Host     string
	Port     int
	Database string
	Enabled  bool
}
