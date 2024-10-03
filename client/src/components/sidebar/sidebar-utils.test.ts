import { describe, expect, it } from 'vitest'
import {
    concatParentUrl,
    sidebarCollapsableRouteMatcher,
} from './sidebar-utils'

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

describe('Collapsable Route matcher function for sidebar', () => {
    it('should match baseUrl:/abc/d and url:/abc/a', () => {
        expect(sidebarCollapsableRouteMatcher('/abc/d', '/abc/a')).toBe(false)
    })

    it('should not match baseUrl:/abc/d and url:/bca/d', () => {
        expect(sidebarCollapsableRouteMatcher('/abc/d', '/abc/a')).toBe(false)
    })

    it('should not match baseUrl:/mem and url:/membored/desu', () => {
        expect(sidebarCollapsableRouteMatcher('/mem', '/membor/desu')).toBe(
            false
        )
    })
})
