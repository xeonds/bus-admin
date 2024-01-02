<template>
  <div class="customer-view">
    <el-row>
      <el-col :span="24">
        <h1>财务</h1>
        <el-alert
          title="统计结果说明"
          type="info"
          description="注意：算法仅供参考，请以实际统计量为准"
          show-icon
        />
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="18">
        <el-card class="card-row">
          <template #header>
            <el-text>成本情况</el-text>
          </template>
          <el-row>
            <el-col :span="6">
              <el-statistic :value="price_purchase">
                <template #title>
                  <div style="display: inline-flex; align-items: center">
                    总进货额
                  </div>
                </template>
                <template #suffix>元</template>
              </el-statistic>
            </el-col>
            <el-col :span="6">
              <el-statistic :value="price_net_income">
                <template #title>
                  <div style="display: inline-flex; align-items: center">
                    盈利情况
                  </div>
                </template>
                <template #suffix>元</template>
              </el-statistic>
            </el-col>
            <el-col :span="6">
              <el-statistic :value="price_debt">
                <template #title>
                  <div style="display: inline-flex; align-items: center">
                    欠款情况
                  </div>
                </template>
                <template #suffix>元</template>
              </el-statistic>
            </el-col>
            <el-col :span="6">
              <el-statistic :value="price_income">
                <template #title>
                  <div style="display: inline-flex; align-items: center">
                    总收入
                  </div>
                </template>
                <template #suffix>元</template>
              </el-statistic>
            </el-col>
          </el-row>
        </el-card>
        <el-card class="card-row">
          <template #header>
            <el-text>营收额统计</el-text>
          </template>
          <BarChart :orders="orderList" />
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="card-row">
          <template #header>
            <el-text>商品总览</el-text>
          </template>
          <el-row>
            <el-col :span="24">
              <el-statistic title="商品总量" :value="productList.length" />
            </el-col>
          </el-row>
        </el-card>
        <el-card class="card-row">
          <template #header>
            <el-text>顾客总览</el-text>
          </template>
          <el-row>
            <el-col :span="12">
              <el-statistic title="所有顾客" :value="customerList.length" />
            </el-col>
            <el-col :span="12">
              <el-statistic title="会员顾客" :value="premierCustomers" />
            </el-col>
          </el-row>
        </el-card>
        <el-card class="card-row">
          <template #header>
            <el-text>热销商品</el-text>
          </template>
          <PieChart :orders="orderList" />
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import customerAPI from '../../api/customer'
import productAPI from '../../api/product'
import purchaseAPI from '../../api/purchase'
import orderAPI from '../../api/order'
import BarChart from '../../components/BarChart.vue'
import PieChart from '../../components/PieChart.vue'

export default {
  components: { BarChart, PieChart },
  data() {
    return {
      customer: {},
      product: {},
      customerList: [],
      productList: [],
      orderList: [],
      purchaseList: [],
      income: {
        labels: ['总收入', '日收入', '月收入'],
        datasets: [
          {
            data: [this.income_totally, this.income_daily, this.income_monthly],
          },
        ],
      },
    }
  },
  computed: {
    filteredCustomerList() {
      return this.customerList.filter((customer) => {
        return (
          customer.name.toLowerCase().includes(this.search.toLowerCase()) ||
          customer.email.toLowerCase().includes(this.search.toLowerCase()) ||
          customer.phone.toLowerCase().includes(this.search.toLowerCase()) ||
          customer.id.toString().includes(this.search)
        )
      })
    },
    premierCustomers() {
      return this.customerList.filter((customer) => customer.is_vip == true)
        .length
    },
    income_totally() {
      return this.orderList.reduce((sum, order) => sum + order.paid, 0)
    },
    income_monthly() {
      return this.orderList
        .filter((order) => {
          const date = new Date()
          const created_at = new Date(order.CreatedAt)
          return created_at.getMonth() == date.getMonth()
        })
        .reduce((sum, order) => sum + order.paid, 0)
    },
    income_daily() {
      return this.orderList
        .filter((order) => {
          const date = new Date()
          const created_at = new Date(order.CreatedAt)
          return created_at.getDate() == date.getDate()
        })
        .reduce((sum, order) => sum + order.paid, 0)
    },
    price_purchase() {
      return this.purchaseList.reduce((sum, purchase) => sum + purchase.paid, 0)
    },
    price_net_income() {
      return (
        this.orderList.reduce((sum, order) => sum + order.paid, 0) -
        this.purchaseList.reduce((sum, purchase) => sum + purchase.debt, 0)
      )
    },
    price_debt() {
      const res =
        this.purchaseList.reduce((sum, purchase) => sum + purchase.debt, 0) -
        this.orderList.reduce((sum, order) => sum + order.paid, 0)
      return res > 0 ? res : 0
    },
    price_income() {
      return this.orderList.reduce((sum, order) => sum + order.paid, 0)
    },
  },
  mounted() {
    this.showCustomerList()
    this.showProdctList()
    this.showOrderList()
    this.showPurchaseList()
  },
  methods: {
    async showCustomerList() {
      await customerAPI
        .getCustomers()
        .then((res) => {
          this.customerList = res
        })
        .catch((err) => {
          this.$message.error('顾客数据加载失败')
        })
    },
    async showProdctList() {
      await productAPI
        .getProducts()
        .then((res) => {
          this.productList = res
        })
        .catch((err) => {
          this.$message.error('商品数据加载失败')
        })
    },
    async showOrderList() {
      await orderAPI
        .getOrders()
        .then((res) => {
          this.orderList = res
        })
        .catch((err) => {
          this.$message.error('订单数据加载失败')
        })
    },
    async showPurchaseList() {
      await purchaseAPI
        .getPurchases()
        .then((res) => {
          this.purchaseList = res
        })
        .catch((err) => {
          this.$message.error('进货数据加载失败')
        })
    },
    querySearch(queryString, cb) {
      const results = queryString
        ? this.productList.filter((product) =>
            product.name.toLowerCase().includes(queryString.toLowerCase())
          )
        : this.productList
      cb(results)
    },
    handleSelect(item) {
      this.product = item
    },
  },
}
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

  .el-row {
    margin-bottom: 2rem;
  }
}
</style>
