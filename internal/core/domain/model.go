package domain

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type User struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string    `json:"firstname" binding:"required" gorm:"type:varchar(32)"`
	LastName  string    `json:"lastname" binding:"required" gorm:"type:varchar(32)"`
	CreatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Commerce struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `json:"commerce_name" binding:"required" gorm:"type:varchar(32)"`
	CreatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Branches struct {
	ID         uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name       string    `json:"branche_name" binding:"required" gorm:"type:varchar(32)"`
	Commerce   Commerce  `json:"Commerce_branche" binding:"required" gorm:"foreignkey:CommerceID"`
	CommerceID uint64    `json:"-"`
	CreatedAt  time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Campaign struct {
	ID               uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Name             string     `json:"campaign_name" binding:"required" gorm:"type:varchar(32)"`
	StartDate        *time.Time `json:"start_date" binding:"required" gorm:"type:time()"`
	EndDate          *time.Time `json:"end_date" binding:"required" gorm:"type:time()"`
	CampaignEquation string     `json:"campaign_equation" binding:"required" gorm:"type:varchar(256)"`
	Commerce         Commerce   `json:"Campaign_commerce" binding:"required" gorm:"foreignkey:CommerceID"`
	CommerceID       uint64     `json:"-"`
	Branches         Branches   `json:"campaign_commerce_branche" binding:"required" gorm:"foreignkey:BranchesID"`
	BranchesID       uint64     `json:"-"`
	CreatedAt        time.Time  `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time  `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type PurchaseTransaction struct {
	ID              uint64    `gorm:"primary_key;auto_increment" json:"id"`
	TransactionType string    `json:"transaction_type" binding:"required" gorm:"type:varchar(32)"`
	ItemType        string    `json:"item_type" binding:"required" gorm:"type:varchar(32)"`
	PurchaseValue   int8      `json:"purchase_value" binding:"required"`
	Commerce        Commerce  `json:"purchase_commerce" binding:"required" gorm:"foreignkey:CommerceID"`
	CommerceID      uint64    `json:"-"`
	Branches        Branches  `json:"purchase_commerce_branche" binding:"required" gorm:"foreignkey:BranchesID"`
	BranchesID      uint64    `json:"-"`
	User            User      `json:"user_iD" binding:"required" gorm:"foreignkey:UserID"`
	UserID          uint64    `json:"-"`
	CreatedAt       time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Awards struct {
	ID             uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name           string    `json:"awards_commerce_name" binding:"required" gorm:"type:varchar(32)"`
	AwardsEquation string    `json:"awards_equation" binding:"required" gorm:"type:varchar(256)"`
	Commerce       Commerce  `json:"awards_commerce" binding:"required" gorm:"foreignkey:CommerceID"`
	CommerceID     uint64    `json:"-"`
	Branches       Branches  `json:"awards_commerce_branche" binding:"required" gorm:"foreignkey:BranchesID"`
	BranchesID     uint64    `json:"-"`
	CreatedAt      time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type LealPoints struct {
	ID           uint64            `gorm:"primary_key;auto_increment" json:"id"`
	Value        int8              `json:"value" binding:"required"`
	VirtualValue pgtype.JSONBCodec `gorm:"type:jsonb;json:"virtual_value" binding:"required"`
	Commerce     Commerce          `json:"Campaign_commerce" binding:"required" gorm:"foreignkey:CommerceID"`
	CommerceID   uint64            `json:"-"`
	Branches     Branches          `json:"campaign_commerce_branche" binding:"required" gorm:"foreignkey:BranchesID"`
	BranchesID   uint64            `json:"-"`
	User         User              `json:"user_iD" binding:"required" gorm:"foreignkey:UserID"`
	UserID       uint64            `json:"-"`
	CreatedAt    time.Time         `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time         `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type CashBack struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Value     int8      `json:"value" binding:"required"`
	User      User      `json:"user_iD" binding:"required" gorm:"foreignkey:UserID"`
	UserID    uint64    `json:"-"`
	CreatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
