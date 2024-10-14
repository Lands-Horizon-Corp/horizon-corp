package repositories

type UserRepository struct {
	admin    *AdminRepository
	owner    *OwnerRepository
	employee *EmployeeRepository
	member   *MemberRepository
}

func NewUserRepository(
	admin *AdminRepository,
	owner *OwnerRepository,
	employee *EmployeeRepository,
	member *MemberRepository,
) *UserRepository {
	return &UserRepository{
		admin:    admin,
		owner:    owner,
		employee: employee,
		member:   member,
	}
}

func (r *UserRepository) Create(accountType string) {

}

func (r *UserRepository) GetAll(accountType string) {

}

func (r *UserRepository) GetByID(accountType string) {

}

func (r *UserRepository) Update(accountType string) {

}

func (r *UserRepository) Delete(accountType string) {

}

func (r *UserRepository) FindByEmailUsernameOrContact(accountType string) {

}

func (r *UserRepository) UpdateColumns(accountType string) {

}
