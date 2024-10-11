import * as handlebars from 'handlebars'
import otpVerication from '../assets/email-templates/account-otp-verication.html?raw'
import accountVerification from '../assets/email-templates/account-verication.html?raw'

type Templates = 'otp' | 'verification'

const TEMPLATE_MAP: Record<Templates, string> = {
    otp: otpVerication,
    verification: accountVerification,
}

export const getTemplates = async (template: Templates) => {
    const templateFile = TEMPLATE_MAP[template]
    if (!templateFile) {
        throw new Error(`${template} is not a valid template`)
    }
    return await getEmailTemplate(templateFile)
}

const getEmailTemplate = async (templateFile: string) => {
    try {
        const compiledTemplate = handlebars.compile(templateFile)
        // to do - add params
        return compiledTemplate({})
    } catch (error) {
        console.error('Error generating email template:', error)
        throw new Error('Failed to generate email template')
    }
}
