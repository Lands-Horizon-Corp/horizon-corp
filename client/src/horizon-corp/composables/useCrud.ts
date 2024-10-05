import getSimpleProperties from '@/lib/utils'
import UseServer from './useServer'
import { create } from 'zustand'

/**
 * Interface representing the CRUD store state.
 *
 * @template T - The type of the data entries.
 */
export interface CrudState<T> {
    // Data
    entries: T[]

    // Mutations
    insert: (data: T) => void
    merge: (data: T[]) => void
    remove: (data: T) => void
    removeMany: (data: T[]) => void
    clear: () => void

    // Test
    connection: () => Promise<boolean>

    // Http methods
    // index
    // store
    // edit
    // update
    //
}

/**
 * A custom Zustand hook for managing CRUD operations on a generic data type.
 *
 * @template T - The type of the data entries.
 * @param {object} resource - The resource metadata for identifying entries.
 * @param {keyof T} resource.identifier - The key used to identify unique entries.
 * @param {string} resource.path - The path for the resource.
 * @param {boolean | object | Array<string>} [allowComplexProps] - Determines whether complex properties should be allowed or which properties to allow.
 * @returns {CrudState<T>} A Zustand store containing the state and actions for CRUD operations.
 */
const useCrud = <T extends object>(
    resource: { identifier: keyof T; path: string },
    allowComplexProps?: boolean | object | Array<string>
): (() => CrudState<T>) => {
    return create<CrudState<T>>((set, get) => ({
        /**
         * List of current data entries in the store.
         */
        entries: [],

        /**
         * Inserts a new entry or updates an existing one based on the identifier.
         *
         * @param {T} data - The entry data to insert or update.
         */
        insert: (data: T) => {
            const currentEntries = get().entries
            const index = currentEntries.findIndex(
                (entry) =>
                    entry[resource.identifier] === data[resource.identifier]
            )

            if (index !== -1) {
                const updatedData = { ...currentEntries[index] }
                delete updatedData[resource.identifier]

                set((state) => ({
                    entries: state.entries.map((entry, i) =>
                        i === index
                            ? {
                                  ...entry,
                                  ...(allowComplexProps === true
                                      ? updatedData
                                      : getSimpleProperties<T>(
                                            updatedData,
                                            allowComplexProps
                                        )),
                              }
                            : entry
                    ),
                }))
            } else {
                set((state) => ({
                    entries: [...state.entries, data],
                }))
            }
        },

        /**
         * Merges an array of new data entries into the existing entries, updating any matching entries by identifier.
         *
         * @param {T[]} data - The array of entries to merge into the store.
         */
        merge: (data: T[]) => {
            const currentEntries = get().entries
            const dataMap = new Map<string, T>(
                data.map((item) => [item[resource.identifier] as string, item])
            )

            const updatedEntries = currentEntries.map((entry) => {
                const key = entry[resource.identifier] as string
                if (dataMap.has(key)) {
                    const updatedEntry = {
                        ...entry,
                        ...dataMap.get(key),
                    }
                    dataMap.delete(key)
                    return updatedEntry
                }
                return entry
            })

            set(() => ({
                entries: [...updatedEntries, ...Array.from(dataMap.values())],
            }))
        },

        /**
         * Removes a single entry from the store based on its identifier.
         *
         * @param {T} data - The entry to remove.
         */
        remove: (data: T) => {
            const currentEntries = get().entries
            const updatedEntries = currentEntries.filter(
                (entry) =>
                    entry[resource.identifier] !== data[resource.identifier]
            )

            set(() => ({
                entries: updatedEntries,
            }))
        },

        /**
         * Removes multiple entries from the store based on their identifiers.
         *
         * @param {T[]} data - The array of entries to remove.
         */
        removeMany: (data: T[]) => {
            const currentEntries = get().entries
            const identifiersToRemove = new Set(
                data.map((item) => item[resource.identifier] as string)
            )

            const updatedEntries = currentEntries.filter(
                (entry) =>
                    !identifiersToRemove.has(
                        entry[resource.identifier] as string
                    )
            )

            set(() => ({
                entries: updatedEntries,
            }))
        },

        /**
         * Clears all entries from the store.
         */
        clear: () => {
            set(() => ({
                entries: [],
            }))
        },

        /**
         * Test backend connection
         */
        connection: async (): Promise<boolean> => {
            const server = new UseServer()
            try {
                await server.get('/')
                return true
            } catch {
                return false
            }
        },
    }))
}

export default useCrud
