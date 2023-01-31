<script>
	import { onMount } from 'svelte';

	import { Chart, Title, BarElement, CategoryScale, LinearScale } from 'chart.js';
	import Bar from 'svelte-chartjs/Bar.svelte';

	Chart.register(Title, BarElement, CategoryScale, LinearScale);

	export let timePeriod = '24h';

	let data = [0, 0, 0, 0, 0];

	let chartOptions = {
        responsive: true,
        maintainAspectRatio: false
    };

	$: dataset = {
		labels: ['No CVSS', 'Low', 'Medium', 'High', 'Critical'],
		datasets: [
			{
				data: data,
				backgroundColor: ['#a3a3a3', '#fde047', '#fb923c', '#dc2626', '#7f1d1d']
			}
		]
	};

	onMount(fetchStats);

	async function fetchStats() {
		await fetch('http://localhost:8077/stats/' + timePeriod)
			.then((r) => r.json())
			.then((obj) => {
				data = [obj.unknown, obj.low, obj.medium, obj.high, obj.critical];
			});
	}
</script>

<Bar data={dataset} options={chartOptions} />
