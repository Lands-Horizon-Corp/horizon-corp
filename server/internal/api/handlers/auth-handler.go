package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rotisserie/eris"
)

type AuthHandler struct {
	repository    *models.ModelRepository
	transformer   *models.ModelTransformer
	currentUser   *CurrentUser
	otpService    *providers.OTPService
	tokenProvider *providers.TokenService
	cfg           *config.AppConfig
	helpers       *helpers.HelpersFunction
	smsProvder    *providers.SMSService
	emailProvider *providers.EmailService
}

func NewAuthHandler(
	repository *models.ModelRepository,
	transformer *models.ModelTransformer,
	currentUser *CurrentUser,
	otpService *providers.OTPService,
	tokenProvider *providers.TokenService,
	cfg *config.AppConfig,
	helpers *helpers.HelpersFunction,
	smsProvder *providers.SMSService,
	emailProvider *providers.EmailService,
) *AuthHandler {
	return &AuthHandler{
		repository:    repository,
		transformer:   transformer,
		currentUser:   currentUser,
		otpService:    otpService,
		tokenProvider: tokenProvider,
		cfg:           cfg,
		helpers:       helpers,
		smsProvder:    smsProvder,
		emailProvider: emailProvider,
	}
}
func (ah *AuthHandler) ChangePassword(ctx *gin.Context, accountType, oldPassword, newPassword, confirmPassword string) (interface{}, error) {
	if newPassword != confirmPassword {
		return nil, eris.New("new password and confirm password do not match")
	}

	switch accountType {
	case "Admin":
		admin, err := ah.currentUser.Admin(ctx)
		if err != nil {
			return nil, eris.Wrap(err, "failed to retrieve current admin user")
		}
		updated, err := ah.repository.AdminChangePassword(
			admin.ID.String(),
			oldPassword,
			newPassword,
			ah.helpers.GetPreload(ctx)...,
		)
		if err != nil {
			return nil, eris.Wrap(err, "failed to change password for admin")
		}
		return ah.transformer.AdminToResource(updated), nil

	case "Member":
		member, err := ah.currentUser.Member(ctx)
		if err != nil {
			return nil, eris.Wrap(err, "failed to retrieve current member user")
		}
		updated, err := ah.repository.MemberChangePassword(
			member.ID.String(),
			oldPassword,
			newPassword,
			ah.helpers.GetPreload(ctx)...,
		)
		if err != nil {
			return nil, eris.Wrap(err, "failed to change password for member")
		}
		return ah.transformer.MemberToResource(updated), nil
	case "Employee":
		employee, err := ah.currentUser.Employee(ctx)
		if err != nil {
			return nil, eris.Wrap(err, "failed to retrieve current employee user")
		}
		updated, err := ah.repository.EmployeeChangePassword(
			employee.ID.String(),
			oldPassword,
			newPassword,
			ah.helpers.GetPreload(ctx)...,
		)
		if err != nil {
			return nil, eris.Wrap(err, "failed to change password for employee")
		}
		return ah.transformer.EmployeeToResource(updated), nil
	case "Owner":
		owner, err := ah.currentUser.Owner(ctx)
		if err != nil {
			return nil, eris.Wrap(err, "failed to retrieve current owner user")
		}
		updated, err := ah.repository.OwnerChangePassword(
			owner.ID.String(),
			oldPassword,
			newPassword,
			ah.helpers.GetPreload(ctx)...,
		)
		if err != nil {
			return nil, eris.Wrap(err, "failed to change password for owner")
		}
		return ah.transformer.OwnerToResource(updated), nil
	default:
		return nil, eris.New("invalid account type")
	}
}

func (ah *AuthHandler) Create(
	ctx *gin.Context,
	accountType, firstName, lastName, middleName, permanentAddress, description string,
	birthDate time.Time, username, email, password, contactNumber string, mediaID, roleID, genderID *uuid.UUID,
	emailTemplate, contactTemplate string,
) (interface{}, error) {
	switch accountType {
	case "Admin":
		admin, err := ah.repository.AdminCreate(&models.Admin{
			FirstName:          firstName,
			LastName:           lastName,
			MiddleName:         middleName,
			PermanentAddress:   permanentAddress,
			Description:        description,
			BirthDate:          birthDate,
			Username:           username,
			Email:              email,
			Password:           password,
			ContactNumber:      contactNumber,
			IsEmailVerified:    false,
			IsContactVerified:  false,
			IsSkipVerification: false,
			Status:             providers.NotAllowedStatus,
			MediaID:            mediaID,
			RoleID:             roleID,
			GenderID:           genderID,
		}, ah.helpers.GetPreload(ctx)...)
		if err != nil {
			return nil, eris.Wrap(err, "failed to create admin")
		}
		if err := ah.sendVerificationOTPs("Admin", admin.ID.String(), email, contactNumber, firstName, lastName, emailTemplate, contactTemplate); err != nil {
			return nil, err
		}
		return ah.transformer.AdminToResource(admin), nil
	case "Member":
		member, err := ah.repository.MemberCreate(&models.Member{
			FirstName:          firstName,
			LastName:           lastName,
			MiddleName:         middleName,
			PermanentAddress:   permanentAddress,
			Description:        description,
			BirthDate:          birthDate,
			Username:           username,
			Email:              email,
			Password:           password,
			ContactNumber:      contactNumber,
			IsEmailVerified:    false,
			IsContactVerified:  false,
			IsSkipVerification: false,
			Status:             providers.NotAllowedStatus,
			MediaID:            mediaID,
			RoleID:             roleID,
			GenderID:           genderID,
		}, ah.helpers.GetPreload(ctx)...)
		if err != nil {
			return nil, eris.Wrap(err, "failed to create member")
		}
		if err := ah.sendVerificationOTPs("Member", member.ID.String(), email, contactNumber, firstName, lastName, emailTemplate, contactTemplate); err != nil {
			return nil, err
		}
		return ah.transformer.MemberToResource(member), nil
	case "Employee":
		employee, err := ah.repository.EmployeeCreate(&models.Employee{
			FirstName:          firstName,
			LastName:           lastName,
			MiddleName:         middleName,
			PermanentAddress:   permanentAddress,
			Description:        description,
			BirthDate:          birthDate,
			Username:           username,
			Email:              email,
			Password:           password,
			ContactNumber:      contactNumber,
			IsEmailVerified:    false,
			IsContactVerified:  false,
			IsSkipVerification: false,
			Status:             providers.NotAllowedStatus,
			MediaID:            mediaID,
			RoleID:             roleID,
			GenderID:           genderID,
		}, ah.helpers.GetPreload(ctx)...)
		if err != nil {
			return nil, eris.Wrap(err, "failed to create employee")
		}

		if err := ah.sendVerificationOTPs("Employee", employee.ID.String(), email, contactNumber, firstName, lastName, emailTemplate, contactTemplate); err != nil {
			return nil, err
		}
		return ah.transformer.EmployeeToResource(employee), nil

	case "Owner":
		owner, err := ah.repository.OwnerCreate(&models.Owner{
			FirstName:          firstName,
			LastName:           lastName,
			MiddleName:         middleName,
			PermanentAddress:   permanentAddress,
			Description:        description,
			BirthDate:          birthDate,
			Username:           username,
			Email:              email,
			Password:           password,
			ContactNumber:      contactNumber,
			IsEmailVerified:    false,
			IsContactVerified:  false,
			IsSkipVerification: false,
			Status:             providers.NotAllowedStatus,
			MediaID:            mediaID,
			RoleID:             roleID,
			GenderID:           genderID,
		}, ah.helpers.GetPreload(ctx)...)
		if err != nil {
			return nil, eris.Wrap(err, "failed to create owner")
		}
		if err := ah.sendVerificationOTPs("Owner", owner.ID.String(), email, contactNumber, firstName, lastName, emailTemplate, contactTemplate); err != nil {
			return nil, err
		}
		return ah.transformer.OwnerToResource(owner), nil
	default:
		return nil, eris.New("invalid account type")
	}
}

func (ah *AuthHandler) ForgotPasswordResetLink(ctx *gin.Context, accountType, key, emailTemplate, contactTemplate string) (string, error) {
	var userID, firstName, lastName string

	switch accountType {
	case "Admin":
		user, err := ah.repository.AdminSearch(key)
		if err != nil {
			return "", err
		}
		userID = user.ID.String()
		firstName = user.FirstName
		lastName = user.LastName
	case "Member":
		user, err := ah.repository.MemberSearch(key)
		if err != nil {
			return "", err
		}
		userID = user.ID.String()
		firstName = user.FirstName
		lastName = user.LastName
	case "Employee":
		user, err := ah.repository.EmployeeSearch(key)
		if err != nil {
			return "", err
		}
		userID = user.ID.String()
		firstName = user.FirstName
		lastName = user.LastName
	case "Owner":
		user, err := ah.repository.OwnerSearch(key)
		if err != nil {
			return "", err
		}
		userID = user.ID.String()
		firstName = user.FirstName
		lastName = user.LastName
	default:
		return "", eris.New("invalid account type")
	}

	token, err := ah.tokenProvider.GenerateUserToken(providers.UserClaims{
		ID:          userID,
		AccountType: accountType,
		UserStatus:  "active",
	}, time.Minute*10)
	if err != nil {
		return "", eris.Wrap(err, "failed to generate reset token")
	}

	resetLink := fmt.Sprintf("%s/auth/password-reset/%s", ah.cfg.AppClientUrl, *token)
	keyType := ah.helpers.GetKeyType(key)
	switch keyType {
	case "contact":
		if err := ah.sendSMS(key, contactTemplate, firstName, lastName, resetLink); err != nil {
			return "", err
		}
	case "email":
		if err := ah.sendEmail(key, emailTemplate, firstName, lastName, resetLink); err != nil {
			return "", err
		}
	default:
		return "", eris.New("invalid key type")
	}
	return resetLink, nil
}

func (ah *AuthHandler) NewPassword(ctx *gin.Context, accountType, currentPassword, confirmPassword string) (interface{}, error) {
	switch accountType {
	case "Admin":
		admin, err := ah.currentUser.Admin(ctx)
		if err != nil {
			return nil, eris.New("invalid account type")
		}
		updatedUser, err := ah.repository.AdminChangePassword(
			admin.ID.String(), currentPassword, confirmPassword,
			ah.helpers.GetPreload(ctx)...,
		)
		return ah.transformer.AdminToResource(updatedUser), err

	case "Member":
		member, err := ah.currentUser.Member(ctx)
		if err != nil {
			return nil, eris.New("invalid account type")
		}
		updatedUser, err := ah.repository.MemberChangePassword(
			member.ID.String(), currentPassword, confirmPassword,
			ah.helpers.GetPreload(ctx)...,
		)
		return ah.transformer.MemberToResource(updatedUser), err

	case "Employee":
		employee, err := ah.currentUser.Employee(ctx)
		if err != nil {
			return nil, eris.New("invalid account type")
		}
		updatedUser, err := ah.repository.EmployeeChangePassword(
			employee.ID.String(), currentPassword, confirmPassword,
			ah.helpers.GetPreload(ctx)...,
		)
		return ah.transformer.EmployeeToResource(updatedUser), err

	case "Owner":
		owner, err := ah.currentUser.Owner(ctx)
		if err != nil {
			return nil, eris.New("invalid account type")
		}
		updatedUser, err := ah.repository.OwnerChangePassword(
			owner.ID.String(), currentPassword, confirmPassword,
			ah.helpers.GetPreload(ctx)...,
		)
		return ah.transformer.OwnerToResource(updatedUser), err

	default:
		return nil, eris.New("invalid key type")
	}
}

func (ah *AuthHandler) ProfileAccountSetting(ctx *gin.Context, accountType string, birthDate time.Time, firstName, middleName, lastName, description, permanentAddress string) (interface{}, error) {
	switch accountType {
	case "Admin":
		admin, err := ah.currentUser.Admin(ctx)
		if err != nil {
			return nil, eris.New("invalid account type")
		}
		userUpdated, err := ah.repository.AdminUpdateByID(admin.ID.String(), &models.Admin{
			BirthDate:        birthDate,
			MiddleName:       middleName,
			FirstName:        firstName,
			LastName:         lastName,
			Description:      description,
			PermanentAddress: permanentAddress,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.AdminToResource(userUpdated), err

	case "Member":
		member, err := ah.currentUser.Member(ctx)
		if err != nil {
			return nil, eris.New("invalid account type")
		}
		userUpdated, err := ah.repository.MemberUpdateByID(member.ID.String(), &models.Member{
			BirthDate:        birthDate,
			MiddleName:       middleName,
			FirstName:        firstName,
			LastName:         lastName,
			Description:      description,
			PermanentAddress: permanentAddress,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.MemberToResource(userUpdated), err

	case "Employee":
		employee, err := ah.currentUser.Employee(ctx)
		if err != nil {
			return nil, eris.New("invalid account type")
		}
		userUpdated, err := ah.repository.EmployeeUpdateByID(employee.ID.String(), &models.Employee{
			BirthDate:        birthDate,
			MiddleName:       middleName,
			FirstName:        firstName,
			LastName:         lastName,
			Description:      description,
			PermanentAddress: permanentAddress,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.EmployeeToResource(userUpdated), err

	case "Owner":
		owner, err := ah.currentUser.Owner(ctx)
		if err != nil {
			return nil, eris.New("invalid account type")
		}
		userUpdated, err := ah.repository.OwnerUpdateByID(owner.ID.String(), &models.Owner{
			BirthDate:        birthDate,
			MiddleName:       middleName,
			FirstName:        firstName,
			LastName:         lastName,
			Description:      description,
			PermanentAddress: permanentAddress,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.OwnerToResource(userUpdated), err

	default:
		return nil, eris.New("invalid account type")
	}
}

func (ah *AuthHandler) ProfileChangeContactNumber(ctx *gin.Context, accountType string, password, contactNumber string) (interface{}, error) {
	switch accountType {
	case "Admin":
		user, err := ah.currentUser.Admin(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}
		if !ah.repository.AdminVerifyPassword(user.ID.String(), password) {
			return nil, eris.New("Wrong password")
		}
		updatedAdmin, err := ah.repository.AdminUpdateByID(user.ID.String(), &models.Admin{
			ContactNumber:     contactNumber,
			IsContactVerified: false,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.AdminToResource(updatedAdmin), err

	case "Member":
		member, err := ah.currentUser.Member(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}
		if !ah.repository.MemberVerifyPassword(member.ID.String(), password) {
			return nil, eris.New("Wrong password")
		}
		updatedMember, err := ah.repository.MemberUpdateByID(member.ID.String(), &models.Member{
			ContactNumber:     contactNumber,
			IsContactVerified: false,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.MemberToResource(updatedMember), err

	case "Employee":
		employee, err := ah.currentUser.Employee(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}
		if !ah.repository.EmployeeVerifyPassword(employee.ID.String(), password) {
			return nil, eris.New("Wrong password")
		}
		updatedEmployee, err := ah.repository.EmployeeUpdateByID(employee.ID.String(), &models.Employee{
			ContactNumber:     contactNumber,
			IsContactVerified: false,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.EmployeeToResource(updatedEmployee), err

	case "Owner":
		owner, err := ah.currentUser.Owner(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}
		if !ah.repository.OwnerVerifyPassword(owner.ID.String(), password) {
			return nil, eris.New("Wrong password")
		}
		updatedOwner, err := ah.repository.OwnerUpdateByID(owner.ID.String(), &models.Owner{
			ContactNumber:     contactNumber,
			IsContactVerified: false,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.OwnerToResource(updatedOwner), err

	default:
		return nil, eris.New("Invalid account type")
	}
}

func (ah *AuthHandler) ProfileChangeEmail(ctx *gin.Context, accountType string, password, email string) (interface{}, error) {
	switch accountType {
	case "Admin":
		user, err := ah.currentUser.Admin(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}
		if !ah.repository.AdminVerifyPassword(user.ID.String(), password) {
			return nil, eris.New("Wrong password")
		}
		updatedUser, err := ah.repository.AdminUpdateByID(user.ID.String(), &models.Admin{
			Email:           email,
			IsEmailVerified: false,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.AdminToResource(updatedUser), err

	case "Member":
		member, err := ah.currentUser.Member(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}
		if !ah.repository.MemberVerifyPassword(member.ID.String(), password) {
			return nil, eris.New("Wrong password")
		}
		updatedMember, err := ah.repository.MemberUpdateByID(member.ID.String(), &models.Member{
			Email:           email,
			IsEmailVerified: false,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.MemberToResource(updatedMember), err
	case "Employee":
		employee, err := ah.currentUser.Employee(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}
		if !ah.repository.EmployeeVerifyPassword(employee.ID.String(), password) {
			return nil, eris.New("Wrong password")
		}
		updatedEmployee, err := ah.repository.EmployeeUpdateByID(employee.ID.String(), &models.Employee{
			Email:           email,
			IsEmailVerified: false,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.EmployeeToResource(updatedEmployee), err

	case "Owner":
		owner, err := ah.currentUser.Owner(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}
		if !ah.repository.OwnerVerifyPassword(owner.ID.String(), password) {
			return nil, eris.New("Wrong password")
		}
		updatedOwner, err := ah.repository.OwnerUpdateByID(owner.ID.String(), &models.Owner{
			Email:           email,
			IsEmailVerified: false,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.OwnerToResource(updatedOwner), err

	default:
		return nil, eris.New("Invalid account type")
	}
}

func (ah *AuthHandler) ProfileChangeUsername(ctx *gin.Context, accountType string, password, username string) (interface{}, error) {
	switch accountType {
	case "Admin":
		user, err := ah.currentUser.Admin(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}
		if !ah.repository.AdminVerifyPassword(user.ID.String(), password) {
			return nil, eris.New("Wrong password")
		}
		updatedUser, err := ah.repository.AdminUpdateByID(user.ID.String(), &models.Admin{
			Username: username,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.AdminToResource(updatedUser), err
	case "Member":
		user, err := ah.currentUser.Member(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}
		if !ah.repository.MemberVerifyPassword(user.ID.String(), password) {
			return nil, eris.New("Wrong password")
		}
		updatedUser, err := ah.repository.MemberUpdateByID(user.ID.String(), &models.Member{
			Username: username,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.MemberToResource(updatedUser), err
	case "Employee":
		user, err := ah.currentUser.Employee(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}
		if !ah.repository.EmployeeVerifyPassword(user.ID.String(), password) {
			return nil, eris.New("Wrong password")
		}
		updatedUser, err := ah.repository.EmployeeUpdateByID(user.ID.String(), &models.Employee{
			Username: username,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.EmployeeToResource(updatedUser), err
	case "Owner":
		user, err := ah.currentUser.Owner(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}
		if !ah.repository.OwnerVerifyPassword(user.ID.String(), password) {
			return nil, eris.New("Wrong password")
		}
		updatedUser, err := ah.repository.OwnerUpdateByID(user.ID.String(), &models.Owner{
			Username: username,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.OwnerToResource(updatedUser), err
	default:
		return nil, eris.New("Invalid account type")
	}
}

func (ah *AuthHandler) ProfilePicture(ctx *gin.Context, accountType string, id *uuid.UUID) (interface{}, error) {
	switch accountType {
	case "Admin":
		user, err := ah.currentUser.Admin(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}
		updatedAdmin, err := ah.repository.AdminUpdateByID(user.ID.String(), &models.Admin{
			MediaID: id,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.AdminToResource(updatedAdmin), err
	case "Member":
		user, err := ah.currentUser.Member(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}
		updatedMember, err := ah.repository.MemberUpdateByID(user.ID.String(), &models.Member{
			MediaID: id,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.MemberToResource(updatedMember), err

	case "Employee":
		user, err := ah.currentUser.Employee(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}
		updatedEmployee, err := ah.repository.EmployeeUpdateByID(user.ID.String(), &models.Employee{
			MediaID: id,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.EmployeeToResource(updatedEmployee), err
	case "Owner":
		user, err := ah.currentUser.Owner(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}
		updatedOwner, err := ah.repository.OwnerUpdateByID(user.ID.String(), &models.Owner{
			MediaID: id,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.OwnerToResource(updatedOwner), err
	default:
		return nil, eris.New("Invalid account type")
	}
}

func (ah *AuthHandler) SendContactNumberVerification(ctx *gin.Context, accountType, contactTemplate string) error {
	switch accountType {
	case "Admin":
		user, err := ah.currentUser.Admin(ctx)
		if err != nil {
			return eris.New("User not authenticated")
		}
		otpMessage := providers.OTPMessage{
			AccountType: "Admin",
			ID:          user.ID.String(),
			MediumType:  providers.Email,
			Reference:   "send-contact-number-verification",
		}
		contactReq := providers.SMSRequest{
			To:   user.ContactNumber,
			Body: contactTemplate,
			Vars: &map[string]string{
				"name": fmt.Sprintf("%s %s", user.FirstName, user.LastName),
			},
		}
		return ah.otpService.SendContactNumberOTP(otpMessage, contactReq)
	case "Member":

		user, err := ah.currentUser.Member(ctx)
		if err != nil {
			return eris.New("User not authenticated")
		}
		otpMessage := providers.OTPMessage{
			AccountType: "Member",
			ID:          user.ID.String(),
			MediumType:  providers.Email,
			Reference:   "send-contact-number-verification",
		}
		contactReq := providers.SMSRequest{
			To:   user.ContactNumber,
			Body: contactTemplate,
			Vars: &map[string]string{
				"name": fmt.Sprintf("%s %s", user.FirstName, user.LastName),
			},
		}
		return ah.otpService.SendContactNumberOTP(otpMessage, contactReq)
	case "Employee":
		user, err := ah.currentUser.Employee(ctx)
		if err != nil {
			return eris.New("User not authenticated")
		}
		otpMessage := providers.OTPMessage{
			AccountType: "Employee",
			ID:          user.ID.String(),
			MediumType:  providers.Email,
			Reference:   "send-contact-number-verification",
		}
		contactReq := providers.SMSRequest{
			To:   user.ContactNumber,
			Body: contactTemplate,
			Vars: &map[string]string{
				"name": fmt.Sprintf("%s %s", user.FirstName, user.LastName),
			},
		}
		return ah.otpService.SendContactNumberOTP(otpMessage, contactReq)

	case "Owner":

		user, err := ah.currentUser.Owner(ctx)
		if err != nil {
			return eris.New("User not authenticated")
		}
		otpMessage := providers.OTPMessage{
			AccountType: "Owner",
			ID:          user.ID.String(),
			MediumType:  providers.Email,
			Reference:   "send-contact-number-verification",
		}
		contactReq := providers.SMSRequest{
			To:   user.ContactNumber,
			Body: contactTemplate,
			Vars: &map[string]string{
				"name": fmt.Sprintf("%s %s", user.FirstName, user.LastName),
			},
		}
		return ah.otpService.SendContactNumberOTP(otpMessage, contactReq)

	default:
		return eris.New("Invalid account type")
	}
}

func (ah *AuthHandler) SendEmailVerification(ctx *gin.Context, accountType string, emailTemplate string) error {
	switch accountType {
	case "Admin":
		admin, err := ah.currentUser.Admin(ctx)
		if err != nil {
			return eris.New("User not authenticated")
		}
		otpMessage := providers.OTPMessage{
			AccountType: "Admin",
			ID:          admin.ID.String(),
			MediumType:  providers.Email,
			Reference:   "send-email-verification",
		}
		emailRequest := providers.EmailRequest{
			To:      admin.Email,
			Subject: "ECOOP: Email Verification",
			Body:    emailTemplate,
		}
		return ah.otpService.SendEmailOTP(otpMessage, emailRequest)
	case "Member":
		member, err := ah.currentUser.Member(ctx)
		if err != nil {
			return eris.New("User not authenticated")
		}
		otpMessage := providers.OTPMessage{
			AccountType: "Member",
			ID:          member.ID.String(),
			MediumType:  providers.Email,
			Reference:   "send-email-verification",
		}
		emailRequest := providers.EmailRequest{
			To:      member.Email,
			Subject: "ECOOP: Email Verification",
			Body:    emailTemplate,
		}
		return ah.otpService.SendEmailOTP(otpMessage, emailRequest)

	case "Employee":
		employee, err := ah.currentUser.Employee(ctx)
		if err != nil {
			return eris.New("User not authenticated")
		}
		otpMessage := providers.OTPMessage{
			AccountType: "Employee",
			ID:          employee.ID.String(),
			MediumType:  providers.Email,
			Reference:   "send-email-verification",
		}
		emailRequest := providers.EmailRequest{
			To:      employee.Email,
			Subject: "ECOOP: Email Verification",
			Body:    emailTemplate,
		}
		return ah.otpService.SendEmailOTP(otpMessage, emailRequest)

	case "Owner":
		owner, err := ah.currentUser.Owner(ctx)
		if err != nil {
			return eris.New("User not authenticated")
		}
		otpMessage := providers.OTPMessage{
			AccountType: "Owner",
			ID:          owner.ID.String(),
			MediumType:  providers.Email,
			Reference:   "send-email-verification",
		}
		emailRequest := providers.EmailRequest{
			To:      owner.Email,
			Subject: "ECOOP: Email Verification",
			Body:    emailTemplate,
		}
		return ah.otpService.SendEmailOTP(otpMessage, emailRequest)
	default:
		return eris.New("Invalid account type")
	}
}

func (ah *AuthHandler) SkipVerification(ctx *gin.Context, accountType string) (interface{}, error) {
	switch accountType {
	case "Admin":
		user, err := ah.currentUser.Admin(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}

		updatedUser, err := ah.repository.AdminUpdateByID(user.ID.String(), &models.Admin{
			IsSkipVerification: true,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.AdminToResource(updatedUser), err
	case "Member":

		user, err := ah.currentUser.Member(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}

		updatedUser, err := ah.repository.MemberUpdateByID(user.ID.String(), &models.Member{
			IsSkipVerification: true,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.MemberToResource(updatedUser), err

	case "Employee":
		user, err := ah.currentUser.Employee(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}

		updatedUser, err := ah.repository.EmployeeUpdateByID(user.ID.String(), &models.Employee{
			IsSkipVerification: true,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.EmployeeToResource(updatedUser), err

	case "Owner":
		user, err := ah.currentUser.Owner(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}

		updatedUser, err := ah.repository.OwnerUpdateByID(user.ID.String(), &models.Owner{
			IsSkipVerification: true,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.OwnerToResource(updatedUser), err

	default:
		return nil, eris.New("Invalid account type")
	}
}

func (ah *AuthHandler) VerifyContactNumber(ctx *gin.Context, accountType string, otp string) (interface{}, error) {
	switch accountType {
	case "Admin":
		user, err := ah.currentUser.Admin(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}
		otpMessage := providers.OTPMessage{
			AccountType: "Admin",
			ID:          user.ID.String(),
			MediumType:  providers.Email,
			Reference:   "send-contact-number-verification",
		}
		isValid, err := ah.otpService.ValidateOTP(otpMessage, otp)
		if err != nil {
			return nil, eris.New("OTP validation error")
		}
		if !isValid {
			return nil, eris.New("Invalid or expired OTP")
		}
		updatedUser, err := ah.repository.AdminUpdateByID(user.ID.String(), &models.Admin{
			IsContactVerified: true,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.AdminToResource(updatedUser), err
	case "Member":
		user, err := ah.currentUser.Member(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}
		otpMessage := providers.OTPMessage{
			AccountType: "Member",
			ID:          user.ID.String(),
			MediumType:  providers.Email,
			Reference:   "send-contact-number-verification",
		}
		isValid, err := ah.otpService.ValidateOTP(otpMessage, otp)
		if err != nil {
			return nil, eris.New("OTP validation error")
		}
		if !isValid {
			return nil, eris.New("Invalid or expired OTP")
		}
		updatedUser, err := ah.repository.MemberUpdateByID(user.ID.String(), &models.Member{
			IsContactVerified: true,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.MemberToResource(updatedUser), err

	case "Employee":
		user, err := ah.currentUser.Employee(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}
		otpMessage := providers.OTPMessage{
			AccountType: "Employee",
			ID:          user.ID.String(),
			MediumType:  providers.Email,
			Reference:   "send-contact-number-verification",
		}
		isValid, err := ah.otpService.ValidateOTP(otpMessage, otp)
		if err != nil {
			return nil, eris.New("OTP validation error")
		}
		if !isValid {
			return nil, eris.New("Invalid or expired OTP")
		}
		updatedUser, err := ah.repository.EmployeeUpdateByID(user.ID.String(), &models.Employee{
			IsContactVerified: true,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.EmployeeToResource(updatedUser), err

	case "Owner":
		user, err := ah.currentUser.Owner(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}
		otpMessage := providers.OTPMessage{
			AccountType: "Owner",
			ID:          user.ID.String(),
			MediumType:  providers.Email,
			Reference:   "send-contact-number-verification",
		}
		isValid, err := ah.otpService.ValidateOTP(otpMessage, otp)
		if err != nil {
			return nil, eris.New("OTP validation error")
		}
		if !isValid {
			return nil, eris.New("Invalid or expired OTP")
		}
		updatedUser, err := ah.repository.OwnerUpdateByID(user.ID.String(), &models.Owner{
			IsContactVerified: true,
		}, ah.helpers.GetPreload(ctx)...)
		return ah.transformer.OwnerToResource(updatedUser), err

	default:
		return nil, eris.New("Invalid account type")
	}
}

func (ah *AuthHandler) VerifyEmail(ctx *gin.Context, accountType string, otp string) (interface{}, error) {
	switch accountType {
	case "Admin":
		user, err := ah.currentUser.Admin(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}
		otpMessage := providers.OTPMessage{
			AccountType: "Admin",
			ID:          user.ID.String(),
			MediumType:  providers.Email,
			Reference:   "send-email-verification",
		}
		isValid, err := ah.otpService.ValidateOTP(otpMessage, otp)
		if err != nil {
			return nil, eris.New("Failed to validate OTP")
		}
		if !isValid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired OTP"})
			return nil, eris.New("Invalid or expired OTP")
		}
		updatedUser, err := ah.repository.AdminUpdateByID(user.ID.String(), &models.Admin{
			IsEmailVerified: true,
		}, ah.helpers.GetPreload(ctx)...)
		if err != nil {
			return nil, eris.New("Failed to update admin details")
		}
		return ah.transformer.AdminToResource(updatedUser), nil
	case "Member":

		user, err := ah.currentUser.Member(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}
		otpMessage := providers.OTPMessage{
			AccountType: "Member",
			ID:          user.ID.String(),
			MediumType:  providers.Email,
			Reference:   "send-email-verification",
		}
		isValid, err := ah.otpService.ValidateOTP(otpMessage, otp)
		if err != nil {
			return nil, eris.New("Failed to validate OTP")
		}
		if !isValid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired OTP"})
			return nil, eris.New("Invalid or expired OTP")
		}
		updatedUser, err := ah.repository.MemberUpdateByID(user.ID.String(), &models.Member{
			IsEmailVerified: true,
		}, ah.helpers.GetPreload(ctx)...)
		if err != nil {
			return nil, eris.New("Failed to update member details")
		}
		return ah.transformer.MemberToResource(updatedUser), nil
	case "Employee":
		user, err := ah.currentUser.Employee(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}
		otpMessage := providers.OTPMessage{
			AccountType: "Employee",
			ID:          user.ID.String(),
			MediumType:  providers.Email,
			Reference:   "send-email-verification",
		}
		isValid, err := ah.otpService.ValidateOTP(otpMessage, otp)
		if err != nil {
			return nil, eris.New("Failed to validate OTP")
		}
		if !isValid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired OTP"})
			return nil, eris.New("Invalid or expired OTP")
		}
		updatedUser, err := ah.repository.EmployeeUpdateByID(user.ID.String(), &models.Employee{
			IsEmailVerified: true,
		}, ah.helpers.GetPreload(ctx)...)
		if err != nil {
			return nil, eris.New("Failed to update employee details")
		}
		return ah.transformer.EmployeeToResource(updatedUser), nil
	case "Owner":
		user, err := ah.currentUser.Owner(ctx)
		if err != nil {
			return nil, eris.New("User not authenticated")
		}
		otpMessage := providers.OTPMessage{
			AccountType: "Owner",
			ID:          user.ID.String(),
			MediumType:  providers.Email,
			Reference:   "send-email-verification",
		}
		isValid, err := ah.otpService.ValidateOTP(otpMessage, otp)
		if err != nil {
			return nil, eris.New("Failed to validate OTP")
		}
		if !isValid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired OTP"})
			return nil, eris.New("Invalid or expired OTP")
		}
		updatedUser, err := ah.repository.OwnerUpdateByID(user.ID.String(), &models.Owner{
			IsEmailVerified: true,
		}, ah.helpers.GetPreload(ctx)...)
		if err != nil {
			return nil, eris.New("Failed to update owner details")
		}
		return ah.transformer.OwnerToResource(updatedUser), nil
	default:
		return nil, eris.New("Invalid account type")
	}
}

func (ah *AuthHandler) SignIn(ctx *gin.Context, accountType string, key, password string) (interface{}, *string, error) {
	switch accountType {
	case "Admin":
		user, err := ah.repository.AdminSignIn(key, password, ah.helpers.GetPreload(ctx)...)
		if err != nil {
			return nil, nil, err
		}
		token, err := ah.tokenProvider.GenerateUserToken(providers.UserClaims{
			ID:          user.ID.String(),
			AccountType: accountType,
		}, 24*time.Hour)
		if err != nil {
			return nil, nil, eris.New("failed to generate token")
		}
		return ah.transformer.AdminToResource(user), token, nil
	case "Member":
		user, err := ah.repository.MemberSignIn(key, password, ah.helpers.GetPreload(ctx)...)
		if err != nil {
			return nil, nil, err
		}
		token, err := ah.tokenProvider.GenerateUserToken(providers.UserClaims{
			ID:          user.ID.String(),
			AccountType: accountType,
		}, 24*time.Hour)
		if err != nil {
			return nil, nil, eris.New("failed to generate token")
		}
		return ah.transformer.MemberToResource(user), token, nil
	case "Employee":
		user, err := ah.repository.EmployeeSignIn(key, password, ah.helpers.GetPreload(ctx)...)
		if err != nil {
			return nil, nil, err
		}
		token, err := ah.tokenProvider.GenerateUserToken(providers.UserClaims{
			ID:          user.ID.String(),
			AccountType: accountType,
		}, 24*time.Hour)
		if err != nil {
			return nil, nil, eris.New("failed to generate token")
		}
		return ah.transformer.EmployeeToResource(user), token, nil
	case "Owner":
		user, err := ah.repository.OwnerSignIn(key, password, ah.helpers.GetPreload(ctx)...)
		if err != nil {
			return nil, nil, err
		}
		token, err := ah.tokenProvider.GenerateUserToken(providers.UserClaims{
			ID:          user.ID.String(),
			AccountType: accountType,
		}, 24*time.Hour)
		if err != nil {
			return nil, nil, eris.New("failed to generate token")
		}
		return ah.transformer.OwnerToResource(user), token, nil
	default:
		return nil, nil, eris.New("Invalid account type")
	}
}

func (ah *AuthHandler) ForceChangePassword(ctx *gin.Context, accountType, key, newPassword string) (interface{}, error) {
	switch accountType {
	case "Admin":
		user, err := ah.repository.AdminForceChangePassword(key, newPassword, ah.helpers.GetPreload(ctx)...)
		if err != nil {
			return nil, err
		}
		return ah.transformer.AdminToResource(user), nil
	case "Member":
		user, err := ah.repository.MemberForceChangePassword(key, newPassword, ah.helpers.GetPreload(ctx)...)
		if err != nil {
			return nil, err
		}
		return ah.transformer.MemberToResource(user), nil
	case "Employee":
		user, err := ah.repository.EmployeeForceChangePassword(key, newPassword, ah.helpers.GetPreload(ctx)...)
		if err != nil {
			return nil, err
		}
		return ah.transformer.EmployeeToResource(user), nil
	case "Owner":
		user, err := ah.repository.OwnerForceChangePassword(key, newPassword, ah.helpers.GetPreload(ctx)...)
		if err != nil {
			return nil, err
		}
		return ah.transformer.OwnerToResource(user), nil
	default:
		return nil, eris.New("Invalid account type")
	}
}

func (ah *AuthHandler) sendVerificationOTPs(
	accountType, accountID, email, contactNumber, firstName, lastName, emailTemplate, contactTemplate string,
) error {
	if err := ah.otpService.SendEmailOTP(providers.OTPMessage{
		AccountType: accountType,
		ID:          accountID,
		MediumType:  "email",
		Reference:   "email-verification",
	}, providers.EmailRequest{
		To:      email,
		Subject: "ECOOP: Email Verification",
		Body:    emailTemplate,
	}); err != nil {
		return eris.Wrap(err, "failed to send email verification")
	}

	if err := ah.otpService.SendContactNumberOTP(providers.OTPMessage{
		AccountType: accountType,
		ID:          accountID,
		MediumType:  "sms",
		Reference:   "sms-verification",
	}, providers.SMSRequest{
		To:   contactNumber,
		Body: contactTemplate,
		Vars: &map[string]string{
			"name": fmt.Sprintf("%s %s", firstName, lastName),
		},
	}); err != nil {
		return eris.Wrap(err, "failed to send contact number verification")
	}

	return nil
}

func (ah *AuthHandler) sendSMS(to, template, firstName, lastName, resetLink string) error {
	contactReq := providers.SMSRequest{
		To:   to,
		Body: template,
		Vars: &map[string]string{
			"name":      fmt.Sprintf("%s %s", firstName, lastName),
			"eventLink": resetLink,
		},
	}
	return ah.smsProvder.SendSMS(contactReq)
}

// Helper method to send email
func (ah *AuthHandler) sendEmail(to, template, firstName, lastName, resetLink string) error {
	emailReq := providers.EmailRequest{
		To:      to,
		Subject: "ECOOP: Change Password Request",
		Body:    template,
		Vars: &map[string]string{
			"name":      fmt.Sprintf("%s %s", firstName, lastName),
			"eventLink": resetLink,
		},
	}
	return ah.emailProvider.SendEmail(emailReq)
}
