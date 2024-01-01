<template>
  <el-container>
    <el-header>
      <el-menu id="nav" mode="horizontal" :ellipsis="false">
        <el-menu-item route="/" index="0">
          <h1>公交调度管理系统</h1>
        </el-menu-item>
      </el-menu>
    </el-header>

    <el-container>
      <el-aside width="200px">
        <el-menu mode="vertical" default-active="0" router>
          <template v-for="(item, index) in items" :key="index">
            <el-menu-item
              :index="`${index}`"
              :route="`/${item.key}`"
              >{{ item.title }}</el-menu-item
            >
          </template>
        </el-menu>
      </el-aside>

      <el-main>
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
const items = reactive([
  {
    key: '',
    title: '首页',
    label: '首页',
  },
  {
    key: 'company',
    title: '公司管理',
    label: '公司管理',
  },
  {
    key: 'vehicle',
    title: '车辆管理',
    label: '车辆管理',
  },
  {
    key: 'team',
    title: '车队管理',
    label: '车队管理',
  },
  {
    key: 'route',
    title: '线路管理',
    label: '线路管理',
  },
  {
    key: 'driver',
    title: '司机管理',
    label: '司机管理',
  },
  {
    key: 'violation',
    title: '违章管理',
    label: '违章管理',
  },
])
</script>

<style>
body {
  margin: 0px;
  padding: 0px;
}
</style>
