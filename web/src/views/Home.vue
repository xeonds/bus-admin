<template>
	<el-scrollbar>
		<el-affix position="top">
			<el-header>
				<el-menu mode="horizontal" :ellipsis="false" :default-active="'1'" router>
					<el-menu-item class="title">商店管理</el-menu-item>
					<el-menu-item index="1" route="/main/dashboard">总览</el-menu-item>
					<el-menu-item index="2" route="/main/customer">客户管理</el-menu-item>
					<el-menu-item index="3" route="/main/product">产品管理</el-menu-item>
					<el-menu-item index="4" route="/main/finance">财务</el-menu-item>
					<div class="flex-grow" />
					<el-menu-item index="5" @click="logout">登出系统</el-menu-item>
				</el-menu>
			</el-header>
		</el-affix>
		<el-main>
			<router-view v-slot="{ Component }">
				<transition>
					<component :is="Component" />
				</transition>
			</router-view>
		</el-main>
		<el-footer>
			<el-divider>
				商店管理系统 &copy; 2023
			</el-divider>
		</el-footer>
	</el-scrollbar>
</template>

<script>
import isLogin from '../utils/isLogin';

export default {
	name: 'Home',
	data() {
		return {
			activeIndex: '1'
		}
	},
	methods: {
		handleSelect(key, keyPath) {
			console.log(key, keyPath);
		},
		logout() {
			localStorage.removeItem('token')
			this.$router.push({ path: '/login' })
		}
	},
	mounted() {
		if (isLogin()) {
			if (this.$route.path == '/main') {
				this.$router.push({ path: '/main/dashboard' })
			}
		} else {
			this.$router.push({ path: '/login' })
			this.$message.error('未登录，请先登录')
		}
	}
}
</script>

<style lang="less">
.card-header {
	display: flex;
	justify-content: space-between;
	align-items: center;
}

.flex-grow {
	flex-grow: 1;
}
</style>

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

