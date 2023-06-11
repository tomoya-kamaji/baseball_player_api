import { createContext } from 'react'
import { SearchParam } from '../types'

const searchParam: SearchParam = {}
export const SearchContext = createContext(searchParam)


