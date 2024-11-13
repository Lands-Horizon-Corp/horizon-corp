import { createContext, useContext, useState, ReactNode } from 'react';

type TSearchFilter = {
    type: 'string' | 'number' | 'date';
    value: string | number | Date;
    mode: 'starts with' | 'ends with' | 'include' | 'not include' | 'specific';
};

type TRangeFilter = {
    type: 'number' | 'date';
    from: number | Date;
    to: number | Date;
};

// Define a type for filters where any string key is allowed
type FilterObject = {
    [key: string]: TSearchFilter | TRangeFilter | undefined;
};

interface DataTableFilterContextType {
    filters: FilterObject;
    setFilter: (field: string, filter: TSearchFilter | TRangeFilter) => void;
    removeFilter: (field: string) => void;
    resetFilter: () => void;
}

const DataTableFilterContext = createContext<DataTableFilterContextType | undefined>(undefined);

export const useDataTableFilter = () => {
    const context = useContext(DataTableFilterContext);

    if (!context) {
        throw new Error('useDataTableFilter must be used within a DataTableFilterProvider');
    }

    return context;
};

// Example Provider
export const DataTableFilterProvider = ({ children }: { children: ReactNode }) => {
    const [filters, setFilters] = useState<FilterObject>({});

    const setFilter = (field: string, filter: TSearchFilter | TRangeFilter) => {
        setFilters(prev => ({ ...prev, [field]: filter }));
    };

    const removeFilter = (field: string) => {
        setFilters(prev => {
            const newFilters = { ...prev };
            delete newFilters[field];
            return newFilters;
        });
    };

    const resetFilter = () => {
        setFilters({});
    };

    return (
        <DataTableFilterContext.Provider value={{ filters, setFilter, removeFilter, resetFilter }}>
            {children}
        </DataTableFilterContext.Provider>
    );
};
