package controllers

// GET: /
// Supports filtering for pagination or no pagination. Results can be converted to records.
// Admin: Can retrieve all members in the current branch and across the company.
// Employee: Can retrieve all members of the current branch.
// Owner: Can retrieve all members across all owned branches.
// Member: Not allowed, except for members in the current branch.
