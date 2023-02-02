<script>
	// @ts-nocheck // to disable error that chartOptions is incompatible with Chart.options due to plugins field
	import { Chart, BarController, BarElement, LinearScale, CategoryScale } from 'chart.js';
	import ChartDataLabels from 'chartjs-plugin-datalabels';
	import { onMount } from 'svelte';

	Chart.register(BarController, BarElement, LinearScale, CategoryScale, ChartDataLabels);

	export let allData = {};
	export let timePeriod = '';

	let chartCtx;
	let chart;

	let chartOptions = {
		responsive: true,
		maintainAspectRatio: false,
		animations: false,
		layout: {
			padding: {
				top: 24 // adjust for datalabels
			}
		},
		plugins: {
			datalabels: {
				anchor: 'end',
				align: 'top'
			}
		}
	};

	let dataset = {
		labels: ['No CVSS', 'Low', 'Medium', 'High', 'Critical'],
		datasets: [
			{
				data: allData[timePeriod],
				backgroundColor: ['#a3a3a3', '#fde047', '#fb923c', '#dc2626', '#7f1d1d']
			}
		]
	};

	$: allData && timePeriod && updateChart();

	onMount(() => {
		chart = new Chart(chartCtx, {
			type: 'bar',
			data: dataset,
			options: chartOptions
		});
	});

	function updateChart() {
		if (chart) {
			chart.data.datasets[0].data = allData[timePeriod];
			chart.update();
		}
	}
</script>

<div class="w-full h-72">
	<canvas bind:this={chartCtx} class="max-w-full" />
</div>
