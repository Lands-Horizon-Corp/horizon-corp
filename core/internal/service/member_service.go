// package service

package service

import (
	"horizon-core/internal/events"
	"horizon-core/internal/models"
	"horizon-core/internal/repository"
	"horizon-core/internal/resources"

	"gorm.io/gorm"
)

type MemberService struct {
	*repository.ModelRepository[models.Member]
}

func NewMemberService(db *gorm.DB) *MemberService {
	return &MemberService{
		ModelRepository: repository.NewModelRepository[models.Member](db),
	}
}

func (r *MemberService) GetMemberByID(id string) (resources.MemberResource, error) {
	eagerLoads := []string{
		"ProfilePicture",
		"Roles.Permissions",
		"Branch.ProfilePicture",
	}

	member, err := r.FindByID(id, eagerLoads)
	if err != nil {
		return resources.MemberResource{}, err
	}

	memberResource := resources.NewMemberResource(member)
	return memberResource, nil
}

func (r *MemberService) ListMembers(req repository.ListRequest) (repository.ListResponse[resources.MemberResource], error) {
	eagerLoads := []string{
		"ProfilePicture",
		"Roles.Permissions",
		"Branch.ProfilePicture",
	}

	listResponse, err := r.List(req, eagerLoads)
	if err != nil {
		return repository.ListResponse[resources.MemberResource]{}, err
	}

	memberResources := make([]resources.MemberResource, len(listResponse.Data))
	for i, member := range listResponse.Data {
		memberResources[i] = resources.NewMemberResource(member)
	}

	response := repository.ListResponse[resources.MemberResource]{
		Data:       memberResources,
		Pagination: listResponse.Pagination,
	}

	return response, nil
}

func (r *MemberService) CreateMember(memberInput models.Member) (resources.MemberResource, error) {
	if err := r.Create(&memberInput); err != nil {
		return resources.MemberResource{}, err
	}
	dispatcher := events.GetDispatcher()
	dispatcher.Dispatch(events.Event{
		Type:    events.MemberCreated,
		Payload: memberInput,
	})
	return r.GetMemberByID(memberInput.ID)
}

func (r *MemberService) UpdateMember(memberInput models.Member) (resources.MemberResource, error) {
	if err := r.Update(&memberInput); err != nil {
		return resources.MemberResource{}, err
	}
	dispatcher := events.GetDispatcher()
	dispatcher.Dispatch(events.Event{
		Type:    events.MemberUpdated,
		Payload: memberInput,
	})
	return r.GetMemberByID(memberInput.ID)
}

func (r *MemberService) DeleteMember(id string) error {
	member, err := r.FindByID(id, nil)
	if err != nil {
		return err
	}

	dispatcher := events.GetDispatcher()
	dispatcher.Dispatch(events.Event{
		Type:    events.MemberDeleted,
		Payload: member,
	})
	return r.Delete(&member)
}
