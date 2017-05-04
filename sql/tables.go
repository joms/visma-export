package sql

// http://jinzhu.me/gorm/models.html#conventions

type HeaderLine struct {
    CustomerNo string `gorm:"column:CustomerNo;primary_key"`
    Name string `gorm:"column:Name"`
    Address string `gorm:"column:MailingAddress"`
    Zip string `gorm:"column:MailingZip"`
    City string `gorm:"column:MailingCity"`
    YourRef string `gorm:"column:YourRef"`
    Phone string `gorm:"column:Phone"`
    Email string `gorm:"column:Email"`
    EmployeeID int `gorm:"column:EmployeeID"`
    Days string `gorm:"column:Days"`
    ProjectNo string `gorm:"column:ProjectNo"`
}

type Article struct {
    Name string `gorm:"column:Name"`
    ArticleID int `gorm:"column:ArticleID;primary_key"`
}

func (Article) TableName() string {
  return "Articles"
}
