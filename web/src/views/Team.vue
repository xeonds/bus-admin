<template>
  <div>
    <h1>车队管理</h1>
    <el-row>
      <el-button @click="addTeamVisible = true">车队信息录入</el-button>
    </el-row>
    <el-dialog v-model="addTeamVisible" title="车队信息录入">
      <Form :col="teamCol" @submit="(data) => { addTeam(data) }" />
    </el-dialog>
    <el-dialog v-model="teamDetailVisible" title="车队详细信息">
      <Form :col="teamCol" @submit="(data) => { addTeam(data) }" />
    </el-dialog>
    <el-table :data="teamData">
      <el-table-column v-for="item in column" :key="item.prop" :prop="item.prop" :label="item.label"
        :width="item.width" />
      <el-table-column prop="actions" label="操作">
        <template #default="{ row }">
          <el-button @click="deleteTeam(row)">删除</el-button>
          <el-button>详情</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script lang="ts" setup>
// data action
import { deleteData, dialogPost, fetchData } from '@/utils/http'
import { addCascade } from '@/utils/common'
// main table
import type { Column } from 'element-plus'
const api = '/team'
const column: Array<Column> = reactive([
  { prop: 'ID', label: 'ID', width: 0 },
  { prop: 'Name', label: '车队名称', width: 0 },
  { prop: 'CompanyID', label: '所属公司ID', width: 0 },
  { prop: 'ManagerName', label: '管理员姓名', width: 0 },
])
const teamData = ref([])
// add team form
const teamCol: Array<Column> = reactive([
  { width: 0, label: '车队名称', prop: 'Name', type: 'string' },
  { width: 0, label: '管理员姓名', prop: 'ManagerName', type: 'string' },
])
const addTeamVisible = ref(false)
const addTeam = async (_data: any) =>
  dialogPost(api, _data, addTeamVisible, teamData)
const deleteTeam = async (_data: any) => deleteData(api, _data, teamData)
// team details dialog
const teamDetailVisible = ref(false)

onMounted(() => {
  fetchData(api, teamData)
  addCascade('/company', teamCol, '所属公司', 'Company', (item: any) => {
    return { label: item.Name, value: item }
  })
})
</script>
