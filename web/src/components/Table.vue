<template>
    <el-table :data="data">
        <el-table-column v-for="col in column" :key="col.prop" :prop="col.prop" :label="col.label">
        </el-table-column>
        <template #default="{ row }">
            <slot name="end" v-bind:row="row"></slot>
        </template>
    </el-table>
</template>

<script lang="ts" setup>
import { http } from '@/utils/http';
import type { Column } from 'element-plus';
const { get } = http
const data: any = ref(null)
const props = defineProps({
    column: {
        type: Array<Column>,
        default: []
    },
    api: {
        type: String,
        default: ''
    }
})
onMounted(() => {
    const { data, err } = get(props.api)
    if (err) console.log(err)
    update(data.value)
})
const update = (_data: any) => { data.value = _data }
</script>
