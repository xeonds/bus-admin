import { fetchData } from './http'

export const addCascade = async (
  api: string,
  cols: any,
  label: string,
  prop: string,
  itemMap: any
) => {
  const item = ref([])
  await fetchData(api, item)
  cols.push({
    width: 0,
    label,
    prop,
    type: 'single-select',
    options: item.value.map(itemMap),
  })
}
