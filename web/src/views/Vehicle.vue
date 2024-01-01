<template>
    <div>
        <h1>车辆管理</h1>
        <el-button @click="addCompanyVisible = true">车辆信息录入</el-button>
        <el-dialog v-model="addCompanyVisible">
            <Form :col="companyCol" :on-submit="(_, data) => { addCompany(data) }"></Form>
        </el-dialog>
        <el-table>
            <el-table-column v-for="item in column" :key="item.prop" :prop="item.prop" :label="item.label"
                :width="item.width"></el-table-column>
            <el-table-column prop="actions" label="操作">
                <template #default="{ row }">
                    <el-button>删除</el-button>
                    <el-button>编辑</el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script lang="ts" setup>
// data action
import { http } from '@/utils/http';
const { get, post } = http
const api = "/vehicle"
const companyData = ref([])
const update = (_data: any) => { companyData.value = _data }
const addCompany = (_data: any) => {
    const { data, err } = post(api, _data)
    if (err.value != null) update(data.value)
}
// table
import type { Column } from 'element-plus';
const column: Array<Column> = reactive([
    { prop: 'ID', label: 'ID', width: 0 },
    { prop: 'Name', label: '车辆名称', width: 0 },
])
// add company form
const companyCol = reactive([{ label: "车辆名称", prop: "Name", type: "string" }])
const addCompanyVisible = ref(false)

onMounted(() => {
    const { data, err } = get(api)
    if (err.value != null) update(data.value)
})

</script>
