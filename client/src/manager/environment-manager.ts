import { decode, encode } from "@/helpers";

// Singleton class for managing environment variables
export class EnvironmentManager {
  private static encryptedEnv: Record<string, string> = {};

  // Static initialization to encrypt environment variables
  public static async initialize(): Promise<void> {
    await this.storeEncryptedEnvVariable('VITE_CLIENT_SECRET_KEY', import.meta.env.VITE_CLIENT_SECRET_KEY);
    await this.storeEncryptedEnvVariable('VITE_CLIENT_PORT', import.meta.env.VITE_CLIENT_PORT);
    await this.storeEncryptedEnvVariable('VITE_CLIENT_SERVER_URL', import.meta.env.VITE_CLIENT_SERVER_URL);
    await this.storeEncryptedEnvVariable('VITE_CLIENT_DOCUMENT_URL', import.meta.env.VITE_CLIENT_DOCUMENT_URL);
    await this.storeEncryptedEnvVariable('VITE_CLIENT_GOOGLE_API_KEY', import.meta.env.VITE_CLIENT_GOOGLE_API_KEY);
  }

  // Store encrypted environment variable (static)
  private static async storeEncryptedEnvVariable(key: string, value: string): Promise<void> {
    if (!value) throw new Error(`Environment variable ${key} is not defined.`);
    const encryptedValue = await encode(value);
    this.encryptedEnv[key] = encryptedValue;
  }

  // Decrypt and retrieve environment variable (static)
  public static async getDecryptedEnvVariable(key: string): Promise<string> {
    const encryptedValue = this.encryptedEnv[key];
    if (!encryptedValue) throw new Error(`Encrypted environment variable ${key} is not found.`);
    return await decode(encryptedValue);
  }
}
