package config

// GetDatabaseConfig возвращает строку подключения к базе данных.
func GetDatabaseConfig() string {
	// Получите конфигурацию из переменных окружения или конфигурационного файла.
	return "host=localhost user=your_user_name dbname=your_db_name password=your_password port=5432 sslmode=disable"
}
