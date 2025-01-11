package controllers

// GET: /
// Retrieve branches with optional filtering for pagination or no pagination. Results can be converted to records.
//     Admin: Allowed
//     Employee: Only the company of current branch
//     Owner: Only own company
//     Member: Only own company

// GET: /:id
//     Admin: Allowed
//     Employee: Only the company of current branch
//     Owner: Only own company
//     Member: Only own company

// DELETE: /:id
// Admin: Allowed but if no members on branches, no branches, and no employees on branches.
// Owner: Only own company but if no members on branches, no branches, and no employees on branches.

// POST: /
//     Admin: Can create company but need to assign owner
//     Owner: Can create company and automaticall assign owner
//     Employee: Not allowed
//     Member: Only own company
