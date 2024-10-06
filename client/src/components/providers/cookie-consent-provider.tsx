import { useEffect } from 'react'
import CookieConsent from 'react-cookie-consent'

function App() {
    const handleAccept = () => {
        localStorage.setItem('cookieConsent', 'true')
        document.cookie =
            'your_cookie_name=your_cookie_value; SameSite=None; Secure; path=/; domain=yourdomain.com'
    }

    useEffect(() => {
        const consent = localStorage.getItem('cookieConsent')
        if (consent) {
            document.cookie =
                'your_cookie_name=your_cookie_value; SameSite=None; Secure; path=/; domain=yourdomain.com'
        }
    }, [])

    return (
        <CookieConsent
            location="bottom"
            buttonText="I understand"
            cookieName="userCookieConsent"
            style={{ background: '#2B373B' }}
            buttonStyle={{ color: '#4e503b', fontSize: '13px' }}
            onAccept={handleAccept}
            expires={150}
        >
            This website uses cookies to enhance the user experience.
        </CookieConsent>
    )
}

export default App
