package postgres

import (
	"context"
	"demonstration-leal-test/internal/core/domain"
)

/**
 * CategoryRepository implements port.CategoryRepository interface
 * and provides an access to the postgres database
 */
type LealRepository struct {
	db *PostgresDB
}

// NewCategoryRepository creates a new category repository instance
func NewLealRepository(db *PostgresDB) *LealRepository {
	return &LealRepository{
		db,
	}
}

func (lr *LealRepository) SaveUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := lr.db.QueryBuilder.Insert("user").
		Columns("firstname", "lastname").
		Values(user.FirstName, user.LastName).
		Suffix("RETURNING *")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = lr.db.QueryRow(ctx, sql, args...).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errCode := lr.db.ErrorCode(err); errCode == "23505" {
			return nil, domain.ErrConflictingData
		}
		return nil, err
	}

	return user, nil
}

func (lr *LealRepository) FindUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

//func (lr *LealRepository) SaveCommerce(commerce domain.Commerce) (*domain.Commerce, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (lr *LealRepository) SaveCommerceBranches(branches domain.Branches) (*domain.Branches, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (lr *LealRepository) FindCommerce(commerce domain.Commerce) (*domain.Commerce, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (lr *LealRepository) SavePurchaseTransaction(transaction domain.PurchaseTransaction) (*domain.PurchaseTransaction, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (lr *LealRepository) FindPurchaseTransaction(transaction domain.PurchaseTransaction) (*domain.PurchaseTransaction, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (lr *LealRepository) SaveCampaign(campaign domain.Campaign) (*domain.Campaign, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (lr *LealRepository) FindCampaign(campaign domain.Campaign) (*domain.Campaign, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (lr *LealRepository) SaveRedeem(points domain.LealPoints) (*domain.LealPoints, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (lr *LealRepository) FindRedeem(points domain.LealPoints) (*domain.LealPoints, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (lr *LealRepository) SaveCampaignCommerce(campaign domain.Campaign) (*domain.Campaign, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (lr *LealRepository) FindCampaignCommerce(campaign domain.Campaign) (*domain.Campaign, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (lr *LealRepository) Clean() error {
//	//TODO implement me
//	panic("implement me")
//}
//
