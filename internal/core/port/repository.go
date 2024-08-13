package port

import (
	"context"
	"demonstration-leal-test/internal/core/domain"
)

type LealRepository interface {
	SaveUser(ctx context.Context, user *domain.User) (*domain.User, error)
	FindUser(ctx context.Context, user *domain.User) (*domain.User, error)
	//SaveCommerce(ctx context.Context, commerce *domain.Commerce) (*domain.Commerce, error)
	//SaveCommerceBranches(ctx context.Context, branches *domain.Branches) (*domain.Branches, error)
	//FindCommerce(ctx context.Context, commerce *domain.Commerce) (*domain.Commerce, error)
	//SavePurchaseTransaction(ctx context.Context, purchaseTransaction *domain.PurchaseTransaction) (*domain.PurchaseTransaction, error)
	//FindPurchaseTransaction(ctx context.Context, purchaseTransaction *domain.PurchaseTransaction) (*domain.PurchaseTransaction, error)
	//SaveCampaign(ctx context.Context, campaign *domain.Campaign) (*domain.Campaign, error)
	//FindCampaign(ctx context.Context, campaign *domain.Campaign) (*domain.Campaign, error)
	//SaveRedeem(ctx context.Context, lealPoints *domain.LealPoints) (*domain.LealPoints, error)
	//FindRedeem(ctx context.Context, lealPoints *domain.LealPoints) (*domain.LealPoints, error)
	//SaveCampaignCommerce(ctx context.Context, campaign *domain.Campaign) (*domain.Campaign, error)
	//FindCampaignCommerce(ctx context.Context, campaign *domain.Campaign) (*domain.Campaign, error)
	//FindAll() ([]domain.User, error)
	//Clean() error
}
