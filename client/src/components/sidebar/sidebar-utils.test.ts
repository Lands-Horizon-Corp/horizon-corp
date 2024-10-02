import { describe, expect, it } from 'vitest'
import { concatParentUrl } from './sidebar-utils'

describe('testing concatParentUrl', () => {
    it('should concatenate baseUrl:/member and url:/task', () => {
        const res = concatParentUrl({ baseUrl: '/member', url: '/task' })
        expect(res).toBe('/member/task')
    })

    it('should return only the url when baseUrl is empty', () => {
        const res = concatParentUrl({ baseUrl: '', url: '/task' })
        expect(res).toBe('/task')
    })

    it('should return only the baseUrl when url is empty', () => {
        const res = concatParentUrl({ baseUrl: '/member', url: '' })
        expect(res).toBe('/member')
    })

    it('should handle undefined values for baseUrl and url by defaulting to empty strings', () => {
        const res = concatParentUrl({})
        expect(res).toBe('')
    })

    it('should concatenate baseUrl:/member and url:/profile', () => {
        const res = concatParentUrl({ baseUrl: '/member', url: '/profile' })
        expect(res).toBe('/member/profile')
    })

    it('should concatenate baseUrl:/member and url:/settings', () => {
        const res = concatParentUrl({ baseUrl: '/member', url: '/settings' })
        expect(res).toBe('/member/settings')
    })
})
