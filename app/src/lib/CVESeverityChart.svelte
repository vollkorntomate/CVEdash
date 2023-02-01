<script>
	import { Chart, Title, BarElement, CategoryScale, LinearScale } from 'chart.js';
	import ChartDataLabels from 'chartjs-plugin-datalabels';
	import Bar from 'svelte-chartjs/Bar.svelte';

	Chart.register(Title, BarElement, CategoryScale, LinearScale, ChartDataLabels);

	export let allData = {};
	export let timePeriod = '24h';

	let chartOptions = {
		responsive: true,
		maintainAspectRatio: true,
		animations: false,
		plugins: {
			datalabels: {
				anchor: 'end',
				align: 'top'
			}
		}
	};

	$: dataset = {
		labels: ['No CVSS', 'Low', 'Medium', 'High', 'Critical'],
		datasets: [
			{
				data: allData[timePeriod],
				backgroundColor: ['#a3a3a3', '#fde047', '#fb923c', '#dc2626', '#7f1d1d']
			}
		]
	};
</script>

<Bar data={dataset} options={chartOptions} />
