import UseServer from "../../request/server";

export default class Connection {
  private static readonly BASE_ENDPOINT = '/'

  /**
   * Test backend connection
   *
   * @returns {Promise<boolean>}
   */
  public static async test(): Promise<boolean> {
    const response = await UseServer.get(Connection.BASE_ENDPOINT)
    return response.status === 200
  }
}