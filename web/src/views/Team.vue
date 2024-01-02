<template>
    <div>
        <h1>车队管理</h1>
        <el-row>
            <el-button @click="addTeamVisible = true">车队信息录入</el-button>
        </el-row>
        <el-dialog title="车队信息录入" v-model="addTeamVisible">
            <Form :col="teamCol" @submit="(data) => { addTeam(data) }"></Form>
        </el-dialog>
        <el-dialog title="车队详细信息" v-model="teamDetailVisible">
            <Form :col="teamCol" @submit="(data) => { addTeam(data) }"></Form>
        </el-dialog>
        <el-table>
            <el-table-column v-for="item in column" :key="item.prop" :prop="item.prop" :label="item.label"
                :width="item.width"></el-table-column>
            <el-table-column prop="actions" label="操作">
                <template #default="{ row }">
                    <el-button>删除</el-button>
                    <el-button>编辑</el-button>
                    <el-button>详情</el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script lang="ts" setup>
// data action
import { http, dialogPost } from '@/utils/http';
const { get } = http
const api = "/team"
// main table
import type { Column } from 'element-plus';
const column: Array<Column> = reactive([
    { prop: 'ID', label: 'ID', width: 0 },
    { prop: 'Name', label: '车队名称', width: 0 },
])
const teamData = ref([])
const update = (_data: any) => { teamData.value = _data }
// add team form
const teamCol = reactive([
    { label: "车队名称", prop: "Name", type: "string" },
    { label: "所属公司", prop: "Company", type: "string" },
    { label: "管理员姓名", prop: "ManagerName", type: "string" }])
const addTeamVisible = ref(false)
const addTeam = (_data: any) => dialogPost(api, _data, addTeamVisible)
// team details dialog
const teamDetailVisible = ref(false)

onMounted(() => {
    const { data, err } = get(api)
    if (err.value != null) update(data.value)
})

</script>
