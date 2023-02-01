<script>
	import { onMount } from 'svelte';
	import CveSummaryView from '../lib/CveSummaryView.svelte';
	import CveChartTabsView from '../lib/CveChartTabsView.svelte';

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

<div class="container sm:mx-auto py-2">
	<div class="flex flex-row flex-wrap md:flex-nowrap">
		<div class="md:basis-1/3 lg:basis-1/4 order-2 md:order-1 m-2">
			<div class="dashboard-card max-h-screen-padded overflow-y-scroll no-scrollbar">
				{#if latestCVEs.length == 0}
					<CveSummaryView
						cveID="NO CONTENT"
						cveDescription="Failed to fetch the newest CVEs. Is the API running?"
						published={new Date().toISOString()}
						cvssScore={10.0}
						cvssSeverity="CRITICAL"
						cvssVector="-/-"
						cvssSource="vollkorntomate"
					/>
				{/if}
				{#each latestCVEs as cve}
					<CveSummaryView
						cveID={cve.id}
						cveDescription={cve.description}
						published={cve.published}
						cvssScore={cve.cvssBaseScore}
						cvssSeverity={cve.cvssBaseSeverity}
						cvssVector={cve.cvssVector}
						cvssSource={cve.cvssSource}
					/>
				{/each}
				<div class="text-center">
					<button on:click={loadMoreLatest} class="p-2 text-neutral-100 bg-blue-600 rounded-md">
						Load more
					</button>
				</div>
			</div>
		</div>

		<div class="md:basis-1/3 lg:basis-1/2 order-1 md:order-2 m-2">
			<div class="dashboard-card">
				<CveChartTabsView />
			</div>
		</div>

		<div class="md:basis-1/3 lg:basis-1/4 order-3 m-2">
			<div class="dashboard-card">
				<CveSummaryView
					cveID="twitter"
					cveDescription=""
					published=""
					cvssScore={0.0}
					cvssSeverity=""
					cvssVector=""
					cvssSource=""
				/>
				<CveSummaryView
					cveID="mastodon"
					cveDescription=""
					published=""
					cvssScore={0.0}
					cvssSeverity=""
					cvssVector=""
					cvssSource=""
				/>
				<CveSummaryView
					cveID="and so forth..."
					cveDescription="But of course not as CVESummaryViews..."
					published=""
					cvssScore={0.0}
					cvssSeverity=""
					cvssVector=""
					cvssSource=""
				/>
			</div>
		</div>
	</div>
</div>
