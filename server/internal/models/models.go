package models

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/google/uuid"
	"github.com/rotisserie/eris"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type UserStatus string

const (
	PendingStatus    UserStatus = "Pending"
	VerifiedStatus   UserStatus = "Verified"
	NotAllowedStatus UserStatus = "Not allowed"
)

type UserAccountType string

const (
	AdminrAccountType    UserAccountType = "Admin"
	EmployeerAccountType UserAccountType = "Employee"
	MemberAccountType    UserAccountType = "Member"
	OwnerrAccountType    UserAccountType = "Owner"
)

type ModelTransformer struct {
	storage *providers.StorageProvider
}

type ModelRepository struct {
	db            *providers.DatabaseService
	helpers       *helpers.HelpersFunction
	cryptoHelpers *helpers.HelpersCryptography
}

var Module = fx.Module(
	"models",
	fx.Provide(
		func(storage *providers.StorageProvider) *ModelTransformer {
			return &ModelTransformer{storage: storage}
		},
		func(db *providers.DatabaseService, helpers *helpers.HelpersFunction, cryptoHelpers *helpers.HelpersCryptography) *ModelRepository {
			return &ModelRepository{db: db, helpers: helpers, cryptoHelpers: cryptoHelpers}
		},
	),
	fx.Invoke(func(
		db *providers.DatabaseService,
		logger *providers.LoggerService,
	) {
		err := db.Client.AutoMigrate(
			&Admin{},
			&Branch{},
			&Company{},
			&Contact{},
			&Employee{},
			&Feedback{},
			&Footstep{},
			&Gender{},
			&Media{},
			&Owner{},
			&Role{},

			// Member
			&Member{},
			&MemberProfile{},
			&MemberClassification{},
			&MemberClassificationHistory{},
			&MemberGender{},
			&MemberGenderHistory{},
			&MemberCenter{},
			&MemberCenterHistory{},
			&MemberGroup{},
			&MemberGroupHistory{},
			&MemberEducationalAttainment{},
			&MemberEducationalAttainmentHistory{},
			&MemberOccupation{},
			&MemberOccupationHistory{},
			&MemberType{},
			&MemberTypeHistory{},
			&MemberDescription{},
			&MemberRecruits{},
			&MemberContactNumberReferences{},
			&MemberWallet{},
			&MemberIncome{},
			&MemberExpenses{},
			&MemberCloseRemarks{},
			&MemberJointAccounts{},
			&MemberAddress{},
			&MemberGovernmentBenefits{},
			&MemberMutualFundsHistory{},
			&MemberAssets{},
			&MemberRelativeAccounts{},
			&MemberBranchRegistration{},
		)
		if err != nil {
			logger.Fatal("failed to migrate database")
		}
		logger.Info("Database migration completed successfully")
	}),
)

type GenericRepository[T any] struct {
	db *gorm.DB
}

// NewGenericRepository creates a new generic repository instance
func NewGenericRepository[T any](db *gorm.DB) *GenericRepository[T] {
	return &GenericRepository[T]{db: db}
}

// GetByID fetches a record by ID with optional preloads
func (r *GenericRepository[T]) GetByID(id string, preloads ...string) (*T, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, eris.Wrap(err, "invalid UUID")
	}

	var entity T
	query := r.db.Model(&entity).Where("id = ?", id)

	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	if err := query.First(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

// GetByColumn fetches a record by a specific column
func (r *GenericRepository[T]) GetByColumn(column string, value interface{}, preloads ...string) (*T, error) {
	var entity T
	query := r.db.Model(&entity).Where(column+" = ?", value)

	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	if err := query.First(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *GenericRepository[T]) UpdateByID(id string, column string, value interface{}, preloads ...string) (*T, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, eris.Wrap(err, "invalid UUID")
	}

	var entity T
	if err := r.db.Model(&entity).Where("id = ?", id).Update(column, value).Error; err != nil {
		return nil, err
	}
	query := r.db.Where("id = ?", id)
	for _, preload := range preloads {
		query = query.Preload(preload)
	}
	if err := query.First(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// DeleteByID deletes a record by ID
func (r *GenericRepository[T]) DeleteByID(id string) error {
	if err := r.db.Delete(new(T), id).Error; err != nil {
		return err
	}
	return nil
}

// Create inserts a new record and returns the created entity
func (r *GenericRepository[T]) Create(entity *T, preloads ...string) (*T, error) {
	if err := r.db.Create(entity).Error; err != nil {
		return nil, err
	}
	query := r.db.Model(entity)
	for _, preload := range preloads {
		query = query.Preload(preload)
	}
	if err := query.First(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

// GetAll fetches all records with optional preloads
func (r *GenericRepository[T]) GetAll(preloads ...string) ([]*T, error) {
	var entities []*T
	query := r.db.Model(new(T))

	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	if err := query.Find(&entities).Error; err != nil {
		return nil, err
	}

	return entities, nil
}

func (r *GenericRepository[T]) Update(entity *T, preloads ...string) (*T, error) {
	if err := r.db.Save(entity).Error; err != nil {
		return nil, err
	}

	query := r.db.Model(entity)
	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	if err := query.First(entity).Error; err != nil {
		return nil, err
	}

	return entity, nil
}
