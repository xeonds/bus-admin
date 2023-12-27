<template>
    <div>
        <h1>公司管理</h1>
        <Table :column="column" :api="api" />
    </div>
</template>

<script lang="ts" setup>
import type { Column } from 'element-plus';
const api = "/company"
const column: Array<Column> = reactive([
    { prop: 'ID', label: 'ID', width: 4 },
    { prop: 'Name', label: '公司名称', width: 4 },
])
</script>
