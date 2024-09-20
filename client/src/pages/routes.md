# public

/
/about
/contact
/pricing
/docs (future feature implement)
/terms-condition
/auth/sign-in
/auth/sign-up
/auth/verification
/auth/verification/email
/auth/verification/mobile/

# private

## user

/user/
/user/profile

## company

/$company/
/$company/employee
/$company/member
/$company/admin
/$company/branch
/$company/branch/$branchId
/$company/branch/$branchId/employee

## admin

{
role : "admin" | "user" | "company"
}

{
user : details...
path : "/company" | "/admin" | "/member"
}
