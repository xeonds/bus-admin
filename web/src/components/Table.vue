<template>
  <el-table :data="data">
    <el-table-column
      v-for="col in column"
      :key="col.prop"
      :prop="col.prop"
      :label="col.label"
    />
  </el-table>
</template>

<script lang="ts" setup>
import { http } from '@/utils/http'
import type { Column } from 'element-plus'
const { get } = http
const data: any = ref(null)
const props = defineProps({
  column: {
    type: Array<Column>,
    default: [],
  },
  api: {
    type: String,
    default: '',
  },
})
onMounted(async () => {
  const { data } = await get(props.api)
  update(data.value)
})
const update = (_data: any) => {
  data.value = _data
}
</script>
