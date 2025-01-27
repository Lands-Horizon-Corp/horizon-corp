Endpoints Needed for Owner user

/api/v1/owner/:id 
    GET 
        - ✅ Working

/api/v1/owner/1/company
    GET 
        - 🚫 Not-working 
        - returns a single company owned by current signed in owner

/api/v1/company - 
    POST (CREATE) 
        - 🚫 Not Working in Owner
        - 🚫 Not Working in Admin
        - returns CompanyResource (also include owner, no preload params required)
    GET 
        - 🚫 Not Working 
        - Should support Filter | Paginated | Preloads
        - returns Paginated<CompanyResource>[] 

Enpoints below rely on api/v1/owner/1/company, before I can test the endpoints below
/api/v1/company/:id - Get | Update | Delete
    GET 
        - Not yet tested in owner
        - ✅ Working in Admin 
        - Should support Preloads
        - returns CompanyResource of specific company ID
    UPDATE 
        - Not yet tested in owner  
        - ✅ Working in Admin
    DELETE 
        - Not yet tested in owner | ✅ Working in Admin

/api/v1/company/:id/branch (filter, paginated, preloads)
    GET 
        - 🚫 Not Working 
        - Should Support filter | paginated | preloads
        - return Paginated<BranchResource>[] but only for the specific company id

/api/v1/branch (filter, paginated, preloads)
    GET 
        - 🚫 Not Working
        - Should Support filter | paginated | preloads
        - return Paginated<BranchResource>[] 


Owner User
- Getting Owner's Own Company
    - Should support Preloads?

- Getting 