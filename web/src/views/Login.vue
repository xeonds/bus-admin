<template>
	<el-container class="container">
		<el-card class="login-card">
			<template #header>
				<el-text>登录</el-text>
			</template>
			<el-form :model="form" :rules="rules" ref="form">
				<el-form-item label="用户名" prop="username">
					<el-input placeholder="请输入用户名" v-model="form.username"></el-input>
				</el-form-item>
				<el-form-item label="密码" prop="password">
					<el-input placeholder="请输入密码" v-model="form.password" type="password"></el-input>
				</el-form-item>
				<el-form-item>
					<el-button type="primary" plain @click="onSubmit">登录</el-button>
				</el-form-item>
			</el-form>
		</el-card>
	</el-container>
</template>

<script>
import authAPI from '../api/auth';
import isLogin from '../utils/isLogin';

export default {
	data() {
		return {
			form: {
				username: '',
				password: ''
			},
			rules: {
				username: [
					{ required: true, message: '请输入用户名', trigger: 'blur' },
					{ min: 3, max: 10, message: '长度在 3 到 10 个字符', trigger: 'blur' }
				],
				password: [
					{ required: true, message: '请输入密码', trigger: 'blur' },
					{ min: 3, max: 20, message: '长度在 6 到 20 个字符', trigger: 'blur' }
				]
			}
		};
	},
	methods: {
		async onSubmit() {
			try {
				await this.$refs.form.validate()
				const res = await authAPI.login(this.form.username, this.form.password)
				this.$store.commit('SET_TOKEN', res.token)
				localStorage.setItem('token', res.token)
				this.$message.success('登录成功')
				this.$router.push({ path: '/main/dashboard' })
			} catch (error) {
				this.$message.error('Login failed: ' + error.message)
			}
		}
	},
	mounted() {
		window.addEventListener('keydown', (e) => {
			if (e.key == 'Enter') {
				this.onSubmit()
			}
		})
		if (isLogin()) {
			this.$router.push({ path: '/main/dashboard' })
			this.$message.success('欢迎回来')
		}
	}
};
</script>

<style lang="less" scoped>
.el-container {
	display: flex;
	justify-content: center;
	align-items: center;
	height: 100vh;
	background-image: url('../assets/bg.png');
	background-size: cover;
	background-repeat: no-repeat;
}
</style>
