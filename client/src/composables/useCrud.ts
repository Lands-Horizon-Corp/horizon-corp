import { create } from 'zustand';

interface CrudState<T> {
  entries: T[];
}


export default <T>() => {
  return create<CrudState<T>>(() => ({
    entries: [],
    // insert:(T)
    // merge(T[])
    // remove (identifier)
    // removeMany (identifier[])
    // restore
    // clear
  }));
};