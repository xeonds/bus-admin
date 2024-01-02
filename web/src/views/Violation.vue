<template>
  <div>
    <h1>违章管理</h1>
    <el-row>
      <el-button @click="insertViolationVisible = true">违章信息录入</el-button>
      <el-button @click="queryDriverViolationVisible = true"
        >司机时段违章查询</el-button
      >
      <el-button @click="queryTeamViolationVisible = true"
        >车队时段违章查询</el-button
      >
    </el-row>
    <el-dialog v-model="insertViolationVisible" title="违章信息录入">
      <Form
        :col="violationCol"
        @submit="
          (data) => {
            insertViolation(data)
          }
        "
      />
    </el-dialog>
    <el-dialog v-model="queryDriverViolationVisible" title="司机违章查询">
      <el-form :inline="true">
        <el-form-item label="司机姓名">
          <el-input v-model="driverName" />
        </el-form-item>
        <el-form-item label="起止时间选择">
          <el-date-picker
            v-model="driverViolationTime"
            type="datetimerange"
            start-placeholder="Start date"
            end-placeholder="End date"
            format="YYYY-MM-DD HH:mm:ss"
            date-format="YYYY/MM/DD ddd"
            time-format="A hh:mm:ss"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="queryDriverViolation"
            >查询</el-button
          >
        </el-form-item>
      </el-form>
      <el-table style="width: 100%" :data="driverViolationRes">
        <el-table-column
          v-for="item in column"
          :key="item.prop"
          :prop="item.prop"
          :label="item.label"
          :width="item.width"
        />
      </el-table>
    </el-dialog>
    <el-dialog v-model="queryTeamViolationVisible" title="车队违章查询">
      <el-form :inline="true">
        <el-form-item label="车队名称">
          <el-input v-model="teamName" />
        </el-form-item>
        <el-form-item label="起止时间选择">
          <el-date-picker
            v-model="teamViolationTime"
            type="datetimerange"
            start-placeholder="Start date"
            end-placeholder="End date"
            format="YYYY-MM-DD HH:mm:ss"
            date-format="YYYY/MM/DD ddd"
            time-format="A hh:mm:ss"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="queryTeamViolation">查询</el-button>
        </el-form-item>
      </el-form>
      <el-table style="width: 100%" :data="teamViolationRes">
        <el-table-column
          v-for="item in column"
          :key="item.prop"
          :prop="item.prop"
          :label="item.label"
          :width="item.width"
        />
      </el-table>
    </el-dialog>
    <el-table :data="violationData">
      <el-table-column
        v-for="item in column"
        :key="item.prop"
        :prop="item.prop"
        :label="item.label"
        :width="item.width"
      />
      <el-table-column prop="actions" label="操作">
        <template #default="{ row }">
          <el-button @click="deleteViolation(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script lang="ts" setup>
// data action
import { type Column, ElMessage } from 'element-plus'
import { addCascade } from '@/utils/common'
import { deleteData, dialogPost, fetchData, http } from '@/utils/http'
const { post } = http
const api = '/violation'
// violation main table
const column: Array<Column> = reactive([
  { prop: 'Driver', label: '司机', width: 0 },
  { prop: 'Vehicle', label: '车辆', width: 0 },
  { prop: 'Team', label: '车队', width: 0 },
  { prop: 'Route', label: '路线', width: 0 },
  { prop: 'OccurredAt', label: '违章时间', width: 0 },
  { prop: 'ViolationType', label: '违章类型', width: 0 },
])
const violationData = ref([])
// insert violation info form
const violationCol = reactive([
  { label: '违章时间', prop: 'OccurredAt', type: 'datetime' },
  { label: '违章类型', prop: 'ViolationType', type: 'string' },
])
const insertViolationVisible = ref(false)
const insertViolation = async (_data: any) =>
  dialogPost(api, _data, insertViolationVisible, violationData)
const deleteViolation = async (_data: any) =>
  deleteData(api, _data.ID, violationData)
// query violation of drivers
const queryDriverViolationVisible = ref(false)
const driverViolationTime = ref('')
const driverName = ref('')
const driverViolationRes = ref([])
const queryDriverViolation = async () => {
  const { data, err } = await post('query/violation/driver', {
    Name: driverName.value,
    Time: driverViolationTime.value,
  })
  if (err.value != null) {
    ElMessage.error(err.value)
  } else driverViolationRes.value = data.value as any
}
//query violation of teams
const queryTeamViolationVisible = ref(false)
const teamViolationTime = ref('')
const teamName = ref('')
const teamViolationRes = ref([])
const queryTeamViolation = async () => {
  const { data, err } = await post('query/violation/team', {
    Name: teamName.value,
    Time: teamViolationTime.value,
  })
  if (err.value != null) {
    ElMessage.error(err.value)
  } else teamViolationRes.value = data.value as any
}

onMounted(() => {
  fetchData(api, violationData)
  addCascade('/driver', violationCol, '司机', 'Driver', (item: any) => {
    return { label: item.Name, value: item }
  })
  addCascade('/vehicle', violationCol, '车辆信息', 'Vehicle', (item: any) => {
    return { label: item.VIN, value: item }
  })
  addCascade('/team', violationCol, '车队信息', 'Team', (item: any) => {
    return { label: item.Name, value: item }
  })
  addCascade('/route', violationCol, '路线', 'Route', (item: any) => {
    return { label: item.Name, value: item }
  })
})
</script>
