package controllers

// GET: /
// Retrieve all my media files along with the total size of all media.
// Supports filtering with or without pagination. Results can be converted to records.

// GET: /:id
// Retrieve current media

// GET: /team
// Retrieve media files based on role:
//     Member: Access not allowed.
//     Admin: Can retrieve all media, including files from all branches, owners, employees, and members.
//     Owner: Can retrieve all media, including files from all owned branches, employees, and members.
//     Employee: Can retrieve all media from employees and members within their branch.
//     Supports filtering with or without pagination. Results can be converted to records.

// POST: /upload
// Upload a new file.

// DELETE: /:id
// Delete a specific file by its ID and the ID was the same as the current user id logged in.

// PUT: /:id
// Update a file's name and description only.
