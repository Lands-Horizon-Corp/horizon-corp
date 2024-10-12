import otpVerificationRaw from '../assets/email-templates/account-otp-verification.html?raw'
import accountVerificationRaw from '../assets/email-templates/account-verification.html?raw'

// Define the types of templates available
type Templates = 'otp' | 'verification'

// Precompile the Handlebars templates
const TEMPLATE_MAP: Record<Templates, string> = {
  otp: otpVerificationRaw,
  verification: accountVerificationRaw,
}

/**
 * Retrieves and generates the email content based on the template type.
 * @param template - The type of template to use ('otp' or 'verification').
 * @param params - An object containing parameters to inject into the template.
 * @returns The generated email content as a string.
 * @throws Will throw an error if the template type is invalid or generation fails.
 */
export const getEmailContent = (
  template: Templates,
): string => {
  const compiledTemplate = TEMPLATE_MAP[template]
  if (!compiledTemplate) {
    throw new Error(`"${template}" is not a valid template type.`)
  }
  try {
    return compiledTemplate
  } catch (error) {
    console.error('Error generating email content:', error)
    throw new Error('Failed to generate email content.')
  }
}
