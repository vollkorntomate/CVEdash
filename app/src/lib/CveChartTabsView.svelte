<script>
	import { onMount } from 'svelte';
	import CveSeverityChart from '$lib/CveSeverityChart.svelte';

	// import tw-elements with no SSR (https://github.com/mdbootstrap/Tailwind-Elements/issues/1053)
	onMount(async () => {
		await import('tw-elements');
	});

	let timePeriods = ['24h', '7d', '30d', 'ytd', '1y', 'alltime'];
	let tabTitles = ['24 hours', '7 days', '30 days', 'YTD', '1 year', 'All Time'];

	let allData = {};
	onMount(fetchStats);

	async function fetchStats() {
		await fetch('http://localhost:8077/stats')
			.then((r) => r.json())
			.then((obj) => {
				Object.keys(obj).forEach((k, i) => {
					allData[k] = [obj[k].unknown, obj[k].low, obj[k].medium, obj[k].high, obj[k].critical];
				});
			});
	}
</script>

<div class="bg-bg3 dark:bg-bg3-dark rounded-lg px-2">
	<ul
		class="nav nav-tabs flex flex-wrap justify-center list-none mb-4"
		id="tabs-tab"
		role="tablist"
	>
		{#each timePeriods as timePeriod, i}
			<li class="nav-item" role="presentation">
				<a
					href="#tabs-{timePeriod}"
					class="nav-link block text-xs uppercase border-b-2 border-transparent px-2 py-2 lg:px-5 lg:py-3 hover:bg-bg2 hover:dark:bg-bg3-dark focus:border-transparent
					{i === 0 ? 'active' : ''}"
					id="tabs-{timePeriod}-tab"
					data-bs-toggle="pill"
					data-bs-target="#tabs-{timePeriod}"
					role="tab"
					aria-controls="tabs-{timePeriod}"
					aria-selected={i === 0 ? 'true' : 'false'}>{tabTitles[i]}</a
				>
			</li>
		{/each}
	</ul>
	<div class="tab-content" id="tabs-tabContent">
		{#each timePeriods as timePeriod, i}
			<div
				class="tab-pane fade {i === 0 ? 'show active' : ''}"
				id="tabs-{timePeriod}"
				role="tabpanel"
				aria-labelledby="tabs-{timePeriod}-tab"
			>
				<CveSeverityChart {allData} {timePeriod} />
			</div>
		{/each}
	</div>
</div>
