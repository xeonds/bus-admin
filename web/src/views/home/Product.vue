<template>
  <div class="product-view">
    <el-row>
      <el-col :span="24">
        <h1>商品管理</h1>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="24">
        <el-card>
          <template #header>
            <div class="card-header">
              <el-text>全部商品</el-text>
              <el-button plain type="primary" @click="handleCreate">添加商品</el-button>
            </div>
          </template>
          <el-table :data="filteredProductList" style="width: 100%" max-height="512" show-summary
            :summary-method="getProductSummary">
            <el-table-column prop="id" label="ID"></el-table-column>
            <el-table-column prop="name" label="产品名称"></el-table-column>
            <el-table-column prop="brand" label="品牌"></el-table-column>
            <el-table-column prop="location" label="产地"></el-table-column>
            <el-table-column prop="description" label="描述"></el-table-column>
            <el-table-column prop="price" label="价格（元/斤）"></el-table-column>
            <el-table-column prop="purchase" label="进货量（斤）"></el-table-column>
            <el-table-column prop="sale" label="销售量（斤）"></el-table-column>
            <el-table-column prop="quantity" label="库存（斤）"></el-table-column>
            <el-table-column prop="note" label="备注"></el-table-column>
            <el-table-column label="操作">
              <template #header>
                <el-input v-model="search.product" placeholder="搜索" />
              </template>
              <template #default="{ row }">
                <el-button link type="primary" @click="handleEdit(row)">编辑</el-button>
                <el-button link type="primary" @click="handleDelete(row)">删除</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="24">
        <el-card>
          <template #header>
            <div class="card-header">
              <el-text>销售记录</el-text>
            </div>
          </template>
          <el-table :data="filteredOrderList" style="width: 100%" max-height="512" show-summary
            :summary-method="getOrderSummary">
            <el-table-column prop="id" label="ID"></el-table-column>
            <el-table-column prop="product.name" label="货物名称"></el-table-column>
            <el-table-column prop="product.brand" label="品牌"></el-table-column>
            <el-table-column prop="quantity" label="销售数量（斤）"></el-table-column>
            <el-table-column prop="product.price" label="销售单价（元/斤）"></el-table-column>
            <el-table-column prop="total_price" label="金额"></el-table-column>
            <el-table-column prop="discount" label="优惠金额"></el-table-column>
            <el-table-column prop="paid" label="实收"></el-table-column>
            <el-table-column prop="payment" label="支付方式" :filters="[
              { text: '现金', value: '现金' },
              { text: '微信', value: '微信' },
              { text: '支付宝', value: '支付宝' },
              { text: '会员卡', value: '会员卡' },
            ]" :filter-method="(value, row) => row.payment == value"></el-table-column>
            <el-table-column prop="note" label="备注"></el-table-column>
            <el-table-column label="操作">
              <template #header>
                <el-input v-model="search.order" placeholder="搜索" />
              </template>
              <template #default="{ row }">
                <el-button link type="primary" @click="handleRecallOrder(row)" :disabled="row.recall">撤单</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="24">
        <el-card>
          <template #header>
            <div class="card-header">
              <el-text>采购记录</el-text>
              <div style="display: inline-flex; align-items: center">
                <el-button plain type="primary" @click="handlePurchase">进货</el-button>
                <el-button plain type="primary" @click="() => is_table_show = !is_table_show">切换显示</el-button>
              </div>
            </div>
          </template>
          <el-table :data="filteredPurchaseList" style="width: 100%" :max-height="512" show-summary
            :summary-method="getPurchaseSummary" v-if="is_table_show">
            <el-table-column prop="id" label="ID"></el-table-column>
            <el-table-column prop="product.name" label="货物名称"></el-table-column>
            <el-table-column prop="product.brand" label="品牌"></el-table-column>
            <el-table-column prop="product.location" label="产地"></el-table-column>
            <el-table-column prop="supplier" label="供货单位"></el-table-column>
            <el-table-column prop="quantity" label="采购数量（斤）"></el-table-column>
            <el-table-column prop="unit_price" label="采购单价（元/斤）"></el-table-column>
            <el-table-column prop="total_price" label="金额"></el-table-column>
            <el-table-column prop="paid" label="已付金额"></el-table-column>
            <el-table-column prop="debt" label="欠款"></el-table-column>
            <el-table-column prop="note" label="备注"></el-table-column>
            <el-table-column label="操作">
              <template #header>
                <el-input v-model="search.purchase" placeholder="搜索" />
              </template>
              <template #default="{ row }">
                <el-button link type="primary" @click="handleRecallPurchase(row)" :disabled="row.recall">退货</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>
    <el-dialog v-model="dialogVisible.create" :title="dialogVisible.edit ? '编辑商品' : '添加商品'">
      <el-form :model="product" :rules="rules" label-width="120px" ref="product">
        <el-form-item prop="name" label="产品名称">
          <el-input v-model="product.name"></el-input>
        </el-form-item>
        <el-form-item prop="brand" label="品牌">
          <el-input v-model="product.brand"></el-input>
        </el-form-item>
        <el-form-item prop="location" label="产地">
          <el-input v-model="product.location"></el-input>
        </el-form-item>
        <el-form-item prop="description" label="描述">
          <el-input v-model="product.description" type="textarea"></el-input>
        </el-form-item>
        <el-form-item prop="price" label="价格">
          <el-input-number v-model.number="product.price" :min="0" :step="0.10" :precision="2"></el-input-number>
        </el-form-item>
        <el-form-item prop="quantity" label="库存">
          <el-input-number v-model.number="product.quantity" :min="0" :precision="2"></el-input-number>
        </el-form-item>
        <el-form-item prop="note" label="备注">
          <el-input v-model="product.note" type="textarea"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm('product')">Save</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
    <el-dialog v-model="dialogVisible.purchase" title="进货">
      <el-form :model="purchase" :rules="rules" label-width="120px" ref="purchase">
        <el-form-item label="商品">
          <el-select v-model="purchase.product" value-key="id" placeholder="请选择进货的商品">
            <el-option v-for="product in productList" :key="product.id" :label="product.name" :value="product" />
          </el-select>
        </el-form-item>
        <el-form-item prop="supplier" label="供货单位">
          <el-input v-model="purchase.supplier"></el-input>
        </el-form-item>
        <el-form-item prop="unit_price" label="采购单价">
          <el-input-number v-model.number="purchase.unit_price" :min="0" :step="0.10" :precision="2"></el-input-number>
        </el-form-item>
        <el-form-item prop="quantity" label="采购量">
          <el-input-number v-model.number="purchase.quantity" :min="0" :precision="2"></el-input-number>
        </el-form-item>
        <el-form-item prop="total_price" label="总计金额">
          <el-input-number v-model.number="purchase.total_price" :min="0" :precision="2"></el-input-number>
        </el-form-item>
        <el-form-item prop="total_price" label="已付金额">
          <el-input-number v-model.number="purchase.paid" :min="0" :precision="2"></el-input-number>
        </el-form-item>
        <el-form-item prop="total_price" label="欠款">
          <el-input-number v-model.number="purchase.debt" :min="0" :precision="2"></el-input-number>
        </el-form-item>
        <el-form-item prop=" note" label="备注">
          <el-input v-model="purchase.note" type="textarea"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm('purchase')">Save</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
import productAPI from '../../api/product';
import purchaseAPI from '../../api/purchase';
import orderAPI from '../../api/order';

export default {
  data() {
    return {
      product: {
        name: '',
        description: '',
        price: 0.5,
        count: 0
      },
      order: {
        id: 0,
        time: '',
        product: {
          name: '',
          description: '',
          price: 0.5,
          count: 0
        },
        count: 0,
        cost: 0
      },
      purchase: {
        id: 0,
        time: '',
        product: {},
        count: 0,
        cost: 0
      },
      rules: {
        name: [
          { required: true, message: '请输入客户姓名', trigger: 'blur' }
        ],
        phone: [
          { required: true, message: '输入手机号', trigger: 'blur' },
          { pattern: /^[0-9]*$/, message: '请输入正确的手机号', trigger: ['blur', 'change'] }
        ]
      },
      productList: [],
      orderList: [],
      purchaseList: [],
      dialogVisible: {
        create: false,
        edit: false,
        purchase: false
      },
      search: {
        product: '',
        order: '',
        purchase: ''
      },
      is_table_show: false
    }
  },
  computed: {
    filteredProductList: function () {
      return this.productList.filter((product) => {
        return (
          product.name.toLowerCase().includes(this.search.product.toLowerCase()) ||
          product.description.toLowerCase().includes(this.search.product.toLowerCase()) ||
          product.price.toString().includes(this.search.product) ||
          product.location.toLowerCase().includes(this.search.product.toLowerCase())
        );
      });
    },
    filteredOrderList: function () {
      return this.orderList.filter((order) => {
        return (
          order.product.name.toLowerCase().includes(this.search.order.toLowerCase()) ||
          order.payment.toLowerCase().includes(this.search.order.toLowerCase())
        );
      });
    },
    filteredPurchaseList: function () {
      return this.purchaseList.filter((purchase) => {
        return (
          purchase.product.name.toLowerCase().includes(this.search.purchase.toLowerCase()) ||
          purchase.product.brand.toLowerCase().includes(this.search.purchase.toLowerCase()) ||
          purchase.supplier.toLowerCase().includes(this.search.purchase.toLowerCase())
        );
      });
    },
  },
  methods: {
    async showProductList() {
      await productAPI.getProducts().then((res) => {
        this.productList = res
      }).catch((err) => {
        this.$message.error('获取商品列表失败')
      })
    },
    async showOrderList() {
      await orderAPI.getOrders().then((res) => {
        this.orderList = res
      }).catch((err) => {
        this.$message.error('获取订单列表失败')
      })
    },
    async showPurchaseList() {
      await purchaseAPI.getPurchases().then((res) => {
        this.purchaseList = res
      }).catch((err) => {
        this.$message.error('获取采购列表失败')
      })
    },
    async submitForm(formName) {
      try {
        await this.$refs[formName].validate()
        switch (formName) {
          case 'product':
            if (this.dialogVisible.edit) {
              await productAPI.updateProduct(this.product.id, this.product)
              this.$message.success('产品更新成功')
            } else {
              await productAPI.createProduct(this.product)
              this.$message.success('产品创建成功')
            }
            await this.showProductList()
            break
          case 'purchase':
            await purchaseAPI.createPurchase(this.purchase)
            this.$message.success('采购成功')
            await this.showPurchaseList()
            await this.showProductList()
            break
        }
      } catch (error) {
        this.$message.error('产品创建失败: ' + error || '未知错误')
      } finally {
        this.resetForm(formName)
        this.dialogVisible.create = false
        this.dialogVisible.edit = false
        this.dialogVisible.purchase = false
      }
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
    },
    handleCreate() {
      this.dialogVisible.create = true
    },
    handleEdit(product) {
      this.product = product
      this.dialogVisible.create = true
      this.dialogVisible.edit = true
    },
    handleDelete(product) {
      this.$confirm(`是否要删除 ${product.name}?`, '注意', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async () => {
        try {
          await productAPI.deleteProduct(product.id)
          this.$message.success('删除成功')
          await this.showProductList()
        } catch (error) {
          this.$message.error('删除失败: ' + error || '未知错误')
        }
      }).catch(() => {
        this.$message.info('已取消删除')
      })
    },
    handlePurchase() {
      this.dialogVisible.purchase = true
    },
    handleRecallPurchase(purchase) {
      this.$confirm(`是否要撤销采购 ${purchase.product.name}?`, '注意', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async () => {
        try {
          await purchaseAPI.updatePurchase(purchase.id, {})
          this.$message.success('撤销采购成功')
          await this.showPurchaseList()
          await this.showProductList()
        } catch (error) {
          this.$message.error('撤销采购失败: ' + error || '未知错误')
        }
      }).catch(() => {
        this.$message.info('已取消撤销采购')
      })
    },
    handleRecallOrder(order) {
      this.$confirm(`是否要撤销订单 ${order.id}?`, '注意', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async () => {
        try {
          await orderAPI.updateOrder(order.id, {})
          this.$message.success('撤销订单成功')
          await this.showOrderList()
          await this.showProductList()
        } catch (error) {
          this.$message.error('撤销订单失败: ' + error || '未知错误')
        }
      }).catch(() => {
        this.$message.info('已取消撤销订单')
      })
    },
    getProductSummary(param) {
      const { columns, data } = param
      const sums = []
      columns.forEach((column, index) => {
        switch (index) {
          case 0:
            sums[index] = '合计'
            return
          case 5:
          case 6:
          case 7:
          case 8:
            sums[index] = data.reduce((sum, item) => sum + item[column.property], 0).toFixed(2)
            return
        }
      })
      return sums
    },
    getOrderSummary(param) {
      const { columns, data } = param
      const sums = []
      columns.forEach((column, index) => {
        switch (index) {
          case 0:
            sums[index] = '合计'
            return
          case 3:
          case 5:
          case 6:
          case 7:
            sums[index] = data.reduce((sum, item) => sum + item[column.property], 0).toFixed(2)
            return
        }
      })
      return sums
    },
    getPurchaseSummary(param) {
      const { columns, data } = param
      const sums = []
      columns.forEach((column, index) => {
        switch (index) {
          case 0:
            sums[index] = '合计'
            return
          case 5:
          case 6:
          case 7:
          case 8:
          case 9:
            sums[index] = data.reduce((sum, item) => sum + item[column.property], 0).toFixed(2)
            return
        }
      })
      return sums
    }
  },
  mounted() {
    this.showProductList()
    this.showPurchaseList()
    this.showOrderList()
  }
}
</script>

<style lang="less" scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.product-view {
  margin: 0 auto;

  .el-row {
    margin-bottom: 2rem;
  }
}
</style>

