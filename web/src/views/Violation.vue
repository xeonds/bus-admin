<template>
    <div>
        <h1>违章管理</h1>
        <el-row>
            <el-button @click="insertViolationVisible = true">违章信息录入</el-button>
            <el-button @click="queryDriverViolationVisible = true">司机时段违章查询</el-button>
            <el-button @click="queryTeamViolationVisible = true">车队时段违章查询</el-button>
        </el-row>
        <el-dialog title="违章信息录入" v-model="insertViolationVisible">
            <Form :col="violationCol" @submit="(data) => { insertViolation(data) }"></Form>
        </el-dialog>
        <el-dialog title="司机违章查询" v-model="queryDriverViolationVisible">
            <el-form :inline="true">
                <el-form-item label="司机姓名">
                    <el-input v-model="driverName"></el-input>
                </el-form-item>
                <el-form-item label="起止时间选择">
                    <el-date-picker v-model="driverViolationTime" type="datetimerange" start-placeholder="Start date"
                        end-placeholder="End date" format="YYYY-MM-DD HH:mm:ss" date-format="YYYY/MM/DD ddd"
                        time-format="A hh:mm:ss" />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="queryDriverViolation">查询</el-button>
                </el-form-item>
            </el-form>
            <el-table style="width:100%" :data="driverViolationRes">
            </el-table>
        </el-dialog>
        <el-dialog title="车队违章查询" v-model="queryTeamViolationVisible">
            <el-form :inline="true">
                <el-form-item label="车队名称">
                    <el-input v-model="driverName"></el-input>
                </el-form-item>
                <el-form-item label="起止时间选择">
                    <el-date-picker v-model="driverViolationTime" type="datetimerange" start-placeholder="Start date"
                        end-placeholder="End date" format="YYYY-MM-DD HH:mm:ss" date-format="YYYY/MM/DD ddd"
                        time-format="A hh:mm:ss" />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="queryDriverViolation">查询</el-button>
                </el-form-item>
            </el-form>
            <el-table style="width:100%" :data="driverViolationRes">
            </el-table>
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
import { ElMessage, type Column } from 'element-plus';
const { get, post } = http
const api = "/violation"
const updateRef = (refData: any, _data: any) => { refData.value = _data }
// violation main table
const column: Array<Column> = reactive([
    { prop: 'ID', label: 'ID', width: 0 },
    { prop: 'Name', label: '公司名称', width: 0 },
])
const violationData = ref([])
const fetchViolationData = () => {
    const { data, err } = get(api)
    if (err.value != null) ElMessage.error(err.value)
    else updateRef(violationData, data.value)
}
// insert violation info form
const violationCol = reactive([
    { label: "司机名称", prop: "Name", type: "string" },
    { label: "车辆信息", prop: "Vehicle", type: "string" },
    { label: "车队信息", prop: "Team", type: "string" },
    { label: "路线", prop: "Route", type: "string" },
    { label: "违章时间", prop: "OccurredAt", type: "date" }])
const insertViolationVisible = ref(false)
const insertViolation = (_data: any) => {
    const { err } = post(api, _data)
    if (err.value != null) ElMessage.error(err.value)
    else fetchViolationData()

}
// query violation of drivers
const queryDriverViolationVisible = ref(false)
const driverViolationTime = ref('')
const driverName = ref('')
const driverViolationRes = ref([])
const queryDriverViolation = () => {
    const { data, err } = post('/query/violation/driver', { Name: driverName.value, Time: driverViolationTime.value })
    if (err.value != null) {
        ElMessage.error(err.value)
        return
    }
    driverViolationRes.value = data.value as any
}

//query violation of teams
const queryTeamViolationVisible = ref(false)

onMounted(() => {
    fetchViolationData()
})

</script>
