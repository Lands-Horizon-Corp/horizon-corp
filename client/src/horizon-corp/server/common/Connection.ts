import UseServer from "../../request/server";

export default class Connection {
  private server: UseServer;

  constructor() {
    this.server = new UseServer();
  }
  /**
  * Test backend connection
  *
  * @returns {Promise<boolean>} 
  */
  async test(): Promise<boolean> {
    const result = await this.server.get('/');
    return result.status === 200
  }

}