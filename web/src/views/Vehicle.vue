<template>
    <div>
        <h1>车辆管理</h1>
        <el-button @click="addVehicleVisible = true">车辆信息录入</el-button>
        <el-dialog title="车辆信息录入" v-model="addVehicleVisible">
            <Form :col="vehicleCol" :on-submit="(_, data) => { addVehicle(data) }"></Form>
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
const vehicleData = ref([])
const update = (_data: any) => { vehicleData.value = _data }
const addVehicle = (_data: any) => {
    const { data, err } = post(api, _data)
    if (err.value != null) update(data.value)
}
// table
import type { Column } from 'element-plus';
const column: Array<Column> = reactive([
    { prop: 'ID', label: 'ID', width: 0 },
    { prop: 'VIN', label: '车牌号', width: 0 },
])
// add vehicle form
const vehicleCol = reactive([
    { label: "车牌号", prop: "VIN", type: "string" },
    { label: "所属公司", prop: "company", type: "string" },
    { label: "路线", prop: "route", type: "string" }])
const addVehicleVisible = ref(false)

onMounted(() => {
    const { data, err } = get(api)
    if (err.value != null) ElMessage.error(err.value)
    else update(data.value)
})

</script>
