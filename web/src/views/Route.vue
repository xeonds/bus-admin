<template>
  <div>
    <h1>线路管理</h1>
    <el-button @click="addRouteVisible = true">线路信息录入</el-button>
    <el-dialog v-model="addRouteVisible" title="线路信息录入">
      <Form :col="routeCol" @submit="(data) => {
          addRoute(data)
        }
        " />
    </el-dialog>
    <el-table :data="routeData">
      <el-table-column v-for="item in column" :key="item.prop" :prop="item.prop" :label="item.label"
        :width="item.width" />
      <el-table-column prop="actions" label="操作">
        <template #default="{ row }">
          <el-button @click="deleteRoute(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script lang="ts" setup>
// data action
import { addCascade } from '@/utils/common'
import { deleteData, dialogPost, fetchData } from '@/utils/http'
// main table
import type { Column } from 'element-plus'
const api = '/route'
const column: Array<Column> = reactive([
  { prop: 'ID', label: 'ID', width: 0 },
  { prop: 'Name', label: '线路名称', width: 0 },
  { prop: 'TeamID', label: '所属车队ID', width: 0 },
])
const routeData = ref([])
// add route form
const routeCol = reactive([{ label: '线路名称', prop: 'Name', type: 'string' }])
const addRouteVisible = ref(false)
const addRoute = async (_data: any) =>
  dialogPost(api, _data, addRouteVisible, routeData)
const deleteRoute = async (_data: any) => deleteData(api, _data.ID, routeData)

onMounted(() => {
  fetchData(api, routeData)
  addCascade('/team', routeCol, '所属车队', 'Team', (item: any) => {
    return { label: item.Name, value: item }
  })
})
</script>
