<script>
	import 'tw-elements';
	import CveSeverityChart from '../lib/CveSeverityChart.svelte';

	let timePeriods = ['24h', '7d', '30d', 'ytd', '1y'];
	let tabTitles = ['24 hours', '7 days', '30 days', 'YTD', '1 year'];
</script>

<ul
	class="nav nav-tabs flex flex-col md:flex-row flex-wrap list-none mb-4"
	id="tabs-tab"
	role="tablist"
>
	{#each timePeriods as timePeriod, i}
		<li class="nav-item" role="presentation">
			<a
				href="#tabs-{timePeriod}"
				class="nav-link block text-xs uppercase border-b-2 border-transparent px-5 py-3 hover:bg-bg3 hover:dark:bg-bg3-dark focus:border-transparent
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
			<CveSeverityChart {timePeriod} />
		</div>
	{/each}
</div>
