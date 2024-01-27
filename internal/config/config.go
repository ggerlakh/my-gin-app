package config

// GetDatabaseConfig возвращает строку подключения к базе данных.
func GetDatabaseConfig() string {
	// Получите конфигурацию из переменных окружения или конфигурационного файла.
	return "host=localhost user=new_user dbname=test password=783901 port=5432 sslmode=disable"
}
