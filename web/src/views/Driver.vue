<template>
  <div>
    <h1>司机管理</h1>
    <el-row>
      <el-button @click="addDriverVisible = true">司机信息录入</el-button>
    </el-row>
    <el-dialog v-model="addDriverVisible" title="司机信息录入">
      <Form
        :col="driverCol"
        @submit="
          (data) => {
            addDriver(data)
          }
        "
      />
    </el-dialog>
    <el-table :data="driverData">
      <el-table-column
        v-for="item in column"
        :key="item.prop"
        :prop="item.prop"
        :label="item.label"
        :width="item.width"
      />
      <el-table-column prop="actions" label="操作">
        <template #default="{ row }">
          <el-button @click="deleteDriver(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script lang="ts" setup>
// data action
import { addCascade } from '@/utils/common'
import { deleteData, dialogPost, fetchData } from '@/utils/http'
// table
import type { Column } from 'element-plus'
const api = '/driver'
const column: Array<Column> = reactive([
  { prop: 'ID', label: 'ID', width: 0 },
  { prop: 'Name', label: '司机姓名', width: 0 },
  { prop: 'RouteID', label: '所属路线ID', width: 0 },
])
const driverData = ref([])
// add driver form
const driverCol = reactive([
  { label: '司机名称', prop: 'Name', type: 'string' },
])
const addDriverVisible = ref(false)
const addDriver = async (_data: any) =>
  dialogPost(api, _data, addDriverVisible, driverData)
const deleteDriver = async (_data: Driver) =>
  deleteData(api, _data.ID, driverData)

onMounted(() => {
  fetchData(api, driverData)
  addCascade('/route', driverCol, '所属路线', 'Route', (item: any) => {
    return { label: item.Name, value: item }
  })
})
</script>
