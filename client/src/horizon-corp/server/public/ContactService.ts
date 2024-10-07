import UseServer from "../../request/server";
import { AxiosResponse } from 'axios';
import type { ContactsRequest, ContactResource } from '../../types'


class ContactService {
  private server: UseServer;

  constructor() {
    this.server = new UseServer();
  }

  /**
   * Gets all contacts.
   *
   * @returns {Promise<ContactResource[]>} - An array of contact resources.
   */
  async getAll(): Promise<ContactResource[]> {
    const response: AxiosResponse<ContactResource[]> = await this.server.get('/contacts');
    return response.data;
  }

  /**
   * Creates a new contact.
   *
   * @param {ContactsRequest} contactData - The contact data to create.
   * @returns {Promise<ContactResource>} - The created contact resource.
   */
  async create(contactData: ContactsRequest): Promise<ContactResource> {
    const response: AxiosResponse<ContactResource> = await this.server.post('/contacts', contactData);
    return response.data;
  }


  /**
   * Deletes a contact by ID.
   *
   * @param {number} id - The ID of the contact to delete.
   * @returns {Promise<void>} - A promise that resolves when the contact is deleted.
   */
  async delete(id: number): Promise<void> {
    await this.server.delete(`/contacts/${id}`);
  }

  /**
   * Updates an existing contact by ID.
   *
   * @param {number} id - The ID of the contact to update.
   * @param {ContactsRequest} contactData - The updated contact data.
   * @returns {Promise<ContactResource>} - The updated contact resource.
   */
  async update(id: number, contactData: ContactsRequest): Promise<ContactResource> {
    const response: AxiosResponse<ContactResource> = await this.server.put(`/contacts/${id}`, contactData);
    return response.data;
  }

  /**
  * Gets a contact by ID.
  *
  * @param {number} id - The ID of the contact to retrieve.
  * @returns {Promise<ContactResource>} - The retrieved contact resource.
  */
  async getById(id: number): Promise<ContactResource> {
    const response: AxiosResponse<ContactResource> = await this.server.get(`/contacts/${id}`);
    return response.data;
  }
}

export default ContactService