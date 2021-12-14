package database

func Migrate() {
	DB.AutoMigrate()
}
