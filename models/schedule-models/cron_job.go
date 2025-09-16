package schedulemodels

type CronJob struct {
	Endpoint string `bson:"endpoint" json:"endpoint"`
	IsActive bool   `bson:"isActive" json:"isActive"`
	Cron     string `bson:"cron" json:"cron"`
	Name     string `bson:"name" json:"name"`
}
