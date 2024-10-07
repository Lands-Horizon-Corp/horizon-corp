export * from './admin'
export * from './employee'
export * from './member'
export * from './owner'
export * from './auth'
import common from './common'

const service = {
  common: common()
}

export default service