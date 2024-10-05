import getSimpleProperties from '@/lib/utils';
import { create } from 'zustand';


interface CrudState<T> {
  entries: T[];
  insert: (data: T) => void;
}


const useCrudStore = <T extends { [key: string]: unknown }>(
  resource: { identifier: string; path: string },
  allowComplexProps?: boolean | object | Array<string> | boolean
) => {

  return create<CrudState<T>>((set, get) => ({
    entries: [],


    insert: (data: T) => {
      const currentEntries = get().entries;

      const index = currentEntries.findIndex(
        (entry) => entry[resource.identifier] === data[resource.identifier]
      );

      if (index !== -1) {
        const duplicate = { ...currentEntries[index] };
        delete duplicate[resource.identifier];

        set((state) => ({
          entries: state.entries.map((entry, i) =>
            i === index
              ? {
                ...entry,
                ...(allowComplexProps === true
                  ? duplicate
                  : getSimpleProperties(duplicate, allowComplexProps)),
              }
              : entry
          ),
        }));
      } else {
        set((state) => ({
          entries: [...state.entries, data],
        }));
      }
    }



  }))
}
// entries: [],
// insert:(T)
// merge(T[])
// remove (identifier)
// removeMany (identifier[])
// restore
// clear

// get
// post
// patch
// storeMany
// destroy
// destroyMany
export default useCrudStore