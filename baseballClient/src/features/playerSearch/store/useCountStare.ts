import { create } from 'zustand'
import { SearchParam } from '../types/index'

type SearchParamState = {
  searchParam: SearchParam
  setSearchParam: (newParam: SearchParam) => void
}

export const useSearchParamStore = create<SearchParamState>((set) => ({
  searchParam: {
    limit: 20
  },
  setSearchParam: (newParam) => set(() => ({ searchParam: newParam }))
}))
