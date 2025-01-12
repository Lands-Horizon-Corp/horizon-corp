package controllers

type AdminController struct{}

func NewAdminController() *AdminController {
	return &AdminController{}
}

// GET: /
// Retrieve admins with optional filtering for pagination or no pagination. Results can be converted to records.
//     Admin: Can retrieve all admins if admin  status is verified
//     Employee: not allowed
//     Owner: not allowed
//     Member: not allowed

// GET: /:ID
//     Admin: if admin  status is verified
//     Employee: not allowed
//     Owner: not allowed
//     Member: not allowed

// POST: /
//     Admin: Can create admin and automatically verified and also if admin  status is verified
//     Employee: not allowed
//     Owner: not allowed
//     Member: not allowed

// POST: /:id
// Verifiy admin
//     Admin: Can create verify admin but only if admin  status is verified
//     Employee: not allowed
//     Owner: not allowed
//     Member: not allowed

// DELETE: /:id
// Verifiy admin
//     Admin: only if admin  status is verified
//     Employee: not allowed
//     Owner: not allowed
//     Member: not allowed

// PUT: /verify/:id account type
// Retrieve admins with optional filtering for pagination or no pagination. Results can be converted to records.
//     Admin: Can retrieve all admin, employee, owner, and employee for verification. changing the status
//     Employee: not allowed
//     Owner: not allowed
//     Member: not allowed
