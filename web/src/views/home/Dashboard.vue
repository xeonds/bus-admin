<template>
  <div class="customer-view">
    <el-row>
      <el-col :span="24">
        <h1>总览</h1>
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="18">
        <el-card class="card-row">
          <template #header>
            <el-text>创建订单</el-text>
          </template>
          <el-form :model="order" ref="orderCreate" :label-width="120">
            <el-row>
              <el-col :span="12">
                <el-form-item label="客户姓名">
                  <el-select v-model="order.customer" filterable remote reserve-keyword placeholder="输入以快速查找"
                    :remote-method="remoteMethod" value-key="id" clearable>
                    <el-option v-for="item in customerList" :key="item.id" :label="item.name" :value="item">
                      <el-row :gutter="40">
                        <el-col :span="8">姓名：{{ item.name }}</el-col>
                        <el-col :span="8">电话：{{ item.phone }}</el-col>
                        <el-col :span="8">余额：{{ item.vip_card.balance.toFixed(2) }}</el-col>
                      </el-row>
                    </el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="商品名称">
                  <el-select v-model="order.product" value-key="id" placeholder="请选择购买的商品">
                    <el-option v-for="product in productList" :key="product.id" :label="product.name" :value="product">
                      <el-row :gutter="40">
                        <el-col :span="8">{{ product.name }}</el-col>
                        <el-col :span="8">库存：{{ product.quantity.toFixed(2) }}</el-col>
                      </el-row>
                    </el-option>
                  </el-select>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="12">
                <el-form-item label="品牌">
                  <el-input v-model="order.product.brand" disabled />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="单价">
                  <el-input v-model.number="order.product.price" :precision="2" disabled />
                </el-form-item>
              </el-col>
            </el-row>
            <el-form-item label="购买数量（斤）">
              <el-input-number v-model.number="order.quantity" :min="0" :precision="2" :step="0.10" />
            </el-form-item>
            <el-form-item label="总价">
              <el-text> {{ (order.product.price * order.quantity).toFixed(2) }} 元</el-text>
            </el-form-item>
            <el-form-item label="优惠金额">
              <el-text>
                {{ (order.product.price * order.quantity * (1 - order.customer.vip_card.discount)).toFixed(2) }}
                元</el-text>
            </el-form-item>
            <el-form-item label="实付">
              <el-text> {{ (order.product.price * order.quantity * order.customer.vip_card.discount).toFixed(2) }}
                元</el-text>
            </el-form-item>
            <el-form-item label="支付方式">
              <el-radio-group v-model="order.payment">
                <el-radio-button label="现金" value="现金" />
                <el-radio-button label="微信" value="微信" />
                <el-radio-button label="支付宝" value="支付宝" />
                <el-radio-button label="会员卡" value="会员卡" />
              </el-radio-group>
            </el-form-item>
            <el-form-item prop="note" label="备注">
              <el-input v-model="order.note" type="textarea" />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="submitForm('orderCreate')" plain>保存</el-button>
            </el-form-item>
          </el-form>
        </el-card>
        <el-card class="card-row">
          <template #header>
            <el-text>客户充值</el-text>
          </template>
          <el-form :inline="true" :model="customer" ref="customerRecharge">
            <el-form-item label="客户姓名">
              <el-select v-model="customer" filterable remote reserve-keyword placeholder="输入以快速查找"
                :remote-method="remoteMethod" value-key="id" clearable>
                <el-option v-for="item  in  customerList" :key="item.id" :label="item.name" :value="item"
                  class="customer-selector">
                  <span class="left">姓名 {{ item.name }}</span>
                  <span class="right">手机号 {{ item.phone }}</span>
                </el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="充值金额">
              <el-input-number v-model="customer.balance" :min="0" :precision="2" :step="100" />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="submitForm('customerRecharge')" plain>保存</el-button>
            </el-form-item>
          </el-form>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="card-row">
          <template #header>
            <el-text>订单总览</el-text>
          </template>
          <el-row>
            <el-col :span="12">
              <el-statistic title="今日订单数" :value="ordersCountToday" />
            </el-col>
            <el-col :span="12">
              <el-statistic :value="incomeToday" :precision="2">
                <template #title>
                  <div style="display: inline-flex; align-items: center">
                    总进货额
                  </div>
                </template>
                <template #suffix>元</template>
              </el-statistic>
            </el-col>
          </el-row>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import customerAPI from "../../api/customer";
import orderAPI from "../../api/order";
import productAPI from "../../api/product";

export default {
  data() {
    return {
      customer: {},
      product: {},
      order: {
        customer: {
          id: 0,
          name: "",
          phone: "",
          vip_card: {
            id: 0,
            balance: 0,
            discount: 0,
            note: "",
          },
          note: "",
        },
        customer_id: 0,
        product: {},
        quantity: 0,
        total_price: 0,
        discount: 0,
        paid: 0,
        payment: "",
        note: "",
      },
      customerList: [],
      productList: [],
      ordersToday: [],
    };
  },
  computed: {
    ordersCountToday() {
      return this.ordersToday.length
    },
    incomeToday() {
      return this.ordersToday.reduce((sum, order) => sum + order.paid, 0)
    },
  },
  methods: {
    async showCustomerList() {
      await customerAPI.getCustomers().then((res) => {
        this.customerList = res;
      }).catch((err) => {
        this.$message.error("顾客数据加载失败");
      });
    },
    async showProdctList() {
      await productAPI.getProducts().then((res) => {
        this.productList = res;
      }).catch((err) => {
        this.$message.error("商品数据加载失败");
      });
    },
    async showOrderList() {
      await orderAPI.getOrders().then((res) => {
        this.orderList = res;
      }).catch((err) => {
        this.$message.error("订单数据加载失败");
      });
    },
    async submitForm(formName) {
      try {
        await this.$refs[formName].validate()
        switch (formName) {
          case 'orderCreate':
            this.order.customer_id = this.order.customer.id
            this.order.total_price = ((this.order.product.price * this.order.quantity) * 100).toFixed(0) / 100
            this.order.discount = ((this.order.total_price * (1 - this.order.customer.vip_card.discount)) * 100).toFixed(0) / 100
            this.order.paid = ((this.order.total_price * this.order.customer.vip_card.discount) * 100).toFixed(0) / 100
            await orderAPI.createOrder(this.order)
            await this.showCustomerList()
            await this.showProdctList()
            await this.getOrdersToday()
            this.$message.success('订单创建成功')
            break
          case 'customerRecharge':
            this.customer.vip_card.balance += this.customer.balance
            await customerAPI.updateCustomer(this.customer.id, this.customer)
            await this.showCustomerList()
            this.$message.success('充值成功，余额为 ' + this.customer.vip_card.balance + ' 元')
            break
        }
      } catch (error) {
        this.$message.error('操作失败: ' + error || '未知错误')
      } finally {
        this.resetForm(formName)
      }
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
    },
    async getOrdersToday() {
      await orderAPI.getOrdersToday().then((res) => {
        this.ordersToday = res;
      }).catch((err) => {
        this.$message.error("今日订单数据加载失败");
      });
    },
  },
  mounted() {
    this.showCustomerList();
    this.showProdctList();
    this.getOrdersToday();
  },
};
</script>

<style lang="less" scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.card-row {
  margin-bottom: 2rem;
}

.customer-view {
  margin: 0 auto;
}

.customer-selector {
  .left {
    float: left;
  }

  .right {
    float: right;
    color: var(--el-text-color-secondary);
    font-size: 13px;
  }
}
</style>
