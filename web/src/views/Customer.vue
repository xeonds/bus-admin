<template>
  <div class="customer-view">
    <el-row>
      <el-col :span="24">
        <h1>客户管理</h1>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="24">
        <el-card>
          <template #header>
            <div class="card-header">
              <el-text>全部客户</el-text>
              <el-button plain type="primary" @click="handleCreateCustomer"
                >添加客户</el-button
              >
            </div>
          </template>
          <el-table :data="filteredCustomerList" style="width: 100%">
            <el-table-column type="selection" />
            <el-table-column prop="id" label="ID" />
            <el-table-column prop="name" label="姓名" />
            <!-- <el-table-column prop="gender" label="性别"></el-table-column> -->
            <el-table-column prop="phone" label="手机号">
              <template #default="{ row }">
                <el-text>{{
                  row.phone
                    .split('')
                    .map((_, index) => (index < 3 || index > 6 ? _ : '*'))
                    .join('')
                }}</el-text>
              </template>
            </el-table-column>
            <el-table-column prop="order" label="订单历史">
              <template #default="{ row }">
                <el-button type="primary" link @click="handleShowOrder(row)"
                  >查看</el-button
                >
              </template>
            </el-table-column>
            <el-table-column prop="vip_card" label="VIP信息">
              <template #default="{ row }">
                <el-button
                  v-if="row.is_vip"
                  type="primary"
                  link
                  @click="handleEditCustomer(row)"
                  >查看</el-button
                >
                <el-text v-else>非VIP</el-text>
              </template>
            </el-table-column>
            <el-table-column prop="vip_card" label="会员卡余额">
              <template #default="{ row }">
                <el-text v-if="row.is_vip">{{ row.vip_card.balance }}</el-text>
                <el-text v-else>非VIP</el-text>
              </template>
            </el-table-column>
            <el-table-column label="Operation">
              <template #header>
                <el-input v-model="search" placeholder="搜索" />
              </template>
              <template #default="{ row }">
                <el-button link type="primary" @click="handleEditCustomer(row)"
                  >编辑</el-button
                >
                <el-button
                  link
                  type="primary"
                  @click="handleDeleteCustomer(row)"
                  >删除</el-button
                >
              </template>
            </el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>
    <el-dialog
      v-model="dialogVisible.create"
      :title="dialogVisible.edit ? '编辑客户' : '添加客户'"
    >
      <el-form
        ref="customer"
        :model="customer"
        :rules="rules"
        label-width="120px"
      >
        <el-form-item label="姓名" prop="name">
          <el-input v-model="customer.name" />
        </el-form-item>
        <el-form-item label="性别" prop="gender">
          <el-radio-group v-model="customer.gender">
            <el-radio label="男" />
            <el-radio label="女" />
          </el-radio-group>
        </el-form-item>
        <el-form-item label="手机号" prop="phone">
          <el-input v-model="customer.phone" />
        </el-form-item>
        <el-form-item label="是否VIP" prop="is_vip">
          <el-switch v-model="customer.is_vip" />
        </el-form-item>
        <el-form-item v-if="customer.is_vip" label="卡号" prop="id">
          <el-input v-model="customer.vip_card.id" disabled />
        </el-form-item>
        <el-form-item v-if="customer.is_vip" label="会员卡余额" prop="balance">
          <el-input-number
            v-model.number="customer.vip_card.balance"
            :precision="2"
            :step="100"
          />
        </el-form-item>
        <el-form-item v-if="customer.is_vip" label="会员卡折扣" prop="discount">
          <el-input-number
            v-model.number="customer.vip_card.discount"
            :precision="2"
            :step="0.01"
            :min="0"
            :max="1"
          />
        </el-form-item>
        <el-form-item v-if="customer.is_vip" label="会员卡备注" prop="note">
          <el-input v-model="customer.vip_card.note" type="textarea" />
        </el-form-item>
        <el-form-item label="客户备注" prop="note">
          <el-input v-model="customer.note" type="textarea" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm('customer')"
            >保存</el-button
          >
        </el-form-item>
      </el-form>
    </el-dialog>
    <el-dialog v-model="dialogVisible.order" size="large" title="历史订单">
      <el-table :data="customer.order">
        <el-table-column prop="id" label="订单ID" />
        <el-table-column prop="CreatedAt" label="订单时间" />
        <el-table-column prop="product.name" label="商品名称" />
        <el-table-column prop="quantity" label="商品数量" />
        <el-table-column prop="total_price" label="总计花费" />
        <el-table-column prop="paid" label="实付" />
      </el-table>
    </el-dialog>
  </div>
</template>

<script>
import customerAPI from '../../api/customer'

export default {
  data() {
    return {
      customer: {
        name: '',
        email: '',
        phone: '',
        vip_card: {
          id: 0,
          balance: 0,
          discount: 1,
          note: '已禁用',
        },
      },
      rules: {
        name: [{ required: true, message: '请输入客户姓名', trigger: 'blur' }],
        phone: [
          { required: true, message: '输入手机号', trigger: 'blur' },
          {
            pattern: /^\d*$/,
            message: '请输入正确的手机号',
            trigger: ['blur', 'change'],
          },
        ],
      },
      customerList: [],
      dialogVisible: {
        create: false,
        edit: false,
        order: false,
      },
      search: '',
    }
  },
  computed: {
    filteredCustomerList() {
      return this.customerList.filter((customer) => {
        return (
          customer.name.toLowerCase().includes(this.search.toLowerCase()) ||
          customer.phone.toLowerCase().includes(this.search.toLowerCase()) ||
          customer.id.toString().includes(this.search)
        )
      })
    },
    men2Women() {
      return `${
        this.customerList.filter((customer) => {
          customer.sex == 'man'
        }).length
      }/${
        this.customerList.filter((customer) => {
          customer.sex == 'man'
        }).length
      }`
    },
    premierCustomers() {
      return this.customerList.filter((customer) => {
        customer.premier == true
      }).length
    },
  },
  mounted() {
    this.showCustomerList()
  },
  methods: {
    async showCustomerList() {
      await customerAPI
        .getCustomers()
        .then((res) => {
          this.customerList = res
        })
        .catch((error) => {
          this.$message.error('获取客户数据失败')
        })
    },
    async submitForm(formName) {
      try {
        await this.$refs[formName].validate()
        if (this.dialogVisible.edit) {
          await customerAPI.updateCustomer(this.customer.id, this.customer)
          this.$message.success('修改顾客信息成功')
        } else {
          await customerAPI.createCustomer(this.customer)
          this.$message.success('创建顾客成功')
        }
        await this.showCustomerList()
      } catch (error) {
        this.$message.error(`创建顾客失败: ${error}` || '未知错误')
      } finally {
        this.resetForm(formName)
        this.dialogVisible.create = false
        this.dialogVisible.edit = false
      }
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
    },
    handleCreateCustomer() {
      this.customer = Object.assign({}, this.customer)
      this.dialogVisible.create = true
    },
    handleEditCustomer(customer) {
      this.customer = customer
      this.dialogVisible.create = true
      this.dialogVisible.edit = true
    },
    handleDeleteCustomer(customer) {
      this.$confirm(`确认删除顾客 ${customer.name}?`, '注意', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          try {
            await customerAPI.deleteCustomer(customer.id)
            this.$message.success('删除成功')
            await this.showCustomerList()
          } catch (error) {
            this.$message.error(`删除顾客失败: ${error}` || '未知错误')
          }
        })
        .catch(() => {
          this.$message.info('取消删除')
        })
    },
    handleShowOrder(customer) {
      this.customer = customer
      this.dialogVisible.order = true
    },
  },
}
</script>

<style lang="less" scoped>
.customer-view {
  margin: 0 auto;

  .el-row {
    margin-bottom: 2rem;
  }
}
</style>
