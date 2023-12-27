<template>
  <el-container>
    <!-- 顶部栏 -->
    <el-header>
      <el-menu id="nav" mode="horizontal" :ellipsis="false">
        <el-menu-item route="/" index="0">
          <h1>公交调度管理系统</h1>
        </el-menu-item>
      </el-menu>
    </el-header>

    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <el-menu mode="vertical" default-active="1">
          <template v-for="(item, index) in items" :key="index">
            <el-menu-item
              v-if="!item.children"
              :index="`${index}`"
              :route="`/admin/${item.key}`"
              >{{ item.title }}</el-menu-item
            >
            <el-sub-menu v-else :index="`${index}`">
              <template #title>
                <span slot="title">{{ item.title }}</span>
              </template>
              <el-menu-item
                v-for="(child, index2) in item.children"
                :key="index2"
                :index="`${index}-${index2}`"
                :route="`/admin/${child.key}`"
              >
                {{ child.title }}
              </el-menu-item>
            </el-sub-menu>
          </template>
        </el-menu>
      </el-aside>

      <!-- 主要内容区域 -->
      <el-main>
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
const router = useRouter()
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
    children: [
      {
        key: 'car',
        title: '汽车管理',
        label: '汽车管理',
      },
      {
        key: 'list',
        title: '车队管理',
        label: '车队管理',
      },
    ],
  },
  {
    key: 'lists',
    title: '车队管理',
    label: '车队管理',
    children: [
      {
        key: 'leader',
        title: '队长管理',
        label: '队长管理',
      },
    ],
  },
  {
    key: 'roads',
    title: '线路管理',
    label: '线路管理',
    children: [
      {
        key: 'driver',
        title: '司机管理',
        label: '司机管理',
      },
    ],
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
