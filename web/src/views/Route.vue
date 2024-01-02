<template>
    <div>
    <h1>线路管理</h1>
        <el-button @click="addRouteVisible = true">线路信息录入</el-button>
        <el-dialog title="线路信息录入" v-model="addRouteVisible">
            <Form :col="routeCol" @submit="(data) => { addRoute(data) }"></Form>
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
const api = "/route"
// main table
import type { Column } from 'element-plus';
const column: Array<Column> = reactive([
    { prop: 'ID', label: 'ID', width: 0 },
    { prop: 'VIN', label: '车牌号', width: 0 },
])
const routeData = ref([])
const update = (_data: any) => { routeData.value = _data }
// add route form
const routeCol = reactive([
    { label: "线路名称", prop: "Name", type: "string" },
    { label: "所属车队", prop: "Team", type: "string" }])
const addRouteVisible = ref(false)
const addRoute = (_data: any) => dialogPost(api, _data, addRouteVisible)

onMounted(() => {
    const { data, err } = get(api)
    if (err.value != null) ElMessage.error(err.value)
    else update(data.value)
})

</script>

