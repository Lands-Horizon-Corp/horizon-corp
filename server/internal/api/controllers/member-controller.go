package controllers

// GET: /
// Retrieve members with optional filtering for pagination or no pagination. Results can be converted to records.
//     Admin: Can retrieve all members in the current branch and across the company.
//     Employee: Can retrieve all members of the current branch.
//     Owner: Can retrieve all members across all owned branches.
//     Member: Access not allowed, except for members in the current branch.

// GET: /member-applications
// Retrieve members with optional filtering for pagination or no pagination. Results can be converted to records.
//     Admin: Can get all member applications all of it
//     Employee: Can get all member application from it's branch only
//     Owner: Can get all member applications from all of its branch
//     Member: not allowed

// GET: /member-application/:id
// 	Admin: Can retrieve the specified member application.
//	Employee: Can retrieve the specified member application.
// 	Owner: Can retrieve the specified member application.
// 	Member: Can retrieve only their own member application, not applications belonging to others.

// PUT: /member-application/:id
// Verifiy member application:
//     Admin: Can verify but must assign employee
//     Owner: Can verify but must assign employee
//     Employee: Can verify but must be on the same branch
// 		 Member: not allowed

// POST: /member-application
// Create a new member:
//     Admin: Must assign both company and branch (uses AdminCreateMemberApplicationRequest).
//     Owner: Must assign a branch only (uses OwnerCreateMemberApplicationRequest).
//     Employee: No need to assign; simply create (uses CreateMemberApplicationRequest).
// 		 Member: Must be pending and can choose different company and branches based on nearest

// DELETE: /:id
// Delete a member based on role permissions:
//     Member: Not allowed.
//     Employee: Can delete members only within their current branch.
//     Admin: Fully allowed.
//     Owner: Can delete members only within their owned branches.

// PUT: /:id
// Update member details based on role permissions:
//     Employee: Cannot edit members but can provide a reset password link.
//     Admin: Cannot edit members but can provide a reset password link.
//     Owner: Cannot edit members but can provide a reset password link.
//     Member: Can edit their own details but not others'.
