export const useFetch = async (url: string, init?: RequestInit | undefined) => {
  const data = ref(null)
  const err = ref(null)
  await fetch(url, init)
    .then((res) => res.json())
    .then((json) => (data.value = json))
    .catch((err) => (err.value = err))

  return { data, err }
}

const useHttp = (baseUrl: string, token: string) => {
  const get = (url: string) => {
    return useFetch(`${baseUrl}${url}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
  }
  const post = (url: string, data: any) => {
    return useFetch(`${baseUrl}${url}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(data),
    })
  }
  const del = (url: string, id: number) => {
    return useFetch(`${baseUrl}${url}/${id}`, {
      method: 'DELETE',
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
  }
  const put = (url: string, data: any) => {
    return useFetch(`${baseUrl}${url}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(data),
    })
  }
  return { get, post, del, put }
}

const baseUrl = '/api/v1'

export const http = useHttp(baseUrl, '')

export const dialogPost = async (
  api: string,
  _data: any,
  visibleRef: any,
  dataSrc: any
) => {
  const { post } = http
  const { err } = await post(api, _data)
  if (err.value != null) ElMessage.error(err.value)
  else ElMessage.success('添加成功')
  visibleRef.value = false // 关闭弹窗
  await fetchData(api, dataSrc)
  return err
}

export const fetchData = async (api: string, dataSrc: any) => {
  const { get } = http
  const { data, err } = await get(api)
  if (err.value != null) ElMessage.warning(err.value)
  dataSrc.value = data.value as any
}

export const deleteData = async (api: string, id: number, dataSrc: any) => {
  const { del } = http
  const { err } = await del(api, id)
  if (err.value != null) ElMessage.error(err.value)
  else {
    ElMessage.success('删除成功')
    fetchData(api, dataSrc)
  }
}
