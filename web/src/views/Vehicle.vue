<template>
  <div>
    <h1>车辆管理</h1>
    <el-button @click="addVehicleVisible = true">车辆信息录入</el-button>
    <el-dialog v-model="addVehicleVisible" title="车辆信息录入">
      <Form
        :col="vehicleCol"
        @submit="
          (data) => {
            addVehicle(data)
          }
        "
      />
    </el-dialog>
    <el-table :data="vehicleData">
      <el-table-column
        v-for="item in column"
        :key="item.prop"
        :prop="item.prop"
        :label="item.label"
        :width="item.width"
      />
      <el-table-column prop="actions" label="操作">
        <template #default="{ row }">
          <el-button @click="deleteVehicle(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script lang="ts" setup>
// data action
import { deleteData, dialogPost, fetchData } from '@/utils/http'
// main table
import type { Column } from 'element-plus'
const api = '/vehicle'
const vehicleData = ref([])
const column: Array<Column> = reactive([
  { prop: 'ID', label: 'ID', width: 0 },
  { prop: 'VIN', label: '车牌号', width: 0 },
])
// add vehicle form
const vehicleCol = reactive([
  { label: '车牌号', prop: 'VIN', type: 'string' },
  { label: '所属公司', prop: 'Company', type: 'string' },
  { label: '路线', prop: 'Route', type: 'string' },
])
const addVehicleVisible = ref(false)
const addVehicle = async (_data: any) =>
  dialogPost(api, _data, addVehicleVisible, vehicleData)
const deleteVehicle = async (_data: any) =>
  deleteData(api, _data.ID, vehicleData)

onMounted(() => {
  fetchData(api, vehicleData)
})
</script>
