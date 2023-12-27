<template>
  <el-scrollbar>
    <el-affix position="top">
      <el-header>
        <el-menu mode="horizontal" :default-active="'1'" router>
          <el-menu-item class="title">Shop Manager</el-menu-item>
          <el-menu-item index="1" route="/main/dashboard"
            >Dashboard</el-menu-item
          >
          <el-menu-item index="2" route="/main/customer"
            >Customer Manage</el-menu-item
          >
          <el-menu-item index="3" route="/main/store"
            >Store Manage</el-menu-item
          >
        </el-menu>
      </el-header>
    </el-affix>
    <el-main>
      <el-text>404 - 你要找的页面不见了</el-text>
      <el-text>{{ count }}秒后自动跳回首页...</el-text>
    </el-main>
    <el-footer>
      <el-divider> Shop Manager &copy; 2023 </el-divider>
    </el-footer>
  </el-scrollbar>
</template>

<script>
export default {
  name: 'Home',
  data() {
    return {
      activeIndex: '1',
      count: 5,
    }
  },
  mounted() {
    this.timer = setInterval(() => {
      this.count--
      if (this.count <= 0) {
        clearInterval(this.timer)
        this.$router.push('/main/dashboard')
      }
    }, 1000)
  },
}
</script>

<style lang="less" scoped>
.el-scrollbar {
  height: 100vh;
  width: 100vw;

  .el-header {
    padding: 0;
  }
}

.title {
  font-size: 20px;
  font-weight: bold;
}
</style>
