package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberWallet struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	MembersProfileID uuid.UUID `gorm:"unsigned" json:"members_profile_id"`
	Debit            float64   `gorm:"type:decimal(12,2);default:0" json:"debit"`
	Credit           float64   `gorm:"type:decimal(12,2);default:0" json:"credit"`
	Date             time.Time `gorm:"type:date" json:"date"`
	Description      string    `gorm:"type:text" json:"description"`

	MembersProfile *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}

type MemberWalletResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	MembersProfileID uuid.UUID              `json:"membersProfileID"`
	Debit            float64                `json:"debit"`
	Credit           float64                `json:"credit"`
	Date             string                 `json:"date"`
	Description      string                 `json:"description"`
	MembersProfile   *MemberProfileResource `json:"membersProfile,omitempty"`
}

func (m *ModelTransformer) MemberWalletToResource(wallet *MemberWallet) *MemberWalletResource {
	if wallet == nil {
		return nil
	}

	return &MemberWalletResource{

		ID:        wallet.ID,
		CreatedAt: wallet.CreatedAt.Format(time.RFC3339),
		UpdatedAt: wallet.UpdatedAt.Format(time.RFC3339),
		DeletedAt: wallet.DeletedAt.Time.Format(time.RFC3339),

		MembersProfileID: wallet.MembersProfileID,
		Debit:            wallet.Debit,
		Credit:           wallet.Credit,
		Date:             wallet.Date.Format("2006-01-02"),
		Description:      wallet.Description,
		MembersProfile:   m.MemberProfileToResource(wallet.MembersProfile),
	}
}

func (m *ModelTransformer) MemberWalletToResourceList(walletList []*MemberWallet) []*MemberWalletResource {
	if walletList == nil {
		return nil
	}

	var walletResources []*MemberWalletResource
	for _, wallet := range walletList {
		walletResources = append(walletResources, m.MemberWalletToResource(wallet))
	}
	return walletResources
}
