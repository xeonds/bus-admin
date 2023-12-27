<template>
    <Bar id="my-chart-id" :options="chartOptions" :data="chartData" />
</template>

<script>
import { Bar } from 'vue-chartjs'
import { Chart as ChartJS, Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale } from 'chart.js'

ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale)

export default {
    props: {
        orders: {
            type: Object,
            required: true,
            default: () => {
                return {}
            }
        },
    },
    components: { Bar },
    data() {
        return {
            chartData: {
                labels: ['January', 'February', 'March'],
                datasets: [{ data: [40, 20, 12] }]
            },
            chartOptions: {
                responsive: true
            }
        }
    },
    mounted() {
        this.renderChart();
    },
    watch: {
        orders() {
            this.renderChart();
        },
    },
    methods: {
        renderChart() {
            const data = {};
            //将订单按照月份分类
            this.orders.forEach(order => {
                const month = order.CreatedAt.split('-')[1];
                if (data[month]) {
                    data[month] += order.paid;
                } else {
                    data[month] = order.paid;
                }
            });
            let labels = Object.keys(data);
            let values = Object.values(data);
            this.chartData = {
                labels,
                datasets: [
                    {
                        label: '月营收额',
                        data: values,
                    },
                ],
            };
        },
    },
}
</script>