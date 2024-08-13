package service

import (
	"context"
	"demonstration-leal-test/internal/core/domain"
	"demonstration-leal-test/internal/core/port"
)

type LealService struct {
	repo port.LealRepository
}

// NewCategoryService creates a new category service instance
func NewLealService(repo port.LealRepository) *LealService {
	return &LealService{
		repo,
	}
}

func (l *LealService) SaveUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	user, err := l.repo.SaveUser(ctx, user)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}
	return user, nil
}

func (l *LealService) FindUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	user, err := l.repo.FindUser(ctx, user)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}
	return user, nil
}

//func (l *LealService) SaveCommerce(commerce domain.Commerce) (domain.Commerce, error) {
//	err := l.LealRepository.SaveCommerce(commerce)
//	return commerce, err
//}
//
//func (l *LealService) SaveCommerceBranches(branches domain.Branches) (domain.Branches, error) {
//	err := l.LealRepository.SaveCommerceBranches(branches)
//	return branches, err
//}
//
//func (l *LealService) FindCommerce(commerce domain.Commerce) (domain.Commerce, error) {
//	err := l.LealRepository.FindCommerce(commerce)
//	return commerce, err
//}
//
//func (l *LealService) SavePurchaseTransaction(purchaseTransaction domain.PurchaseTransaction) (domain.PurchaseTransaction, error) {
//	err := l.LealRepository.SavePurchaseTransaction(purchaseTransaction)
//	return purchaseTransaction, err
//}
//
//func (l *LealService) FindPurchaseTransaction(purchaseTransaction domain.PurchaseTransaction) (domain.PurchaseTransaction, error) {
//	err := l.LealRepository.FindPurchaseTransaction(purchaseTransaction)
//	return purchaseTransaction, err
//}
//
//func (l *LealService) SaveCampaign(campaign domain.Campaign) (domain.Campaign, error) {
//	err := l.LealRepository.SaveCampaign(campaign)
//	return campaign, err
//}
//
//func (l *LealService) FindCampaign(campaign domain.Campaign) (domain.Campaign, error) {
//	err := l.LealRepository.FindCampaign(campaign)
//	return campaign, err
//}
//
//func (l *LealService) SaveRedeem(lealPoints domain.LealPoints) (domain.LealPoints, error) {
//	err := l.LealRepository.SaveRedeem(lealPoints)
//	return lealPoints, err
//}
//
//func (l *LealService) FindRedeem(lealPoints domain.LealPoints) (domain.LealPoints, error) {
//	err := l.LealRepository.FindRedeem(lealPoints)
//	return lealPoints, err
//}
//
//func (l *LealService) SaveCampaignCommerce(campaign domain.Campaign) (domain.Campaign, error) {
//	err := l.LealRepository.SaveCampaignCommerce(campaign)
//	return campaign, err
//}
//
//func (l *LealService) FindCampaignCommerce(campaign domain.Campaign) (domain.Campaign, error) {
//	err := l.LealRepository.FindCampaignCommerce(campaign)
//	return campaign, err
//}
