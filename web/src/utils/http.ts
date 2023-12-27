export const useFetch = (url: string, init?: RequestInit | undefined) => {
  const data = ref(null)
  const err = ref(null)
  fetch(url, init)
  .then((res)=>res.json())
  .then((json)=>data.value = json)
  .catch((err)=>err.value = err)

  return {data,err}
}

const useHttp = (baseUrl: string, token: string) => {
  const get = (url: string) => {
    return useFetch(`${baseUrl}${url}`, {
      headers: {
        'Authorization': `Bearer ${token}`
        }
    })
  }
  const post = (url: string, data: any) => {
    return useFetch(`${baseUrl}${url}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
      body: JSON.stringify(data)
    })
  }
  return {get,post}
}

const baseUrl = '/api/v1/'

export const http = useHttp(baseUrl, import.meta.env.VITE_API_TOKEN)