<template>
  <Bar id="my-chart-id" :options="chartOptions" :data="chartData" />
</template>

<script>
import { Bar } from 'vue-chartjs'
import {
  BarElement,
  CategoryScale,
  Chart as ChartJS,
  Legend,
  LinearScale,
  Title,
  Tooltip,
} from 'chart.js'

ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale)

export default {
  components: { Bar },
  props: {
    orders: {
      type: Object,
      required: true,
      default: () => {
        return {}
      },
    },
  },
  data() {
    return {
      chartData: {
        labels: ['January', 'February', 'March'],
        datasets: [{ data: [40, 20, 12] }],
      },
      chartOptions: {
        responsive: true,
      },
    }
  },
  watch: {
    orders() {
      this.renderChart()
    },
  },
  mounted() {
    this.renderChart()
  },
  methods: {
    renderChart() {
      const data = {}
      //将订单按照月份分类
      this.orders.forEach((order) => {
        const month = order.CreatedAt.split('-')[1]
        if (data[month]) {
          data[month] += order.paid
        } else {
          data[month] = order.paid
        }
      })
      const labels = Object.keys(data)
      const values = Object.values(data)
      this.chartData = {
        labels,
        datasets: [
          {
            label: '月营收额',
            data: values,
          },
        ],
      }
    },
  },
}
</script>
