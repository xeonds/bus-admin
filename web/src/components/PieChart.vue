<template>
    <Pie id="chart-id" :options="chartOptions" :data="chartData" />
</template>

<script>
import { Pie } from 'vue-chartjs';
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js'

ChartJS.register(ArcElement, Tooltip, Legend)

export default {
    props: {
        orders: {
            type: Array,
            required: true,
        },
    },
    components: { Pie },
    data() {
        return {
            chartData: {
                labels: [],
                datasets: []
            },
            chartOptions: {
                responsive: true,
                maintainAspectRatio: false,
            },
        };
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
            const products = {};
            this.orders.forEach(order => {
                const productId = order.product.name;
                if (products[productId]) {
                    products[productId] += order.quantity;
                } else {
                    products[productId] = order.quantity;
                }
            });
            let labels = Object.keys(products);
            let values = Object.values(products);
            //截取前四个
            if (labels.length > 4) {
                labels.splice(4);
                values.splice(4);
            }
            this.chartData = {
                labels,
                datasets: [
                    {
                        label: '总销售量',
                        data: values,
                        backgroundColor: ['#41B883', '#E46651', '#00D8FF', '#DD1B16'],
                        borderWidth: 0,
                    },
                ],
            };
        },
    },
};
</script>
