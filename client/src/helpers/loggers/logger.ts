/* eslint-disable no-console */
type LogMethod = (...args: unknown[]) => void

class Logger {
    private static instance: Logger
    private isDevelopment: boolean

    public log: LogMethod
    public warn: LogMethod
    public error: LogMethod
    public info: LogMethod
    public debug: LogMethod

    private constructor() {
        console.log(
            '\n                  ......                                    \n            .,,,,,,,,,,,,,,,,,,,                             \n        ,,,,,,,,,,,,,,,,,,,,,,,,,,                          \n      ,,,,,,,,,,,,,,  .,,,,,,,,,,,,,                        \n    ,,,,,,,,,,           ,,,,,,,,,,,,                       \n      ,,,,,,,          .,,,,,,,,,,,                          \n  ,*,,,,,,          ,,,,,,,,,,,,                             \n.**,,,,.**      .,,,,,,,,,,,                                \n.,,,,,,,**    ,,,,,,,,,,,                                   \n  .,,,,.**       ,,,,,,                                      \n    *******       ,                                         \n    **********              **,                             \n      ************,,  ,,*********,                          \n        **************************                          \n            ********************                             \n                  ******.\n'
        )

        this.isDevelopment =
            typeof import.meta.env !== 'undefined' &&
            import.meta.env.VITE_CLIENT_APP_ENV === 'development'

        if (typeof document !== 'undefined') {
            document.addEventListener('contextmenu', (event) =>
                event.preventDefault()
            )
        }

        if (this.isDevelopment) {
            this.log = console.log.bind(console)
            this.warn = console.warn.bind(console)
            this.error = console.error.bind(console)
            this.info = console.info.bind(console)
            this.debug = console.debug.bind(console)
        } else {
            this.log = () => {}
            this.warn = () => {}
            this.error = () => {}
            this.info = () => {}
            this.debug = () => {}

            console.log = () => {}
            console.warn = () => {}
            console.error = () => {}
            console.info = () => {}
            console.debug = () => {}
        }
    }

    public static getInstance(): Logger {
        if (!Logger.instance) {
            Logger.instance = new Logger()
        }
        return Logger.instance
    }
}

export default Logger.getInstance()
