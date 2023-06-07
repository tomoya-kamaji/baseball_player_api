import { SWRConfiguration } from 'swr'
import useSWR from 'swr'
import { ApiError, ApiPathParam, ApiQueryParam, ApiResponse, ExactPathByHttpMethod } from '../axios/schemaUtil'
import { RequestResponse, createApiClient } from '../axios/axiosClient'
import { useCallback, useMemo } from 'react'

type GetPath = ExactPathByHttpMethod<'get'>

type UseFetchOption<Path extends GetPath> = {
  path: Path
  params?: {
    paths?: ApiPathParam<Path, 'get'>
    query?: ApiQueryParam<Path, 'get'>
  }
  shouldCancel?: boolean
  onSuccess?: (data: ApiResponse<Path, 'get'>) => void
  onError?: (error: ApiError<Path, 'get'>) => void
} & Omit<SWRConfiguration, 'onSuccess' | 'onError'> // revalidateOnFocus などのSWRのオプションも渡せるようにしておく

export const useFetch = <Path extends GetPath>(options: UseFetchOption<Path>) => {
  const { path, params, onSuccess, onError, shouldCancel, ...swrOptions } = options

  const api = createApiClient({ path, httpMethod: 'get', params })

  const {
    data: swrData,
    mutate: swrMutate,
    isValidating
  } = useSWR<RequestResponse<Path, 'get'>>(shouldCancel ? null : api.path(), () => api.request(), {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    onSuccess: (res: any) => {
      if (res.result === 'success' && !!onSuccess) {
        // eslint-disable-next-line @typescript-eslint/no-unsafe-argument
        onSuccess(res.data)
        return
      }
      if (res.result === 'error' && !!onError) {
        // eslint-disable-next-line @typescript-eslint/no-unsafe-argument
        onError(res.error)
        return
      }
    },
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    onError: (e: any) => {
      // unexpected error をハンドリング
      console.error(e.message)
    },
    ...swrOptions
  })

  const data = useMemo(() => {
    if (!swrData) return undefined
    if (swrData.result !== 'success') return undefined
    return swrData.data
  }, [swrData])

  const error = useMemo(() => {
    if (!swrData) return undefined
    if (swrData.result !== 'error') return undefined
    return swrData.error
  }, [swrData])

  const mutate = useCallback(
    (data?: ApiResponse<Path, 'get'>, shouldRevalidate?: boolean) => {
      void swrMutate(data ? { result: 'success', data } : undefined, shouldRevalidate)
    },
    [swrMutate]
  )

  return { data, error, mutate, isValidating }
}
