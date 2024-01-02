<template>
    <div>
        <h1>司机管理</h1>
        <el-row>
            <el-button @click="addDriverVisible = true">司机信息录入</el-button>
        </el-row>
        <el-dialog title="司机信息录入" v-model="addDriverVisible">
            <Form :col="companyCol" @submit="(data) => { addCompany(data) }"></Form>
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
import { dialogPost, http } from '@/utils/http';
const { get } = http
const api = "/company"
// table
import type { Column } from 'element-plus';
const column: Array<Column> = reactive([
    { prop: 'ID', label: 'ID', width: 0 },
    { prop: 'Name', label: '公司名称', width: 0 },
])
const companyData = ref([])
const update = (_data: any) => { companyData.value = _data }
// add company form
const companyCol = reactive([{ label: "司机名称", prop: "Name", type: "string" }])
const addDriverVisible = ref(false)
const addCompany = (_data: any) => dialogPost(api, _data, addDriverVisible)

onMounted(() => {
    const { data, err } = get(api)
    if (err.value != null) update(data.value)
})

</script>
