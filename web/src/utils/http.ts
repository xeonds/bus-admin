export const useFetch = (url: string, init?: RequestInit | undefined) => {
  const data = ref(null)
  const err = ref(null)
  fetch(url, init)
    .then((res) => res.json())
    .then((json) => data.value = json)
    .catch((err) => err.value = err)

  return { data, err }
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
  return { get, post }
}

const baseUrl = '/api/v1'

export const http = useHttp(baseUrl, "")

export const dialogPost = (api: string, _data: any, visibleRef: any) => {
  const { post } = http
  const { err } = post(api, _data)
  if (err.value != null) ElMessage.error(err.value)
  else ElMessage.success("添加成功")
  visibleRef.value = false // 关闭弹窗
  return err
}