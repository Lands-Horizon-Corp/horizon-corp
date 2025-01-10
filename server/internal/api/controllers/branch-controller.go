package controllers

// GET: /
// Retrieve branches with optional filtering for pagination or no pagination. Results can be converted to records.
//     Admin: Can can all
//     Employee: Only current branch
//     Owner: All of it's branches
//     Member: only current branch

// GET: /member-application
// Retrieve branches with optional filtering for pagination or no pagination. Results can be converted to records.
// Can can get all close to radius or place based on PostalCode, Province, City, Barangay, Region
// this will return all nearest furthest branches

// POST: /
//     Admin: Allowed but must select company first
//     Owner: Allowed
//     Employee: Not Allowed
//     Member: Not Allowed

// PUT: /:id
//     Admin: Allowed
//     Owner: Allowed
//     Employee: not Allowed
//     Member: not Allowed

// DELETE: /:id
//     Admin: Allowed but there must be not members
//     Owner: Allowed but there must be not members
//     Employee: not Allowed
//     Member: not Allowed
