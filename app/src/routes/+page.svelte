<script>
	import { onMount } from 'svelte';
	import CveSummaryView from '$lib/CveSummaryView.svelte';
	import CveChartTabsView from '$lib/CveChartTabsView.svelte';
	import { CveSummary } from '$lib/types';

	let latestPage = 1;
	let latestCVEs = [];

	onMount(fetchLatest);

	async function fetchLatest() {
		await fetch('http://localhost:8077/latest/' + latestPage)
			.then((r) => r.json())
			.then((data) => {
				latestCVEs = latestCVEs.concat(data);
			});
	}

	async function loadMoreLatest() {
		latestPage += 1;
		await fetchLatest();
	}
</script>

<svelte:head>
	<title>CVE Dashboard</title>
</svelte:head>

<div class="flex flex-wrap md:flex-nowrap gap-4">
	<div class="flex md:basis-1/2 lg:basis-2/5 order-2 md:order-1">
		<!-- use "h-screen md:h-full" here if wanting "double scrolling" on mobile -->
		<div class="flex flex-col dashboard-card">
			<div class="dashboard-card-title">
				<span>Latest published CVEs</span>
			</div>
			<div class="overflow-y-scroll no-scrollbar">
				{#if latestCVEs.length > 0}
					{#each latestCVEs as cve}
						<CveSummaryView {cve} />
					{/each}
					<div class="text-center">
						<button on:click={loadMoreLatest} class="p-2 text-neutral-100 bg-blue-600 rounded-md">
							Load more
						</button>
					</div>
				{:else}
					<CveSummaryView
						cve={new CveSummary(
							'NO CONTENT',
							'Failed to fetch the newest CVEs. Is the API running?',
							new Date().toISOString(),
							'CRITICAL',
							10.0,
							'-/-',
							'vollkorntomate'
						)}
					/>
				{/if}
			</div>
		</div>
	</div>

	<div class="md:basis-1/2 lg:basis-3/5 w-full order-1 md:order-2">
		<div class="dashboard-card">
			<div class="dashboard-card-title">
				<span>CVSS Score Statistics</span>
			</div>
			<CveChartTabsView />
		</div>
	</div>
</div>
