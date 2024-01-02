<template>
  <div>
    <h1>公司管理</h1>
    <el-button @click="addCompanyVisible = true">添加新公司</el-button>
    <el-dialog v-model="addCompanyVisible" title="添加新公司">
      <Form
        :col="companyCol"
        @submit="
          (data) => {
            addCompany(data)
          }
        "
      />
    </el-dialog>
    <el-table :data="companyData">
      <el-table-column
        v-for="item in column"
        :key="item.prop"
        :prop="item.prop"
        :label="item.label"
        :width="item.width"
      />
      <el-table-column prop="actions" label="操作">
        <template #default="{ row }">
          <el-button @click="deleteCompany(row)">删除</el-button>
          <el-button @click="ElMessage.info('开发中')">编辑</el-button>
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
const api = '/company'
const companyData = ref([])
const column: Array<Column> = reactive([
  { prop: 'ID', label: 'ID', width: 0 },
  { prop: 'Name', label: '公司名称', width: 0 },
])
// add company form
const companyCol = reactive([
  { label: '公司名称', prop: 'Name', type: 'string' },
])
const addCompanyVisible = ref(false)
const addCompany = async (_data: Company) =>
  dialogPost(api, _data, addCompanyVisible, companyData)
const deleteCompany = async (_data: Company) =>
  deleteData(api, _data.ID, companyData)

onMounted(() => {
  fetchData(api, companyData)
})
</script>
