import UseServer from '../../request/server'
import { ContactsResource, ContactsRequest } from '../../types'

/**
 * Service class to handle CRUD operations for contacts.
 */
export default class ContactService {
    private static readonly BASE_ENDPOINT = '/contacts'

    /**
     * Retrieves all contacts.
     *
     * @returns {Promise<ContactsResource[]>} - A promise that resolves to an array of contact resources.
     */
    public static async getAll(): Promise<ContactsResource[]> {
        const response = await UseServer.get<ContactsResource[]>(
            ContactService.BASE_ENDPOINT
        )
        return response.data
    }

    /**
     * Creates a new contact.
     *
     * @param {ContactsRequest} contactData - The data for the new contact.
     * @returns {Promise<ContactsResource>} - A promise that resolves to the created contact resource.
     */
    public static async create(
        contactData: ContactsRequest
    ): Promise<ContactsResource> {
        const response = await UseServer.post<
            ContactsRequest,
            ContactsResource
        >(ContactService.BASE_ENDPOINT, contactData)
        return response.data
    }

    /**
     * Deletes a contact by its ID.
     *
     * @param {number} id - The ID of the contact to delete.
     * @returns {Promise<void>} - A promise that resolves when the deletion is complete.
     */
    public static async delete(id: number): Promise<void> {
        const endpoint = `${ContactService.BASE_ENDPOINT}/${id}`
        await UseServer.delete<void>(endpoint)
    }

    /**
     * Updates an existing contact by its ID.
     *
     * @param {number} id - The ID of the contact to update.
     * @param {ContactsRequest} contactData - The updated data for the contact.
     * @returns {Promise<ContactsResource>} - A promise that resolves to the updated contact resource.
     */
    public static async update(
        id: number,
        contactData: ContactsRequest
    ): Promise<ContactsResource> {
        const endpoint = `${ContactService.BASE_ENDPOINT}/${id}`
        const response = await UseServer.put<ContactsRequest, ContactsResource>(
            endpoint,
            contactData
        )
        return response.data
    }

    /**
     * Retrieves a contact by its ID.
     *
     * @param {number} id - The ID of the contact to retrieve.
     * @returns {Promise<ContactsResource>} - A promise that resolves to the contact resource.
     */
    public static async getById(id: number): Promise<ContactsResource> {
        const endpoint = `${ContactService.BASE_ENDPOINT}/${id}`
        const response = await UseServer.get<ContactsResource>(endpoint)
        return response.data
    }
}
